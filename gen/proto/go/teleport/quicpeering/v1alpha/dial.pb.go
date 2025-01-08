// Teleport
// Copyright (C) 2024 Gravitational, Inc.
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
// 	protoc-gen-go v1.36.2
// 	protoc        (unknown)
// source: teleport/quicpeering/v1alpha/dial.proto

package quicpeeringv1alpha

import (
	status "google.golang.org/genproto/googleapis/rpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Sent from a proxy to a peer proxy in a fresh QUIC stream to dial a Teleport
// resource through a QUIC proxy peering connection. The message is sent in
// protobuf binary format, prefixed by its length encoded as a little endian
// 32-bit unsigned integer.
type DialRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The target agent for the connection attempt; should be "<host id>.<cluster name>".
	TargetHostId string `protobuf:"bytes,1,opt,name=target_host_id,json=targetHostId,proto3" json:"target_host_id,omitempty"`
	// The type of the connection as defined by api/types.TunnelType ("node",
	// "app", "kube"...).
	ConnectionType string `protobuf:"bytes,2,opt,name=connection_type,json=connectionType,proto3" json:"connection_type,omitempty"`
	// The source of the connection, the network address of the user for whom the
	// connection is being tunneled, as seen from the proxy sending the request.
	Source *Addr `protobuf:"bytes,3,opt,name=source,proto3" json:"source,omitempty"`
	// The destination of the connection, used as a weak hint and as something to
	// put in the "local address" of the connection object handled by the agent.
	Destination *Addr `protobuf:"bytes,4,opt,name=destination,proto3" json:"destination,omitempty"`
	// The time of the client, must be provided and within 5 minutes of the local
	// server time for 0-RTT requests.
	Timestamp *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// A random id for each dial request, must be provided and unique among dial
	// requests recently received by the server.
	Nonce uint64 `protobuf:"fixed64,6,opt,name=nonce,proto3" json:"nonce,omitempty"`
	// The ping flag should be set if the request is actually for a reachability
	// and authentication check. If set, all other fields are functionally
	// ignored, although nonce and timestamp should still be set so they can be
	// logged.
	Ping          bool `protobuf:"varint,7,opt,name=ping,proto3" json:"ping,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DialRequest) Reset() {
	*x = DialRequest{}
	mi := &file_teleport_quicpeering_v1alpha_dial_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DialRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DialRequest) ProtoMessage() {}

func (x *DialRequest) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_quicpeering_v1alpha_dial_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DialRequest.ProtoReflect.Descriptor instead.
func (*DialRequest) Descriptor() ([]byte, []int) {
	return file_teleport_quicpeering_v1alpha_dial_proto_rawDescGZIP(), []int{0}
}

func (x *DialRequest) GetTargetHostId() string {
	if x != nil {
		return x.TargetHostId
	}
	return ""
}

func (x *DialRequest) GetConnectionType() string {
	if x != nil {
		return x.ConnectionType
	}
	return ""
}

func (x *DialRequest) GetSource() *Addr {
	if x != nil {
		return x.Source
	}
	return nil
}

func (x *DialRequest) GetDestination() *Addr {
	if x != nil {
		return x.Destination
	}
	return nil
}

func (x *DialRequest) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *DialRequest) GetNonce() uint64 {
	if x != nil {
		return x.Nonce
	}
	return 0
}

func (x *DialRequest) GetPing() bool {
	if x != nil {
		return x.Ping
	}
	return false
}

// A stringy Go net.Addr. Can be converted to and from a lib/utils.NetAddr.
type Addr struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Likely always "tcp".
	Network string `protobuf:"bytes,1,opt,name=network,proto3" json:"network,omitempty"`
	// Depending on the network, likely "<ip address>:<port>".
	Addr          string `protobuf:"bytes,2,opt,name=addr,proto3" json:"addr,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Addr) Reset() {
	*x = Addr{}
	mi := &file_teleport_quicpeering_v1alpha_dial_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Addr) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Addr) ProtoMessage() {}

func (x *Addr) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_quicpeering_v1alpha_dial_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Addr.ProtoReflect.Descriptor instead.
func (*Addr) Descriptor() ([]byte, []int) {
	return file_teleport_quicpeering_v1alpha_dial_proto_rawDescGZIP(), []int{1}
}

func (x *Addr) GetNetwork() string {
	if x != nil {
		return x.Network
	}
	return ""
}

func (x *Addr) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

// Sent from the server to the client as a response to a DialRequest. The
// message is likewise sent in protobuf binary format, prefixed by its length
// encoded as a little endian uint32.
type DialResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The success or failure of the dial. If the dial is successful, the stream
	// will continue with the data of the connection.
	Status        *status.Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DialResponse) Reset() {
	*x = DialResponse{}
	mi := &file_teleport_quicpeering_v1alpha_dial_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DialResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DialResponse) ProtoMessage() {}

func (x *DialResponse) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_quicpeering_v1alpha_dial_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DialResponse.ProtoReflect.Descriptor instead.
func (*DialResponse) Descriptor() ([]byte, []int) {
	return file_teleport_quicpeering_v1alpha_dial_proto_rawDescGZIP(), []int{2}
}

func (x *DialResponse) GetStatus() *status.Status {
	if x != nil {
		return x.Status
	}
	return nil
}

var File_teleport_quicpeering_v1alpha_dial_proto protoreflect.FileDescriptor

var file_teleport_quicpeering_v1alpha_dial_proto_rawDesc = []byte{
	0x0a, 0x27, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x71, 0x75, 0x69, 0x63, 0x70,
	0x65, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x64,
	0x69, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1c, 0x74, 0x65, 0x6c, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x2e, 0x71, 0x75, 0x69, 0x63, 0x70, 0x65, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x72, 0x70, 0x63, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xc2, 0x02, 0x0a, 0x0b, 0x44, 0x69, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x24, 0x0a, 0x0e, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x68, 0x6f, 0x73, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x74, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x48, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x63, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x3a, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x22, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x71, 0x75, 0x69, 0x63,
	0x70, 0x65, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e,
	0x41, 0x64, 0x64, 0x72, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x44, 0x0a, 0x0b,
	0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x22, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x71, 0x75, 0x69,
	0x63, 0x70, 0x65, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x2e, 0x41, 0x64, 0x64, 0x72, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x14, 0x0a, 0x05,
	0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x06, 0x52, 0x05, 0x6e, 0x6f, 0x6e,
	0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x69, 0x6e, 0x67, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x04, 0x70, 0x69, 0x6e, 0x67, 0x22, 0x34, 0x0a, 0x04, 0x41, 0x64, 0x64, 0x72, 0x12, 0x18,
	0x0a, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x64, 0x64, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x64, 0x64, 0x72, 0x22, 0x3a, 0x0a, 0x0c,
	0x44, 0x69, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x60, 0x5a, 0x5e, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x72, 0x61, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x67, 0x65,
	0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x2f, 0x74, 0x65, 0x6c, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x2f, 0x71, 0x75, 0x69, 0x63, 0x70, 0x65, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x2f,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x3b, 0x71, 0x75, 0x69, 0x63, 0x70, 0x65, 0x65, 0x72,
	0x69, 0x6e, 0x67, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_teleport_quicpeering_v1alpha_dial_proto_rawDescOnce sync.Once
	file_teleport_quicpeering_v1alpha_dial_proto_rawDescData = file_teleport_quicpeering_v1alpha_dial_proto_rawDesc
)

func file_teleport_quicpeering_v1alpha_dial_proto_rawDescGZIP() []byte {
	file_teleport_quicpeering_v1alpha_dial_proto_rawDescOnce.Do(func() {
		file_teleport_quicpeering_v1alpha_dial_proto_rawDescData = protoimpl.X.CompressGZIP(file_teleport_quicpeering_v1alpha_dial_proto_rawDescData)
	})
	return file_teleport_quicpeering_v1alpha_dial_proto_rawDescData
}

var file_teleport_quicpeering_v1alpha_dial_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_teleport_quicpeering_v1alpha_dial_proto_goTypes = []any{
	(*DialRequest)(nil),           // 0: teleport.quicpeering.v1alpha.DialRequest
	(*Addr)(nil),                  // 1: teleport.quicpeering.v1alpha.Addr
	(*DialResponse)(nil),          // 2: teleport.quicpeering.v1alpha.DialResponse
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
	(*status.Status)(nil),         // 4: google.rpc.Status
}
var file_teleport_quicpeering_v1alpha_dial_proto_depIdxs = []int32{
	1, // 0: teleport.quicpeering.v1alpha.DialRequest.source:type_name -> teleport.quicpeering.v1alpha.Addr
	1, // 1: teleport.quicpeering.v1alpha.DialRequest.destination:type_name -> teleport.quicpeering.v1alpha.Addr
	3, // 2: teleport.quicpeering.v1alpha.DialRequest.timestamp:type_name -> google.protobuf.Timestamp
	4, // 3: teleport.quicpeering.v1alpha.DialResponse.status:type_name -> google.rpc.Status
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_teleport_quicpeering_v1alpha_dial_proto_init() }
func file_teleport_quicpeering_v1alpha_dial_proto_init() {
	if File_teleport_quicpeering_v1alpha_dial_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_teleport_quicpeering_v1alpha_dial_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_teleport_quicpeering_v1alpha_dial_proto_goTypes,
		DependencyIndexes: file_teleport_quicpeering_v1alpha_dial_proto_depIdxs,
		MessageInfos:      file_teleport_quicpeering_v1alpha_dial_proto_msgTypes,
	}.Build()
	File_teleport_quicpeering_v1alpha_dial_proto = out.File
	file_teleport_quicpeering_v1alpha_dial_proto_rawDesc = nil
	file_teleport_quicpeering_v1alpha_dial_proto_goTypes = nil
	file_teleport_quicpeering_v1alpha_dial_proto_depIdxs = nil
}
