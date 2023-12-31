// Code generated by protoc-gen-goext. DO NOT EDIT.

package clickhouse

import (
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
)

func (m *GetFormatSchemaRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *GetFormatSchemaRequest) SetFormatSchemaName(v string) {
	m.FormatSchemaName = v
}

func (m *ListFormatSchemasRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *ListFormatSchemasRequest) SetPageSize(v int64) {
	m.PageSize = v
}

func (m *ListFormatSchemasRequest) SetPageToken(v string) {
	m.PageToken = v
}

func (m *ListFormatSchemasResponse) SetFormatSchemas(v []*FormatSchema) {
	m.FormatSchemas = v
}

func (m *ListFormatSchemasResponse) SetNextPageToken(v string) {
	m.NextPageToken = v
}

func (m *CreateFormatSchemaRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *CreateFormatSchemaRequest) SetFormatSchemaName(v string) {
	m.FormatSchemaName = v
}

func (m *CreateFormatSchemaRequest) SetType(v FormatSchemaType) {
	m.Type = v
}

func (m *CreateFormatSchemaRequest) SetUri(v string) {
	m.Uri = v
}

func (m *CreateFormatSchemaMetadata) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *CreateFormatSchemaMetadata) SetFormatSchemaName(v string) {
	m.FormatSchemaName = v
}

func (m *UpdateFormatSchemaRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *UpdateFormatSchemaRequest) SetFormatSchemaName(v string) {
	m.FormatSchemaName = v
}

func (m *UpdateFormatSchemaRequest) SetUpdateMask(v *fieldmaskpb.FieldMask) {
	m.UpdateMask = v
}

func (m *UpdateFormatSchemaRequest) SetUri(v string) {
	m.Uri = v
}

func (m *UpdateFormatSchemaMetadata) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *UpdateFormatSchemaMetadata) SetFormatSchemaName(v string) {
	m.FormatSchemaName = v
}

func (m *DeleteFormatSchemaRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *DeleteFormatSchemaRequest) SetFormatSchemaName(v string) {
	m.FormatSchemaName = v
}

func (m *DeleteFormatSchemaMetadata) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *DeleteFormatSchemaMetadata) SetFormatSchemaName(v string) {
	m.FormatSchemaName = v
}
