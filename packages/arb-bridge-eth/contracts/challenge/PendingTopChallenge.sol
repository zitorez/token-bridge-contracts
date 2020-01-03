/*
 * Copyright 2019, Offchain Labs, Inc.
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

pragma solidity ^0.5.3;

import "./BisectionChallenge.sol";
import "./ChallengeUtils.sol";

import "../arch/Protocol.sol";


contract PendingTopChallenge is BisectionChallenge {

    event Bisected(
        bytes32[] chainHashes,
        uint totalLength,
        uint256 deadlineTicks
    );

    event OneStepProofCompleted();

    // Proof was incorrect
    string constant HC_OSP_PROOF = "HC_OSP_PROOF";

    function bisect(
        bytes32[] memory _chainHashes,
        uint _chainLength
    )
        public
        asserterAction
    {
        uint bisectionCount = _chainHashes.length - 1;

        requireMatchesPrevState(
            ChallengeUtils.pendingTopHash(
                _chainHashes[0],
                _chainHashes[bisectionCount],
                _chainLength
            )
        );

        require(_chainLength > 1, "Can't bisect chain of less than 2");
        bytes32[] memory hashes = new bytes32[](bisectionCount);
        hashes[0] = ChallengeUtils.pendingTopHash(
            _chainHashes[0],
            _chainHashes[1],
            firstSegmentSize(_chainLength, bisectionCount)
        );
        for (uint i = 1; i < bisectionCount; i++) {
            hashes[i] = ChallengeUtils.pendingTopHash(
                _chainHashes[i],
                _chainHashes[i + 1],
                otherSegmentSize(_chainLength, bisectionCount)
            );
        }

        commitToSegment(hashes);
        asserterResponded();
        emit Bisected(
            _chainHashes,
            _chainLength,
            deadlineTicks
        );
    }

    function oneStepProof(
        bytes32 _lowerHash,
        bytes32 _topHash,
        bytes32 _value
    )
        public
        asserterAction
    {
        requireMatchesPrevState(
            ChallengeUtils.pendingTopHash(
                _lowerHash,
                _topHash,
                1
            )
        );

        require(Protocol.addMessageToPending(_lowerHash, _value) == _topHash, HC_OSP_PROOF);

        emit OneStepProofCompleted();
        _asserterWin();
    }
}
