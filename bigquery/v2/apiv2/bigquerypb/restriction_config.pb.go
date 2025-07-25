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
// source: google/cloud/bigquery/v2/restriction_config.proto

package bigquerypb

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// RestrictionType specifies the type of dataset/table restriction.
type RestrictionConfig_RestrictionType int32

const (
	// Should never be used.
	RestrictionConfig_RESTRICTION_TYPE_UNSPECIFIED RestrictionConfig_RestrictionType = 0
	// Restrict data egress. See [Data
	// egress](https://cloud.google.com/bigquery/docs/analytics-hub-introduction#data_egress)
	// for more details.
	RestrictionConfig_RESTRICTED_DATA_EGRESS RestrictionConfig_RestrictionType = 1
)

// Enum value maps for RestrictionConfig_RestrictionType.
var (
	RestrictionConfig_RestrictionType_name = map[int32]string{
		0: "RESTRICTION_TYPE_UNSPECIFIED",
		1: "RESTRICTED_DATA_EGRESS",
	}
	RestrictionConfig_RestrictionType_value = map[string]int32{
		"RESTRICTION_TYPE_UNSPECIFIED": 0,
		"RESTRICTED_DATA_EGRESS":       1,
	}
)

func (x RestrictionConfig_RestrictionType) Enum() *RestrictionConfig_RestrictionType {
	p := new(RestrictionConfig_RestrictionType)
	*p = x
	return p
}

func (x RestrictionConfig_RestrictionType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RestrictionConfig_RestrictionType) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_bigquery_v2_restriction_config_proto_enumTypes[0].Descriptor()
}

func (RestrictionConfig_RestrictionType) Type() protoreflect.EnumType {
	return &file_google_cloud_bigquery_v2_restriction_config_proto_enumTypes[0]
}

func (x RestrictionConfig_RestrictionType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RestrictionConfig_RestrictionType.Descriptor instead.
func (RestrictionConfig_RestrictionType) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_bigquery_v2_restriction_config_proto_rawDescGZIP(), []int{0, 0}
}

type RestrictionConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. Specifies the type of dataset/table restriction.
	Type RestrictionConfig_RestrictionType `protobuf:"varint,1,opt,name=type,proto3,enum=google.cloud.bigquery.v2.RestrictionConfig_RestrictionType" json:"type,omitempty"`
}

func (x *RestrictionConfig) Reset() {
	*x = RestrictionConfig{}
	mi := &file_google_cloud_bigquery_v2_restriction_config_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RestrictionConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RestrictionConfig) ProtoMessage() {}

func (x *RestrictionConfig) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_bigquery_v2_restriction_config_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RestrictionConfig.ProtoReflect.Descriptor instead.
func (*RestrictionConfig) Descriptor() ([]byte, []int) {
	return file_google_cloud_bigquery_v2_restriction_config_proto_rawDescGZIP(), []int{0}
}

func (x *RestrictionConfig) GetType() RestrictionConfig_RestrictionType {
	if x != nil {
		return x.Type
	}
	return RestrictionConfig_RESTRICTION_TYPE_UNSPECIFIED
}

var File_google_cloud_bigquery_v2_restriction_config_proto protoreflect.FileDescriptor

var file_google_cloud_bigquery_v2_restriction_config_proto_rawDesc = []byte{
	0x0a, 0x31, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x62,
	0x69, 0x67, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2f, 0x76, 0x32, 0x2f, 0x72, 0x65, 0x73, 0x74, 0x72,
	0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x18, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x62, 0x69, 0x67, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x76, 0x32, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f,
	0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xba,
	0x01, 0x0a, 0x11, 0x52, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x12, 0x54, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x3b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x62, 0x69, 0x67, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x76, 0x32, 0x2e, 0x52, 0x65,
	0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e,
	0x52, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x42,
	0x03, 0xe0, 0x41, 0x03, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x4f, 0x0a, 0x0f, 0x52, 0x65,
	0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x20, 0x0a,
	0x1c, 0x52, 0x45, 0x53, 0x54, 0x52, 0x49, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12,
	0x1a, 0x0a, 0x16, 0x52, 0x45, 0x53, 0x54, 0x52, 0x49, 0x43, 0x54, 0x45, 0x44, 0x5f, 0x44, 0x41,
	0x54, 0x41, 0x5f, 0x45, 0x47, 0x52, 0x45, 0x53, 0x53, 0x10, 0x01, 0x42, 0x75, 0x0a, 0x1c, 0x63,
	0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x62, 0x69, 0x67, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x76, 0x32, 0x42, 0x16, 0x52, 0x65, 0x73,
	0x74, 0x72, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3b, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x62, 0x69, 0x67, 0x71, 0x75,
	0x65, 0x72, 0x79, 0x2f, 0x76, 0x32, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x62, 0x69, 0x67,
	0x71, 0x75, 0x65, 0x72, 0x79, 0x70, 0x62, 0x3b, 0x62, 0x69, 0x67, 0x71, 0x75, 0x65, 0x72, 0x79,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_bigquery_v2_restriction_config_proto_rawDescOnce sync.Once
	file_google_cloud_bigquery_v2_restriction_config_proto_rawDescData = file_google_cloud_bigquery_v2_restriction_config_proto_rawDesc
)

func file_google_cloud_bigquery_v2_restriction_config_proto_rawDescGZIP() []byte {
	file_google_cloud_bigquery_v2_restriction_config_proto_rawDescOnce.Do(func() {
		file_google_cloud_bigquery_v2_restriction_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_bigquery_v2_restriction_config_proto_rawDescData)
	})
	return file_google_cloud_bigquery_v2_restriction_config_proto_rawDescData
}

var file_google_cloud_bigquery_v2_restriction_config_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_cloud_bigquery_v2_restriction_config_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_cloud_bigquery_v2_restriction_config_proto_goTypes = []any{
	(RestrictionConfig_RestrictionType)(0), // 0: google.cloud.bigquery.v2.RestrictionConfig.RestrictionType
	(*RestrictionConfig)(nil),              // 1: google.cloud.bigquery.v2.RestrictionConfig
}
var file_google_cloud_bigquery_v2_restriction_config_proto_depIdxs = []int32{
	0, // 0: google.cloud.bigquery.v2.RestrictionConfig.type:type_name -> google.cloud.bigquery.v2.RestrictionConfig.RestrictionType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_google_cloud_bigquery_v2_restriction_config_proto_init() }
func file_google_cloud_bigquery_v2_restriction_config_proto_init() {
	if File_google_cloud_bigquery_v2_restriction_config_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_cloud_bigquery_v2_restriction_config_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_bigquery_v2_restriction_config_proto_goTypes,
		DependencyIndexes: file_google_cloud_bigquery_v2_restriction_config_proto_depIdxs,
		EnumInfos:         file_google_cloud_bigquery_v2_restriction_config_proto_enumTypes,
		MessageInfos:      file_google_cloud_bigquery_v2_restriction_config_proto_msgTypes,
	}.Build()
	File_google_cloud_bigquery_v2_restriction_config_proto = out.File
	file_google_cloud_bigquery_v2_restriction_config_proto_rawDesc = nil
	file_google_cloud_bigquery_v2_restriction_config_proto_goTypes = nil
	file_google_cloud_bigquery_v2_restriction_config_proto_depIdxs = nil
}
