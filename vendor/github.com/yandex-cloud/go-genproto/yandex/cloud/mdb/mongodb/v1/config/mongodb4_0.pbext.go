// Code generated by protoc-gen-goext. DO NOT EDIT.

package mongodb

import (
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

func (m *MongodConfig4_0) SetStorage(v *MongodConfig4_0_Storage) {
	m.Storage = v
}

func (m *MongodConfig4_0) SetOperationProfiling(v *MongodConfig4_0_OperationProfiling) {
	m.OperationProfiling = v
}

func (m *MongodConfig4_0) SetNet(v *MongodConfig4_0_Network) {
	m.Net = v
}

func (m *MongodConfig4_0_Storage) SetWiredTiger(v *MongodConfig4_0_Storage_WiredTiger) {
	m.WiredTiger = v
}

func (m *MongodConfig4_0_Storage) SetJournal(v *MongodConfig4_0_Storage_Journal) {
	m.Journal = v
}

func (m *MongodConfig4_0_Storage_WiredTiger) SetEngineConfig(v *MongodConfig4_0_Storage_WiredTiger_EngineConfig) {
	m.EngineConfig = v
}

func (m *MongodConfig4_0_Storage_WiredTiger) SetCollectionConfig(v *MongodConfig4_0_Storage_WiredTiger_CollectionConfig) {
	m.CollectionConfig = v
}

func (m *MongodConfig4_0_Storage_WiredTiger_EngineConfig) SetCacheSizeGb(v *wrapperspb.DoubleValue) {
	m.CacheSizeGb = v
}

func (m *MongodConfig4_0_Storage_WiredTiger_CollectionConfig) SetBlockCompressor(v MongodConfig4_0_Storage_WiredTiger_CollectionConfig_Compressor) {
	m.BlockCompressor = v
}

func (m *MongodConfig4_0_Storage_Journal) SetCommitInterval(v *wrapperspb.Int64Value) {
	m.CommitInterval = v
}

func (m *MongodConfig4_0_OperationProfiling) SetMode(v MongodConfig4_0_OperationProfiling_Mode) {
	m.Mode = v
}

func (m *MongodConfig4_0_OperationProfiling) SetSlowOpThreshold(v *wrapperspb.Int64Value) {
	m.SlowOpThreshold = v
}

func (m *MongodConfig4_0_Network) SetMaxIncomingConnections(v *wrapperspb.Int64Value) {
	m.MaxIncomingConnections = v
}

func (m *MongoCfgConfig4_0) SetStorage(v *MongoCfgConfig4_0_Storage) {
	m.Storage = v
}

func (m *MongoCfgConfig4_0) SetOperationProfiling(v *MongoCfgConfig4_0_OperationProfiling) {
	m.OperationProfiling = v
}

func (m *MongoCfgConfig4_0) SetNet(v *MongoCfgConfig4_0_Network) {
	m.Net = v
}

func (m *MongoCfgConfig4_0_Storage) SetWiredTiger(v *MongoCfgConfig4_0_Storage_WiredTiger) {
	m.WiredTiger = v
}

func (m *MongoCfgConfig4_0_Storage_WiredTiger) SetEngineConfig(v *MongoCfgConfig4_0_Storage_WiredTiger_EngineConfig) {
	m.EngineConfig = v
}

func (m *MongoCfgConfig4_0_Storage_WiredTiger_EngineConfig) SetCacheSizeGb(v *wrapperspb.DoubleValue) {
	m.CacheSizeGb = v
}

func (m *MongoCfgConfig4_0_OperationProfiling) SetMode(v MongoCfgConfig4_0_OperationProfiling_Mode) {
	m.Mode = v
}

func (m *MongoCfgConfig4_0_OperationProfiling) SetSlowOpThreshold(v *wrapperspb.Int64Value) {
	m.SlowOpThreshold = v
}

func (m *MongoCfgConfig4_0_Network) SetMaxIncomingConnections(v *wrapperspb.Int64Value) {
	m.MaxIncomingConnections = v
}

func (m *MongosConfig4_0) SetNet(v *MongosConfig4_0_Network) {
	m.Net = v
}

func (m *MongosConfig4_0_Network) SetMaxIncomingConnections(v *wrapperspb.Int64Value) {
	m.MaxIncomingConnections = v
}

func (m *MongodConfigSet4_0) SetEffectiveConfig(v *MongodConfig4_0) {
	m.EffectiveConfig = v
}

func (m *MongodConfigSet4_0) SetUserConfig(v *MongodConfig4_0) {
	m.UserConfig = v
}

func (m *MongodConfigSet4_0) SetDefaultConfig(v *MongodConfig4_0) {
	m.DefaultConfig = v
}

func (m *MongoCfgConfigSet4_0) SetEffectiveConfig(v *MongoCfgConfig4_0) {
	m.EffectiveConfig = v
}

func (m *MongoCfgConfigSet4_0) SetUserConfig(v *MongoCfgConfig4_0) {
	m.UserConfig = v
}

func (m *MongoCfgConfigSet4_0) SetDefaultConfig(v *MongoCfgConfig4_0) {
	m.DefaultConfig = v
}

func (m *MongosConfigSet4_0) SetEffectiveConfig(v *MongosConfig4_0) {
	m.EffectiveConfig = v
}

func (m *MongosConfigSet4_0) SetUserConfig(v *MongosConfig4_0) {
	m.UserConfig = v
}

func (m *MongosConfigSet4_0) SetDefaultConfig(v *MongosConfig4_0) {
	m.DefaultConfig = v
}
