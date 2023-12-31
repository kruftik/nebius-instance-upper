// Code generated by protoc-gen-goext. DO NOT EDIT.

package kms

import (
	operation "github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
)

func (m *CreateAsymmetricEncryptionKeyRequest) SetFolderId(v string) {
	m.FolderId = v
}

func (m *CreateAsymmetricEncryptionKeyRequest) SetName(v string) {
	m.Name = v
}

func (m *CreateAsymmetricEncryptionKeyRequest) SetDescription(v string) {
	m.Description = v
}

func (m *CreateAsymmetricEncryptionKeyRequest) SetLabels(v map[string]string) {
	m.Labels = v
}

func (m *CreateAsymmetricEncryptionKeyRequest) SetEncryptionAlgorithm(v AsymmetricEncryptionAlgorithm) {
	m.EncryptionAlgorithm = v
}

func (m *CreateAsymmetricEncryptionKeyRequest) SetDeletionProtection(v bool) {
	m.DeletionProtection = v
}

func (m *CreateAsymmetricEncryptionKeyMetadata) SetKeyId(v string) {
	m.KeyId = v
}

func (m *GetAsymmetricEncryptionKeyRequest) SetKeyId(v string) {
	m.KeyId = v
}

func (m *ListAsymmetricEncryptionKeysRequest) SetFolderId(v string) {
	m.FolderId = v
}

func (m *ListAsymmetricEncryptionKeysRequest) SetPageSize(v int64) {
	m.PageSize = v
}

func (m *ListAsymmetricEncryptionKeysRequest) SetPageToken(v string) {
	m.PageToken = v
}

func (m *ListAsymmetricEncryptionKeysResponse) SetKeys(v []*AsymmetricEncryptionKey) {
	m.Keys = v
}

func (m *ListAsymmetricEncryptionKeysResponse) SetNextPageToken(v string) {
	m.NextPageToken = v
}

func (m *UpdateAsymmetricEncryptionKeyRequest) SetKeyId(v string) {
	m.KeyId = v
}

func (m *UpdateAsymmetricEncryptionKeyRequest) SetUpdateMask(v *fieldmaskpb.FieldMask) {
	m.UpdateMask = v
}

func (m *UpdateAsymmetricEncryptionKeyRequest) SetName(v string) {
	m.Name = v
}

func (m *UpdateAsymmetricEncryptionKeyRequest) SetDescription(v string) {
	m.Description = v
}

func (m *UpdateAsymmetricEncryptionKeyRequest) SetStatus(v AsymmetricEncryptionKey_Status) {
	m.Status = v
}

func (m *UpdateAsymmetricEncryptionKeyRequest) SetLabels(v map[string]string) {
	m.Labels = v
}

func (m *UpdateAsymmetricEncryptionKeyRequest) SetDeletionProtection(v bool) {
	m.DeletionProtection = v
}

func (m *UpdateAsymmetricEncryptionKeyMetadata) SetKeyId(v string) {
	m.KeyId = v
}

func (m *DeleteAsymmetricEncryptionKeyRequest) SetKeyId(v string) {
	m.KeyId = v
}

func (m *DeleteAsymmetricEncryptionKeyMetadata) SetKeyId(v string) {
	m.KeyId = v
}

func (m *ListAsymmetricEncryptionKeyOperationsRequest) SetKeyId(v string) {
	m.KeyId = v
}

func (m *ListAsymmetricEncryptionKeyOperationsRequest) SetPageSize(v int64) {
	m.PageSize = v
}

func (m *ListAsymmetricEncryptionKeyOperationsRequest) SetPageToken(v string) {
	m.PageToken = v
}

func (m *ListAsymmetricEncryptionKeyOperationsResponse) SetOperations(v []*operation.Operation) {
	m.Operations = v
}

func (m *ListAsymmetricEncryptionKeyOperationsResponse) SetNextPageToken(v string) {
	m.NextPageToken = v
}
