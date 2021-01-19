// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.6.1
// source: yandex/cloud/kms/v1/symmetric_key.proto

package kms

import (
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Supported symmetric encryption algorithms.
type SymmetricAlgorithm int32

const (
	SymmetricAlgorithm_SYMMETRIC_ALGORITHM_UNSPECIFIED SymmetricAlgorithm = 0
	// AES algorithm with 128-bit keys.
	SymmetricAlgorithm_AES_128 SymmetricAlgorithm = 1
	// AES algorithm with 192-bit keys.
	SymmetricAlgorithm_AES_192 SymmetricAlgorithm = 2
	// AES algorithm with 256-bit keys.
	SymmetricAlgorithm_AES_256 SymmetricAlgorithm = 3
)

// Enum value maps for SymmetricAlgorithm.
var (
	SymmetricAlgorithm_name = map[int32]string{
		0: "SYMMETRIC_ALGORITHM_UNSPECIFIED",
		1: "AES_128",
		2: "AES_192",
		3: "AES_256",
	}
	SymmetricAlgorithm_value = map[string]int32{
		"SYMMETRIC_ALGORITHM_UNSPECIFIED": 0,
		"AES_128":                         1,
		"AES_192":                         2,
		"AES_256":                         3,
	}
)

func (x SymmetricAlgorithm) Enum() *SymmetricAlgorithm {
	p := new(SymmetricAlgorithm)
	*p = x
	return p
}

func (x SymmetricAlgorithm) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SymmetricAlgorithm) Descriptor() protoreflect.EnumDescriptor {
	return file_yandex_cloud_kms_v1_symmetric_key_proto_enumTypes[0].Descriptor()
}

func (SymmetricAlgorithm) Type() protoreflect.EnumType {
	return &file_yandex_cloud_kms_v1_symmetric_key_proto_enumTypes[0]
}

func (x SymmetricAlgorithm) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SymmetricAlgorithm.Descriptor instead.
func (SymmetricAlgorithm) EnumDescriptor() ([]byte, []int) {
	return file_yandex_cloud_kms_v1_symmetric_key_proto_rawDescGZIP(), []int{0}
}

type SymmetricKey_Status int32

const (
	SymmetricKey_STATUS_UNSPECIFIED SymmetricKey_Status = 0
	// The key is being created.
	SymmetricKey_CREATING SymmetricKey_Status = 1
	// The key is active and can be used for encryption and decryption.
	// Can be set to INACTIVE using the [SymmetricKeyService.Update] method.
	SymmetricKey_ACTIVE SymmetricKey_Status = 2
	// The key is inactive and unusable.
	// Can be set to ACTIVE using the [SymmetricKeyService.Update] method.
	SymmetricKey_INACTIVE SymmetricKey_Status = 3
)

// Enum value maps for SymmetricKey_Status.
var (
	SymmetricKey_Status_name = map[int32]string{
		0: "STATUS_UNSPECIFIED",
		1: "CREATING",
		2: "ACTIVE",
		3: "INACTIVE",
	}
	SymmetricKey_Status_value = map[string]int32{
		"STATUS_UNSPECIFIED": 0,
		"CREATING":           1,
		"ACTIVE":             2,
		"INACTIVE":           3,
	}
)

func (x SymmetricKey_Status) Enum() *SymmetricKey_Status {
	p := new(SymmetricKey_Status)
	*p = x
	return p
}

func (x SymmetricKey_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SymmetricKey_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_yandex_cloud_kms_v1_symmetric_key_proto_enumTypes[1].Descriptor()
}

func (SymmetricKey_Status) Type() protoreflect.EnumType {
	return &file_yandex_cloud_kms_v1_symmetric_key_proto_enumTypes[1]
}

func (x SymmetricKey_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SymmetricKey_Status.Descriptor instead.
func (SymmetricKey_Status) EnumDescriptor() ([]byte, []int) {
	return file_yandex_cloud_kms_v1_symmetric_key_proto_rawDescGZIP(), []int{0, 0}
}

// Possible version status.
type SymmetricKeyVersion_Status int32

const (
	SymmetricKeyVersion_STATUS_UNSPECIFIED SymmetricKeyVersion_Status = 0
	// The version is active and can be used for encryption and decryption.
	SymmetricKeyVersion_ACTIVE SymmetricKeyVersion_Status = 1
	// The version is scheduled for destruction, the time when it will be destroyed
	// is specified in the [SymmetricKeyVersion.destroy_at] field.
	SymmetricKeyVersion_SCHEDULED_FOR_DESTRUCTION SymmetricKeyVersion_Status = 2
	// The version is destroyed and cannot be recovered.
	SymmetricKeyVersion_DESTROYED SymmetricKeyVersion_Status = 3
)

// Enum value maps for SymmetricKeyVersion_Status.
var (
	SymmetricKeyVersion_Status_name = map[int32]string{
		0: "STATUS_UNSPECIFIED",
		1: "ACTIVE",
		2: "SCHEDULED_FOR_DESTRUCTION",
		3: "DESTROYED",
	}
	SymmetricKeyVersion_Status_value = map[string]int32{
		"STATUS_UNSPECIFIED":        0,
		"ACTIVE":                    1,
		"SCHEDULED_FOR_DESTRUCTION": 2,
		"DESTROYED":                 3,
	}
)

func (x SymmetricKeyVersion_Status) Enum() *SymmetricKeyVersion_Status {
	p := new(SymmetricKeyVersion_Status)
	*p = x
	return p
}

func (x SymmetricKeyVersion_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SymmetricKeyVersion_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_yandex_cloud_kms_v1_symmetric_key_proto_enumTypes[2].Descriptor()
}

func (SymmetricKeyVersion_Status) Type() protoreflect.EnumType {
	return &file_yandex_cloud_kms_v1_symmetric_key_proto_enumTypes[2]
}

func (x SymmetricKeyVersion_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SymmetricKeyVersion_Status.Descriptor instead.
func (SymmetricKeyVersion_Status) EnumDescriptor() ([]byte, []int) {
	return file_yandex_cloud_kms_v1_symmetric_key_proto_rawDescGZIP(), []int{1, 0}
}

// A symmetric KMS key that may contain several versions of the cryptographic material.
type SymmetricKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the key.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// ID of the folder that the key belongs to.
	FolderId string `protobuf:"bytes,2,opt,name=folder_id,json=folderId,proto3" json:"folder_id,omitempty"`
	// Time when the key was created.
	CreatedAt *timestamp.Timestamp `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// Name of the key.
	Name string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	// Description of the key.
	Description string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	// Custom labels for the key as `key:value` pairs. Maximum 64 per key.
	Labels map[string]string `protobuf:"bytes,6,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Current status of the key.
	Status SymmetricKey_Status `protobuf:"varint,7,opt,name=status,proto3,enum=yandex.cloud.kms.v1.SymmetricKey_Status" json:"status,omitempty"`
	// Primary version of the key, used as the default for all encrypt/decrypt operations,
	// when no version ID is specified.
	PrimaryVersion *SymmetricKeyVersion `protobuf:"bytes,8,opt,name=primary_version,json=primaryVersion,proto3" json:"primary_version,omitempty"`
	// Default encryption algorithm to be used with new versions of the key.
	DefaultAlgorithm SymmetricAlgorithm `protobuf:"varint,9,opt,name=default_algorithm,json=defaultAlgorithm,proto3,enum=yandex.cloud.kms.v1.SymmetricAlgorithm" json:"default_algorithm,omitempty"`
	// Time of the last key rotation (time when the last version was created).
	// Empty if the key does not have versions yet.
	RotatedAt *timestamp.Timestamp `protobuf:"bytes,10,opt,name=rotated_at,json=rotatedAt,proto3" json:"rotated_at,omitempty"`
	// Time period between automatic key rotations.
	RotationPeriod *duration.Duration `protobuf:"bytes,11,opt,name=rotation_period,json=rotationPeriod,proto3" json:"rotation_period,omitempty"`
}

func (x *SymmetricKey) Reset() {
	*x = SymmetricKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_kms_v1_symmetric_key_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SymmetricKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SymmetricKey) ProtoMessage() {}

func (x *SymmetricKey) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_kms_v1_symmetric_key_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SymmetricKey.ProtoReflect.Descriptor instead.
func (*SymmetricKey) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_kms_v1_symmetric_key_proto_rawDescGZIP(), []int{0}
}

func (x *SymmetricKey) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *SymmetricKey) GetFolderId() string {
	if x != nil {
		return x.FolderId
	}
	return ""
}

func (x *SymmetricKey) GetCreatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *SymmetricKey) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SymmetricKey) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *SymmetricKey) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *SymmetricKey) GetStatus() SymmetricKey_Status {
	if x != nil {
		return x.Status
	}
	return SymmetricKey_STATUS_UNSPECIFIED
}

func (x *SymmetricKey) GetPrimaryVersion() *SymmetricKeyVersion {
	if x != nil {
		return x.PrimaryVersion
	}
	return nil
}

func (x *SymmetricKey) GetDefaultAlgorithm() SymmetricAlgorithm {
	if x != nil {
		return x.DefaultAlgorithm
	}
	return SymmetricAlgorithm_SYMMETRIC_ALGORITHM_UNSPECIFIED
}

func (x *SymmetricKey) GetRotatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.RotatedAt
	}
	return nil
}

func (x *SymmetricKey) GetRotationPeriod() *duration.Duration {
	if x != nil {
		return x.RotationPeriod
	}
	return nil
}

// Symmetric KMS key version: metadata about actual cryptographic data.
type SymmetricKeyVersion struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the key version.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// ID of the symmetric KMS key that the version belongs to.
	KeyId string `protobuf:"bytes,2,opt,name=key_id,json=keyId,proto3" json:"key_id,omitempty"`
	// Status of the key version.
	Status SymmetricKeyVersion_Status `protobuf:"varint,3,opt,name=status,proto3,enum=yandex.cloud.kms.v1.SymmetricKeyVersion_Status" json:"status,omitempty"`
	// Encryption algorithm that should be used when using the key version to encrypt plaintext.
	Algorithm SymmetricAlgorithm `protobuf:"varint,4,opt,name=algorithm,proto3,enum=yandex.cloud.kms.v1.SymmetricAlgorithm" json:"algorithm,omitempty"`
	// Time when the key version was created.
	CreatedAt *timestamp.Timestamp `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// Indication of a primary version, that is to be used by default for all cryptographic
	// operations that don't have a key version explicitly specified.
	Primary bool `protobuf:"varint,6,opt,name=primary,proto3" json:"primary,omitempty"`
	// Time when the key version is going to be destroyed. Empty unless the status
	// is `SCHEDULED_FOR_DESTRUCTION`.
	DestroyAt *timestamp.Timestamp `protobuf:"bytes,7,opt,name=destroy_at,json=destroyAt,proto3" json:"destroy_at,omitempty"`
}

func (x *SymmetricKeyVersion) Reset() {
	*x = SymmetricKeyVersion{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_kms_v1_symmetric_key_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SymmetricKeyVersion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SymmetricKeyVersion) ProtoMessage() {}

func (x *SymmetricKeyVersion) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_kms_v1_symmetric_key_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SymmetricKeyVersion.ProtoReflect.Descriptor instead.
func (*SymmetricKeyVersion) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_kms_v1_symmetric_key_proto_rawDescGZIP(), []int{1}
}

func (x *SymmetricKeyVersion) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *SymmetricKeyVersion) GetKeyId() string {
	if x != nil {
		return x.KeyId
	}
	return ""
}

func (x *SymmetricKeyVersion) GetStatus() SymmetricKeyVersion_Status {
	if x != nil {
		return x.Status
	}
	return SymmetricKeyVersion_STATUS_UNSPECIFIED
}

func (x *SymmetricKeyVersion) GetAlgorithm() SymmetricAlgorithm {
	if x != nil {
		return x.Algorithm
	}
	return SymmetricAlgorithm_SYMMETRIC_ALGORITHM_UNSPECIFIED
}

func (x *SymmetricKeyVersion) GetCreatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *SymmetricKeyVersion) GetPrimary() bool {
	if x != nil {
		return x.Primary
	}
	return false
}

func (x *SymmetricKeyVersion) GetDestroyAt() *timestamp.Timestamp {
	if x != nil {
		return x.DestroyAt
	}
	return nil
}

var File_yandex_cloud_kms_v1_symmetric_key_proto protoreflect.FileDescriptor

var file_yandex_cloud_kms_v1_symmetric_key_proto_rawDesc = []byte{
	0x0a, 0x27, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6b,
	0x6d, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x79, 0x6d, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x5f,
	0x6b, 0x65, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x79, 0x61, 0x6e, 0x64, 0x65,
	0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6b, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xe2, 0x05, 0x0a, 0x0c, 0x53, 0x79, 0x6d, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x4b, 0x65, 0x79,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x1b, 0x0a, 0x09, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x39, 0x0a,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x45,
	0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2d,
	0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6b, 0x6d,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x79, 0x6d, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x4b, 0x65,
	0x79, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x6c,
	0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x40, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x28, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6b, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x79, 0x6d, 0x6d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x51, 0x0a, 0x0f, 0x70, 0x72, 0x69, 0x6d, 0x61,
	0x72, 0x79, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x28, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x6b, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x79, 0x6d, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63,
	0x4b, 0x65, 0x79, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x0e, 0x70, 0x72, 0x69, 0x6d,
	0x61, 0x72, 0x79, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x54, 0x0a, 0x11, 0x64, 0x65,
	0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x61, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x27, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6b, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x79, 0x6d, 0x6d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x41, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x52, 0x10,
	0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x41, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d,
	0x12, 0x39, 0x0a, 0x0a, 0x72, 0x6f, 0x74, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x09, 0x72, 0x6f, 0x74, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x42, 0x0a, 0x0f, 0x72,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x0e, 0x72, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x1a,
	0x39, 0x0a, 0x0b, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x48, 0x0a, 0x06, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x12, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55,
	0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08,
	0x43, 0x52, 0x45, 0x41, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x41, 0x43,
	0x54, 0x49, 0x56, 0x45, 0x10, 0x02, 0x12, 0x0c, 0x0a, 0x08, 0x49, 0x4e, 0x41, 0x43, 0x54, 0x49,
	0x56, 0x45, 0x10, 0x03, 0x22, 0xb8, 0x03, 0x0a, 0x13, 0x53, 0x79, 0x6d, 0x6d, 0x65, 0x74, 0x72,
	0x69, 0x63, 0x4b, 0x65, 0x79, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x15, 0x0a, 0x06,
	0x6b, 0x65, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6b, 0x65,
	0x79, 0x49, 0x64, 0x12, 0x47, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x2f, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x6b, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x79, 0x6d, 0x6d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x45, 0x0a, 0x09,
	0x61, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x27, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6b,
	0x6d, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x79, 0x6d, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x41,
	0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x52, 0x09, 0x61, 0x6c, 0x67, 0x6f, 0x72, 0x69,
	0x74, 0x68, 0x6d, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x18,
	0x0a, 0x07, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x07, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x39, 0x0a, 0x0a, 0x64, 0x65, 0x73, 0x74,
	0x72, 0x6f, 0x79, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x64, 0x65, 0x73, 0x74, 0x72, 0x6f,
	0x79, 0x41, 0x74, 0x22, 0x5a, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a,
	0x12, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46,
	0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x41, 0x43, 0x54, 0x49, 0x56, 0x45, 0x10,
	0x01, 0x12, 0x1d, 0x0a, 0x19, 0x53, 0x43, 0x48, 0x45, 0x44, 0x55, 0x4c, 0x45, 0x44, 0x5f, 0x46,
	0x4f, 0x52, 0x5f, 0x44, 0x45, 0x53, 0x54, 0x52, 0x55, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x02,
	0x12, 0x0d, 0x0a, 0x09, 0x44, 0x45, 0x53, 0x54, 0x52, 0x4f, 0x59, 0x45, 0x44, 0x10, 0x03, 0x2a,
	0x60, 0x0a, 0x12, 0x53, 0x79, 0x6d, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x41, 0x6c, 0x67, 0x6f,
	0x72, 0x69, 0x74, 0x68, 0x6d, 0x12, 0x23, 0x0a, 0x1f, 0x53, 0x59, 0x4d, 0x4d, 0x45, 0x54, 0x52,
	0x49, 0x43, 0x5f, 0x41, 0x4c, 0x47, 0x4f, 0x52, 0x49, 0x54, 0x48, 0x4d, 0x5f, 0x55, 0x4e, 0x53,
	0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x41, 0x45,
	0x53, 0x5f, 0x31, 0x32, 0x38, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x41, 0x45, 0x53, 0x5f, 0x31,
	0x39, 0x32, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x41, 0x45, 0x53, 0x5f, 0x32, 0x35, 0x36, 0x10,
	0x03, 0x42, 0x56, 0x0a, 0x17, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6b, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x5a, 0x3b, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2d,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6b,
	0x6d, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x6b, 0x6d, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_yandex_cloud_kms_v1_symmetric_key_proto_rawDescOnce sync.Once
	file_yandex_cloud_kms_v1_symmetric_key_proto_rawDescData = file_yandex_cloud_kms_v1_symmetric_key_proto_rawDesc
)

func file_yandex_cloud_kms_v1_symmetric_key_proto_rawDescGZIP() []byte {
	file_yandex_cloud_kms_v1_symmetric_key_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_kms_v1_symmetric_key_proto_rawDescData = protoimpl.X.CompressGZIP(file_yandex_cloud_kms_v1_symmetric_key_proto_rawDescData)
	})
	return file_yandex_cloud_kms_v1_symmetric_key_proto_rawDescData
}

var file_yandex_cloud_kms_v1_symmetric_key_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_yandex_cloud_kms_v1_symmetric_key_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_yandex_cloud_kms_v1_symmetric_key_proto_goTypes = []interface{}{
	(SymmetricAlgorithm)(0),         // 0: yandex.cloud.kms.v1.SymmetricAlgorithm
	(SymmetricKey_Status)(0),        // 1: yandex.cloud.kms.v1.SymmetricKey.Status
	(SymmetricKeyVersion_Status)(0), // 2: yandex.cloud.kms.v1.SymmetricKeyVersion.Status
	(*SymmetricKey)(nil),            // 3: yandex.cloud.kms.v1.SymmetricKey
	(*SymmetricKeyVersion)(nil),     // 4: yandex.cloud.kms.v1.SymmetricKeyVersion
	nil,                             // 5: yandex.cloud.kms.v1.SymmetricKey.LabelsEntry
	(*timestamp.Timestamp)(nil),     // 6: google.protobuf.Timestamp
	(*duration.Duration)(nil),       // 7: google.protobuf.Duration
}
var file_yandex_cloud_kms_v1_symmetric_key_proto_depIdxs = []int32{
	6,  // 0: yandex.cloud.kms.v1.SymmetricKey.created_at:type_name -> google.protobuf.Timestamp
	5,  // 1: yandex.cloud.kms.v1.SymmetricKey.labels:type_name -> yandex.cloud.kms.v1.SymmetricKey.LabelsEntry
	1,  // 2: yandex.cloud.kms.v1.SymmetricKey.status:type_name -> yandex.cloud.kms.v1.SymmetricKey.Status
	4,  // 3: yandex.cloud.kms.v1.SymmetricKey.primary_version:type_name -> yandex.cloud.kms.v1.SymmetricKeyVersion
	0,  // 4: yandex.cloud.kms.v1.SymmetricKey.default_algorithm:type_name -> yandex.cloud.kms.v1.SymmetricAlgorithm
	6,  // 5: yandex.cloud.kms.v1.SymmetricKey.rotated_at:type_name -> google.protobuf.Timestamp
	7,  // 6: yandex.cloud.kms.v1.SymmetricKey.rotation_period:type_name -> google.protobuf.Duration
	2,  // 7: yandex.cloud.kms.v1.SymmetricKeyVersion.status:type_name -> yandex.cloud.kms.v1.SymmetricKeyVersion.Status
	0,  // 8: yandex.cloud.kms.v1.SymmetricKeyVersion.algorithm:type_name -> yandex.cloud.kms.v1.SymmetricAlgorithm
	6,  // 9: yandex.cloud.kms.v1.SymmetricKeyVersion.created_at:type_name -> google.protobuf.Timestamp
	6,  // 10: yandex.cloud.kms.v1.SymmetricKeyVersion.destroy_at:type_name -> google.protobuf.Timestamp
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_yandex_cloud_kms_v1_symmetric_key_proto_init() }
func file_yandex_cloud_kms_v1_symmetric_key_proto_init() {
	if File_yandex_cloud_kms_v1_symmetric_key_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_yandex_cloud_kms_v1_symmetric_key_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SymmetricKey); i {
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
		file_yandex_cloud_kms_v1_symmetric_key_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SymmetricKeyVersion); i {
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
			RawDescriptor: file_yandex_cloud_kms_v1_symmetric_key_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_yandex_cloud_kms_v1_symmetric_key_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_kms_v1_symmetric_key_proto_depIdxs,
		EnumInfos:         file_yandex_cloud_kms_v1_symmetric_key_proto_enumTypes,
		MessageInfos:      file_yandex_cloud_kms_v1_symmetric_key_proto_msgTypes,
	}.Build()
	File_yandex_cloud_kms_v1_symmetric_key_proto = out.File
	file_yandex_cloud_kms_v1_symmetric_key_proto_rawDesc = nil
	file_yandex_cloud_kms_v1_symmetric_key_proto_goTypes = nil
	file_yandex_cloud_kms_v1_symmetric_key_proto_depIdxs = nil
}
