//
// Teleport
// Copyright (C) 2023  Gravitational, Inc.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        (unknown)
// source: teleport/lib/teleterm/v1/server.proto

package teletermv1

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

// Server describes connected Server
type Server struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// uri is the server uri
	Uri string `protobuf:"bytes,1,opt,name=uri,proto3" json:"uri,omitempty"`
	// tunnel indicates if this server is connected over a reverse tunnel
	Tunnel bool `protobuf:"varint,2,opt,name=tunnel,proto3" json:"tunnel,omitempty"`
	// name is the server name
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// hostname is this server hostname
	Hostname string `protobuf:"bytes,4,opt,name=hostname,proto3" json:"hostname,omitempty"`
	// addr is this server ip address
	Addr string `protobuf:"bytes,5,opt,name=addr,proto3" json:"addr,omitempty"`
	// labels is this server list of labels
	Labels []*Label `protobuf:"bytes,6,rep,name=labels,proto3" json:"labels,omitempty"`
	// node sub kind: teleport, openssh, openssh-ec2-ice
	SubKind string `protobuf:"bytes,7,opt,name=sub_kind,json=subKind,proto3" json:"sub_kind,omitempty"`
}

func (x *Server) Reset() {
	*x = Server{}
	mi := &file_teleport_lib_teleterm_v1_server_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Server) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Server) ProtoMessage() {}

func (x *Server) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_lib_teleterm_v1_server_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Server.ProtoReflect.Descriptor instead.
func (*Server) Descriptor() ([]byte, []int) {
	return file_teleport_lib_teleterm_v1_server_proto_rawDescGZIP(), []int{0}
}

func (x *Server) GetUri() string {
	if x != nil {
		return x.Uri
	}
	return ""
}

func (x *Server) GetTunnel() bool {
	if x != nil {
		return x.Tunnel
	}
	return false
}

func (x *Server) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Server) GetHostname() string {
	if x != nil {
		return x.Hostname
	}
	return ""
}

func (x *Server) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *Server) GetLabels() []*Label {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *Server) GetSubKind() string {
	if x != nil {
		return x.SubKind
	}
	return ""
}

var File_teleport_lib_teleterm_v1_server_proto protoreflect.FileDescriptor

var file_teleport_lib_teleterm_v1_server_proto_rawDesc = []byte{
	0x0a, 0x25, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x6c, 0x69, 0x62, 0x2f, 0x74,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x72, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x2e, 0x6c, 0x69, 0x62, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x72, 0x6d, 0x2e, 0x76,
	0x31, 0x1a, 0x24, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x6c, 0x69, 0x62, 0x2f,
	0x74, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x72, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x61, 0x62, 0x65,
	0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xca, 0x01, 0x0a, 0x06, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x75, 0x72, 0x69, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x74, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x61, 0x64, 0x64, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x64, 0x64, 0x72,
	0x12, 0x37, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1f, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x6c, 0x69, 0x62, 0x2e,
	0x74, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x61, 0x62, 0x65,
	0x6c, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x75, 0x62,
	0x5f, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x62,
	0x4b, 0x69, 0x6e, 0x64, 0x42, 0x54, 0x5a, 0x52, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x67, 0x72, 0x61, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c,
	0x2f, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x2f, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f,
	0x6c, 0x69, 0x62, 0x2f, 0x74, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x72, 0x6d, 0x2f, 0x76, 0x31, 0x3b,
	0x74, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x72, 0x6d, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_teleport_lib_teleterm_v1_server_proto_rawDescOnce sync.Once
	file_teleport_lib_teleterm_v1_server_proto_rawDescData = file_teleport_lib_teleterm_v1_server_proto_rawDesc
)

func file_teleport_lib_teleterm_v1_server_proto_rawDescGZIP() []byte {
	file_teleport_lib_teleterm_v1_server_proto_rawDescOnce.Do(func() {
		file_teleport_lib_teleterm_v1_server_proto_rawDescData = protoimpl.X.CompressGZIP(file_teleport_lib_teleterm_v1_server_proto_rawDescData)
	})
	return file_teleport_lib_teleterm_v1_server_proto_rawDescData
}

var file_teleport_lib_teleterm_v1_server_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_teleport_lib_teleterm_v1_server_proto_goTypes = []any{
	(*Server)(nil), // 0: teleport.lib.teleterm.v1.Server
	(*Label)(nil),  // 1: teleport.lib.teleterm.v1.Label
}
var file_teleport_lib_teleterm_v1_server_proto_depIdxs = []int32{
	1, // 0: teleport.lib.teleterm.v1.Server.labels:type_name -> teleport.lib.teleterm.v1.Label
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_teleport_lib_teleterm_v1_server_proto_init() }
func file_teleport_lib_teleterm_v1_server_proto_init() {
	if File_teleport_lib_teleterm_v1_server_proto != nil {
		return
	}
	file_teleport_lib_teleterm_v1_label_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_teleport_lib_teleterm_v1_server_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_teleport_lib_teleterm_v1_server_proto_goTypes,
		DependencyIndexes: file_teleport_lib_teleterm_v1_server_proto_depIdxs,
		MessageInfos:      file_teleport_lib_teleterm_v1_server_proto_msgTypes,
	}.Build()
	File_teleport_lib_teleterm_v1_server_proto = out.File
	file_teleport_lib_teleterm_v1_server_proto_rawDesc = nil
	file_teleport_lib_teleterm_v1_server_proto_goTypes = nil
	file_teleport_lib_teleterm_v1_server_proto_depIdxs = nil
}
