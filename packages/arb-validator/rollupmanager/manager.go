/*
* Copyright 2020, Offchain Labs, Inc.
*
* Licensed under the Apache License, Version 2.0 (the "License");
* you may not use this file except in compliance with the License.
* You may obtain a copy of the License at
*
*    http://www.apache.org/licenses/LICENSE-2.0
*
* Unless required by applicable law or agreed to in writing, software
* distributed under the License is distributed on an "AS IS" BASIS,
* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
* See the License for the specific language governing permissions and
* limitations under the License.
 */

package rollupmanager

import (
	"context"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/observer"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/chainlistener"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/chainobserver"
	errors2 "github.com/pkg/errors"
	"log"
	"math/big"
	"sync"
	"time"

	"github.com/offchainlabs/arbitrum/packages/arb-checkpointer/checkpointing"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/arbbridge"
)

type Manager struct {
	// The mutex should be held whenever listeres or reorgCache are accessed or
	// set and whenever activeChain is going to be set (but not when it is accessed)
	sync.Mutex

	listeners   []chainlistener.ChainListener
	activeChain *chainobserver.ChainObserver

	// These variables are only written by the constructor
	RollupAddress common.Address
	checkpointer  checkpointing.RollupCheckpointer
}

const defaultMaxReorgDepth = 100

const assumedValidThreshold = 2

func CreateManager(
	ctx context.Context,
	rollupAddr common.Address,
	clnt arbbridge.ArbClient,
	aoFilePath string,
	dbPath string,
) (*Manager, error) {
	checkpointer, err := checkpointing.NewIndexedCheckpointer(
		rollupAddr,
		dbPath,
		big.NewInt(defaultMaxReorgDepth),
		false,
	)
	if err != nil {
		return nil, err
	}
	return CreateManagerAdvanced(
		ctx,
		rollupAddr,
		true,
		clnt,
		checkpointer,
		aoFilePath,
	)
}

func CreateManagerAdvanced(
	ctx context.Context,
	rollupAddr common.Address,
	updateOpinion bool,
	clnt arbbridge.ArbClient,
	checkpointer checkpointing.RollupCheckpointer,
	aoFilePath string,
) (*Manager, error) {
	if !checkpointer.Initialized() {
		if err := checkpointer.Initialize(aoFilePath); err != nil {
			return nil, err
		}
	}
	initialMachine, err := checkpointer.GetInitialMachine()
	if err != nil {
		return nil, err
	}

	rollupWatcher, err := clnt.NewRollupWatcher(rollupAddr)
	if err != nil {
		return nil, err
	}

	if err := rollupWatcher.VerifyArbChain(ctx, initialMachine.Hash()); err != nil {
		return nil, err
	}

	man := &Manager{
		RollupAddress: rollupAddr,
		checkpointer:  checkpointer,
	}
	go func() {
		for {
			runCtx, cancelFunc := context.WithCancel(ctx)

			rollupWatcher, err := clnt.NewRollupWatcher(rollupAddr)
			if err != nil {
				log.Fatal(err)
			}

			inboxAddr, err := rollupWatcher.InboxAddress(runCtx)
			if err != nil {
				log.Fatal(err)
			}

			inboxWatcher, err := clnt.NewGlobalInboxWatcher(inboxAddr, rollupAddr)
			if err != nil {
				log.Fatal(err)
			}

			chain, err := chainobserver.NewChainObserver(
				runCtx,
				rollupAddr,
				updateOpinion,
				clnt,
				rollupWatcher,
				checkpointer,
				assumedValidThreshold,
			)
			if err != nil {
				log.Fatal(err)
			}

			man.Lock()
			man.activeChain = chain
			// Add manager's listeners
			for _, listener := range man.listeners {
				man.activeChain.AddListener(runCtx, listener)
			}
			man.Unlock()

			time.Sleep(time.Second) // give time for things to settle, post-reorg, before restarting stuff

			log.Println("Starting validator from", man.activeChain.CurrentEventId().BlockId)

			man.activeChain.RestartFromLatestValid(runCtx)

			man.activeChain.Start(runCtx)

			caughtUpToL1 := false

			err = func() error {
				// If the local chain is significantly behind the L1, catch up
				// more efficiently. We process `MaxReorgHeight` blocks at a
				// time up to `MaxReorgHeight` blocks before the current head
				// and and assume that no reorg will occur affecting the blocks
				// we are processing
				maxReorg := checkpointer.MaxReorgHeight()
				for {
					if err := man.activeChain.UpdateAssumedValidBlock(runCtx, clnt, assumedValidThreshold); err != nil {
						return err
					}
					nextEventId := man.activeChain.CurrentEventId()
					startHeight := nextEventId.BlockId.Height.AsInt()

					fetchEnd, err := observer.CalculateCatchupFetch(runCtx, startHeight, clnt, maxReorg)
					if err != nil {
						return errors2.Wrap(err, "error calculating fast catchup")
					}
					if fetchEnd == nil {
						break
					}

					log.Println("Getting events between", startHeight, "and", fetchEnd)
					inboxDeliveredEvents, err := inboxWatcher.GetDeliveredEvents(runCtx, startHeight, fetchEnd)
					if err != nil {
						return errors2.Wrap(err, "Manager hit error doing fast catchup")
					}
					inboxEvents := make([]arbbridge.Event, 0, len(inboxDeliveredEvents))
					for _, ev := range inboxDeliveredEvents {
						inboxEvents = append(inboxEvents, ev)
					}
					events, err := rollupWatcher.GetAllEvents(runCtx, startHeight, fetchEnd)
					if err != nil {
						return errors2.Wrap(err, "Manager hit error doing fast catchup")
					}
					allEvents := arbbridge.MergeEventsUnsafe(inboxEvents, events)
					for _, ev := range allEvents {
						if ev.GetChainInfo().Cmp(nextEventId) < 0 {
							// Start at the matching event id
							continue
						}
						blockId := ev.GetChainInfo().BlockId
						if blockId.Height.AsInt().Cmp(startHeight) > 0 {
							man.activeChain.NotifyNewBlock(blockId.Clone())
							startHeight = blockId.Height.AsInt()
						}
						err := man.activeChain.HandleNotification(runCtx, ev)
						if err != nil {
							return errors2.Wrap(err, "Manager hit error processing event during fast catchup")
						}
					}
					if fetchEnd.Cmp(startHeight) > 0 {
						endBlockId, err := clnt.BlockIdForHeight(runCtx, common.NewTimeBlocks(fetchEnd))
						if err != nil {
							return err
						}
						man.activeChain.NotifyNewBlock(endBlockId)
					}
				}

				startEventId := man.activeChain.CurrentEventId()
				headersChan, err := clnt.SubscribeBlockHeaders(runCtx, startEventId.BlockId)
				if err != nil {
					return errors2.Wrap(err, "Error subscribing to block headers")
				}

				lastDebugPrint := time.Now()
				for maybeBlockId := range headersChan {
					if maybeBlockId.Err != nil {
						return errors2.Wrap(maybeBlockId.Err, "Error getting new header")
					}

					if err := man.activeChain.UpdateAssumedValidBlock(runCtx, clnt, assumedValidThreshold); err != nil {
						return err
					}

					blockId := maybeBlockId.BlockId
					timestamp := maybeBlockId.Timestamp

					currentOnChain, err := clnt.CurrentBlockId(runCtx)
					if err != nil {
						return err
					}

					if !caughtUpToL1 && blockId.Height.Cmp(currentOnChain.Height) >= 0 {
						caughtUpToL1 = true
						man.activeChain.NowAtHead()
						log.Println("Now at head")
					}

					man.activeChain.NotifyNewBlock(blockId.Clone())

					if caughtUpToL1 || time.Since(lastDebugPrint).Seconds() > 5 {
						log.Print(man.activeChain.DebugString("== "))
						lastDebugPrint = time.Now()
					}

					inboxEvents, err := inboxWatcher.GetEvents(runCtx, blockId, timestamp)
					if err != nil {
						return errors2.Wrapf(err, "Manager hit error getting inbox events with block %v", blockId)
					}
					// It's safe to process inbox events out of order with rollup events as long
					// as the inbox events are ahead of the rollup ones
					for _, ev := range inboxEvents {
						if ev.GetChainInfo().Cmp(startEventId) < 0 {
							// Start at the matching event id
							continue
						}
						err := man.activeChain.HandleNotification(runCtx, ev)
						if err != nil {
							return errors2.Wrap(err, "Manager hit error processing event")
						}
					}

					events, err := rollupWatcher.GetEvents(runCtx, blockId, timestamp)
					if err != nil {
						return errors2.Wrapf(err, "Manager hit error getting rollup events with block %v", blockId)
					}

					for _, ev := range events {
						if ev.GetChainInfo().Cmp(startEventId) < 0 {
							// Start at the matching event id
							continue
						}
						err := man.activeChain.HandleNotification(runCtx, ev)
						if err != nil {
							return errors2.Wrap(err, "Manager hit error processing event")
						}
					}
				}
				return nil
			}()

			if err != nil {
				log.Println(err)
			}

			man.Lock()
			man.activeChain = nil
			man.Unlock()

			cancelFunc()

			select {
			case <-ctx.Done():
				return
			default:

			}
		}
	}()

	return man, nil
}

func (man *Manager) AddListener(listener chainlistener.ChainListener) {
	man.Lock()
	defer man.Unlock()
	man.listeners = append(man.listeners, listener)
	if man.activeChain != nil {
		man.activeChain.AddListener(context.Background(), listener)
	}
}

func (man *Manager) GetCheckpointer() checkpointing.RollupCheckpointer {
	return man.checkpointer
}
