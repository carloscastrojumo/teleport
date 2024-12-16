// Copyright 2024 Gravitational, Inc.
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
// source: teleport/userprovisioning/v2/statichostuser.proto

package userprovisioningv2

import (
	v1 "github.com/gravitational/teleport/api/gen/proto/go/teleport/header/v1"
	v11 "github.com/gravitational/teleport/api/gen/proto/go/teleport/label/v1"
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

// StaticHostUser is a resource that represents host users that should be
// created on matching nodes.
type StaticHostUser struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// kind is a resource kind.
	Kind string `protobuf:"bytes,1,opt,name=kind,proto3" json:"kind,omitempty"`
	// sub_kind is an optional resource sub kind, used in some resources.
	SubKind string `protobuf:"bytes,2,opt,name=sub_kind,json=subKind,proto3" json:"sub_kind,omitempty"`
	// version is the resource version. It must be specified.
	// Supported values are: `v2`.
	Version string `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	// metadata is resource metadata.
	Metadata *v1.Metadata `protobuf:"bytes,4,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// spec is the static host user spec.
	Spec          *StaticHostUserSpec `protobuf:"bytes,5,opt,name=spec,proto3" json:"spec,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *StaticHostUser) Reset() {
	*x = StaticHostUser{}
	mi := &file_teleport_userprovisioning_v2_statichostuser_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StaticHostUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StaticHostUser) ProtoMessage() {}

func (x *StaticHostUser) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_userprovisioning_v2_statichostuser_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StaticHostUser.ProtoReflect.Descriptor instead.
func (*StaticHostUser) Descriptor() ([]byte, []int) {
	return file_teleport_userprovisioning_v2_statichostuser_proto_rawDescGZIP(), []int{0}
}

func (x *StaticHostUser) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *StaticHostUser) GetSubKind() string {
	if x != nil {
		return x.SubKind
	}
	return ""
}

func (x *StaticHostUser) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *StaticHostUser) GetMetadata() *v1.Metadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *StaticHostUser) GetSpec() *StaticHostUserSpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

// Matcher is a matcher for nodes to add the user to.
type Matcher struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// node_labels is a map of node labels that will create a user from this
	// resource.
	NodeLabels []*v11.Label `protobuf:"bytes,1,rep,name=node_labels,json=nodeLabels,proto3" json:"node_labels,omitempty"`
	// node_labels_expression is a predicate expression to create a user from
	// this resource.
	NodeLabelsExpression string `protobuf:"bytes,2,opt,name=node_labels_expression,json=nodeLabelsExpression,proto3" json:"node_labels_expression,omitempty"`
	// groups is a list of additional groups to add the user to.
	Groups []string `protobuf:"bytes,3,rep,name=groups,proto3" json:"groups,omitempty"`
	// sudoers is a list of sudoer entries to add.
	Sudoers []string `protobuf:"bytes,4,rep,name=sudoers,proto3" json:"sudoers,omitempty"`
	// uid is the new user's uid.
	Uid int64 `protobuf:"varint,5,opt,name=uid,proto3" json:"uid,omitempty"`
	// gid is the new user's gid.
	Gid int64 `protobuf:"varint,6,opt,name=gid,proto3" json:"gid,omitempty"`
	// default_shell is the new user's default shell
	DefaultShell string `protobuf:"bytes,7,opt,name=default_shell,json=defaultShell,proto3" json:"default_shell,omitempty"`
	// take_ownership_if_user_exists will take ownership of existing, unmanaged users
	TakeOwnershipIfUserExists bool `protobuf:"varint,8,opt,name=take_ownership_if_user_exists,json=takeOwnershipIfUserExists,proto3" json:"take_ownership_if_user_exists,omitempty"`
	unknownFields             protoimpl.UnknownFields
	sizeCache                 protoimpl.SizeCache
}

func (x *Matcher) Reset() {
	*x = Matcher{}
	mi := &file_teleport_userprovisioning_v2_statichostuser_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Matcher) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Matcher) ProtoMessage() {}

func (x *Matcher) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_userprovisioning_v2_statichostuser_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Matcher.ProtoReflect.Descriptor instead.
func (*Matcher) Descriptor() ([]byte, []int) {
	return file_teleport_userprovisioning_v2_statichostuser_proto_rawDescGZIP(), []int{1}
}

func (x *Matcher) GetNodeLabels() []*v11.Label {
	if x != nil {
		return x.NodeLabels
	}
	return nil
}

func (x *Matcher) GetNodeLabelsExpression() string {
	if x != nil {
		return x.NodeLabelsExpression
	}
	return ""
}

func (x *Matcher) GetGroups() []string {
	if x != nil {
		return x.Groups
	}
	return nil
}

func (x *Matcher) GetSudoers() []string {
	if x != nil {
		return x.Sudoers
	}
	return nil
}

func (x *Matcher) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *Matcher) GetGid() int64 {
	if x != nil {
		return x.Gid
	}
	return 0
}

func (x *Matcher) GetDefaultShell() string {
	if x != nil {
		return x.DefaultShell
	}
	return ""
}

func (x *Matcher) GetTakeOwnershipIfUserExists() bool {
	if x != nil {
		return x.TakeOwnershipIfUserExists
	}
	return false
}

// StaticHostUserSpec is the static host user spec.
type StaticHostUserSpec struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Matchers      []*Matcher             `protobuf:"bytes,1,rep,name=matchers,proto3" json:"matchers,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *StaticHostUserSpec) Reset() {
	*x = StaticHostUserSpec{}
	mi := &file_teleport_userprovisioning_v2_statichostuser_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StaticHostUserSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StaticHostUserSpec) ProtoMessage() {}

func (x *StaticHostUserSpec) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_userprovisioning_v2_statichostuser_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StaticHostUserSpec.ProtoReflect.Descriptor instead.
func (*StaticHostUserSpec) Descriptor() ([]byte, []int) {
	return file_teleport_userprovisioning_v2_statichostuser_proto_rawDescGZIP(), []int{2}
}

func (x *StaticHostUserSpec) GetMatchers() []*Matcher {
	if x != nil {
		return x.Matchers
	}
	return nil
}

var File_teleport_userprovisioning_v2_statichostuser_proto protoreflect.FileDescriptor

var file_teleport_userprovisioning_v2_statichostuser_proto_rawDesc = []byte{
	0x0a, 0x31, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x70,
	0x72, 0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x32, 0x2f, 0x73,
	0x74, 0x61, 0x74, 0x69, 0x63, 0x68, 0x6f, 0x73, 0x74, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x1c, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x76,
	0x32, 0x1a, 0x21, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x68, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x6c,
	0x61, 0x62, 0x65, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xd9, 0x01, 0x0a, 0x0e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x63, 0x48, 0x6f,
	0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x75,
	0x62, 0x5f, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75,
	0x62, 0x4b, 0x69, 0x6e, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x38, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x68, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52,
	0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x44, 0x0a, 0x04, 0x73, 0x70, 0x65,
	0x63, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e,
	0x69, 0x6e, 0x67, 0x2e, 0x76, 0x32, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x63, 0x48, 0x6f, 0x73,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x53, 0x70, 0x65, 0x63, 0x52, 0x04, 0x73, 0x70, 0x65, 0x63, 0x22,
	0xb7, 0x02, 0x0a, 0x07, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x12, 0x39, 0x0a, 0x0b, 0x6e,
	0x6f, 0x64, 0x65, 0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x18, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x6c, 0x61, 0x62, 0x65,
	0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x52, 0x0a, 0x6e, 0x6f, 0x64, 0x65,
	0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x34, 0x0a, 0x16, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x6c,
	0x61, 0x62, 0x65, 0x6c, 0x73, 0x5f, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x6e, 0x6f, 0x64, 0x65, 0x4c, 0x61, 0x62, 0x65,
	0x6c, 0x73, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x64, 0x6f, 0x65, 0x72, 0x73, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x64, 0x6f, 0x65, 0x72, 0x73, 0x12, 0x10,
	0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x75, 0x69, 0x64,
	0x12, 0x10, 0x0a, 0x03, 0x67, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x67,
	0x69, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x73, 0x68,
	0x65, 0x6c, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x64, 0x65, 0x66, 0x61, 0x75,
	0x6c, 0x74, 0x53, 0x68, 0x65, 0x6c, 0x6c, 0x12, 0x40, 0x0a, 0x1d, 0x74, 0x61, 0x6b, 0x65, 0x5f,
	0x6f, 0x77, 0x6e, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x5f, 0x69, 0x66, 0x5f, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x65, 0x78, 0x69, 0x73, 0x74, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x19,
	0x74, 0x61, 0x6b, 0x65, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x49, 0x66, 0x55,
	0x73, 0x65, 0x72, 0x45, 0x78, 0x69, 0x73, 0x74, 0x73, 0x22, 0x57, 0x0a, 0x12, 0x53, 0x74, 0x61,
	0x74, 0x69, 0x63, 0x48, 0x6f, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x53, 0x70, 0x65, 0x63, 0x12,
	0x41, 0x0a, 0x08, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x25, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x32,
	0x2e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x52, 0x08, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65,
	0x72, 0x73, 0x42, 0x64, 0x5a, 0x62, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x67, 0x72, 0x61, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x2f, 0x74,
	0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x65, 0x6e, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x2f, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x69,
	0x6e, 0x67, 0x2f, 0x76, 0x32, 0x3b, 0x75, 0x73, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x73,
	0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x76, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_teleport_userprovisioning_v2_statichostuser_proto_rawDescOnce sync.Once
	file_teleport_userprovisioning_v2_statichostuser_proto_rawDescData = file_teleport_userprovisioning_v2_statichostuser_proto_rawDesc
)

func file_teleport_userprovisioning_v2_statichostuser_proto_rawDescGZIP() []byte {
	file_teleport_userprovisioning_v2_statichostuser_proto_rawDescOnce.Do(func() {
		file_teleport_userprovisioning_v2_statichostuser_proto_rawDescData = protoimpl.X.CompressGZIP(file_teleport_userprovisioning_v2_statichostuser_proto_rawDescData)
	})
	return file_teleport_userprovisioning_v2_statichostuser_proto_rawDescData
}

var file_teleport_userprovisioning_v2_statichostuser_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_teleport_userprovisioning_v2_statichostuser_proto_goTypes = []any{
	(*StaticHostUser)(nil),     // 0: teleport.userprovisioning.v2.StaticHostUser
	(*Matcher)(nil),            // 1: teleport.userprovisioning.v2.Matcher
	(*StaticHostUserSpec)(nil), // 2: teleport.userprovisioning.v2.StaticHostUserSpec
	(*v1.Metadata)(nil),        // 3: teleport.header.v1.Metadata
	(*v11.Label)(nil),          // 4: teleport.label.v1.Label
}
var file_teleport_userprovisioning_v2_statichostuser_proto_depIdxs = []int32{
	3, // 0: teleport.userprovisioning.v2.StaticHostUser.metadata:type_name -> teleport.header.v1.Metadata
	2, // 1: teleport.userprovisioning.v2.StaticHostUser.spec:type_name -> teleport.userprovisioning.v2.StaticHostUserSpec
	4, // 2: teleport.userprovisioning.v2.Matcher.node_labels:type_name -> teleport.label.v1.Label
	1, // 3: teleport.userprovisioning.v2.StaticHostUserSpec.matchers:type_name -> teleport.userprovisioning.v2.Matcher
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_teleport_userprovisioning_v2_statichostuser_proto_init() }
func file_teleport_userprovisioning_v2_statichostuser_proto_init() {
	if File_teleport_userprovisioning_v2_statichostuser_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_teleport_userprovisioning_v2_statichostuser_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_teleport_userprovisioning_v2_statichostuser_proto_goTypes,
		DependencyIndexes: file_teleport_userprovisioning_v2_statichostuser_proto_depIdxs,
		MessageInfos:      file_teleport_userprovisioning_v2_statichostuser_proto_msgTypes,
	}.Build()
	File_teleport_userprovisioning_v2_statichostuser_proto = out.File
	file_teleport_userprovisioning_v2_statichostuser_proto_rawDesc = nil
	file_teleport_userprovisioning_v2_statichostuser_proto_goTypes = nil
	file_teleport_userprovisioning_v2_statichostuser_proto_depIdxs = nil
}
