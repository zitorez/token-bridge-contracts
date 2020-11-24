//
// Copyright 2019, Offchain Labs, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: checkpointing.proto

package checkpointing

import (
	proto "github.com/golang/protobuf/proto"
	ckptcontext "github.com/offchainlabs/arbitrum/packages/arb-checkpointer/ckptcontext"
	common "github.com/offchainlabs/arbitrum/packages/arb-util/common"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type CheckpointMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FormatVersion uint64             `protobuf:"varint,1,opt,name=formatVersion,proto3" json:"formatVersion,omitempty"`
	Oldest        *common.BlockIdBuf `protobuf:"bytes,2,opt,name=oldest,proto3" json:"oldest,omitempty"`
	Newest        *common.BlockIdBuf `protobuf:"bytes,3,opt,name=newest,proto3" json:"newest,omitempty"`
}

func (x *CheckpointMetadata) Reset() {
	*x = CheckpointMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_checkpointing_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckpointMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckpointMetadata) ProtoMessage() {}

func (x *CheckpointMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_checkpointing_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckpointMetadata.ProtoReflect.Descriptor instead.
func (*CheckpointMetadata) Descriptor() ([]byte, []int) {
	return file_checkpointing_proto_rawDescGZIP(), []int{0}
}

func (x *CheckpointMetadata) GetFormatVersion() uint64 {
	if x != nil {
		return x.FormatVersion
	}
	return 0
}

func (x *CheckpointMetadata) GetOldest() *common.BlockIdBuf {
	if x != nil {
		return x.Oldest
	}
	return nil
}

func (x *CheckpointMetadata) GetNewest() *common.BlockIdBuf {
	if x != nil {
		return x.Newest
	}
	return nil
}

type CheckpointLinks struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Next *common.BlockIdBuf `protobuf:"bytes,1,opt,name=next,proto3" json:"next,omitempty"`
	Prev *common.BlockIdBuf `protobuf:"bytes,2,opt,name=prev,proto3" json:"prev,omitempty"`
}

func (x *CheckpointLinks) Reset() {
	*x = CheckpointLinks{}
	if protoimpl.UnsafeEnabled {
		mi := &file_checkpointing_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckpointLinks) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckpointLinks) ProtoMessage() {}

func (x *CheckpointLinks) ProtoReflect() protoreflect.Message {
	mi := &file_checkpointing_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckpointLinks.ProtoReflect.Descriptor instead.
func (*CheckpointLinks) Descriptor() ([]byte, []int) {
	return file_checkpointing_proto_rawDescGZIP(), []int{1}
}

func (x *CheckpointLinks) GetNext() *common.BlockIdBuf {
	if x != nil {
		return x.Next
	}
	return nil
}

func (x *CheckpointLinks) GetPrev() *common.BlockIdBuf {
	if x != nil {
		return x.Prev
	}
	return nil
}

type HeightBoundsBuf struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Lo *common.TimeBlocksBuf `protobuf:"bytes,1,opt,name=lo,proto3" json:"lo,omitempty"`
	Hi *common.TimeBlocksBuf `protobuf:"bytes,2,opt,name=hi,proto3" json:"hi,omitempty"`
}

func (x *HeightBoundsBuf) Reset() {
	*x = HeightBoundsBuf{}
	if protoimpl.UnsafeEnabled {
		mi := &file_checkpointing_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HeightBoundsBuf) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeightBoundsBuf) ProtoMessage() {}

func (x *HeightBoundsBuf) ProtoReflect() protoreflect.Message {
	mi := &file_checkpointing_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeightBoundsBuf.ProtoReflect.Descriptor instead.
func (*HeightBoundsBuf) Descriptor() ([]byte, []int) {
	return file_checkpointing_proto_rawDescGZIP(), []int{2}
}

func (x *HeightBoundsBuf) GetLo() *common.TimeBlocksBuf {
	if x != nil {
		return x.Lo
	}
	return nil
}

func (x *HeightBoundsBuf) GetHi() *common.TimeBlocksBuf {
	if x != nil {
		return x.Hi
	}
	return nil
}

type CheckpointWithManifest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Contents []byte                          `protobuf:"bytes,1,opt,name=contents,proto3" json:"contents,omitempty"`
	Manifest *ckptcontext.CheckpointManifest `protobuf:"bytes,2,opt,name=manifest,proto3" json:"manifest,omitempty"`
}

func (x *CheckpointWithManifest) Reset() {
	*x = CheckpointWithManifest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_checkpointing_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckpointWithManifest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckpointWithManifest) ProtoMessage() {}

func (x *CheckpointWithManifest) ProtoReflect() protoreflect.Message {
	mi := &file_checkpointing_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckpointWithManifest.ProtoReflect.Descriptor instead.
func (*CheckpointWithManifest) Descriptor() ([]byte, []int) {
	return file_checkpointing_proto_rawDescGZIP(), []int{3}
}

func (x *CheckpointWithManifest) GetContents() []byte {
	if x != nil {
		return x.Contents
	}
	return nil
}

func (x *CheckpointWithManifest) GetManifest() *ckptcontext.CheckpointManifest {
	if x != nil {
		return x.Manifest
	}
	return nil
}

var File_checkpointing_proto protoreflect.FileDescriptor

var file_checkpointing_proto_rawDesc = []byte{
	0x0a, 0x13, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65,
	0x73, 0x1a, 0x1c, 0x61, 0x72, 0x62, 0x2d, 0x75, 0x74, 0x69, 0x6c, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x2e, 0x61, 0x72, 0x62, 0x2d, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x2f, 0x63, 0x6b, 0x70, 0x74, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x2f, 0x63, 0x6b,
	0x70, 0x74, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x92, 0x01, 0x0a, 0x12, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x24, 0x0a, 0x0d, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74,
	0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0d, 0x66,
	0x6f, 0x72, 0x6d, 0x61, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x2a, 0x0a, 0x06,
	0x6f, 0x6c, 0x64, 0x65, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x64, 0x42, 0x75, 0x66,
	0x52, 0x06, 0x6f, 0x6c, 0x64, 0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x06, 0x6e, 0x65, 0x77, 0x65,
	0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x64, 0x42, 0x75, 0x66, 0x52, 0x06, 0x6e, 0x65,
	0x77, 0x65, 0x73, 0x74, 0x22, 0x61, 0x0a, 0x0f, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x73, 0x12, 0x26, 0x0a, 0x04, 0x6e, 0x65, 0x78, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x42,
	0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x64, 0x42, 0x75, 0x66, 0x52, 0x04, 0x6e, 0x65, 0x78, 0x74, 0x12,
	0x26, 0x0a, 0x04, 0x70, 0x72, 0x65, 0x76, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x64, 0x42, 0x75,
	0x66, 0x52, 0x04, 0x70, 0x72, 0x65, 0x76, 0x22, 0x5f, 0x0a, 0x0f, 0x48, 0x65, 0x69, 0x67, 0x68,
	0x74, 0x42, 0x6f, 0x75, 0x6e, 0x64, 0x73, 0x42, 0x75, 0x66, 0x12, 0x25, 0x0a, 0x02, 0x6c, 0x6f,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x42, 0x75, 0x66, 0x52, 0x02, 0x6c,
	0x6f, 0x12, 0x25, 0x0a, 0x02, 0x68, 0x69, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x42, 0x6c, 0x6f, 0x63, 0x6b,
	0x73, 0x42, 0x75, 0x66, 0x52, 0x02, 0x68, 0x69, 0x22, 0x71, 0x0a, 0x16, 0x43, 0x68, 0x65, 0x63,
	0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x57, 0x69, 0x74, 0x68, 0x4d, 0x61, 0x6e, 0x69, 0x66, 0x65,
	0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x3b,
	0x0a, 0x08, 0x6d, 0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1f, 0x2e, 0x63, 0x6b, 0x70, 0x74, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x2e, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x4d, 0x61, 0x6e, 0x69, 0x66, 0x65, 0x73,
	0x74, 0x52, 0x08, 0x6d, 0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x42, 0x4a, 0x5a, 0x48, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x66, 0x66, 0x63, 0x68, 0x61,
	0x69, 0x6e, 0x6c, 0x61, 0x62, 0x73, 0x2f, 0x61, 0x72, 0x62, 0x69, 0x74, 0x72, 0x75, 0x6d, 0x2f,
	0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x61, 0x72, 0x62, 0x2d, 0x63, 0x68, 0x65,
	0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_checkpointing_proto_rawDescOnce sync.Once
	file_checkpointing_proto_rawDescData = file_checkpointing_proto_rawDesc
)

func file_checkpointing_proto_rawDescGZIP() []byte {
	file_checkpointing_proto_rawDescOnce.Do(func() {
		file_checkpointing_proto_rawDescData = protoimpl.X.CompressGZIP(file_checkpointing_proto_rawDescData)
	})
	return file_checkpointing_proto_rawDescData
}

var file_checkpointing_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_checkpointing_proto_goTypes = []interface{}{
	(*CheckpointMetadata)(nil),             // 0: structures.CheckpointMetadata
	(*CheckpointLinks)(nil),                // 1: structures.CheckpointLinks
	(*HeightBoundsBuf)(nil),                // 2: structures.HeightBoundsBuf
	(*CheckpointWithManifest)(nil),         // 3: structures.CheckpointWithManifest
	(*common.BlockIdBuf)(nil),              // 4: common.BlockIdBuf
	(*common.TimeBlocksBuf)(nil),           // 5: common.TimeBlocksBuf
	(*ckptcontext.CheckpointManifest)(nil), // 6: ckptcontext.CheckpointManifest
}
var file_checkpointing_proto_depIdxs = []int32{
	4, // 0: structures.CheckpointMetadata.oldest:type_name -> common.BlockIdBuf
	4, // 1: structures.CheckpointMetadata.newest:type_name -> common.BlockIdBuf
	4, // 2: structures.CheckpointLinks.next:type_name -> common.BlockIdBuf
	4, // 3: structures.CheckpointLinks.prev:type_name -> common.BlockIdBuf
	5, // 4: structures.HeightBoundsBuf.lo:type_name -> common.TimeBlocksBuf
	5, // 5: structures.HeightBoundsBuf.hi:type_name -> common.TimeBlocksBuf
	6, // 6: structures.CheckpointWithManifest.manifest:type_name -> ckptcontext.CheckpointManifest
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_checkpointing_proto_init() }
func file_checkpointing_proto_init() {
	if File_checkpointing_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_checkpointing_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckpointMetadata); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_checkpointing_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckpointLinks); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_checkpointing_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HeightBoundsBuf); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_checkpointing_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckpointWithManifest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_checkpointing_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_checkpointing_proto_goTypes,
		DependencyIndexes: file_checkpointing_proto_depIdxs,
		MessageInfos:      file_checkpointing_proto_msgTypes,
	}.Build()
	File_checkpointing_proto = out.File
	file_checkpointing_proto_rawDesc = nil
	file_checkpointing_proto_goTypes = nil
	file_checkpointing_proto_depIdxs = nil
}
