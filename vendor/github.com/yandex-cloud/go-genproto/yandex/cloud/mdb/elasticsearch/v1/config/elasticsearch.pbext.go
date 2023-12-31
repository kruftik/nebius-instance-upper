// Code generated by protoc-gen-goext. DO NOT EDIT.

package elasticsearch

import (
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

func (m *ElasticsearchConfig7) SetMaxClauseCount(v *wrapperspb.Int64Value) {
	m.MaxClauseCount = v
}

func (m *ElasticsearchConfig7) SetFielddataCacheSize(v string) {
	m.FielddataCacheSize = v
}

func (m *ElasticsearchConfig7) SetReindexRemoteWhitelist(v string) {
	m.ReindexRemoteWhitelist = v
}

func (m *ElasticsearchConfig7) SetReindexSslCaPath(v string) {
	m.ReindexSslCaPath = v
}

func (m *ElasticsearchConfigSet7) SetEffectiveConfig(v *ElasticsearchConfig7) {
	m.EffectiveConfig = v
}

func (m *ElasticsearchConfigSet7) SetUserConfig(v *ElasticsearchConfig7) {
	m.UserConfig = v
}

func (m *ElasticsearchConfigSet7) SetDefaultConfig(v *ElasticsearchConfig7) {
	m.DefaultConfig = v
}
