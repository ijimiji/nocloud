//
//Copyright © 2021-2022 Nikita Ivanovski info@slnt-opp.xyz
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.1
// source: pkg/dns/proto/dns.proto

package proto

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

type Record struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	A     []*Record_A     `protobuf:"bytes,1,rep,name=a,proto3" json:"a,omitempty"`
	Aaaa  []*Record_AAAA  `protobuf:"bytes,2,rep,name=aaaa,proto3" json:"aaaa,omitempty"`
	Cname []*Record_CNAME `protobuf:"bytes,3,rep,name=cname,proto3" json:"cname,omitempty"`
	Txt   []*Record_TXT   `protobuf:"bytes,4,rep,name=txt,proto3" json:"txt,omitempty"`
}

func (x *Record) Reset() {
	*x = Record{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_dns_proto_dns_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Record) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Record) ProtoMessage() {}

func (x *Record) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_dns_proto_dns_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Record.ProtoReflect.Descriptor instead.
func (*Record) Descriptor() ([]byte, []int) {
	return file_pkg_dns_proto_dns_proto_rawDescGZIP(), []int{0}
}

func (x *Record) GetA() []*Record_A {
	if x != nil {
		return x.A
	}
	return nil
}

func (x *Record) GetAaaa() []*Record_AAAA {
	if x != nil {
		return x.Aaaa
	}
	return nil
}

func (x *Record) GetCname() []*Record_CNAME {
	if x != nil {
		return x.Cname
	}
	return nil
}

func (x *Record) GetTxt() []*Record_TXT {
	if x != nil {
		return x.Txt
	}
	return nil
}

type Zone struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string             `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Locations map[string]*Record `protobuf:"bytes,2,rep,name=locations,proto3" json:"locations,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Zone) Reset() {
	*x = Zone{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_dns_proto_dns_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Zone) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Zone) ProtoMessage() {}

func (x *Zone) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_dns_proto_dns_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Zone.ProtoReflect.Descriptor instead.
func (*Zone) Descriptor() ([]byte, []int) {
	return file_pkg_dns_proto_dns_proto_rawDescGZIP(), []int{1}
}

func (x *Zone) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Zone) GetLocations() map[string]*Record {
	if x != nil {
		return x.Locations
	}
	return nil
}

type ListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListRequest) Reset() {
	*x = ListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_dns_proto_dns_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRequest) ProtoMessage() {}

func (x *ListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_dns_proto_dns_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRequest.ProtoReflect.Descriptor instead.
func (*ListRequest) Descriptor() ([]byte, []int) {
	return file_pkg_dns_proto_dns_proto_rawDescGZIP(), []int{2}
}

type ListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Zones []string `protobuf:"bytes,1,rep,name=zones,proto3" json:"zones,omitempty"`
}

func (x *ListResponse) Reset() {
	*x = ListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_dns_proto_dns_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListResponse) ProtoMessage() {}

func (x *ListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_dns_proto_dns_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListResponse.ProtoReflect.Descriptor instead.
func (*ListResponse) Descriptor() ([]byte, []int) {
	return file_pkg_dns_proto_dns_proto_rawDescGZIP(), []int{3}
}

func (x *ListResponse) GetZones() []string {
	if x != nil {
		return x.Zones
	}
	return nil
}

type Result struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result int32 `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *Result) Reset() {
	*x = Result{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_dns_proto_dns_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Result) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Result) ProtoMessage() {}

func (x *Result) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_dns_proto_dns_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Result.ProtoReflect.Descriptor instead.
func (*Result) Descriptor() ([]byte, []int) {
	return file_pkg_dns_proto_dns_proto_rawDescGZIP(), []int{4}
}

func (x *Result) GetResult() int32 {
	if x != nil {
		return x.Result
	}
	return 0
}

type Record_A struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ip  string `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"`
	Ttl int32  `protobuf:"varint,2,opt,name=ttl,proto3" json:"ttl,omitempty"`
}

func (x *Record_A) Reset() {
	*x = Record_A{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_dns_proto_dns_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Record_A) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Record_A) ProtoMessage() {}

func (x *Record_A) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_dns_proto_dns_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Record_A.ProtoReflect.Descriptor instead.
func (*Record_A) Descriptor() ([]byte, []int) {
	return file_pkg_dns_proto_dns_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Record_A) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *Record_A) GetTtl() int32 {
	if x != nil {
		return x.Ttl
	}
	return 0
}

type Record_AAAA struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ip  string `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"`
	Ttl int32  `protobuf:"varint,2,opt,name=ttl,proto3" json:"ttl,omitempty"`
}

func (x *Record_AAAA) Reset() {
	*x = Record_AAAA{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_dns_proto_dns_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Record_AAAA) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Record_AAAA) ProtoMessage() {}

func (x *Record_AAAA) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_dns_proto_dns_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Record_AAAA.ProtoReflect.Descriptor instead.
func (*Record_AAAA) Descriptor() ([]byte, []int) {
	return file_pkg_dns_proto_dns_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Record_AAAA) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *Record_AAAA) GetTtl() int32 {
	if x != nil {
		return x.Ttl
	}
	return 0
}

type Record_CNAME struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Host string `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	Ttl  int32  `protobuf:"varint,2,opt,name=ttl,proto3" json:"ttl,omitempty"`
}

func (x *Record_CNAME) Reset() {
	*x = Record_CNAME{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_dns_proto_dns_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Record_CNAME) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Record_CNAME) ProtoMessage() {}

func (x *Record_CNAME) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_dns_proto_dns_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Record_CNAME.ProtoReflect.Descriptor instead.
func (*Record_CNAME) Descriptor() ([]byte, []int) {
	return file_pkg_dns_proto_dns_proto_rawDescGZIP(), []int{0, 2}
}

func (x *Record_CNAME) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *Record_CNAME) GetTtl() int32 {
	if x != nil {
		return x.Ttl
	}
	return 0
}

type Record_TXT struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	Ttl  int32  `protobuf:"varint,2,opt,name=ttl,proto3" json:"ttl,omitempty"`
}

func (x *Record_TXT) Reset() {
	*x = Record_TXT{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_dns_proto_dns_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Record_TXT) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Record_TXT) ProtoMessage() {}

func (x *Record_TXT) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_dns_proto_dns_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Record_TXT.ProtoReflect.Descriptor instead.
func (*Record_TXT) Descriptor() ([]byte, []int) {
	return file_pkg_dns_proto_dns_proto_rawDescGZIP(), []int{0, 3}
}

func (x *Record_TXT) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *Record_TXT) GetTtl() int32 {
	if x != nil {
		return x.Ttl
	}
	return 0
}

var File_pkg_dns_proto_dns_proto protoreflect.FileDescriptor

var file_pkg_dns_proto_dns_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x6b, 0x67, 0x2f, 0x64, 0x6e, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x64, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x6e, 0x6f, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x64, 0x6e, 0x73, 0x22, 0xe4, 0x02, 0x0a, 0x06, 0x52, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x12, 0x23, 0x0a, 0x01, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x6e,
	0x6f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x6e, 0x73, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x2e, 0x41, 0x52, 0x01, 0x61, 0x12, 0x2c, 0x0a, 0x04, 0x61, 0x61, 0x61, 0x61, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6e, 0x6f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64,
	0x6e, 0x73, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x2e, 0x41, 0x41, 0x41, 0x41, 0x52, 0x04,
	0x61, 0x61, 0x61, 0x61, 0x12, 0x2f, 0x0a, 0x05, 0x63, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x6e, 0x6f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x6e,
	0x73, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x2e, 0x43, 0x4e, 0x41, 0x4d, 0x45, 0x52, 0x05,
	0x63, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x29, 0x0a, 0x03, 0x74, 0x78, 0x74, 0x18, 0x04, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6e, 0x6f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x6e, 0x73,
	0x2e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x2e, 0x54, 0x58, 0x54, 0x52, 0x03, 0x74, 0x78, 0x74,
	0x1a, 0x25, 0x0a, 0x01, 0x41, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x70, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x74, 0x6c, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x03, 0x74, 0x74, 0x6c, 0x1a, 0x28, 0x0a, 0x04, 0x41, 0x41, 0x41, 0x41, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x70, 0x12,
	0x10, 0x0a, 0x03, 0x74, 0x74, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x74, 0x74,
	0x6c, 0x1a, 0x2d, 0x0a, 0x05, 0x43, 0x4e, 0x41, 0x4d, 0x45, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f,
	0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x10,
	0x0a, 0x03, 0x74, 0x74, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x74, 0x74, 0x6c,
	0x1a, 0x2b, 0x0a, 0x03, 0x54, 0x58, 0x54, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x74,
	0x74, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x74, 0x74, 0x6c, 0x22, 0xad, 0x01,
	0x0a, 0x04, 0x5a, 0x6f, 0x6e, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x3e, 0x0a, 0x09, 0x6c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e,
	0x6e, 0x6f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x6e, 0x73, 0x2e, 0x5a, 0x6f, 0x6e, 0x65,
	0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x09, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x51, 0x0a, 0x0e, 0x4c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x29,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e,
	0x6e, 0x6f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x6e, 0x73, 0x2e, 0x52, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x0d, 0x0a,
	0x0b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x24, 0x0a, 0x0c,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x7a, 0x6f, 0x6e, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x7a, 0x6f, 0x6e,
	0x65, 0x73, 0x22, 0x20, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x32, 0xd0, 0x01, 0x0a, 0x03, 0x44, 0x4e, 0x53, 0x12, 0x2b, 0x0a, 0x03,
	0x47, 0x65, 0x74, 0x12, 0x11, 0x2e, 0x6e, 0x6f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x6e,
	0x73, 0x2e, 0x5a, 0x6f, 0x6e, 0x65, 0x1a, 0x11, 0x2e, 0x6e, 0x6f, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x64, 0x6e, 0x73, 0x2e, 0x5a, 0x6f, 0x6e, 0x65, 0x12, 0x3b, 0x0a, 0x04, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x18, 0x2e, 0x6e, 0x6f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x6e, 0x73, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x6e, 0x6f,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x6e, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x03, 0x50, 0x75, 0x74, 0x12, 0x11, 0x2e,
	0x6e, 0x6f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x6e, 0x73, 0x2e, 0x5a, 0x6f, 0x6e, 0x65,
	0x1a, 0x13, 0x2e, 0x6e, 0x6f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x6e, 0x73, 0x2e, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x30, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12,
	0x11, 0x2e, 0x6e, 0x6f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x6e, 0x73, 0x2e, 0x5a, 0x6f,
	0x6e, 0x65, 0x1a, 0x13, 0x2e, 0x6e, 0x6f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x6e, 0x73,
	0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x42, 0x92, 0x01, 0x0a, 0x0f, 0x63, 0x6f, 0x6d, 0x2e,
	0x6e, 0x6f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x6e, 0x73, 0x42, 0x08, 0x44, 0x6e, 0x73,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6c, 0x6e, 0x74, 0x6f, 0x70, 0x70, 0x2f, 0x6e, 0x6f, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x64, 0x6e, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0xa2, 0x02, 0x03, 0x4e, 0x44, 0x58, 0xaa, 0x02, 0x0b, 0x4e, 0x6f, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x44, 0x6e, 0x73, 0xca, 0x02, 0x0b, 0x4e, 0x6f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5c,
	0x44, 0x6e, 0x73, 0xe2, 0x02, 0x17, 0x4e, 0x6f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x44, 0x6e,
	0x73, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0c,
	0x4e, 0x6f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x44, 0x6e, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_dns_proto_dns_proto_rawDescOnce sync.Once
	file_pkg_dns_proto_dns_proto_rawDescData = file_pkg_dns_proto_dns_proto_rawDesc
)

func file_pkg_dns_proto_dns_proto_rawDescGZIP() []byte {
	file_pkg_dns_proto_dns_proto_rawDescOnce.Do(func() {
		file_pkg_dns_proto_dns_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_dns_proto_dns_proto_rawDescData)
	})
	return file_pkg_dns_proto_dns_proto_rawDescData
}

var file_pkg_dns_proto_dns_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_pkg_dns_proto_dns_proto_goTypes = []interface{}{
	(*Record)(nil),       // 0: nocloud.dns.Record
	(*Zone)(nil),         // 1: nocloud.dns.Zone
	(*ListRequest)(nil),  // 2: nocloud.dns.ListRequest
	(*ListResponse)(nil), // 3: nocloud.dns.ListResponse
	(*Result)(nil),       // 4: nocloud.dns.Result
	(*Record_A)(nil),     // 5: nocloud.dns.Record.A
	(*Record_AAAA)(nil),  // 6: nocloud.dns.Record.AAAA
	(*Record_CNAME)(nil), // 7: nocloud.dns.Record.CNAME
	(*Record_TXT)(nil),   // 8: nocloud.dns.Record.TXT
	nil,                  // 9: nocloud.dns.Zone.LocationsEntry
}
var file_pkg_dns_proto_dns_proto_depIdxs = []int32{
	5,  // 0: nocloud.dns.Record.a:type_name -> nocloud.dns.Record.A
	6,  // 1: nocloud.dns.Record.aaaa:type_name -> nocloud.dns.Record.AAAA
	7,  // 2: nocloud.dns.Record.cname:type_name -> nocloud.dns.Record.CNAME
	8,  // 3: nocloud.dns.Record.txt:type_name -> nocloud.dns.Record.TXT
	9,  // 4: nocloud.dns.Zone.locations:type_name -> nocloud.dns.Zone.LocationsEntry
	0,  // 5: nocloud.dns.Zone.LocationsEntry.value:type_name -> nocloud.dns.Record
	1,  // 6: nocloud.dns.DNS.Get:input_type -> nocloud.dns.Zone
	2,  // 7: nocloud.dns.DNS.List:input_type -> nocloud.dns.ListRequest
	1,  // 8: nocloud.dns.DNS.Put:input_type -> nocloud.dns.Zone
	1,  // 9: nocloud.dns.DNS.Delete:input_type -> nocloud.dns.Zone
	1,  // 10: nocloud.dns.DNS.Get:output_type -> nocloud.dns.Zone
	3,  // 11: nocloud.dns.DNS.List:output_type -> nocloud.dns.ListResponse
	4,  // 12: nocloud.dns.DNS.Put:output_type -> nocloud.dns.Result
	4,  // 13: nocloud.dns.DNS.Delete:output_type -> nocloud.dns.Result
	10, // [10:14] is the sub-list for method output_type
	6,  // [6:10] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_pkg_dns_proto_dns_proto_init() }
func file_pkg_dns_proto_dns_proto_init() {
	if File_pkg_dns_proto_dns_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_dns_proto_dns_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Record); i {
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
		file_pkg_dns_proto_dns_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Zone); i {
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
		file_pkg_dns_proto_dns_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRequest); i {
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
		file_pkg_dns_proto_dns_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListResponse); i {
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
		file_pkg_dns_proto_dns_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Result); i {
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
		file_pkg_dns_proto_dns_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Record_A); i {
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
		file_pkg_dns_proto_dns_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Record_AAAA); i {
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
		file_pkg_dns_proto_dns_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Record_CNAME); i {
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
		file_pkg_dns_proto_dns_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Record_TXT); i {
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
			RawDescriptor: file_pkg_dns_proto_dns_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_dns_proto_dns_proto_goTypes,
		DependencyIndexes: file_pkg_dns_proto_dns_proto_depIdxs,
		MessageInfos:      file_pkg_dns_proto_dns_proto_msgTypes,
	}.Build()
	File_pkg_dns_proto_dns_proto = out.File
	file_pkg_dns_proto_dns_proto_rawDesc = nil
	file_pkg_dns_proto_dns_proto_goTypes = nil
	file_pkg_dns_proto_dns_proto_depIdxs = nil
}
