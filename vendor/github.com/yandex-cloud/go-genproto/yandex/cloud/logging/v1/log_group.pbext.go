// Code generated by protoc-gen-goext. DO NOT EDIT.

package logging

import (
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

func (m *LogGroup) SetId(v string) {
	m.Id = v
}

func (m *LogGroup) SetFolderId(v string) {
	m.FolderId = v
}

func (m *LogGroup) SetCloudId(v string) {
	m.CloudId = v
}

func (m *LogGroup) SetCreatedAt(v *timestamppb.Timestamp) {
	m.CreatedAt = v
}

func (m *LogGroup) SetName(v string) {
	m.Name = v
}

func (m *LogGroup) SetDescription(v string) {
	m.Description = v
}

func (m *LogGroup) SetLabels(v map[string]string) {
	m.Labels = v
}

func (m *LogGroup) SetStatus(v LogGroup_Status) {
	m.Status = v
}

func (m *LogGroup) SetRetentionPeriod(v *durationpb.Duration) {
	m.RetentionPeriod = v
}

func (m *LogGroup) SetDataStream(v string) {
	m.DataStream = v
}