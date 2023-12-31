// Code generated by protoc-gen-goext. DO NOT EDIT.

package loadbalancer

import (
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

func (m *TargetGroup) SetId(v string) {
	m.Id = v
}

func (m *TargetGroup) SetFolderId(v string) {
	m.FolderId = v
}

func (m *TargetGroup) SetCreatedAt(v *timestamppb.Timestamp) {
	m.CreatedAt = v
}

func (m *TargetGroup) SetName(v string) {
	m.Name = v
}

func (m *TargetGroup) SetDescription(v string) {
	m.Description = v
}

func (m *TargetGroup) SetLabels(v map[string]string) {
	m.Labels = v
}

func (m *TargetGroup) SetRegionId(v string) {
	m.RegionId = v
}

func (m *TargetGroup) SetTargets(v []*Target) {
	m.Targets = v
}

func (m *Target) SetSubnetId(v string) {
	m.SubnetId = v
}

func (m *Target) SetAddress(v string) {
	m.Address = v
}
