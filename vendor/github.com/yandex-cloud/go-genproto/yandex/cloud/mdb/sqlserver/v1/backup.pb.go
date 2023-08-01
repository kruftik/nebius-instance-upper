// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.29.0
// 	protoc        v3.17.3
// source: yandex/cloud/mdb/sqlserver/v1/backup.proto

package sqlserver

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

// An SQL Server backup resource.
//
// For more information, see the [Backup](/docs/managed-sqlserver/concepts/backup) section in the documentation.
type Backup struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the backup.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// ID of the folder that the backup belongs to.
	FolderId string `protobuf:"bytes,2,opt,name=folder_id,json=folderId,proto3" json:"folder_id,omitempty"`
	// Time when the backup operation was completed.
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// ID of the SQL Server cluster that the backup was created for.
	SourceClusterId string `protobuf:"bytes,4,opt,name=source_cluster_id,json=sourceClusterId,proto3" json:"source_cluster_id,omitempty"`
	// Time when the backup operation was started.
	StartedAt *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=started_at,json=startedAt,proto3" json:"started_at,omitempty"`
	// List of databases included in the backup.
	Databases []string `protobuf:"bytes,6,rep,name=databases,proto3" json:"databases,omitempty"`
}

func (x *Backup) Reset() {
	*x = Backup{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_mdb_sqlserver_v1_backup_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Backup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Backup) ProtoMessage() {}

func (x *Backup) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_mdb_sqlserver_v1_backup_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Backup.ProtoReflect.Descriptor instead.
func (*Backup) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_mdb_sqlserver_v1_backup_proto_rawDescGZIP(), []int{0}
}

func (x *Backup) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Backup) GetFolderId() string {
	if x != nil {
		return x.FolderId
	}
	return ""
}

func (x *Backup) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Backup) GetSourceClusterId() string {
	if x != nil {
		return x.SourceClusterId
	}
	return ""
}

func (x *Backup) GetStartedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.StartedAt
	}
	return nil
}

func (x *Backup) GetDatabases() []string {
	if x != nil {
		return x.Databases
	}
	return nil
}

var File_yandex_cloud_mdb_sqlserver_v1_backup_proto protoreflect.FileDescriptor

var file_yandex_cloud_mdb_sqlserver_v1_backup_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6d,
	0x64, 0x62, 0x2f, 0x73, 0x71, 0x6c, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f,
	0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1d, 0x79, 0x61,
	0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6d, 0x64, 0x62, 0x2e, 0x73,
	0x71, 0x6c, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf5, 0x01, 0x0a,
	0x06, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x6f, 0x6c, 0x64, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x6f, 0x6c, 0x64,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12,
	0x2a, 0x0a, 0x11, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x39, 0x0a, 0x0a, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61,
	0x73, 0x65, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x64, 0x61, 0x74, 0x61, 0x62,
	0x61, 0x73, 0x65, 0x73, 0x42, 0x75, 0x0a, 0x21, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x64, 0x62, 0x2e, 0x73, 0x71, 0x6c,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x42, 0x03, 0x50, 0x53, 0x42, 0x5a, 0x4b,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65,
	0x78, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2f, 0x6d, 0x64, 0x62, 0x2f, 0x73, 0x71, 0x6c, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x76,
	0x31, 0x3b, 0x73, 0x71, 0x6c, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_yandex_cloud_mdb_sqlserver_v1_backup_proto_rawDescOnce sync.Once
	file_yandex_cloud_mdb_sqlserver_v1_backup_proto_rawDescData = file_yandex_cloud_mdb_sqlserver_v1_backup_proto_rawDesc
)

func file_yandex_cloud_mdb_sqlserver_v1_backup_proto_rawDescGZIP() []byte {
	file_yandex_cloud_mdb_sqlserver_v1_backup_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_mdb_sqlserver_v1_backup_proto_rawDescData = protoimpl.X.CompressGZIP(file_yandex_cloud_mdb_sqlserver_v1_backup_proto_rawDescData)
	})
	return file_yandex_cloud_mdb_sqlserver_v1_backup_proto_rawDescData
}

var file_yandex_cloud_mdb_sqlserver_v1_backup_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_yandex_cloud_mdb_sqlserver_v1_backup_proto_goTypes = []interface{}{
	(*Backup)(nil),                // 0: yandex.cloud.mdb.sqlserver.v1.Backup
	(*timestamppb.Timestamp)(nil), // 1: google.protobuf.Timestamp
}
var file_yandex_cloud_mdb_sqlserver_v1_backup_proto_depIdxs = []int32{
	1, // 0: yandex.cloud.mdb.sqlserver.v1.Backup.created_at:type_name -> google.protobuf.Timestamp
	1, // 1: yandex.cloud.mdb.sqlserver.v1.Backup.started_at:type_name -> google.protobuf.Timestamp
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_yandex_cloud_mdb_sqlserver_v1_backup_proto_init() }
func file_yandex_cloud_mdb_sqlserver_v1_backup_proto_init() {
	if File_yandex_cloud_mdb_sqlserver_v1_backup_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_yandex_cloud_mdb_sqlserver_v1_backup_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Backup); i {
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
			RawDescriptor: file_yandex_cloud_mdb_sqlserver_v1_backup_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_yandex_cloud_mdb_sqlserver_v1_backup_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_mdb_sqlserver_v1_backup_proto_depIdxs,
		MessageInfos:      file_yandex_cloud_mdb_sqlserver_v1_backup_proto_msgTypes,
	}.Build()
	File_yandex_cloud_mdb_sqlserver_v1_backup_proto = out.File
	file_yandex_cloud_mdb_sqlserver_v1_backup_proto_rawDesc = nil
	file_yandex_cloud_mdb_sqlserver_v1_backup_proto_goTypes = nil
	file_yandex_cloud_mdb_sqlserver_v1_backup_proto_depIdxs = nil
}
