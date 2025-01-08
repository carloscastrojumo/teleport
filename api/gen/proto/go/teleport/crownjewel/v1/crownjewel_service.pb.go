// Copyright 2024 Gravitational, Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.2
// 	protoc        (unknown)
// source: teleport/crownjewel/v1/crownjewel_service.proto

package crownjewelv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// CrownJewelRequest is a request to create a new CrownJewel.
type CreateCrownJewelRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	CrownJewel    *CrownJewel            `protobuf:"bytes,2,opt,name=crown_jewel,json=crownJewel,proto3" json:"crown_jewel,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateCrownJewelRequest) Reset() {
	*x = CreateCrownJewelRequest{}
	mi := &file_teleport_crownjewel_v1_crownjewel_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateCrownJewelRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCrownJewelRequest) ProtoMessage() {}

func (x *CreateCrownJewelRequest) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_crownjewel_v1_crownjewel_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCrownJewelRequest.ProtoReflect.Descriptor instead.
func (*CreateCrownJewelRequest) Descriptor() ([]byte, []int) {
	return file_teleport_crownjewel_v1_crownjewel_service_proto_rawDescGZIP(), []int{0}
}

func (x *CreateCrownJewelRequest) GetCrownJewel() *CrownJewel {
	if x != nil {
		return x.CrownJewel
	}
	return nil
}

// GetCrownJewelRequest is a request to get a CrownJewel by name.
type GetCrownJewelRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Name is the name of the CrownJewel to get.
	Name          string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetCrownJewelRequest) Reset() {
	*x = GetCrownJewelRequest{}
	mi := &file_teleport_crownjewel_v1_crownjewel_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetCrownJewelRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCrownJewelRequest) ProtoMessage() {}

func (x *GetCrownJewelRequest) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_crownjewel_v1_crownjewel_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCrownJewelRequest.ProtoReflect.Descriptor instead.
func (*GetCrownJewelRequest) Descriptor() ([]byte, []int) {
	return file_teleport_crownjewel_v1_crownjewel_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetCrownJewelRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// ListCrownJewelsRequest is a request to get a list of CrownJewels.
type ListCrownJewelsRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// page_size is the maximum number of items to return.
	// The server may impose a different page size at its discretion.
	PageSize int64 `protobuf:"varint,1,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// page_token is the next_page_token value returned from a previous List request, if any.
	PageToken     string `protobuf:"bytes,2,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListCrownJewelsRequest) Reset() {
	*x = ListCrownJewelsRequest{}
	mi := &file_teleport_crownjewel_v1_crownjewel_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListCrownJewelsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCrownJewelsRequest) ProtoMessage() {}

func (x *ListCrownJewelsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_crownjewel_v1_crownjewel_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCrownJewelsRequest.ProtoReflect.Descriptor instead.
func (*ListCrownJewelsRequest) Descriptor() ([]byte, []int) {
	return file_teleport_crownjewel_v1_crownjewel_service_proto_rawDescGZIP(), []int{2}
}

func (x *ListCrownJewelsRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListCrownJewelsRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

// ListCrownJewelsResponse is a response to GetCrownJewelsRequest.
type ListCrownJewelsResponse struct {
	state       protoimpl.MessageState `protogen:"open.v1"`
	CrownJewels []*CrownJewel          `protobuf:"bytes,1,rep,name=crown_jewels,json=crownJewels,proto3" json:"crown_jewels,omitempty"`
	// Token to retrieve the next page of results, or empty if there are no
	// more results in the list.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListCrownJewelsResponse) Reset() {
	*x = ListCrownJewelsResponse{}
	mi := &file_teleport_crownjewel_v1_crownjewel_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListCrownJewelsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCrownJewelsResponse) ProtoMessage() {}

func (x *ListCrownJewelsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_crownjewel_v1_crownjewel_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCrownJewelsResponse.ProtoReflect.Descriptor instead.
func (*ListCrownJewelsResponse) Descriptor() ([]byte, []int) {
	return file_teleport_crownjewel_v1_crownjewel_service_proto_rawDescGZIP(), []int{3}
}

func (x *ListCrownJewelsResponse) GetCrownJewels() []*CrownJewel {
	if x != nil {
		return x.CrownJewels
	}
	return nil
}

func (x *ListCrownJewelsResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

// UpdateCrownJewelRequest is a request to update an existing CrownJewel.
type UpdateCrownJewelRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	CrownJewel    *CrownJewel            `protobuf:"bytes,2,opt,name=crown_jewel,json=crownJewel,proto3" json:"crown_jewel,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateCrownJewelRequest) Reset() {
	*x = UpdateCrownJewelRequest{}
	mi := &file_teleport_crownjewel_v1_crownjewel_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateCrownJewelRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCrownJewelRequest) ProtoMessage() {}

func (x *UpdateCrownJewelRequest) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_crownjewel_v1_crownjewel_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCrownJewelRequest.ProtoReflect.Descriptor instead.
func (*UpdateCrownJewelRequest) Descriptor() ([]byte, []int) {
	return file_teleport_crownjewel_v1_crownjewel_service_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateCrownJewelRequest) GetCrownJewel() *CrownJewel {
	if x != nil {
		return x.CrownJewel
	}
	return nil
}

// UpsertCrownJewelRequest is a request to upsert a CrownJewel.
type UpsertCrownJewelRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	CrownJewel    *CrownJewel            `protobuf:"bytes,2,opt,name=crown_jewel,json=crownJewel,proto3" json:"crown_jewel,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpsertCrownJewelRequest) Reset() {
	*x = UpsertCrownJewelRequest{}
	mi := &file_teleport_crownjewel_v1_crownjewel_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpsertCrownJewelRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpsertCrownJewelRequest) ProtoMessage() {}

func (x *UpsertCrownJewelRequest) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_crownjewel_v1_crownjewel_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpsertCrownJewelRequest.ProtoReflect.Descriptor instead.
func (*UpsertCrownJewelRequest) Descriptor() ([]byte, []int) {
	return file_teleport_crownjewel_v1_crownjewel_service_proto_rawDescGZIP(), []int{5}
}

func (x *UpsertCrownJewelRequest) GetCrownJewel() *CrownJewel {
	if x != nil {
		return x.CrownJewel
	}
	return nil
}

// DeleteCrownJewelRequest is a request to delete a CrownJewel.
type DeleteCrownJewelRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Name is the name of the CrownJewel to delete.
	Name          string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteCrownJewelRequest) Reset() {
	*x = DeleteCrownJewelRequest{}
	mi := &file_teleport_crownjewel_v1_crownjewel_service_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteCrownJewelRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCrownJewelRequest) ProtoMessage() {}

func (x *DeleteCrownJewelRequest) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_crownjewel_v1_crownjewel_service_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCrownJewelRequest.ProtoReflect.Descriptor instead.
func (*DeleteCrownJewelRequest) Descriptor() ([]byte, []int) {
	return file_teleport_crownjewel_v1_crownjewel_service_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteCrownJewelRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_teleport_crownjewel_v1_crownjewel_service_proto protoreflect.FileDescriptor

var file_teleport_crownjewel_v1_crownjewel_service_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x63, 0x72, 0x6f, 0x77, 0x6e,
	0x6a, 0x65, 0x77, 0x65, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x72, 0x6f, 0x77, 0x6e, 0x6a, 0x65,
	0x77, 0x65, 0x6c, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x16, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x63, 0x72, 0x6f, 0x77,
	0x6e, 0x6a, 0x65, 0x77, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x27, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x2f, 0x63, 0x72, 0x6f, 0x77, 0x6e, 0x6a, 0x65, 0x77, 0x65, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x63,
	0x72, 0x6f, 0x77, 0x6e, 0x6a, 0x65, 0x77, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x72, 0x0a, 0x17, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x72, 0x6f, 0x77, 0x6e, 0x4a, 0x65,
	0x77, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x43, 0x0a, 0x0b, 0x63, 0x72,
	0x6f, 0x77, 0x6e, 0x5f, 0x6a, 0x65, 0x77, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x22, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x63, 0x72, 0x6f, 0x77, 0x6e,
	0x6a, 0x65, 0x77, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x6f, 0x77, 0x6e, 0x4a, 0x65,
	0x77, 0x65, 0x6c, 0x52, 0x0a, 0x63, 0x72, 0x6f, 0x77, 0x6e, 0x4a, 0x65, 0x77, 0x65, 0x6c, 0x4a,
	0x04, 0x08, 0x01, 0x10, 0x02, 0x52, 0x0c, 0x63, 0x72, 0x6f, 0x77, 0x6e, 0x5f, 0x6a, 0x65, 0x77,
	0x65, 0x6c, 0x73, 0x22, 0x2a, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x43, 0x72, 0x6f, 0x77, 0x6e, 0x4a,
	0x65, 0x77, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22,
	0x54, 0x0a, 0x16, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x72, 0x6f, 0x77, 0x6e, 0x4a, 0x65, 0x77, 0x65,
	0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67,
	0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61,
	0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x88, 0x01, 0x0a, 0x17, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x72,
	0x6f, 0x77, 0x6e, 0x4a, 0x65, 0x77, 0x65, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x45, 0x0a, 0x0c, 0x63, 0x72, 0x6f, 0x77, 0x6e, 0x5f, 0x6a, 0x65, 0x77, 0x65, 0x6c,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x2e, 0x63, 0x72, 0x6f, 0x77, 0x6e, 0x6a, 0x65, 0x77, 0x65, 0x6c, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x72, 0x6f, 0x77, 0x6e, 0x4a, 0x65, 0x77, 0x65, 0x6c, 0x52, 0x0b, 0x63, 0x72, 0x6f,
	0x77, 0x6e, 0x4a, 0x65, 0x77, 0x65, 0x6c, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74,
	0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x22, 0x72, 0x0a, 0x17, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x72, 0x6f, 0x77, 0x6e, 0x4a,
	0x65, 0x77, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x43, 0x0a, 0x0b, 0x63,
	0x72, 0x6f, 0x77, 0x6e, 0x5f, 0x6a, 0x65, 0x77, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x22, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x63, 0x72, 0x6f, 0x77,
	0x6e, 0x6a, 0x65, 0x77, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x6f, 0x77, 0x6e, 0x4a,
	0x65, 0x77, 0x65, 0x6c, 0x52, 0x0a, 0x63, 0x72, 0x6f, 0x77, 0x6e, 0x4a, 0x65, 0x77, 0x65, 0x6c,
	0x4a, 0x04, 0x08, 0x01, 0x10, 0x02, 0x52, 0x0c, 0x63, 0x72, 0x6f, 0x77, 0x6e, 0x5f, 0x6a, 0x65,
	0x77, 0x65, 0x6c, 0x73, 0x22, 0x72, 0x0a, 0x17, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x43, 0x72,
	0x6f, 0x77, 0x6e, 0x4a, 0x65, 0x77, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x43, 0x0a, 0x0b, 0x63, 0x72, 0x6f, 0x77, 0x6e, 0x5f, 0x6a, 0x65, 0x77, 0x65, 0x6c, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e,
	0x63, 0x72, 0x6f, 0x77, 0x6e, 0x6a, 0x65, 0x77, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72,
	0x6f, 0x77, 0x6e, 0x4a, 0x65, 0x77, 0x65, 0x6c, 0x52, 0x0a, 0x63, 0x72, 0x6f, 0x77, 0x6e, 0x4a,
	0x65, 0x77, 0x65, 0x6c, 0x4a, 0x04, 0x08, 0x01, 0x10, 0x02, 0x52, 0x0c, 0x63, 0x72, 0x6f, 0x77,
	0x6e, 0x5f, 0x6a, 0x65, 0x77, 0x65, 0x6c, 0x73, 0x22, 0x2d, 0x0a, 0x17, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x43, 0x72, 0x6f, 0x77, 0x6e, 0x4a, 0x65, 0x77, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x32, 0x82, 0x05, 0x0a, 0x11, 0x43, 0x72, 0x6f, 0x77,
	0x6e, 0x4a, 0x65, 0x77, 0x65, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x67, 0x0a,
	0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x72, 0x6f, 0x77, 0x6e, 0x4a, 0x65, 0x77, 0x65,
	0x6c, 0x12, 0x2f, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x63, 0x72, 0x6f,
	0x77, 0x6e, 0x6a, 0x65, 0x77, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x43, 0x72, 0x6f, 0x77, 0x6e, 0x4a, 0x65, 0x77, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x22, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x63, 0x72,
	0x6f, 0x77, 0x6e, 0x6a, 0x65, 0x77, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x6f, 0x77,
	0x6e, 0x4a, 0x65, 0x77, 0x65, 0x6c, 0x12, 0x61, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x43, 0x72, 0x6f,
	0x77, 0x6e, 0x4a, 0x65, 0x77, 0x65, 0x6c, 0x12, 0x2c, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x2e, 0x63, 0x72, 0x6f, 0x77, 0x6e, 0x6a, 0x65, 0x77, 0x65, 0x6c, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x43, 0x72, 0x6f, 0x77, 0x6e, 0x4a, 0x65, 0x77, 0x65, 0x6c, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x2e, 0x63, 0x72, 0x6f, 0x77, 0x6e, 0x6a, 0x65, 0x77, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x72, 0x6f, 0x77, 0x6e, 0x4a, 0x65, 0x77, 0x65, 0x6c, 0x12, 0x72, 0x0a, 0x0f, 0x4c, 0x69, 0x73,
	0x74, 0x43, 0x72, 0x6f, 0x77, 0x6e, 0x4a, 0x65, 0x77, 0x65, 0x6c, 0x73, 0x12, 0x2e, 0x2e, 0x74,
	0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x63, 0x72, 0x6f, 0x77, 0x6e, 0x6a, 0x65, 0x77,
	0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x72, 0x6f, 0x77, 0x6e, 0x4a,
	0x65, 0x77, 0x65, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x74,
	0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x63, 0x72, 0x6f, 0x77, 0x6e, 0x6a, 0x65, 0x77,
	0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x72, 0x6f, 0x77, 0x6e, 0x4a,
	0x65, 0x77, 0x65, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x67, 0x0a,
	0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x72, 0x6f, 0x77, 0x6e, 0x4a, 0x65, 0x77, 0x65,
	0x6c, 0x12, 0x2f, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x63, 0x72, 0x6f,
	0x77, 0x6e, 0x6a, 0x65, 0x77, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x43, 0x72, 0x6f, 0x77, 0x6e, 0x4a, 0x65, 0x77, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x22, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x63, 0x72,
	0x6f, 0x77, 0x6e, 0x6a, 0x65, 0x77, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x6f, 0x77,
	0x6e, 0x4a, 0x65, 0x77, 0x65, 0x6c, 0x12, 0x67, 0x0a, 0x10, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74,
	0x43, 0x72, 0x6f, 0x77, 0x6e, 0x4a, 0x65, 0x77, 0x65, 0x6c, 0x12, 0x2f, 0x2e, 0x74, 0x65, 0x6c,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x63, 0x72, 0x6f, 0x77, 0x6e, 0x6a, 0x65, 0x77, 0x65, 0x6c,
	0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x43, 0x72, 0x6f, 0x77, 0x6e, 0x4a,
	0x65, 0x77, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x74, 0x65,
	0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x63, 0x72, 0x6f, 0x77, 0x6e, 0x6a, 0x65, 0x77, 0x65,
	0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x6f, 0x77, 0x6e, 0x4a, 0x65, 0x77, 0x65, 0x6c, 0x12,
	0x5b, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x72, 0x6f, 0x77, 0x6e, 0x4a, 0x65,
	0x77, 0x65, 0x6c, 0x12, 0x2f, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x63,
	0x72, 0x6f, 0x77, 0x6e, 0x6a, 0x65, 0x77, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x43, 0x72, 0x6f, 0x77, 0x6e, 0x4a, 0x65, 0x77, 0x65, 0x6c, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x58, 0x5a, 0x56,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x72, 0x61, 0x76, 0x69,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x67, 0x6f, 0x2f, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x63, 0x72, 0x6f, 0x77,
	0x6e, 0x6a, 0x65, 0x77, 0x65, 0x6c, 0x2f, 0x76, 0x31, 0x3b, 0x63, 0x72, 0x6f, 0x77, 0x6e, 0x6a,
	0x65, 0x77, 0x65, 0x6c, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_teleport_crownjewel_v1_crownjewel_service_proto_rawDescOnce sync.Once
	file_teleport_crownjewel_v1_crownjewel_service_proto_rawDescData = file_teleport_crownjewel_v1_crownjewel_service_proto_rawDesc
)

func file_teleport_crownjewel_v1_crownjewel_service_proto_rawDescGZIP() []byte {
	file_teleport_crownjewel_v1_crownjewel_service_proto_rawDescOnce.Do(func() {
		file_teleport_crownjewel_v1_crownjewel_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_teleport_crownjewel_v1_crownjewel_service_proto_rawDescData)
	})
	return file_teleport_crownjewel_v1_crownjewel_service_proto_rawDescData
}

var file_teleport_crownjewel_v1_crownjewel_service_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_teleport_crownjewel_v1_crownjewel_service_proto_goTypes = []any{
	(*CreateCrownJewelRequest)(nil), // 0: teleport.crownjewel.v1.CreateCrownJewelRequest
	(*GetCrownJewelRequest)(nil),    // 1: teleport.crownjewel.v1.GetCrownJewelRequest
	(*ListCrownJewelsRequest)(nil),  // 2: teleport.crownjewel.v1.ListCrownJewelsRequest
	(*ListCrownJewelsResponse)(nil), // 3: teleport.crownjewel.v1.ListCrownJewelsResponse
	(*UpdateCrownJewelRequest)(nil), // 4: teleport.crownjewel.v1.UpdateCrownJewelRequest
	(*UpsertCrownJewelRequest)(nil), // 5: teleport.crownjewel.v1.UpsertCrownJewelRequest
	(*DeleteCrownJewelRequest)(nil), // 6: teleport.crownjewel.v1.DeleteCrownJewelRequest
	(*CrownJewel)(nil),              // 7: teleport.crownjewel.v1.CrownJewel
	(*emptypb.Empty)(nil),           // 8: google.protobuf.Empty
}
var file_teleport_crownjewel_v1_crownjewel_service_proto_depIdxs = []int32{
	7,  // 0: teleport.crownjewel.v1.CreateCrownJewelRequest.crown_jewel:type_name -> teleport.crownjewel.v1.CrownJewel
	7,  // 1: teleport.crownjewel.v1.ListCrownJewelsResponse.crown_jewels:type_name -> teleport.crownjewel.v1.CrownJewel
	7,  // 2: teleport.crownjewel.v1.UpdateCrownJewelRequest.crown_jewel:type_name -> teleport.crownjewel.v1.CrownJewel
	7,  // 3: teleport.crownjewel.v1.UpsertCrownJewelRequest.crown_jewel:type_name -> teleport.crownjewel.v1.CrownJewel
	0,  // 4: teleport.crownjewel.v1.CrownJewelService.CreateCrownJewel:input_type -> teleport.crownjewel.v1.CreateCrownJewelRequest
	1,  // 5: teleport.crownjewel.v1.CrownJewelService.GetCrownJewel:input_type -> teleport.crownjewel.v1.GetCrownJewelRequest
	2,  // 6: teleport.crownjewel.v1.CrownJewelService.ListCrownJewels:input_type -> teleport.crownjewel.v1.ListCrownJewelsRequest
	4,  // 7: teleport.crownjewel.v1.CrownJewelService.UpdateCrownJewel:input_type -> teleport.crownjewel.v1.UpdateCrownJewelRequest
	5,  // 8: teleport.crownjewel.v1.CrownJewelService.UpsertCrownJewel:input_type -> teleport.crownjewel.v1.UpsertCrownJewelRequest
	6,  // 9: teleport.crownjewel.v1.CrownJewelService.DeleteCrownJewel:input_type -> teleport.crownjewel.v1.DeleteCrownJewelRequest
	7,  // 10: teleport.crownjewel.v1.CrownJewelService.CreateCrownJewel:output_type -> teleport.crownjewel.v1.CrownJewel
	7,  // 11: teleport.crownjewel.v1.CrownJewelService.GetCrownJewel:output_type -> teleport.crownjewel.v1.CrownJewel
	3,  // 12: teleport.crownjewel.v1.CrownJewelService.ListCrownJewels:output_type -> teleport.crownjewel.v1.ListCrownJewelsResponse
	7,  // 13: teleport.crownjewel.v1.CrownJewelService.UpdateCrownJewel:output_type -> teleport.crownjewel.v1.CrownJewel
	7,  // 14: teleport.crownjewel.v1.CrownJewelService.UpsertCrownJewel:output_type -> teleport.crownjewel.v1.CrownJewel
	8,  // 15: teleport.crownjewel.v1.CrownJewelService.DeleteCrownJewel:output_type -> google.protobuf.Empty
	10, // [10:16] is the sub-list for method output_type
	4,  // [4:10] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_teleport_crownjewel_v1_crownjewel_service_proto_init() }
func file_teleport_crownjewel_v1_crownjewel_service_proto_init() {
	if File_teleport_crownjewel_v1_crownjewel_service_proto != nil {
		return
	}
	file_teleport_crownjewel_v1_crownjewel_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_teleport_crownjewel_v1_crownjewel_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_teleport_crownjewel_v1_crownjewel_service_proto_goTypes,
		DependencyIndexes: file_teleport_crownjewel_v1_crownjewel_service_proto_depIdxs,
		MessageInfos:      file_teleport_crownjewel_v1_crownjewel_service_proto_msgTypes,
	}.Build()
	File_teleport_crownjewel_v1_crownjewel_service_proto = out.File
	file_teleport_crownjewel_v1_crownjewel_service_proto_rawDesc = nil
	file_teleport_crownjewel_v1_crownjewel_service_proto_goTypes = nil
	file_teleport_crownjewel_v1_crownjewel_service_proto_depIdxs = nil
}
