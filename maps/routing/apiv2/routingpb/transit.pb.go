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
// source: google/maps/routing/v2/transit.proto

package routingpb

import (
	localized_text "google.golang.org/genproto/googleapis/type/localized_text"
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

// The type of vehicles for transit routes.
type TransitVehicle_TransitVehicleType int32

const (
	// Unused.
	TransitVehicle_TRANSIT_VEHICLE_TYPE_UNSPECIFIED TransitVehicle_TransitVehicleType = 0
	// Bus.
	TransitVehicle_BUS TransitVehicle_TransitVehicleType = 1
	// A vehicle that operates on a cable, usually on the ground. Aerial cable
	// cars may be of the type `GONDOLA_LIFT`.
	TransitVehicle_CABLE_CAR TransitVehicle_TransitVehicleType = 2
	// Commuter rail.
	TransitVehicle_COMMUTER_TRAIN TransitVehicle_TransitVehicleType = 3
	// Ferry.
	TransitVehicle_FERRY TransitVehicle_TransitVehicleType = 4
	// A vehicle that is pulled up a steep incline by a cable. A Funicular
	// typically consists of two cars, with each car acting as a counterweight
	// for the other.
	TransitVehicle_FUNICULAR TransitVehicle_TransitVehicleType = 5
	// An aerial cable car.
	TransitVehicle_GONDOLA_LIFT TransitVehicle_TransitVehicleType = 6
	// Heavy rail.
	TransitVehicle_HEAVY_RAIL TransitVehicle_TransitVehicleType = 7
	// High speed train.
	TransitVehicle_HIGH_SPEED_TRAIN TransitVehicle_TransitVehicleType = 8
	// Intercity bus.
	TransitVehicle_INTERCITY_BUS TransitVehicle_TransitVehicleType = 9
	// Long distance train.
	TransitVehicle_LONG_DISTANCE_TRAIN TransitVehicle_TransitVehicleType = 10
	// Light rail transit.
	TransitVehicle_METRO_RAIL TransitVehicle_TransitVehicleType = 11
	// Monorail.
	TransitVehicle_MONORAIL TransitVehicle_TransitVehicleType = 12
	// All other vehicles.
	TransitVehicle_OTHER TransitVehicle_TransitVehicleType = 13
	// Rail.
	TransitVehicle_RAIL TransitVehicle_TransitVehicleType = 14
	// Share taxi is a kind of bus with the ability to drop off and pick up
	// passengers anywhere on its route.
	TransitVehicle_SHARE_TAXI TransitVehicle_TransitVehicleType = 15
	// Underground light rail.
	TransitVehicle_SUBWAY TransitVehicle_TransitVehicleType = 16
	// Above ground light rail.
	TransitVehicle_TRAM TransitVehicle_TransitVehicleType = 17
	// Trolleybus.
	TransitVehicle_TROLLEYBUS TransitVehicle_TransitVehicleType = 18
)

// Enum value maps for TransitVehicle_TransitVehicleType.
var (
	TransitVehicle_TransitVehicleType_name = map[int32]string{
		0:  "TRANSIT_VEHICLE_TYPE_UNSPECIFIED",
		1:  "BUS",
		2:  "CABLE_CAR",
		3:  "COMMUTER_TRAIN",
		4:  "FERRY",
		5:  "FUNICULAR",
		6:  "GONDOLA_LIFT",
		7:  "HEAVY_RAIL",
		8:  "HIGH_SPEED_TRAIN",
		9:  "INTERCITY_BUS",
		10: "LONG_DISTANCE_TRAIN",
		11: "METRO_RAIL",
		12: "MONORAIL",
		13: "OTHER",
		14: "RAIL",
		15: "SHARE_TAXI",
		16: "SUBWAY",
		17: "TRAM",
		18: "TROLLEYBUS",
	}
	TransitVehicle_TransitVehicleType_value = map[string]int32{
		"TRANSIT_VEHICLE_TYPE_UNSPECIFIED": 0,
		"BUS":                              1,
		"CABLE_CAR":                        2,
		"COMMUTER_TRAIN":                   3,
		"FERRY":                            4,
		"FUNICULAR":                        5,
		"GONDOLA_LIFT":                     6,
		"HEAVY_RAIL":                       7,
		"HIGH_SPEED_TRAIN":                 8,
		"INTERCITY_BUS":                    9,
		"LONG_DISTANCE_TRAIN":              10,
		"METRO_RAIL":                       11,
		"MONORAIL":                         12,
		"OTHER":                            13,
		"RAIL":                             14,
		"SHARE_TAXI":                       15,
		"SUBWAY":                           16,
		"TRAM":                             17,
		"TROLLEYBUS":                       18,
	}
)

func (x TransitVehicle_TransitVehicleType) Enum() *TransitVehicle_TransitVehicleType {
	p := new(TransitVehicle_TransitVehicleType)
	*p = x
	return p
}

func (x TransitVehicle_TransitVehicleType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TransitVehicle_TransitVehicleType) Descriptor() protoreflect.EnumDescriptor {
	return file_google_maps_routing_v2_transit_proto_enumTypes[0].Descriptor()
}

func (TransitVehicle_TransitVehicleType) Type() protoreflect.EnumType {
	return &file_google_maps_routing_v2_transit_proto_enumTypes[0]
}

func (x TransitVehicle_TransitVehicleType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TransitVehicle_TransitVehicleType.Descriptor instead.
func (TransitVehicle_TransitVehicleType) EnumDescriptor() ([]byte, []int) {
	return file_google_maps_routing_v2_transit_proto_rawDescGZIP(), []int{3, 0}
}

// A transit agency that operates a transit line.
type TransitAgency struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of this transit agency.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The transit agency's locale-specific formatted phone number.
	PhoneNumber string `protobuf:"bytes,2,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	// The transit agency's URI.
	Uri string `protobuf:"bytes,3,opt,name=uri,proto3" json:"uri,omitempty"`
}

func (x *TransitAgency) Reset() {
	*x = TransitAgency{}
	mi := &file_google_maps_routing_v2_transit_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TransitAgency) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransitAgency) ProtoMessage() {}

func (x *TransitAgency) ProtoReflect() protoreflect.Message {
	mi := &file_google_maps_routing_v2_transit_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransitAgency.ProtoReflect.Descriptor instead.
func (*TransitAgency) Descriptor() ([]byte, []int) {
	return file_google_maps_routing_v2_transit_proto_rawDescGZIP(), []int{0}
}

func (x *TransitAgency) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TransitAgency) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *TransitAgency) GetUri() string {
	if x != nil {
		return x.Uri
	}
	return ""
}

// Contains information about the transit line used in this step.
type TransitLine struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The transit agency (or agencies) that operates this transit line.
	Agencies []*TransitAgency `protobuf:"bytes,1,rep,name=agencies,proto3" json:"agencies,omitempty"`
	// The full name of this transit line, For example, "8 Avenue Local".
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// the URI for this transit line as provided by the transit agency.
	Uri string `protobuf:"bytes,3,opt,name=uri,proto3" json:"uri,omitempty"`
	// The color commonly used in signage for this line. Represented in
	// hexadecimal.
	Color string `protobuf:"bytes,4,opt,name=color,proto3" json:"color,omitempty"`
	// The URI for the icon associated with this line.
	IconUri string `protobuf:"bytes,5,opt,name=icon_uri,json=iconUri,proto3" json:"icon_uri,omitempty"`
	// The short name of this transit line. This name will normally be a line
	// number, such as "M7" or "355".
	NameShort string `protobuf:"bytes,6,opt,name=name_short,json=nameShort,proto3" json:"name_short,omitempty"`
	// The color commonly used in text on signage for this line. Represented in
	// hexadecimal.
	TextColor string `protobuf:"bytes,7,opt,name=text_color,json=textColor,proto3" json:"text_color,omitempty"`
	// The type of vehicle that operates on this transit line.
	Vehicle *TransitVehicle `protobuf:"bytes,8,opt,name=vehicle,proto3" json:"vehicle,omitempty"`
}

func (x *TransitLine) Reset() {
	*x = TransitLine{}
	mi := &file_google_maps_routing_v2_transit_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TransitLine) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransitLine) ProtoMessage() {}

func (x *TransitLine) ProtoReflect() protoreflect.Message {
	mi := &file_google_maps_routing_v2_transit_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransitLine.ProtoReflect.Descriptor instead.
func (*TransitLine) Descriptor() ([]byte, []int) {
	return file_google_maps_routing_v2_transit_proto_rawDescGZIP(), []int{1}
}

func (x *TransitLine) GetAgencies() []*TransitAgency {
	if x != nil {
		return x.Agencies
	}
	return nil
}

func (x *TransitLine) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TransitLine) GetUri() string {
	if x != nil {
		return x.Uri
	}
	return ""
}

func (x *TransitLine) GetColor() string {
	if x != nil {
		return x.Color
	}
	return ""
}

func (x *TransitLine) GetIconUri() string {
	if x != nil {
		return x.IconUri
	}
	return ""
}

func (x *TransitLine) GetNameShort() string {
	if x != nil {
		return x.NameShort
	}
	return ""
}

func (x *TransitLine) GetTextColor() string {
	if x != nil {
		return x.TextColor
	}
	return ""
}

func (x *TransitLine) GetVehicle() *TransitVehicle {
	if x != nil {
		return x.Vehicle
	}
	return nil
}

// Information about a transit stop.
type TransitStop struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the transit stop.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The location of the stop expressed in latitude/longitude coordinates.
	Location *Location `protobuf:"bytes,2,opt,name=location,proto3" json:"location,omitempty"`
}

func (x *TransitStop) Reset() {
	*x = TransitStop{}
	mi := &file_google_maps_routing_v2_transit_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TransitStop) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransitStop) ProtoMessage() {}

func (x *TransitStop) ProtoReflect() protoreflect.Message {
	mi := &file_google_maps_routing_v2_transit_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransitStop.ProtoReflect.Descriptor instead.
func (*TransitStop) Descriptor() ([]byte, []int) {
	return file_google_maps_routing_v2_transit_proto_rawDescGZIP(), []int{2}
}

func (x *TransitStop) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TransitStop) GetLocation() *Location {
	if x != nil {
		return x.Location
	}
	return nil
}

// Information about a vehicle used in transit routes.
type TransitVehicle struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of this vehicle, capitalized.
	Name *localized_text.LocalizedText `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The type of vehicle used.
	Type TransitVehicle_TransitVehicleType `protobuf:"varint,2,opt,name=type,proto3,enum=google.maps.routing.v2.TransitVehicle_TransitVehicleType" json:"type,omitempty"`
	// The URI for an icon associated with this vehicle type.
	IconUri string `protobuf:"bytes,3,opt,name=icon_uri,json=iconUri,proto3" json:"icon_uri,omitempty"`
	// The URI for the icon associated with this vehicle type, based on the local
	// transport signage.
	LocalIconUri string `protobuf:"bytes,4,opt,name=local_icon_uri,json=localIconUri,proto3" json:"local_icon_uri,omitempty"`
}

func (x *TransitVehicle) Reset() {
	*x = TransitVehicle{}
	mi := &file_google_maps_routing_v2_transit_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TransitVehicle) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransitVehicle) ProtoMessage() {}

func (x *TransitVehicle) ProtoReflect() protoreflect.Message {
	mi := &file_google_maps_routing_v2_transit_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransitVehicle.ProtoReflect.Descriptor instead.
func (*TransitVehicle) Descriptor() ([]byte, []int) {
	return file_google_maps_routing_v2_transit_proto_rawDescGZIP(), []int{3}
}

func (x *TransitVehicle) GetName() *localized_text.LocalizedText {
	if x != nil {
		return x.Name
	}
	return nil
}

func (x *TransitVehicle) GetType() TransitVehicle_TransitVehicleType {
	if x != nil {
		return x.Type
	}
	return TransitVehicle_TRANSIT_VEHICLE_TYPE_UNSPECIFIED
}

func (x *TransitVehicle) GetIconUri() string {
	if x != nil {
		return x.IconUri
	}
	return ""
}

func (x *TransitVehicle) GetLocalIconUri() string {
	if x != nil {
		return x.LocalIconUri
	}
	return ""
}

var File_google_maps_routing_v2_transit_proto protoreflect.FileDescriptor

var file_google_maps_routing_v2_transit_proto_rawDesc = []byte{
	0x0a, 0x24, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x6d, 0x61, 0x70, 0x73, 0x2f, 0x72, 0x6f,
	0x75, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x32, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d,
	0x61, 0x70, 0x73, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x32, 0x1a, 0x25,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x6d, 0x61, 0x70, 0x73, 0x2f, 0x72, 0x6f, 0x75, 0x74,
	0x69, 0x6e, 0x67, 0x2f, 0x76, 0x32, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x74, 0x79,
	0x70, 0x65, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x5f, 0x74, 0x65, 0x78,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x58, 0x0a, 0x0d, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x69, 0x74, 0x41, 0x67, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c,
	0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12,
	0x10, 0x0a, 0x03, 0x75, 0x72, 0x69, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72,
	0x69, 0x22, 0xa7, 0x02, 0x0a, 0x0b, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x74, 0x4c, 0x69, 0x6e,
	0x65, 0x12, 0x41, 0x0a, 0x08, 0x61, 0x67, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d, 0x61, 0x70,
	0x73, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x32, 0x2e, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x69, 0x74, 0x41, 0x67, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x08, 0x61, 0x67, 0x65, 0x6e,
	0x63, 0x69, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x69, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x69, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f,
	0x6c, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72,
	0x12, 0x19, 0x0a, 0x08, 0x69, 0x63, 0x6f, 0x6e, 0x5f, 0x75, 0x72, 0x69, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x69, 0x63, 0x6f, 0x6e, 0x55, 0x72, 0x69, 0x12, 0x1d, 0x0a, 0x0a, 0x6e,
	0x61, 0x6d, 0x65, 0x5f, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x6e, 0x61, 0x6d, 0x65, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x65,
	0x78, 0x74, 0x5f, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x74, 0x65, 0x78, 0x74, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x12, 0x40, 0x0a, 0x07, 0x76, 0x65, 0x68,
	0x69, 0x63, 0x6c, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x6d, 0x61, 0x70, 0x73, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67,
	0x2e, 0x76, 0x32, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x74, 0x56, 0x65, 0x68, 0x69, 0x63,
	0x6c, 0x65, 0x52, 0x07, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x22, 0x5f, 0x0a, 0x0b, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x69, 0x74, 0x53, 0x74, 0x6f, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x3c,
	0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x20, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d, 0x61, 0x70, 0x73, 0x2e, 0x72,
	0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x32, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xa0, 0x04, 0x0a,
	0x0e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x74, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x12,
	0x2e, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x4c, 0x6f, 0x63, 0x61,
	0x6c, 0x69, 0x7a, 0x65, 0x64, 0x54, 0x65, 0x78, 0x74, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x4d, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x39, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d, 0x61, 0x70, 0x73, 0x2e, 0x72, 0x6f, 0x75, 0x74,
	0x69, 0x6e, 0x67, 0x2e, 0x76, 0x32, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x74, 0x56, 0x65,
	0x68, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x74, 0x56, 0x65, 0x68,
	0x69, 0x63, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x19,
	0x0a, 0x08, 0x69, 0x63, 0x6f, 0x6e, 0x5f, 0x75, 0x72, 0x69, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x69, 0x63, 0x6f, 0x6e, 0x55, 0x72, 0x69, 0x12, 0x24, 0x0a, 0x0e, 0x6c, 0x6f, 0x63,
	0x61, 0x6c, 0x5f, 0x69, 0x63, 0x6f, 0x6e, 0x5f, 0x75, 0x72, 0x69, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x49, 0x63, 0x6f, 0x6e, 0x55, 0x72, 0x69, 0x22,
	0xcd, 0x02, 0x0a, 0x12, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x74, 0x56, 0x65, 0x68, 0x69, 0x63,
	0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x24, 0x0a, 0x20, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x49,
	0x54, 0x5f, 0x56, 0x45, 0x48, 0x49, 0x43, 0x4c, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55,
	0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03,
	0x42, 0x55, 0x53, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x43, 0x41, 0x42, 0x4c, 0x45, 0x5f, 0x43,
	0x41, 0x52, 0x10, 0x02, 0x12, 0x12, 0x0a, 0x0e, 0x43, 0x4f, 0x4d, 0x4d, 0x55, 0x54, 0x45, 0x52,
	0x5f, 0x54, 0x52, 0x41, 0x49, 0x4e, 0x10, 0x03, 0x12, 0x09, 0x0a, 0x05, 0x46, 0x45, 0x52, 0x52,
	0x59, 0x10, 0x04, 0x12, 0x0d, 0x0a, 0x09, 0x46, 0x55, 0x4e, 0x49, 0x43, 0x55, 0x4c, 0x41, 0x52,
	0x10, 0x05, 0x12, 0x10, 0x0a, 0x0c, 0x47, 0x4f, 0x4e, 0x44, 0x4f, 0x4c, 0x41, 0x5f, 0x4c, 0x49,
	0x46, 0x54, 0x10, 0x06, 0x12, 0x0e, 0x0a, 0x0a, 0x48, 0x45, 0x41, 0x56, 0x59, 0x5f, 0x52, 0x41,
	0x49, 0x4c, 0x10, 0x07, 0x12, 0x14, 0x0a, 0x10, 0x48, 0x49, 0x47, 0x48, 0x5f, 0x53, 0x50, 0x45,
	0x45, 0x44, 0x5f, 0x54, 0x52, 0x41, 0x49, 0x4e, 0x10, 0x08, 0x12, 0x11, 0x0a, 0x0d, 0x49, 0x4e,
	0x54, 0x45, 0x52, 0x43, 0x49, 0x54, 0x59, 0x5f, 0x42, 0x55, 0x53, 0x10, 0x09, 0x12, 0x17, 0x0a,
	0x13, 0x4c, 0x4f, 0x4e, 0x47, 0x5f, 0x44, 0x49, 0x53, 0x54, 0x41, 0x4e, 0x43, 0x45, 0x5f, 0x54,
	0x52, 0x41, 0x49, 0x4e, 0x10, 0x0a, 0x12, 0x0e, 0x0a, 0x0a, 0x4d, 0x45, 0x54, 0x52, 0x4f, 0x5f,
	0x52, 0x41, 0x49, 0x4c, 0x10, 0x0b, 0x12, 0x0c, 0x0a, 0x08, 0x4d, 0x4f, 0x4e, 0x4f, 0x52, 0x41,
	0x49, 0x4c, 0x10, 0x0c, 0x12, 0x09, 0x0a, 0x05, 0x4f, 0x54, 0x48, 0x45, 0x52, 0x10, 0x0d, 0x12,
	0x08, 0x0a, 0x04, 0x52, 0x41, 0x49, 0x4c, 0x10, 0x0e, 0x12, 0x0e, 0x0a, 0x0a, 0x53, 0x48, 0x41,
	0x52, 0x45, 0x5f, 0x54, 0x41, 0x58, 0x49, 0x10, 0x0f, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x55, 0x42,
	0x57, 0x41, 0x59, 0x10, 0x10, 0x12, 0x08, 0x0a, 0x04, 0x54, 0x52, 0x41, 0x4d, 0x10, 0x11, 0x12,
	0x0e, 0x0a, 0x0a, 0x54, 0x52, 0x4f, 0x4c, 0x4c, 0x45, 0x59, 0x42, 0x55, 0x53, 0x10, 0x12, 0x42,
	0xbe, 0x01, 0x0a, 0x1a, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d,
	0x61, 0x70, 0x73, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x32, 0x42, 0x0c,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3a,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x67, 0x6f, 0x2f, 0x6d, 0x61, 0x70, 0x73, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67,
	0x2f, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x70, 0x62,
	0x3b, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x70, 0x62, 0xa2, 0x02, 0x05, 0x47, 0x4d, 0x52,
	0x56, 0x32, 0xaa, 0x02, 0x16, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x4d, 0x61, 0x70, 0x73,
	0x2e, 0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x56, 0x32, 0xca, 0x02, 0x16, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x4d, 0x61, 0x70, 0x73, 0x5c, 0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e,
	0x67, 0x5c, 0x56, 0x32, 0xea, 0x02, 0x19, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x4d,
	0x61, 0x70, 0x73, 0x3a, 0x3a, 0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x3a, 0x3a, 0x56, 0x32,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_maps_routing_v2_transit_proto_rawDescOnce sync.Once
	file_google_maps_routing_v2_transit_proto_rawDescData = file_google_maps_routing_v2_transit_proto_rawDesc
)

func file_google_maps_routing_v2_transit_proto_rawDescGZIP() []byte {
	file_google_maps_routing_v2_transit_proto_rawDescOnce.Do(func() {
		file_google_maps_routing_v2_transit_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_maps_routing_v2_transit_proto_rawDescData)
	})
	return file_google_maps_routing_v2_transit_proto_rawDescData
}

var file_google_maps_routing_v2_transit_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_maps_routing_v2_transit_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_google_maps_routing_v2_transit_proto_goTypes = []any{
	(TransitVehicle_TransitVehicleType)(0), // 0: google.maps.routing.v2.TransitVehicle.TransitVehicleType
	(*TransitAgency)(nil),                  // 1: google.maps.routing.v2.TransitAgency
	(*TransitLine)(nil),                    // 2: google.maps.routing.v2.TransitLine
	(*TransitStop)(nil),                    // 3: google.maps.routing.v2.TransitStop
	(*TransitVehicle)(nil),                 // 4: google.maps.routing.v2.TransitVehicle
	(*Location)(nil),                       // 5: google.maps.routing.v2.Location
	(*localized_text.LocalizedText)(nil),   // 6: google.type.LocalizedText
}
var file_google_maps_routing_v2_transit_proto_depIdxs = []int32{
	1, // 0: google.maps.routing.v2.TransitLine.agencies:type_name -> google.maps.routing.v2.TransitAgency
	4, // 1: google.maps.routing.v2.TransitLine.vehicle:type_name -> google.maps.routing.v2.TransitVehicle
	5, // 2: google.maps.routing.v2.TransitStop.location:type_name -> google.maps.routing.v2.Location
	6, // 3: google.maps.routing.v2.TransitVehicle.name:type_name -> google.type.LocalizedText
	0, // 4: google.maps.routing.v2.TransitVehicle.type:type_name -> google.maps.routing.v2.TransitVehicle.TransitVehicleType
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_google_maps_routing_v2_transit_proto_init() }
func file_google_maps_routing_v2_transit_proto_init() {
	if File_google_maps_routing_v2_transit_proto != nil {
		return
	}
	file_google_maps_routing_v2_location_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_maps_routing_v2_transit_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_maps_routing_v2_transit_proto_goTypes,
		DependencyIndexes: file_google_maps_routing_v2_transit_proto_depIdxs,
		EnumInfos:         file_google_maps_routing_v2_transit_proto_enumTypes,
		MessageInfos:      file_google_maps_routing_v2_transit_proto_msgTypes,
	}.Build()
	File_google_maps_routing_v2_transit_proto = out.File
	file_google_maps_routing_v2_transit_proto_rawDesc = nil
	file_google_maps_routing_v2_transit_proto_goTypes = nil
	file_google_maps_routing_v2_transit_proto_depIdxs = nil
}
