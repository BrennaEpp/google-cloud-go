// Copyright 2025 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v4.25.7
// source: google/maps/addressvalidation/v1/metadata.proto

package addressvalidationpb

import (
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

// The metadata for the post-processed address. `metadata` is not guaranteed to
// be fully populated for every address sent to the Address Validation API.
type AddressMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Indicates that this is the address of a business.
	// If unset, indicates that the value is unknown.
	Business *bool `protobuf:"varint,2,opt,name=business,proto3,oneof" json:"business,omitempty"`
	// Indicates that the address of a PO box.
	// If unset, indicates that the value is unknown.
	PoBox *bool `protobuf:"varint,3,opt,name=po_box,json=poBox,proto3,oneof" json:"po_box,omitempty"`
	// Indicates that this is the address of a residence.
	// If unset, indicates that the value is unknown.
	Residential *bool `protobuf:"varint,6,opt,name=residential,proto3,oneof" json:"residential,omitempty"`
}

func (x *AddressMetadata) Reset() {
	*x = AddressMetadata{}
	mi := &file_google_maps_addressvalidation_v1_metadata_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddressMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddressMetadata) ProtoMessage() {}

func (x *AddressMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_google_maps_addressvalidation_v1_metadata_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddressMetadata.ProtoReflect.Descriptor instead.
func (*AddressMetadata) Descriptor() ([]byte, []int) {
	return file_google_maps_addressvalidation_v1_metadata_proto_rawDescGZIP(), []int{0}
}

func (x *AddressMetadata) GetBusiness() bool {
	if x != nil && x.Business != nil {
		return *x.Business
	}
	return false
}

func (x *AddressMetadata) GetPoBox() bool {
	if x != nil && x.PoBox != nil {
		return *x.PoBox
	}
	return false
}

func (x *AddressMetadata) GetResidential() bool {
	if x != nil && x.Residential != nil {
		return *x.Residential
	}
	return false
}

var File_google_maps_addressvalidation_v1_metadata_proto protoreflect.FileDescriptor

var file_google_maps_addressvalidation_v1_metadata_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x6d, 0x61, 0x70, 0x73, 0x2f, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f,
	0x76, 0x31, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d, 0x61, 0x70, 0x73, 0x2e, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x76, 0x31, 0x22, 0x9d, 0x01, 0x0a, 0x0f, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1f, 0x0a, 0x08, 0x62, 0x75, 0x73, 0x69, 0x6e,
	0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x08, 0x62, 0x75, 0x73,
	0x69, 0x6e, 0x65, 0x73, 0x73, 0x88, 0x01, 0x01, 0x12, 0x1a, 0x0a, 0x06, 0x70, 0x6f, 0x5f, 0x62,
	0x6f, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x48, 0x01, 0x52, 0x05, 0x70, 0x6f, 0x42, 0x6f,
	0x78, 0x88, 0x01, 0x01, 0x12, 0x25, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x69, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x61, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x48, 0x02, 0x52, 0x0b, 0x72, 0x65, 0x73,
	0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x88, 0x01, 0x01, 0x42, 0x0b, 0x0a, 0x09, 0x5f,
	0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x70, 0x6f, 0x5f,
	0x62, 0x6f, 0x78, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x72, 0x65, 0x73, 0x69, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x61, 0x6c, 0x42, 0x87, 0x02, 0x0a, 0x24, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x6d, 0x61, 0x70, 0x73, 0x2e, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x42, 0x0d, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x58, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x67, 0x6f, 0x2f, 0x6d, 0x61, 0x70, 0x73, 0x2f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x2f,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x70, 0x62, 0x3b, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x70, 0x62, 0xa2, 0x02, 0x07, 0x47, 0x4d, 0x50, 0x41, 0x56, 0x56,
	0x31, 0xaa, 0x02, 0x20, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x4d, 0x61, 0x70, 0x73, 0x2e,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x20, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x4d, 0x61,
	0x70, 0x73, 0x5c, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5c, 0x56, 0x31, 0xea, 0x02, 0x23, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x3a, 0x3a, 0x4d, 0x61, 0x70, 0x73, 0x3a, 0x3a, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x56,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_maps_addressvalidation_v1_metadata_proto_rawDescOnce sync.Once
	file_google_maps_addressvalidation_v1_metadata_proto_rawDescData = file_google_maps_addressvalidation_v1_metadata_proto_rawDesc
)

func file_google_maps_addressvalidation_v1_metadata_proto_rawDescGZIP() []byte {
	file_google_maps_addressvalidation_v1_metadata_proto_rawDescOnce.Do(func() {
		file_google_maps_addressvalidation_v1_metadata_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_maps_addressvalidation_v1_metadata_proto_rawDescData)
	})
	return file_google_maps_addressvalidation_v1_metadata_proto_rawDescData
}

var file_google_maps_addressvalidation_v1_metadata_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_maps_addressvalidation_v1_metadata_proto_goTypes = []any{
	(*AddressMetadata)(nil), // 0: google.maps.addressvalidation.v1.AddressMetadata
}
var file_google_maps_addressvalidation_v1_metadata_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_maps_addressvalidation_v1_metadata_proto_init() }
func file_google_maps_addressvalidation_v1_metadata_proto_init() {
	if File_google_maps_addressvalidation_v1_metadata_proto != nil {
		return
	}
	file_google_maps_addressvalidation_v1_metadata_proto_msgTypes[0].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_maps_addressvalidation_v1_metadata_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_maps_addressvalidation_v1_metadata_proto_goTypes,
		DependencyIndexes: file_google_maps_addressvalidation_v1_metadata_proto_depIdxs,
		MessageInfos:      file_google_maps_addressvalidation_v1_metadata_proto_msgTypes,
	}.Build()
	File_google_maps_addressvalidation_v1_metadata_proto = out.File
	file_google_maps_addressvalidation_v1_metadata_proto_rawDesc = nil
	file_google_maps_addressvalidation_v1_metadata_proto_goTypes = nil
	file_google_maps_addressvalidation_v1_metadata_proto_depIdxs = nil
}
