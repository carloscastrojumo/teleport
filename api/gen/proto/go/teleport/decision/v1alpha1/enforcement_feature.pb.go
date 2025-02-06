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
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: teleport/decision/v1alpha1/enforcement_feature.proto

package decisionpb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// EnforcementFeature represents PEP (Policy Enforcement Point) features.
type EnforcementFeature int32

const (
	// ENFORCEMENT_FEATURE_UNSPECIFIED is the default/unspecified value for
	// EnforcementFeature. Asserting this feature has no effect.
	EnforcementFeature_ENFORCEMENT_FEATURE_UNSPECIFIED EnforcementFeature = 0
)

// Enum value maps for EnforcementFeature.
var (
	EnforcementFeature_name = map[int32]string{
		0: "ENFORCEMENT_FEATURE_UNSPECIFIED",
	}
	EnforcementFeature_value = map[string]int32{
		"ENFORCEMENT_FEATURE_UNSPECIFIED": 0,
	}
)

func (x EnforcementFeature) Enum() *EnforcementFeature {
	p := new(EnforcementFeature)
	*p = x
	return p
}

func (x EnforcementFeature) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EnforcementFeature) Descriptor() protoreflect.EnumDescriptor {
	return file_teleport_decision_v1alpha1_enforcement_feature_proto_enumTypes[0].Descriptor()
}

func (EnforcementFeature) Type() protoreflect.EnumType {
	return &file_teleport_decision_v1alpha1_enforcement_feature_proto_enumTypes[0]
}

func (x EnforcementFeature) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EnforcementFeature.Descriptor instead.
func (EnforcementFeature) EnumDescriptor() ([]byte, []int) {
	return file_teleport_decision_v1alpha1_enforcement_feature_proto_rawDescGZIP(), []int{0}
}

var File_teleport_decision_v1alpha1_enforcement_feature_proto protoreflect.FileDescriptor

var file_teleport_decision_v1alpha1_enforcement_feature_proto_rawDesc = string([]byte{
	0x0a, 0x34, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x64, 0x65, 0x63, 0x69, 0x73,
	0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x65, 0x6e, 0x66,
	0x6f, 0x72, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x2e, 0x64, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2a, 0x39, 0x0a, 0x12, 0x45, 0x6e, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x23, 0x0a, 0x1f, 0x45, 0x4e, 0x46, 0x4f,
	0x52, 0x43, 0x45, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x46, 0x45, 0x41, 0x54, 0x55, 0x52, 0x45, 0x5f,
	0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x42, 0x5a, 0x5a,
	0x58, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x72, 0x61, 0x76,
	0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x67, 0x6f, 0x2f, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x64, 0x65, 0x63,
	0x69, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x3b, 0x64,
	0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
})

var (
	file_teleport_decision_v1alpha1_enforcement_feature_proto_rawDescOnce sync.Once
	file_teleport_decision_v1alpha1_enforcement_feature_proto_rawDescData []byte
)

func file_teleport_decision_v1alpha1_enforcement_feature_proto_rawDescGZIP() []byte {
	file_teleport_decision_v1alpha1_enforcement_feature_proto_rawDescOnce.Do(func() {
		file_teleport_decision_v1alpha1_enforcement_feature_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_teleport_decision_v1alpha1_enforcement_feature_proto_rawDesc), len(file_teleport_decision_v1alpha1_enforcement_feature_proto_rawDesc)))
	})
	return file_teleport_decision_v1alpha1_enforcement_feature_proto_rawDescData
}

var file_teleport_decision_v1alpha1_enforcement_feature_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_teleport_decision_v1alpha1_enforcement_feature_proto_goTypes = []any{
	(EnforcementFeature)(0), // 0: teleport.decision.v1alpha1.EnforcementFeature
}
var file_teleport_decision_v1alpha1_enforcement_feature_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_teleport_decision_v1alpha1_enforcement_feature_proto_init() }
func file_teleport_decision_v1alpha1_enforcement_feature_proto_init() {
	if File_teleport_decision_v1alpha1_enforcement_feature_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_teleport_decision_v1alpha1_enforcement_feature_proto_rawDesc), len(file_teleport_decision_v1alpha1_enforcement_feature_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_teleport_decision_v1alpha1_enforcement_feature_proto_goTypes,
		DependencyIndexes: file_teleport_decision_v1alpha1_enforcement_feature_proto_depIdxs,
		EnumInfos:         file_teleport_decision_v1alpha1_enforcement_feature_proto_enumTypes,
	}.Build()
	File_teleport_decision_v1alpha1_enforcement_feature_proto = out.File
	file_teleport_decision_v1alpha1_enforcement_feature_proto_goTypes = nil
	file_teleport_decision_v1alpha1_enforcement_feature_proto_depIdxs = nil
}
