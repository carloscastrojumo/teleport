// Copyright 2022 Gravitational, Inc
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
// 	protoc-gen-go v1.36.1
// 	protoc        (unknown)
// source: teleport/devicetrust/v1/device.proto

package devicetrustv1

import (
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

// AttestationType indicates the degree to which the device credential has
// been attested.
type DeviceAttestationType int32

const (
	// Bare public key which has only verified with proof of ownership.
	// Used on macOS.
	DeviceAttestationType_DEVICE_ATTESTATION_TYPE_UNSPECIFIED DeviceAttestationType = 0
	// Credential was verified through a TPM EK->AK chain on enrollment.
	DeviceAttestationType_DEVICE_ATTESTATION_TYPE_TPM_EKPUB DeviceAttestationType = 1
	// Credential was verified through a TPM EKCert->AK chain on enrollment,
	// but no allow-listed CAs were configured to validate this EKCert against.
	DeviceAttestationType_DEVICE_ATTESTATION_TYPE_TPM_EKCERT DeviceAttestationType = 2
	// Credential was verified through a TPM EKCert->AK chain on enrollment, and
	// the EKCert was signed by a configured allow-listed CA.
	DeviceAttestationType_DEVICE_ATTESTATION_TYPE_TPM_EKCERT_TRUSTED DeviceAttestationType = 3
)

// Enum value maps for DeviceAttestationType.
var (
	DeviceAttestationType_name = map[int32]string{
		0: "DEVICE_ATTESTATION_TYPE_UNSPECIFIED",
		1: "DEVICE_ATTESTATION_TYPE_TPM_EKPUB",
		2: "DEVICE_ATTESTATION_TYPE_TPM_EKCERT",
		3: "DEVICE_ATTESTATION_TYPE_TPM_EKCERT_TRUSTED",
	}
	DeviceAttestationType_value = map[string]int32{
		"DEVICE_ATTESTATION_TYPE_UNSPECIFIED":        0,
		"DEVICE_ATTESTATION_TYPE_TPM_EKPUB":          1,
		"DEVICE_ATTESTATION_TYPE_TPM_EKCERT":         2,
		"DEVICE_ATTESTATION_TYPE_TPM_EKCERT_TRUSTED": 3,
	}
)

func (x DeviceAttestationType) Enum() *DeviceAttestationType {
	p := new(DeviceAttestationType)
	*p = x
	return p
}

func (x DeviceAttestationType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DeviceAttestationType) Descriptor() protoreflect.EnumDescriptor {
	return file_teleport_devicetrust_v1_device_proto_enumTypes[0].Descriptor()
}

func (DeviceAttestationType) Type() protoreflect.EnumType {
	return &file_teleport_devicetrust_v1_device_proto_enumTypes[0]
}

func (x DeviceAttestationType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DeviceAttestationType.Descriptor instead.
func (DeviceAttestationType) EnumDescriptor() ([]byte, []int) {
	return file_teleport_devicetrust_v1_device_proto_rawDescGZIP(), []int{0}
}

// DeviceEnrollStatus represents the enrollment status of a device.
type DeviceEnrollStatus int32

const (
	DeviceEnrollStatus_DEVICE_ENROLL_STATUS_UNSPECIFIED DeviceEnrollStatus = 0
	// Device is registered but not enrolled.
	DeviceEnrollStatus_DEVICE_ENROLL_STATUS_NOT_ENROLLED DeviceEnrollStatus = 1
	// Device is registered and enrolled.
	DeviceEnrollStatus_DEVICE_ENROLL_STATUS_ENROLLED DeviceEnrollStatus = 2
)

// Enum value maps for DeviceEnrollStatus.
var (
	DeviceEnrollStatus_name = map[int32]string{
		0: "DEVICE_ENROLL_STATUS_UNSPECIFIED",
		1: "DEVICE_ENROLL_STATUS_NOT_ENROLLED",
		2: "DEVICE_ENROLL_STATUS_ENROLLED",
	}
	DeviceEnrollStatus_value = map[string]int32{
		"DEVICE_ENROLL_STATUS_UNSPECIFIED":  0,
		"DEVICE_ENROLL_STATUS_NOT_ENROLLED": 1,
		"DEVICE_ENROLL_STATUS_ENROLLED":     2,
	}
)

func (x DeviceEnrollStatus) Enum() *DeviceEnrollStatus {
	p := new(DeviceEnrollStatus)
	*p = x
	return p
}

func (x DeviceEnrollStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DeviceEnrollStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_teleport_devicetrust_v1_device_proto_enumTypes[1].Descriptor()
}

func (DeviceEnrollStatus) Type() protoreflect.EnumType {
	return &file_teleport_devicetrust_v1_device_proto_enumTypes[1]
}

func (x DeviceEnrollStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DeviceEnrollStatus.Descriptor instead.
func (DeviceEnrollStatus) EnumDescriptor() ([]byte, []int) {
	return file_teleport_devicetrust_v1_device_proto_rawDescGZIP(), []int{1}
}

// Device represents a registered device.
// Registered devices may be enrolled. Enrolled devices are allowed to perform
// device-aware actions.
type Device struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// API version of the Device definition, present for compatibility with
	// types.DeviceV1.
	// Always "v1".
	ApiVersion string `protobuf:"bytes,1,opt,name=api_version,json=apiVersion,proto3" json:"api_version,omitempty"`
	// Unique device identifier.
	// System managed.
	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// Device operating system.
	// Required.
	OsType OSType `protobuf:"varint,3,opt,name=os_type,json=osType,proto3,enum=teleport.devicetrust.v1.OSType" json:"os_type,omitempty"`
	// Device inventory identifier.
	// Takes different meanings depending on the device and operating system.
	// For macOS devices it is the device serial number.
	// Required.
	AssetTag string `protobuf:"bytes,4,opt,name=asset_tag,json=assetTag,proto3" json:"asset_tag,omitempty"`
	// Create time.
	// System managed.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Last update time.
	// System managed.
	UpdateTime *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	// Enrollment token for the device.
	// Only present in situations where device creation and enrollment are rolled
	// into a single operation.
	// Transient.
	EnrollToken *DeviceEnrollToken `protobuf:"bytes,7,opt,name=enroll_token,json=enrollToken,proto3" json:"enroll_token,omitempty"`
	// Enrollment status of the device.
	// May be changed to DEVICE_ENROLL_STATUS_NOT_ENROLLED at any time to
	// forcefully unenroll a device (server-side only).
	// System managed.
	EnrollStatus DeviceEnrollStatus `protobuf:"varint,8,opt,name=enroll_status,json=enrollStatus,proto3,enum=teleport.devicetrust.v1.DeviceEnrollStatus" json:"enroll_status,omitempty"`
	// Currently enrolled device credential.
	// Manually unenrolling a device clears the credential.
	// System managed.
	Credential *DeviceCredential `protobuf:"bytes,9,opt,name=credential,proto3" json:"credential,omitempty"`
	// Device data collected during enrollment and device authentication.
	// Enrollment data is always present, while authentication data is capped at N
	// most recent events.
	// Only present in certain read modes.
	// Transient.
	CollectedData []*DeviceCollectedData `protobuf:"bytes,10,rep,name=collected_data,json=collectedData,proto3" json:"collected_data,omitempty"`
	// Source of the device.
	// Devices managed directly via Teleport (`tctl`, Web UI, etc) have no
	// assigned source.
	Source *DeviceSource `protobuf:"bytes,11,opt,name=source,proto3" json:"source,omitempty"`
	// Device information acquired from an external source.
	Profile *DeviceProfile `protobuf:"bytes,12,opt,name=profile,proto3" json:"profile,omitempty"`
	// Device owner.
	// Usually the owner is the same user who performed the enrollment ceremony.
	// May be empty for legacy devices (Teleport v13.2 and older).
	// Manually unenrolling a device clears the owner.
	// System-managed.
	Owner         string `protobuf:"bytes,13,opt,name=owner,proto3" json:"owner,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Device) Reset() {
	*x = Device{}
	mi := &file_teleport_devicetrust_v1_device_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Device) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Device) ProtoMessage() {}

func (x *Device) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_devicetrust_v1_device_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Device.ProtoReflect.Descriptor instead.
func (*Device) Descriptor() ([]byte, []int) {
	return file_teleport_devicetrust_v1_device_proto_rawDescGZIP(), []int{0}
}

func (x *Device) GetApiVersion() string {
	if x != nil {
		return x.ApiVersion
	}
	return ""
}

func (x *Device) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Device) GetOsType() OSType {
	if x != nil {
		return x.OsType
	}
	return OSType_OS_TYPE_UNSPECIFIED
}

func (x *Device) GetAssetTag() string {
	if x != nil {
		return x.AssetTag
	}
	return ""
}

func (x *Device) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *Device) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

func (x *Device) GetEnrollToken() *DeviceEnrollToken {
	if x != nil {
		return x.EnrollToken
	}
	return nil
}

func (x *Device) GetEnrollStatus() DeviceEnrollStatus {
	if x != nil {
		return x.EnrollStatus
	}
	return DeviceEnrollStatus_DEVICE_ENROLL_STATUS_UNSPECIFIED
}

func (x *Device) GetCredential() *DeviceCredential {
	if x != nil {
		return x.Credential
	}
	return nil
}

func (x *Device) GetCollectedData() []*DeviceCollectedData {
	if x != nil {
		return x.CollectedData
	}
	return nil
}

func (x *Device) GetSource() *DeviceSource {
	if x != nil {
		return x.Source
	}
	return nil
}

func (x *Device) GetProfile() *DeviceProfile {
	if x != nil {
		return x.Profile
	}
	return nil
}

func (x *Device) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

// DeviceCredential represents the current enrolled public key of a device.
type DeviceCredential struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Unique identifier of the credential, defined client-side.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Device public key marshaled as a PKIX, ASN.1 DER. Used only on MacOS.
	PublicKeyDer []byte `protobuf:"bytes,2,opt,name=public_key_der,json=publicKeyDer,proto3" json:"public_key_der,omitempty"`
	// The degree to which the device credential is attested.
	DeviceAttestationType DeviceAttestationType `protobuf:"varint,3,opt,name=device_attestation_type,json=deviceAttestationType,proto3,enum=teleport.devicetrust.v1.DeviceAttestationType" json:"device_attestation_type,omitempty"`
	// For TPM devices, the serial number of the TPM endorsement certificate.
	TpmEkcertSerial string `protobuf:"bytes,4,opt,name=tpm_ekcert_serial,json=tpmEkcertSerial,proto3" json:"tpm_ekcert_serial,omitempty"`
	// For TPM devices, the encoded TPMT_PUBLIC structure containing the
	// attestation public key and signing parameters.
	TpmAkPublic   []byte `protobuf:"bytes,5,opt,name=tpm_ak_public,json=tpmAkPublic,proto3" json:"tpm_ak_public,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeviceCredential) Reset() {
	*x = DeviceCredential{}
	mi := &file_teleport_devicetrust_v1_device_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeviceCredential) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeviceCredential) ProtoMessage() {}

func (x *DeviceCredential) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_devicetrust_v1_device_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeviceCredential.ProtoReflect.Descriptor instead.
func (*DeviceCredential) Descriptor() ([]byte, []int) {
	return file_teleport_devicetrust_v1_device_proto_rawDescGZIP(), []int{1}
}

func (x *DeviceCredential) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *DeviceCredential) GetPublicKeyDer() []byte {
	if x != nil {
		return x.PublicKeyDer
	}
	return nil
}

func (x *DeviceCredential) GetDeviceAttestationType() DeviceAttestationType {
	if x != nil {
		return x.DeviceAttestationType
	}
	return DeviceAttestationType_DEVICE_ATTESTATION_TYPE_UNSPECIFIED
}

func (x *DeviceCredential) GetTpmEkcertSerial() string {
	if x != nil {
		return x.TpmEkcertSerial
	}
	return ""
}

func (x *DeviceCredential) GetTpmAkPublic() []byte {
	if x != nil {
		return x.TpmAkPublic
	}
	return nil
}

var File_teleport_devicetrust_v1_device_proto protoreflect.FileDescriptor

var file_teleport_devicetrust_v1_device_proto_rawDesc = []byte{
	0x0a, 0x24, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x74, 0x72, 0x75, 0x73, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x74, 0x72, 0x75, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x33, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x74, 0x72, 0x75, 0x73, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x5f, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x31, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f,
	0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x74, 0x72, 0x75, 0x73, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x64,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x65, 0x6e, 0x72, 0x6f, 0x6c, 0x6c, 0x5f, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2c, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x74, 0x72, 0x75, 0x73, 0x74, 0x2f, 0x76,
	0x31, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2b, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x74, 0x72, 0x75, 0x73, 0x74, 0x2f, 0x76, 0x31, 0x2f,
	0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x25, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x64, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x74, 0x72, 0x75, 0x73, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x73, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe2, 0x05, 0x0a, 0x06, 0x44,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x70, 0x69, 0x5f, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x70, 0x69, 0x56,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x38, 0x0a, 0x07, 0x6f, 0x73, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x74, 0x72, 0x75, 0x73, 0x74, 0x2e, 0x76,
	0x31, 0x2e, 0x4f, 0x53, 0x54, 0x79, 0x70, 0x65, 0x52, 0x06, 0x6f, 0x73, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x1b, 0x0a, 0x09, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x74, 0x61, 0x67, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x73, 0x73, 0x65, 0x74, 0x54, 0x61, 0x67, 0x12, 0x3b, 0x0a,
	0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x3b, 0x0a, 0x0b, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x4d, 0x0a, 0x0c, 0x65, 0x6e, 0x72, 0x6f, 0x6c,
	0x6c, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e,
	0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x74,
	0x72, 0x75, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x45, 0x6e,
	0x72, 0x6f, 0x6c, 0x6c, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x0b, 0x65, 0x6e, 0x72, 0x6f, 0x6c,
	0x6c, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x50, 0x0a, 0x0d, 0x65, 0x6e, 0x72, 0x6f, 0x6c, 0x6c,
	0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2b, 0x2e,
	0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x74,
	0x72, 0x75, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x45, 0x6e,
	0x72, 0x6f, 0x6c, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x0c, 0x65, 0x6e, 0x72, 0x6f,
	0x6c, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x49, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x74,
	0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x74, 0x72,
	0x75, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x43, 0x72, 0x65,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x61, 0x6c, 0x12, 0x53, 0x0a, 0x0e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x65, 0x64,
	0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x74, 0x65,
	0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x74, 0x72, 0x75,
	0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x43, 0x6f, 0x6c, 0x6c,
	0x65, 0x63, 0x74, 0x65, 0x64, 0x44, 0x61, 0x74, 0x61, 0x52, 0x0d, 0x63, 0x6f, 0x6c, 0x6c, 0x65,
	0x63, 0x74, 0x65, 0x64, 0x44, 0x61, 0x74, 0x61, 0x12, 0x3d, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x74, 0x72, 0x75, 0x73, 0x74, 0x2e,
	0x76, 0x31, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52,
	0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x40, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x74, 0x72, 0x75, 0x73, 0x74, 0x2e,
	0x76, 0x31, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x52, 0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x77, 0x6e,
	0x65, 0x72, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x22,
	0x80, 0x02, 0x0a, 0x10, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x61, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x24, 0x0a, 0x0e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x6b,
	0x65, 0x79, 0x5f, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c, 0x70, 0x75,
	0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x44, 0x65, 0x72, 0x12, 0x66, 0x0a, 0x17, 0x64, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x5f, 0x61, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2e, 0x2e, 0x74, 0x65,
	0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x74, 0x72, 0x75,
	0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x41, 0x74, 0x74, 0x65,
	0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x15, 0x64, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x41, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x2a, 0x0a, 0x11, 0x74, 0x70, 0x6d, 0x5f, 0x65, 0x6b, 0x63, 0x65, 0x72, 0x74,
	0x5f, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x74,
	0x70, 0x6d, 0x45, 0x6b, 0x63, 0x65, 0x72, 0x74, 0x53, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x12, 0x22,
	0x0a, 0x0d, 0x74, 0x70, 0x6d, 0x5f, 0x61, 0x6b, 0x5f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b, 0x74, 0x70, 0x6d, 0x41, 0x6b, 0x50, 0x75, 0x62, 0x6c,
	0x69, 0x63, 0x2a, 0xbf, 0x01, 0x0a, 0x15, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x41, 0x74, 0x74,
	0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x27, 0x0a, 0x23,
	0x44, 0x45, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x41, 0x54, 0x54, 0x45, 0x53, 0x54, 0x41, 0x54, 0x49,
	0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46,
	0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x25, 0x0a, 0x21, 0x44, 0x45, 0x56, 0x49, 0x43, 0x45, 0x5f,
	0x41, 0x54, 0x54, 0x45, 0x53, 0x54, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x54, 0x50, 0x4d, 0x5f, 0x45, 0x4b, 0x50, 0x55, 0x42, 0x10, 0x01, 0x12, 0x26, 0x0a, 0x22,
	0x44, 0x45, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x41, 0x54, 0x54, 0x45, 0x53, 0x54, 0x41, 0x54, 0x49,
	0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x54, 0x50, 0x4d, 0x5f, 0x45, 0x4b, 0x43, 0x45,
	0x52, 0x54, 0x10, 0x02, 0x12, 0x2e, 0x0a, 0x2a, 0x44, 0x45, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x41,
	0x54, 0x54, 0x45, 0x53, 0x54, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x54, 0x50, 0x4d, 0x5f, 0x45, 0x4b, 0x43, 0x45, 0x52, 0x54, 0x5f, 0x54, 0x52, 0x55, 0x53, 0x54,
	0x45, 0x44, 0x10, 0x03, 0x2a, 0x84, 0x01, 0x0a, 0x12, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x45,
	0x6e, 0x72, 0x6f, 0x6c, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x24, 0x0a, 0x20, 0x44,
	0x45, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x45, 0x4e, 0x52, 0x4f, 0x4c, 0x4c, 0x5f, 0x53, 0x54, 0x41,
	0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10,
	0x00, 0x12, 0x25, 0x0a, 0x21, 0x44, 0x45, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x45, 0x4e, 0x52, 0x4f,
	0x4c, 0x4c, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x45, 0x4e,
	0x52, 0x4f, 0x4c, 0x4c, 0x45, 0x44, 0x10, 0x01, 0x12, 0x21, 0x0a, 0x1d, 0x44, 0x45, 0x56, 0x49,
	0x43, 0x45, 0x5f, 0x45, 0x4e, 0x52, 0x4f, 0x4c, 0x4c, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53,
	0x5f, 0x45, 0x4e, 0x52, 0x4f, 0x4c, 0x4c, 0x45, 0x44, 0x10, 0x02, 0x42, 0x5a, 0x5a, 0x58, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x72, 0x61, 0x76, 0x69, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67,
	0x6f, 0x2f, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x74, 0x72, 0x75, 0x73, 0x74, 0x2f, 0x76, 0x31, 0x3b, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x74, 0x72, 0x75, 0x73, 0x74, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_teleport_devicetrust_v1_device_proto_rawDescOnce sync.Once
	file_teleport_devicetrust_v1_device_proto_rawDescData = file_teleport_devicetrust_v1_device_proto_rawDesc
)

func file_teleport_devicetrust_v1_device_proto_rawDescGZIP() []byte {
	file_teleport_devicetrust_v1_device_proto_rawDescOnce.Do(func() {
		file_teleport_devicetrust_v1_device_proto_rawDescData = protoimpl.X.CompressGZIP(file_teleport_devicetrust_v1_device_proto_rawDescData)
	})
	return file_teleport_devicetrust_v1_device_proto_rawDescData
}

var file_teleport_devicetrust_v1_device_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_teleport_devicetrust_v1_device_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_teleport_devicetrust_v1_device_proto_goTypes = []any{
	(DeviceAttestationType)(0),    // 0: teleport.devicetrust.v1.DeviceAttestationType
	(DeviceEnrollStatus)(0),       // 1: teleport.devicetrust.v1.DeviceEnrollStatus
	(*Device)(nil),                // 2: teleport.devicetrust.v1.Device
	(*DeviceCredential)(nil),      // 3: teleport.devicetrust.v1.DeviceCredential
	(OSType)(0),                   // 4: teleport.devicetrust.v1.OSType
	(*timestamppb.Timestamp)(nil), // 5: google.protobuf.Timestamp
	(*DeviceEnrollToken)(nil),     // 6: teleport.devicetrust.v1.DeviceEnrollToken
	(*DeviceCollectedData)(nil),   // 7: teleport.devicetrust.v1.DeviceCollectedData
	(*DeviceSource)(nil),          // 8: teleport.devicetrust.v1.DeviceSource
	(*DeviceProfile)(nil),         // 9: teleport.devicetrust.v1.DeviceProfile
}
var file_teleport_devicetrust_v1_device_proto_depIdxs = []int32{
	4,  // 0: teleport.devicetrust.v1.Device.os_type:type_name -> teleport.devicetrust.v1.OSType
	5,  // 1: teleport.devicetrust.v1.Device.create_time:type_name -> google.protobuf.Timestamp
	5,  // 2: teleport.devicetrust.v1.Device.update_time:type_name -> google.protobuf.Timestamp
	6,  // 3: teleport.devicetrust.v1.Device.enroll_token:type_name -> teleport.devicetrust.v1.DeviceEnrollToken
	1,  // 4: teleport.devicetrust.v1.Device.enroll_status:type_name -> teleport.devicetrust.v1.DeviceEnrollStatus
	3,  // 5: teleport.devicetrust.v1.Device.credential:type_name -> teleport.devicetrust.v1.DeviceCredential
	7,  // 6: teleport.devicetrust.v1.Device.collected_data:type_name -> teleport.devicetrust.v1.DeviceCollectedData
	8,  // 7: teleport.devicetrust.v1.Device.source:type_name -> teleport.devicetrust.v1.DeviceSource
	9,  // 8: teleport.devicetrust.v1.Device.profile:type_name -> teleport.devicetrust.v1.DeviceProfile
	0,  // 9: teleport.devicetrust.v1.DeviceCredential.device_attestation_type:type_name -> teleport.devicetrust.v1.DeviceAttestationType
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_teleport_devicetrust_v1_device_proto_init() }
func file_teleport_devicetrust_v1_device_proto_init() {
	if File_teleport_devicetrust_v1_device_proto != nil {
		return
	}
	file_teleport_devicetrust_v1_device_collected_data_proto_init()
	file_teleport_devicetrust_v1_device_enroll_token_proto_init()
	file_teleport_devicetrust_v1_device_profile_proto_init()
	file_teleport_devicetrust_v1_device_source_proto_init()
	file_teleport_devicetrust_v1_os_type_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_teleport_devicetrust_v1_device_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_teleport_devicetrust_v1_device_proto_goTypes,
		DependencyIndexes: file_teleport_devicetrust_v1_device_proto_depIdxs,
		EnumInfos:         file_teleport_devicetrust_v1_device_proto_enumTypes,
		MessageInfos:      file_teleport_devicetrust_v1_device_proto_msgTypes,
	}.Build()
	File_teleport_devicetrust_v1_device_proto = out.File
	file_teleport_devicetrust_v1_device_proto_rawDesc = nil
	file_teleport_devicetrust_v1_device_proto_goTypes = nil
	file_teleport_devicetrust_v1_device_proto_depIdxs = nil
}
