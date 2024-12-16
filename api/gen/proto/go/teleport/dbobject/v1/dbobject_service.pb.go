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
// 	protoc-gen-go v1.36.0
// 	protoc        (unknown)
// source: teleport/dbobject/v1/dbobject_service.proto

package dbobjectv1

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

// The request for CreateDatabaseObject.
type CreateDatabaseObjectRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The database object to create.
	Object        *DatabaseObject `protobuf:"bytes,1,opt,name=object,proto3" json:"object,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateDatabaseObjectRequest) Reset() {
	*x = CreateDatabaseObjectRequest{}
	mi := &file_teleport_dbobject_v1_dbobject_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateDatabaseObjectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDatabaseObjectRequest) ProtoMessage() {}

func (x *CreateDatabaseObjectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_dbobject_v1_dbobject_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDatabaseObjectRequest.ProtoReflect.Descriptor instead.
func (*CreateDatabaseObjectRequest) Descriptor() ([]byte, []int) {
	return file_teleport_dbobject_v1_dbobject_service_proto_rawDescGZIP(), []int{0}
}

func (x *CreateDatabaseObjectRequest) GetObject() *DatabaseObject {
	if x != nil {
		return x.Object
	}
	return nil
}

// The request for GetDatabaseObject.
type GetDatabaseObjectRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The name of the database object to fetch.
	Name          string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetDatabaseObjectRequest) Reset() {
	*x = GetDatabaseObjectRequest{}
	mi := &file_teleport_dbobject_v1_dbobject_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetDatabaseObjectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDatabaseObjectRequest) ProtoMessage() {}

func (x *GetDatabaseObjectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_dbobject_v1_dbobject_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDatabaseObjectRequest.ProtoReflect.Descriptor instead.
func (*GetDatabaseObjectRequest) Descriptor() ([]byte, []int) {
	return file_teleport_dbobject_v1_dbobject_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetDatabaseObjectRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// The request for ListDatabaseObjects.
type ListDatabaseObjectsRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The maximum number of items to return.
	// The server may impose a different page size at its discretion.
	PageSize int32 `protobuf:"varint,1,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// The page_token is the next_page_token value returned from a previous List request, if any.
	PageToken     string `protobuf:"bytes,2,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListDatabaseObjectsRequest) Reset() {
	*x = ListDatabaseObjectsRequest{}
	mi := &file_teleport_dbobject_v1_dbobject_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListDatabaseObjectsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListDatabaseObjectsRequest) ProtoMessage() {}

func (x *ListDatabaseObjectsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_dbobject_v1_dbobject_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListDatabaseObjectsRequest.ProtoReflect.Descriptor instead.
func (*ListDatabaseObjectsRequest) Descriptor() ([]byte, []int) {
	return file_teleport_dbobject_v1_dbobject_service_proto_rawDescGZIP(), []int{2}
}

func (x *ListDatabaseObjectsRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListDatabaseObjectsRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

// The response for ListDatabaseObjects.
type ListDatabaseObjectsResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The page of database objects that matched the request.
	Objects []*DatabaseObject `protobuf:"bytes,1,rep,name=objects,proto3" json:"objects,omitempty"`
	// Token to retrieve the next page of results, or empty if there are no
	// more results in the list.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListDatabaseObjectsResponse) Reset() {
	*x = ListDatabaseObjectsResponse{}
	mi := &file_teleport_dbobject_v1_dbobject_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListDatabaseObjectsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListDatabaseObjectsResponse) ProtoMessage() {}

func (x *ListDatabaseObjectsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_dbobject_v1_dbobject_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListDatabaseObjectsResponse.ProtoReflect.Descriptor instead.
func (*ListDatabaseObjectsResponse) Descriptor() ([]byte, []int) {
	return file_teleport_dbobject_v1_dbobject_service_proto_rawDescGZIP(), []int{3}
}

func (x *ListDatabaseObjectsResponse) GetObjects() []*DatabaseObject {
	if x != nil {
		return x.Objects
	}
	return nil
}

func (x *ListDatabaseObjectsResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

// The request for UpdateDatabaseObject.
type UpdateDatabaseObjectRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The database object to replace.
	Object        *DatabaseObject `protobuf:"bytes,1,opt,name=object,proto3" json:"object,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateDatabaseObjectRequest) Reset() {
	*x = UpdateDatabaseObjectRequest{}
	mi := &file_teleport_dbobject_v1_dbobject_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateDatabaseObjectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateDatabaseObjectRequest) ProtoMessage() {}

func (x *UpdateDatabaseObjectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_dbobject_v1_dbobject_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateDatabaseObjectRequest.ProtoReflect.Descriptor instead.
func (*UpdateDatabaseObjectRequest) Descriptor() ([]byte, []int) {
	return file_teleport_dbobject_v1_dbobject_service_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateDatabaseObjectRequest) GetObject() *DatabaseObject {
	if x != nil {
		return x.Object
	}
	return nil
}

// The request for UpsertDatabaseObject.
type UpsertDatabaseObjectRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The database object to create or replace.
	Object        *DatabaseObject `protobuf:"bytes,1,opt,name=object,proto3" json:"object,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpsertDatabaseObjectRequest) Reset() {
	*x = UpsertDatabaseObjectRequest{}
	mi := &file_teleport_dbobject_v1_dbobject_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpsertDatabaseObjectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpsertDatabaseObjectRequest) ProtoMessage() {}

func (x *UpsertDatabaseObjectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_dbobject_v1_dbobject_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpsertDatabaseObjectRequest.ProtoReflect.Descriptor instead.
func (*UpsertDatabaseObjectRequest) Descriptor() ([]byte, []int) {
	return file_teleport_dbobject_v1_dbobject_service_proto_rawDescGZIP(), []int{5}
}

func (x *UpsertDatabaseObjectRequest) GetObject() *DatabaseObject {
	if x != nil {
		return x.Object
	}
	return nil
}

// The request for DeleteDatabaseObject.
type DeleteDatabaseObjectRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The name of the database object to delete.
	Name          string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteDatabaseObjectRequest) Reset() {
	*x = DeleteDatabaseObjectRequest{}
	mi := &file_teleport_dbobject_v1_dbobject_service_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteDatabaseObjectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteDatabaseObjectRequest) ProtoMessage() {}

func (x *DeleteDatabaseObjectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_dbobject_v1_dbobject_service_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteDatabaseObjectRequest.ProtoReflect.Descriptor instead.
func (*DeleteDatabaseObjectRequest) Descriptor() ([]byte, []int) {
	return file_teleport_dbobject_v1_dbobject_service_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteDatabaseObjectRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_teleport_dbobject_v1_dbobject_service_proto protoreflect.FileDescriptor

var file_teleport_dbobject_v1_dbobject_service_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x64, 0x62, 0x6f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x62, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x74,
	0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x64, 0x62, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x23, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x64, 0x62, 0x6f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x62, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5b, 0x0a, 0x1b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44,
	0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x3c, 0x0a, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e,
	0x64, 0x62, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x61, 0x74, 0x61,
	0x62, 0x61, 0x73, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x06, 0x6f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x22, 0x2e, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73,
	0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x22, 0x58, 0x0a, 0x1a, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61,
	0x73, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1d, 0x0a,
	0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x85, 0x01, 0x0a,
	0x1b, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x4f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a, 0x07,
	0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e,
	0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x64, 0x62, 0x6f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x4f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x52, 0x07, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x12, 0x26, 0x0a, 0x0f,
	0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x5b, 0x0a, 0x1b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x61,
	0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x3c, 0x0a, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x64,
	0x62, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x62,
	0x61, 0x73, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x22, 0x5b, 0x0a, 0x1b, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x44, 0x61, 0x74, 0x61, 0x62,
	0x61, 0x73, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x3c, 0x0a, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x24, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x64, 0x62, 0x6f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65,
	0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x31,
	0x0a, 0x1b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65,
	0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x32, 0xb4, 0x05, 0x0a, 0x15, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x4f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x69, 0x0a, 0x11, 0x47,
	0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x12, 0x2e, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x64, 0x62, 0x6f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x62,
	0x61, 0x73, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x24, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x64, 0x62, 0x6f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65,
	0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x7a, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x61,
	0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x12, 0x30, 0x2e,
	0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x64, 0x62, 0x6f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73,
	0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x31, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x64, 0x62, 0x6f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x62,
	0x61, 0x73, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x6f, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61,
	0x62, 0x61, 0x73, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x31, 0x2e, 0x74, 0x65, 0x6c,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x64, 0x62, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65,
	0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e,
	0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x64, 0x62, 0x6f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x4f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x12, 0x6f, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x61, 0x74,
	0x61, 0x62, 0x61, 0x73, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x31, 0x2e, 0x74, 0x65,
	0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x64, 0x62, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x2e,
	0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73,
	0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24,
	0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x64, 0x62, 0x6f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x4f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x12, 0x6f, 0x0a, 0x14, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x44, 0x61,
	0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x31, 0x2e, 0x74,
	0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x64, 0x62, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61,
	0x73, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x24, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x64, 0x62, 0x6f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x4f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x61, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44,
	0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x31, 0x2e,
	0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x64, 0x62, 0x6f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x62,
	0x61, 0x73, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x54, 0x5a, 0x52, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x72, 0x61, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x2f, 0x74,
	0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x64, 0x62, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x2f, 0x76, 0x31, 0x3b, 0x64, 0x62, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x76, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_teleport_dbobject_v1_dbobject_service_proto_rawDescOnce sync.Once
	file_teleport_dbobject_v1_dbobject_service_proto_rawDescData = file_teleport_dbobject_v1_dbobject_service_proto_rawDesc
)

func file_teleport_dbobject_v1_dbobject_service_proto_rawDescGZIP() []byte {
	file_teleport_dbobject_v1_dbobject_service_proto_rawDescOnce.Do(func() {
		file_teleport_dbobject_v1_dbobject_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_teleport_dbobject_v1_dbobject_service_proto_rawDescData)
	})
	return file_teleport_dbobject_v1_dbobject_service_proto_rawDescData
}

var file_teleport_dbobject_v1_dbobject_service_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_teleport_dbobject_v1_dbobject_service_proto_goTypes = []any{
	(*CreateDatabaseObjectRequest)(nil), // 0: teleport.dbobject.v1.CreateDatabaseObjectRequest
	(*GetDatabaseObjectRequest)(nil),    // 1: teleport.dbobject.v1.GetDatabaseObjectRequest
	(*ListDatabaseObjectsRequest)(nil),  // 2: teleport.dbobject.v1.ListDatabaseObjectsRequest
	(*ListDatabaseObjectsResponse)(nil), // 3: teleport.dbobject.v1.ListDatabaseObjectsResponse
	(*UpdateDatabaseObjectRequest)(nil), // 4: teleport.dbobject.v1.UpdateDatabaseObjectRequest
	(*UpsertDatabaseObjectRequest)(nil), // 5: teleport.dbobject.v1.UpsertDatabaseObjectRequest
	(*DeleteDatabaseObjectRequest)(nil), // 6: teleport.dbobject.v1.DeleteDatabaseObjectRequest
	(*DatabaseObject)(nil),              // 7: teleport.dbobject.v1.DatabaseObject
	(*emptypb.Empty)(nil),               // 8: google.protobuf.Empty
}
var file_teleport_dbobject_v1_dbobject_service_proto_depIdxs = []int32{
	7,  // 0: teleport.dbobject.v1.CreateDatabaseObjectRequest.object:type_name -> teleport.dbobject.v1.DatabaseObject
	7,  // 1: teleport.dbobject.v1.ListDatabaseObjectsResponse.objects:type_name -> teleport.dbobject.v1.DatabaseObject
	7,  // 2: teleport.dbobject.v1.UpdateDatabaseObjectRequest.object:type_name -> teleport.dbobject.v1.DatabaseObject
	7,  // 3: teleport.dbobject.v1.UpsertDatabaseObjectRequest.object:type_name -> teleport.dbobject.v1.DatabaseObject
	1,  // 4: teleport.dbobject.v1.DatabaseObjectService.GetDatabaseObject:input_type -> teleport.dbobject.v1.GetDatabaseObjectRequest
	2,  // 5: teleport.dbobject.v1.DatabaseObjectService.ListDatabaseObjects:input_type -> teleport.dbobject.v1.ListDatabaseObjectsRequest
	0,  // 6: teleport.dbobject.v1.DatabaseObjectService.CreateDatabaseObject:input_type -> teleport.dbobject.v1.CreateDatabaseObjectRequest
	4,  // 7: teleport.dbobject.v1.DatabaseObjectService.UpdateDatabaseObject:input_type -> teleport.dbobject.v1.UpdateDatabaseObjectRequest
	5,  // 8: teleport.dbobject.v1.DatabaseObjectService.UpsertDatabaseObject:input_type -> teleport.dbobject.v1.UpsertDatabaseObjectRequest
	6,  // 9: teleport.dbobject.v1.DatabaseObjectService.DeleteDatabaseObject:input_type -> teleport.dbobject.v1.DeleteDatabaseObjectRequest
	7,  // 10: teleport.dbobject.v1.DatabaseObjectService.GetDatabaseObject:output_type -> teleport.dbobject.v1.DatabaseObject
	3,  // 11: teleport.dbobject.v1.DatabaseObjectService.ListDatabaseObjects:output_type -> teleport.dbobject.v1.ListDatabaseObjectsResponse
	7,  // 12: teleport.dbobject.v1.DatabaseObjectService.CreateDatabaseObject:output_type -> teleport.dbobject.v1.DatabaseObject
	7,  // 13: teleport.dbobject.v1.DatabaseObjectService.UpdateDatabaseObject:output_type -> teleport.dbobject.v1.DatabaseObject
	7,  // 14: teleport.dbobject.v1.DatabaseObjectService.UpsertDatabaseObject:output_type -> teleport.dbobject.v1.DatabaseObject
	8,  // 15: teleport.dbobject.v1.DatabaseObjectService.DeleteDatabaseObject:output_type -> google.protobuf.Empty
	10, // [10:16] is the sub-list for method output_type
	4,  // [4:10] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_teleport_dbobject_v1_dbobject_service_proto_init() }
func file_teleport_dbobject_v1_dbobject_service_proto_init() {
	if File_teleport_dbobject_v1_dbobject_service_proto != nil {
		return
	}
	file_teleport_dbobject_v1_dbobject_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_teleport_dbobject_v1_dbobject_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_teleport_dbobject_v1_dbobject_service_proto_goTypes,
		DependencyIndexes: file_teleport_dbobject_v1_dbobject_service_proto_depIdxs,
		MessageInfos:      file_teleport_dbobject_v1_dbobject_service_proto_msgTypes,
	}.Build()
	File_teleport_dbobject_v1_dbobject_service_proto = out.File
	file_teleport_dbobject_v1_dbobject_service_proto_rawDesc = nil
	file_teleport_dbobject_v1_dbobject_service_proto_goTypes = nil
	file_teleport_dbobject_v1_dbobject_service_proto_depIdxs = nil
}
