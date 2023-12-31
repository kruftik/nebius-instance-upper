// Code generated by protoc-gen-goext. DO NOT EDIT.

package kms

import (
	operation "github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
)

func (m *CreateAsymmetricSignatureKeyRequest) SetFolderId(v string) {
	m.FolderId = v
}

func (m *CreateAsymmetricSignatureKeyRequest) SetName(v string) {
	m.Name = v
}

func (m *CreateAsymmetricSignatureKeyRequest) SetDescription(v string) {
	m.Description = v
}

func (m *CreateAsymmetricSignatureKeyRequest) SetLabels(v map[string]string) {
	m.Labels = v
}

func (m *CreateAsymmetricSignatureKeyRequest) SetSignatureAlgorithm(v AsymmetricSignatureAlgorithm) {
	m.SignatureAlgorithm = v
}

func (m *CreateAsymmetricSignatureKeyRequest) SetDeletionProtection(v bool) {
	m.DeletionProtection = v
}

func (m *CreateAsymmetricSignatureKeyMetadata) SetKeyId(v string) {
	m.KeyId = v
}

func (m *GetAsymmetricSignatureKeyRequest) SetKeyId(v string) {
	m.KeyId = v
}

func (m *ListAsymmetricSignatureKeysRequest) SetFolderId(v string) {
	m.FolderId = v
}

func (m *ListAsymmetricSignatureKeysRequest) SetPageSize(v int64) {
	m.PageSize = v
}

func (m *ListAsymmetricSignatureKeysRequest) SetPageToken(v string) {
	m.PageToken = v
}

func (m *ListAsymmetricSignatureKeysResponse) SetKeys(v []*AsymmetricSignatureKey) {
	m.Keys = v
}

func (m *ListAsymmetricSignatureKeysResponse) SetNextPageToken(v string) {
	m.NextPageToken = v
}

func (m *UpdateAsymmetricSignatureKeyRequest) SetKeyId(v string) {
	m.KeyId = v
}

func (m *UpdateAsymmetricSignatureKeyRequest) SetUpdateMask(v *fieldmaskpb.FieldMask) {
	m.UpdateMask = v
}

func (m *UpdateAsymmetricSignatureKeyRequest) SetName(v string) {
	m.Name = v
}

func (m *UpdateAsymmetricSignatureKeyRequest) SetDescription(v string) {
	m.Description = v
}

func (m *UpdateAsymmetricSignatureKeyRequest) SetStatus(v AsymmetricSignatureKey_Status) {
	m.Status = v
}

func (m *UpdateAsymmetricSignatureKeyRequest) SetLabels(v map[string]string) {
	m.Labels = v
}

func (m *UpdateAsymmetricSignatureKeyRequest) SetDeletionProtection(v bool) {
	m.DeletionProtection = v
}

func (m *UpdateAsymmetricSignatureKeyMetadata) SetKeyId(v string) {
	m.KeyId = v
}

func (m *DeleteAsymmetricSignatureKeyRequest) SetKeyId(v string) {
	m.KeyId = v
}

func (m *DeleteAsymmetricSignatureKeyMetadata) SetKeyId(v string) {
	m.KeyId = v
}

func (m *ListAsymmetricSignatureKeyOperationsRequest) SetKeyId(v string) {
	m.KeyId = v
}

func (m *ListAsymmetricSignatureKeyOperationsRequest) SetPageSize(v int64) {
	m.PageSize = v
}

func (m *ListAsymmetricSignatureKeyOperationsRequest) SetPageToken(v string) {
	m.PageToken = v
}

func (m *ListAsymmetricSignatureKeyOperationsResponse) SetOperations(v []*operation.Operation) {
	m.Operations = v
}

func (m *ListAsymmetricSignatureKeyOperationsResponse) SetNextPageToken(v string) {
	m.NextPageToken = v
}
