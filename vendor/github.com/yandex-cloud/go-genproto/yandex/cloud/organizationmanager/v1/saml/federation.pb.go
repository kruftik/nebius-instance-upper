// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.29.0
// 	protoc        v3.17.3
// source: yandex/cloud/organizationmanager/v1/saml/federation.proto

package saml

import (
	_ "github.com/yandex-cloud/go-genproto/yandex/cloud"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
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

type BindingType int32

const (
	BindingType_BINDING_TYPE_UNSPECIFIED BindingType = 0
	// HTTP POST binding.
	BindingType_POST BindingType = 1
	// HTTP redirect binding.
	BindingType_REDIRECT BindingType = 2
	// HTTP artifact binding.
	BindingType_ARTIFACT BindingType = 3
)

// Enum value maps for BindingType.
var (
	BindingType_name = map[int32]string{
		0: "BINDING_TYPE_UNSPECIFIED",
		1: "POST",
		2: "REDIRECT",
		3: "ARTIFACT",
	}
	BindingType_value = map[string]int32{
		"BINDING_TYPE_UNSPECIFIED": 0,
		"POST":                     1,
		"REDIRECT":                 2,
		"ARTIFACT":                 3,
	}
)

func (x BindingType) Enum() *BindingType {
	p := new(BindingType)
	*p = x
	return p
}

func (x BindingType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BindingType) Descriptor() protoreflect.EnumDescriptor {
	return file_yandex_cloud_organizationmanager_v1_saml_federation_proto_enumTypes[0].Descriptor()
}

func (BindingType) Type() protoreflect.EnumType {
	return &file_yandex_cloud_organizationmanager_v1_saml_federation_proto_enumTypes[0]
}

func (x BindingType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BindingType.Descriptor instead.
func (BindingType) EnumDescriptor() ([]byte, []int) {
	return file_yandex_cloud_organizationmanager_v1_saml_federation_proto_rawDescGZIP(), []int{0}
}

// A federation.
// For more information, see [SAML-compatible identity federations](/docs/iam/concepts/federations).
type Federation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the federation.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// ID of the organization that the federation belongs to.
	OrganizationId string `protobuf:"bytes,2,opt,name=organization_id,json=organizationId,proto3" json:"organization_id,omitempty"`
	// Name of the federation.
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// Description of the federation.
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	// Creation timestamp.
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// Browser cookie lifetime in seconds.
	// If the cookie is still valid, the management console
	// authenticates the user immediately and redirects them to the home page.
	CookieMaxAge *durationpb.Duration `protobuf:"bytes,6,opt,name=cookie_max_age,json=cookieMaxAge,proto3" json:"cookie_max_age,omitempty"`
	// Add new users automatically on successful authentication.
	// The user becomes member of the organization automatically,
	// but you need to grant other roles to them.
	//
	// If the value is `false`, users who aren't added to the organization
	// can't log in, even if they have authenticated on your server.
	AutoCreateAccountOnLogin bool `protobuf:"varint,7,opt,name=auto_create_account_on_login,json=autoCreateAccountOnLogin,proto3" json:"auto_create_account_on_login,omitempty"`
	// ID of the IdP server to be used for authentication.
	// The IdP server also responds to IAM with this ID after the user authenticates.
	Issuer string `protobuf:"bytes,8,opt,name=issuer,proto3" json:"issuer,omitempty"`
	// Single sign-on endpoint binding type. Most Identity Providers support the `POST` binding type.
	//
	// SAML Binding is a mapping of a SAML protocol message onto standard messaging
	// formats and/or communications protocols.
	SsoBinding BindingType `protobuf:"varint,9,opt,name=sso_binding,json=ssoBinding,proto3,enum=yandex.cloud.organizationmanager.v1.saml.BindingType" json:"sso_binding,omitempty"`
	// Single sign-on endpoint URL.
	// Specify the link to the IdP login page here.
	SsoUrl string `protobuf:"bytes,10,opt,name=sso_url,json=ssoUrl,proto3" json:"sso_url,omitempty"`
	// Federation security settings.
	SecuritySettings *FederationSecuritySettings `protobuf:"bytes,11,opt,name=security_settings,json=securitySettings,proto3" json:"security_settings,omitempty"`
	// Use case insensitive Name IDs.
	CaseInsensitiveNameIds bool `protobuf:"varint,12,opt,name=case_insensitive_name_ids,json=caseInsensitiveNameIds,proto3" json:"case_insensitive_name_ids,omitempty"`
	// Resource labels as “ key:value “ pairs. Maximum of 64 per resource.
	Labels map[string]string `protobuf:"bytes,13,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Federation) Reset() {
	*x = Federation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_organizationmanager_v1_saml_federation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Federation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Federation) ProtoMessage() {}

func (x *Federation) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_organizationmanager_v1_saml_federation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Federation.ProtoReflect.Descriptor instead.
func (*Federation) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_organizationmanager_v1_saml_federation_proto_rawDescGZIP(), []int{0}
}

func (x *Federation) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Federation) GetOrganizationId() string {
	if x != nil {
		return x.OrganizationId
	}
	return ""
}

func (x *Federation) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Federation) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Federation) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Federation) GetCookieMaxAge() *durationpb.Duration {
	if x != nil {
		return x.CookieMaxAge
	}
	return nil
}

func (x *Federation) GetAutoCreateAccountOnLogin() bool {
	if x != nil {
		return x.AutoCreateAccountOnLogin
	}
	return false
}

func (x *Federation) GetIssuer() string {
	if x != nil {
		return x.Issuer
	}
	return ""
}

func (x *Federation) GetSsoBinding() BindingType {
	if x != nil {
		return x.SsoBinding
	}
	return BindingType_BINDING_TYPE_UNSPECIFIED
}

func (x *Federation) GetSsoUrl() string {
	if x != nil {
		return x.SsoUrl
	}
	return ""
}

func (x *Federation) GetSecuritySettings() *FederationSecuritySettings {
	if x != nil {
		return x.SecuritySettings
	}
	return nil
}

func (x *Federation) GetCaseInsensitiveNameIds() bool {
	if x != nil {
		return x.CaseInsensitiveNameIds
	}
	return false
}

func (x *Federation) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

// Federation security settings.
type FederationSecuritySettings struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Enable encrypted assertions.
	EncryptedAssertions bool `protobuf:"varint,1,opt,name=encrypted_assertions,json=encryptedAssertions,proto3" json:"encrypted_assertions,omitempty"`
	// Value parameter ForceAuthn in SAMLRequest.
	ForceAuthn bool `protobuf:"varint,2,opt,name=force_authn,json=forceAuthn,proto3" json:"force_authn,omitempty"`
}

func (x *FederationSecuritySettings) Reset() {
	*x = FederationSecuritySettings{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_organizationmanager_v1_saml_federation_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FederationSecuritySettings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FederationSecuritySettings) ProtoMessage() {}

func (x *FederationSecuritySettings) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_organizationmanager_v1_saml_federation_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FederationSecuritySettings.ProtoReflect.Descriptor instead.
func (*FederationSecuritySettings) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_organizationmanager_v1_saml_federation_proto_rawDescGZIP(), []int{1}
}

func (x *FederationSecuritySettings) GetEncryptedAssertions() bool {
	if x != nil {
		return x.EncryptedAssertions
	}
	return false
}

func (x *FederationSecuritySettings) GetForceAuthn() bool {
	if x != nil {
		return x.ForceAuthn
	}
	return false
}

var File_yandex_cloud_organizationmanager_v1_saml_federation_proto protoreflect.FileDescriptor

var file_yandex_cloud_organizationmanager_v1_saml_federation_proto_rawDesc = []byte{
	0x0a, 0x39, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6f,
	0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x61, 0x6d, 0x6c, 0x2f, 0x66, 0x65, 0x64, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x28, 0x79, 0x61, 0x6e,
	0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69,
	0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x73, 0x61, 0x6d, 0x6c, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf0, 0x06, 0x0a, 0x0a, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x0c, 0xe8, 0xc7, 0x31, 0x01, 0x8a, 0xc8, 0x31, 0x04, 0x3c, 0x3d, 0x35, 0x30, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6f, 0x72, 0x67,
	0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x39, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x25, 0xe8, 0xc7, 0x31, 0x01, 0xf2,
	0xc7, 0x31, 0x1d, 0x7c, 0x5b, 0x61, 0x2d, 0x7a, 0x5d, 0x5b, 0x2d, 0x61, 0x2d, 0x7a, 0x30, 0x2d,
	0x39, 0x5d, 0x7b, 0x31, 0x2c, 0x36, 0x31, 0x7d, 0x5b, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x5d,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2b, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0x8a, 0xc8, 0x31,
	0x05, 0x3c, 0x3d, 0x32, 0x35, 0x36, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x4c,
	0x0a, 0x0e, 0x63, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x5f, 0x6d, 0x61, 0x78, 0x5f, 0x61, 0x67, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x42, 0x0b, 0xfa, 0xc7, 0x31, 0x07, 0x31, 0x30, 0x6d, 0x2d, 0x31, 0x32, 0x68, 0x52, 0x0c,
	0x63, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x4d, 0x61, 0x78, 0x41, 0x67, 0x65, 0x12, 0x3e, 0x0a, 0x1c,
	0x61, 0x75, 0x74, 0x6f, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x5f, 0x6f, 0x6e, 0x5f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x18, 0x61, 0x75, 0x74, 0x6f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4f, 0x6e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x26, 0x0a, 0x06,
	0x69, 0x73, 0x73, 0x75, 0x65, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0e, 0xe8, 0xc7,
	0x31, 0x01, 0x8a, 0xc8, 0x31, 0x06, 0x3c, 0x3d, 0x38, 0x30, 0x30, 0x30, 0x52, 0x06, 0x69, 0x73,
	0x73, 0x75, 0x65, 0x72, 0x12, 0x56, 0x0a, 0x0b, 0x73, 0x73, 0x6f, 0x5f, 0x62, 0x69, 0x6e, 0x64,
	0x69, 0x6e, 0x67, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x35, 0x2e, 0x79, 0x61, 0x6e, 0x64,
	0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x73, 0x61, 0x6d, 0x6c, 0x2e, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x0a, 0x73, 0x73, 0x6f, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x27, 0x0a, 0x07,
	0x73, 0x73, 0x6f, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0e, 0xe8,
	0xc7, 0x31, 0x01, 0x8a, 0xc8, 0x31, 0x06, 0x3c, 0x3d, 0x38, 0x30, 0x30, 0x30, 0x52, 0x06, 0x73,
	0x73, 0x6f, 0x55, 0x72, 0x6c, 0x12, 0x71, 0x0a, 0x11, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74,
	0x79, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x44, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x61, 0x6d, 0x6c, 0x2e, 0x46, 0x65, 0x64, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x53, 0x65,
	0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x10, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79,
	0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x39, 0x0a, 0x19, 0x63, 0x61, 0x73, 0x65,
	0x5f, 0x69, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x08, 0x52, 0x16, 0x63, 0x61, 0x73,
	0x65, 0x49, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x76, 0x65, 0x4e, 0x61, 0x6d, 0x65,
	0x49, 0x64, 0x73, 0x12, 0x58, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x0d, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x40, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x61, 0x6d, 0x6c, 0x2e, 0x46,
	0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x1a, 0x39, 0x0a,
	0x0b, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x70, 0x0a, 0x1a, 0x46, 0x65, 0x64, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x53, 0x65,
	0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x31, 0x0a, 0x14, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x73, 0x73, 0x65, 0x72, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x13, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x41,
	0x73, 0x73, 0x65, 0x72, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x66, 0x6f, 0x72,
	0x63, 0x65, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a,
	0x66, 0x6f, 0x72, 0x63, 0x65, 0x41, 0x75, 0x74, 0x68, 0x6e, 0x2a, 0x51, 0x0a, 0x0b, 0x42, 0x69,
	0x6e, 0x64, 0x69, 0x6e, 0x67, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x18, 0x42, 0x49, 0x4e,
	0x44, 0x49, 0x4e, 0x47, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43,
	0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x50, 0x4f, 0x53, 0x54, 0x10,
	0x01, 0x12, 0x0c, 0x0a, 0x08, 0x52, 0x45, 0x44, 0x49, 0x52, 0x45, 0x43, 0x54, 0x10, 0x02, 0x12,
	0x0c, 0x0a, 0x08, 0x41, 0x52, 0x54, 0x49, 0x46, 0x41, 0x43, 0x54, 0x10, 0x03, 0x42, 0x81, 0x01,
	0x0a, 0x2c, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x61, 0x6d, 0x6c, 0x5a, 0x51,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65,
	0x78, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2f, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x61, 0x6d, 0x6c, 0x3b, 0x73, 0x61, 0x6d,
	0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_yandex_cloud_organizationmanager_v1_saml_federation_proto_rawDescOnce sync.Once
	file_yandex_cloud_organizationmanager_v1_saml_federation_proto_rawDescData = file_yandex_cloud_organizationmanager_v1_saml_federation_proto_rawDesc
)

func file_yandex_cloud_organizationmanager_v1_saml_federation_proto_rawDescGZIP() []byte {
	file_yandex_cloud_organizationmanager_v1_saml_federation_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_organizationmanager_v1_saml_federation_proto_rawDescData = protoimpl.X.CompressGZIP(file_yandex_cloud_organizationmanager_v1_saml_federation_proto_rawDescData)
	})
	return file_yandex_cloud_organizationmanager_v1_saml_federation_proto_rawDescData
}

var file_yandex_cloud_organizationmanager_v1_saml_federation_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_yandex_cloud_organizationmanager_v1_saml_federation_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_yandex_cloud_organizationmanager_v1_saml_federation_proto_goTypes = []interface{}{
	(BindingType)(0),                   // 0: yandex.cloud.organizationmanager.v1.saml.BindingType
	(*Federation)(nil),                 // 1: yandex.cloud.organizationmanager.v1.saml.Federation
	(*FederationSecuritySettings)(nil), // 2: yandex.cloud.organizationmanager.v1.saml.FederationSecuritySettings
	nil,                                // 3: yandex.cloud.organizationmanager.v1.saml.Federation.LabelsEntry
	(*timestamppb.Timestamp)(nil),      // 4: google.protobuf.Timestamp
	(*durationpb.Duration)(nil),        // 5: google.protobuf.Duration
}
var file_yandex_cloud_organizationmanager_v1_saml_federation_proto_depIdxs = []int32{
	4, // 0: yandex.cloud.organizationmanager.v1.saml.Federation.created_at:type_name -> google.protobuf.Timestamp
	5, // 1: yandex.cloud.organizationmanager.v1.saml.Federation.cookie_max_age:type_name -> google.protobuf.Duration
	0, // 2: yandex.cloud.organizationmanager.v1.saml.Federation.sso_binding:type_name -> yandex.cloud.organizationmanager.v1.saml.BindingType
	2, // 3: yandex.cloud.organizationmanager.v1.saml.Federation.security_settings:type_name -> yandex.cloud.organizationmanager.v1.saml.FederationSecuritySettings
	3, // 4: yandex.cloud.organizationmanager.v1.saml.Federation.labels:type_name -> yandex.cloud.organizationmanager.v1.saml.Federation.LabelsEntry
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_yandex_cloud_organizationmanager_v1_saml_federation_proto_init() }
func file_yandex_cloud_organizationmanager_v1_saml_federation_proto_init() {
	if File_yandex_cloud_organizationmanager_v1_saml_federation_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_yandex_cloud_organizationmanager_v1_saml_federation_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Federation); i {
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
		file_yandex_cloud_organizationmanager_v1_saml_federation_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FederationSecuritySettings); i {
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
			RawDescriptor: file_yandex_cloud_organizationmanager_v1_saml_federation_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_yandex_cloud_organizationmanager_v1_saml_federation_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_organizationmanager_v1_saml_federation_proto_depIdxs,
		EnumInfos:         file_yandex_cloud_organizationmanager_v1_saml_federation_proto_enumTypes,
		MessageInfos:      file_yandex_cloud_organizationmanager_v1_saml_federation_proto_msgTypes,
	}.Build()
	File_yandex_cloud_organizationmanager_v1_saml_federation_proto = out.File
	file_yandex_cloud_organizationmanager_v1_saml_federation_proto_rawDesc = nil
	file_yandex_cloud_organizationmanager_v1_saml_federation_proto_goTypes = nil
	file_yandex_cloud_organizationmanager_v1_saml_federation_proto_depIdxs = nil
}
