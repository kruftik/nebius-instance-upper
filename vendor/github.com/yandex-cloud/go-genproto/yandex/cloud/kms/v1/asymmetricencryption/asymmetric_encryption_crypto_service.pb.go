// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.29.0
// 	protoc        v3.17.3
// source: yandex/cloud/kms/v1/asymmetricencryption/asymmetric_encryption_crypto_service.proto

package kms

import (
	_ "github.com/yandex-cloud/go-genproto/yandex/cloud"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type AsymmetricDecryptRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the asymmetric KMS key to use for decryption.
	KeyId string `protobuf:"bytes,1,opt,name=key_id,json=keyId,proto3" json:"key_id,omitempty"`
	// Ciphertext to be decrypted.
	// Should be encoded with base64.
	Ciphertext []byte `protobuf:"bytes,2,opt,name=ciphertext,proto3" json:"ciphertext,omitempty"`
}

func (x *AsymmetricDecryptRequest) Reset() {
	*x = AsymmetricDecryptRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_kms_v1_asymmetricencryption_asymmetric_encryption_crypto_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AsymmetricDecryptRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AsymmetricDecryptRequest) ProtoMessage() {}

func (x *AsymmetricDecryptRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_kms_v1_asymmetricencryption_asymmetric_encryption_crypto_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AsymmetricDecryptRequest.ProtoReflect.Descriptor instead.
func (*AsymmetricDecryptRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_kms_v1_asymmetricencryption_asymmetric_encryption_crypto_service_proto_rawDescGZIP(), []int{0}
}

func (x *AsymmetricDecryptRequest) GetKeyId() string {
	if x != nil {
		return x.KeyId
	}
	return ""
}

func (x *AsymmetricDecryptRequest) GetCiphertext() []byte {
	if x != nil {
		return x.Ciphertext
	}
	return nil
}

type AsymmetricDecryptResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the asymmetric KMS key that was used for decryption.
	KeyId string `protobuf:"bytes,1,opt,name=key_id,json=keyId,proto3" json:"key_id,omitempty"`
	// Decrypted plaintext.
	Plaintext []byte `protobuf:"bytes,2,opt,name=plaintext,proto3" json:"plaintext,omitempty"`
}

func (x *AsymmetricDecryptResponse) Reset() {
	*x = AsymmetricDecryptResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_kms_v1_asymmetricencryption_asymmetric_encryption_crypto_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AsymmetricDecryptResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AsymmetricDecryptResponse) ProtoMessage() {}

func (x *AsymmetricDecryptResponse) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_kms_v1_asymmetricencryption_asymmetric_encryption_crypto_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AsymmetricDecryptResponse.ProtoReflect.Descriptor instead.
func (*AsymmetricDecryptResponse) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_kms_v1_asymmetricencryption_asymmetric_encryption_crypto_service_proto_rawDescGZIP(), []int{1}
}

func (x *AsymmetricDecryptResponse) GetKeyId() string {
	if x != nil {
		return x.KeyId
	}
	return ""
}

func (x *AsymmetricDecryptResponse) GetPlaintext() []byte {
	if x != nil {
		return x.Plaintext
	}
	return nil
}

type AsymmetricGetPublicKeyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the asymmetric KMS key to be used for public key retrieval.
	KeyId string `protobuf:"bytes,1,opt,name=key_id,json=keyId,proto3" json:"key_id,omitempty"`
}

func (x *AsymmetricGetPublicKeyRequest) Reset() {
	*x = AsymmetricGetPublicKeyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_kms_v1_asymmetricencryption_asymmetric_encryption_crypto_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AsymmetricGetPublicKeyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AsymmetricGetPublicKeyRequest) ProtoMessage() {}

func (x *AsymmetricGetPublicKeyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_kms_v1_asymmetricencryption_asymmetric_encryption_crypto_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AsymmetricGetPublicKeyRequest.ProtoReflect.Descriptor instead.
func (*AsymmetricGetPublicKeyRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_kms_v1_asymmetricencryption_asymmetric_encryption_crypto_service_proto_rawDescGZIP(), []int{2}
}

func (x *AsymmetricGetPublicKeyRequest) GetKeyId() string {
	if x != nil {
		return x.KeyId
	}
	return ""
}

type AsymmetricGetPublicKeyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the asymmetric KMS key to get public key of.
	KeyId string `protobuf:"bytes,1,opt,name=key_id,json=keyId,proto3" json:"key_id,omitempty"`
	// Public key value.
	// The value is a PEM-encoded X.509 public key, also known as SubjectPublicKeyInfo (SPKI),
	// as defined in RFC 5280.
	PublicKey string `protobuf:"bytes,2,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
}

func (x *AsymmetricGetPublicKeyResponse) Reset() {
	*x = AsymmetricGetPublicKeyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_kms_v1_asymmetricencryption_asymmetric_encryption_crypto_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AsymmetricGetPublicKeyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AsymmetricGetPublicKeyResponse) ProtoMessage() {}

func (x *AsymmetricGetPublicKeyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_kms_v1_asymmetricencryption_asymmetric_encryption_crypto_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AsymmetricGetPublicKeyResponse.ProtoReflect.Descriptor instead.
func (*AsymmetricGetPublicKeyResponse) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_kms_v1_asymmetricencryption_asymmetric_encryption_crypto_service_proto_rawDescGZIP(), []int{3}
}

func (x *AsymmetricGetPublicKeyResponse) GetKeyId() string {
	if x != nil {
		return x.KeyId
	}
	return ""
}

func (x *AsymmetricGetPublicKeyResponse) GetPublicKey() string {
	if x != nil {
		return x.PublicKey
	}
	return ""
}

var File_yandex_cloud_kms_v1_asymmetricencryption_asymmetric_encryption_crypto_service_proto protoreflect.FileDescriptor

var file_yandex_cloud_kms_v1_asymmetricencryption_asymmetric_encryption_crypto_service_proto_rawDesc = []byte{
	0x0a, 0x53, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6b,
	0x6d, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x73, 0x79, 0x6d, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63,
	0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x61, 0x73, 0x79, 0x6d, 0x6d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x5f, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x28, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x6b, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x73, 0x79, 0x6d, 0x6d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x79,
	0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x70, 0x0a, 0x18,
	0x41, 0x73, 0x79, 0x6d, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x44, 0x65, 0x63, 0x72, 0x79, 0x70,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x06, 0x6b, 0x65, 0x79, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0c, 0xe8, 0xc7, 0x31, 0x01, 0x8a, 0xc8,
	0x31, 0x04, 0x3c, 0x3d, 0x35, 0x30, 0x52, 0x05, 0x6b, 0x65, 0x79, 0x49, 0x64, 0x12, 0x2f, 0x0a,
	0x0a, 0x63, 0x69, 0x70, 0x68, 0x65, 0x72, 0x74, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0c, 0x42, 0x0f, 0xe8, 0xc7, 0x31, 0x01, 0x8a, 0xc8, 0x31, 0x07, 0x3c, 0x3d, 0x33, 0x32, 0x37,
	0x36, 0x38, 0x52, 0x0a, 0x63, 0x69, 0x70, 0x68, 0x65, 0x72, 0x74, 0x65, 0x78, 0x74, 0x22, 0x50,
	0x0a, 0x19, 0x41, 0x73, 0x79, 0x6d, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x44, 0x65, 0x63, 0x72,
	0x79, 0x70, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x15, 0x0a, 0x06, 0x6b,
	0x65, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6b, 0x65, 0x79,
	0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x6c, 0x61, 0x69, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x70, 0x6c, 0x61, 0x69, 0x6e, 0x74, 0x65, 0x78, 0x74,
	0x22, 0x44, 0x0a, 0x1d, 0x41, 0x73, 0x79, 0x6d, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x47, 0x65,
	0x74, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x23, 0x0a, 0x06, 0x6b, 0x65, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x0c, 0xe8, 0xc7, 0x31, 0x01, 0x8a, 0xc8, 0x31, 0x04, 0x3c, 0x3d, 0x35, 0x30, 0x52,
	0x05, 0x6b, 0x65, 0x79, 0x49, 0x64, 0x22, 0x56, 0x0a, 0x1e, 0x41, 0x73, 0x79, 0x6d, 0x6d, 0x65,
	0x74, 0x72, 0x69, 0x63, 0x47, 0x65, 0x74, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x15, 0x0a, 0x06, 0x6b, 0x65, 0x79, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6b, 0x65, 0x79, 0x49, 0x64, 0x12,
	0x1d, 0x0a, 0x0a, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x32, 0xd7,
	0x03, 0x0a, 0x21, 0x41, 0x73, 0x79, 0x6d, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x45, 0x6e, 0x63,
	0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0xd0, 0x01, 0x0a, 0x07, 0x44, 0x65, 0x63, 0x72, 0x79, 0x70, 0x74,
	0x12, 0x42, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x6b, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x73, 0x79, 0x6d, 0x6d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x41, 0x73, 0x79, 0x6d,
	0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x44, 0x65, 0x63, 0x72, 0x79, 0x70, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x43, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x6b, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x73, 0x79, 0x6d, 0x6d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x41, 0x73, 0x79, 0x6d, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x44, 0x65, 0x63, 0x72, 0x79, 0x70,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x3c, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x36, 0x3a, 0x01, 0x2a, 0x22, 0x31, 0x2f, 0x6b, 0x6d, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x73,
	0x79, 0x6d, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x4b, 0x65, 0x79, 0x73, 0x2f, 0x7b, 0x6b, 0x65, 0x79, 0x5f, 0x69, 0x64, 0x7d, 0x3a,
	0x64, 0x65, 0x63, 0x72, 0x79, 0x70, 0x74, 0x12, 0xde, 0x01, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x50,
	0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x12, 0x47, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65,
	0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6b, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x61,
	0x73, 0x79, 0x6d, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x41, 0x73, 0x79, 0x6d, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x47, 0x65,
	0x74, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x48, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x6b, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x73, 0x79, 0x6d, 0x6d, 0x65, 0x74, 0x72,
	0x69, 0x63, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x41, 0x73, 0x79,
	0x6d, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x47, 0x65, 0x74, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x4b, 0x65, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x3b, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x35, 0x12, 0x33, 0x2f, 0x6b, 0x6d, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x73, 0x79,
	0x6d, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x4b, 0x65, 0x79, 0x73, 0x2f, 0x7b, 0x6b, 0x65, 0x79, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x70,
	0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x42, 0x6b, 0x0a, 0x17, 0x79, 0x61, 0x6e, 0x64,
	0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6b, 0x6d, 0x73,
	0x2e, 0x76, 0x31, 0x5a, 0x50, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67, 0x6f, 0x2d,
	0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6b, 0x6d, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x73, 0x79,
	0x6d, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x3b, 0x6b, 0x6d, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_yandex_cloud_kms_v1_asymmetricencryption_asymmetric_encryption_crypto_service_proto_rawDescOnce sync.Once
	file_yandex_cloud_kms_v1_asymmetricencryption_asymmetric_encryption_crypto_service_proto_rawDescData = file_yandex_cloud_kms_v1_asymmetricencryption_asymmetric_encryption_crypto_service_proto_rawDesc
)

func file_yandex_cloud_kms_v1_asymmetricencryption_asymmetric_encryption_crypto_service_proto_rawDescGZIP() []byte {
	file_yandex_cloud_kms_v1_asymmetricencryption_asymmetric_encryption_crypto_service_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_kms_v1_asymmetricencryption_asymmetric_encryption_crypto_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_yandex_cloud_kms_v1_asymmetricencryption_asymmetric_encryption_crypto_service_proto_rawDescData)
	})
	return file_yandex_cloud_kms_v1_asymmetricencryption_asymmetric_encryption_crypto_service_proto_rawDescData
}

var file_yandex_cloud_kms_v1_asymmetricencryption_asymmetric_encryption_crypto_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_yandex_cloud_kms_v1_asymmetricencryption_asymmetric_encryption_crypto_service_proto_goTypes = []interface{}{
	(*AsymmetricDecryptRequest)(nil),       // 0: yandex.cloud.kms.v1.asymmetricencryption.AsymmetricDecryptRequest
	(*AsymmetricDecryptResponse)(nil),      // 1: yandex.cloud.kms.v1.asymmetricencryption.AsymmetricDecryptResponse
	(*AsymmetricGetPublicKeyRequest)(nil),  // 2: yandex.cloud.kms.v1.asymmetricencryption.AsymmetricGetPublicKeyRequest
	(*AsymmetricGetPublicKeyResponse)(nil), // 3: yandex.cloud.kms.v1.asymmetricencryption.AsymmetricGetPublicKeyResponse
}
var file_yandex_cloud_kms_v1_asymmetricencryption_asymmetric_encryption_crypto_service_proto_depIdxs = []int32{
	0, // 0: yandex.cloud.kms.v1.asymmetricencryption.AsymmetricEncryptionCryptoService.Decrypt:input_type -> yandex.cloud.kms.v1.asymmetricencryption.AsymmetricDecryptRequest
	2, // 1: yandex.cloud.kms.v1.asymmetricencryption.AsymmetricEncryptionCryptoService.GetPublicKey:input_type -> yandex.cloud.kms.v1.asymmetricencryption.AsymmetricGetPublicKeyRequest
	1, // 2: yandex.cloud.kms.v1.asymmetricencryption.AsymmetricEncryptionCryptoService.Decrypt:output_type -> yandex.cloud.kms.v1.asymmetricencryption.AsymmetricDecryptResponse
	3, // 3: yandex.cloud.kms.v1.asymmetricencryption.AsymmetricEncryptionCryptoService.GetPublicKey:output_type -> yandex.cloud.kms.v1.asymmetricencryption.AsymmetricGetPublicKeyResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() {
	file_yandex_cloud_kms_v1_asymmetricencryption_asymmetric_encryption_crypto_service_proto_init()
}
func file_yandex_cloud_kms_v1_asymmetricencryption_asymmetric_encryption_crypto_service_proto_init() {
	if File_yandex_cloud_kms_v1_asymmetricencryption_asymmetric_encryption_crypto_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_yandex_cloud_kms_v1_asymmetricencryption_asymmetric_encryption_crypto_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AsymmetricDecryptRequest); i {
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
		file_yandex_cloud_kms_v1_asymmetricencryption_asymmetric_encryption_crypto_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AsymmetricDecryptResponse); i {
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
		file_yandex_cloud_kms_v1_asymmetricencryption_asymmetric_encryption_crypto_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AsymmetricGetPublicKeyRequest); i {
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
		file_yandex_cloud_kms_v1_asymmetricencryption_asymmetric_encryption_crypto_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AsymmetricGetPublicKeyResponse); i {
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
			RawDescriptor: file_yandex_cloud_kms_v1_asymmetricencryption_asymmetric_encryption_crypto_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_yandex_cloud_kms_v1_asymmetricencryption_asymmetric_encryption_crypto_service_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_kms_v1_asymmetricencryption_asymmetric_encryption_crypto_service_proto_depIdxs,
		MessageInfos:      file_yandex_cloud_kms_v1_asymmetricencryption_asymmetric_encryption_crypto_service_proto_msgTypes,
	}.Build()
	File_yandex_cloud_kms_v1_asymmetricencryption_asymmetric_encryption_crypto_service_proto = out.File
	file_yandex_cloud_kms_v1_asymmetricencryption_asymmetric_encryption_crypto_service_proto_rawDesc = nil
	file_yandex_cloud_kms_v1_asymmetricencryption_asymmetric_encryption_crypto_service_proto_goTypes = nil
	file_yandex_cloud_kms_v1_asymmetricencryption_asymmetric_encryption_crypto_service_proto_depIdxs = nil
}