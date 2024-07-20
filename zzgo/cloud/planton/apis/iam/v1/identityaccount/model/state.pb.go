// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: cloud/planton/apis/iam/v1/identityaccount/model/state.proto

package model

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	model "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/commons/apiresource/model"
	_ "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/commons/apiresource/options/apiresourcefieldoptions"
	_ "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/commons/apiresource/options/apiresourcemessageoptions"
	identityaccounttype "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/iam/v1/identityaccount/enums/identityaccounttype"
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

// identity-account for user-account or machine-account
type IdentityAccount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// api-version
	ApiVersion string `protobuf:"bytes,1,opt,name=api_version,json=apiVersion,proto3" json:"api_version,omitempty"`
	// resource-kind
	Kind string `protobuf:"bytes,2,opt,name=kind,proto3" json:"kind,omitempty"`
	// metadata for the resource
	// id:
	// (read-only) this is generated by idp provider.
	// for user accounts it can be in different formats.
	// * "auth0|<random-string>" when user signs up using email/password.
	// * "google-oauth|<random-string>" when user signs up using google.
	// for machine accounts it is the client_id of the oauth application created for the
	// m2m app on idp provider and is stored in the planton-cloud database.
	//
	// name:
	//
	// name of the machine account.
	// name must match the regex '^[A-Za-z0-9-]{6,16}$'.
	// It represents a string of 6-16 alphanumeric characters or hyphens, with no spaces or other special characters.
	Metadata *model.ApiResourceMetadata `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// spec
	Spec *IdentityAccountSpec `protobuf:"bytes,4,opt,name=spec,proto3" json:"spec,omitempty"`
	// status
	Status *IdentityAccountStatus `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *IdentityAccount) Reset() {
	*x = IdentityAccount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_iam_v1_identityaccount_model_state_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IdentityAccount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IdentityAccount) ProtoMessage() {}

func (x *IdentityAccount) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_iam_v1_identityaccount_model_state_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IdentityAccount.ProtoReflect.Descriptor instead.
func (*IdentityAccount) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_iam_v1_identityaccount_model_state_proto_rawDescGZIP(), []int{0}
}

func (x *IdentityAccount) GetApiVersion() string {
	if x != nil {
		return x.ApiVersion
	}
	return ""
}

func (x *IdentityAccount) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *IdentityAccount) GetMetadata() *model.ApiResourceMetadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *IdentityAccount) GetSpec() *IdentityAccountSpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

func (x *IdentityAccount) GetStatus() *IdentityAccountStatus {
	if x != nil {
		return x.Status
	}
	return nil
}

// identity-account spec
type IdentityAccountSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// this is only relevant in case of machine identity accounts.
	// machine accounts are associated to organizations on creation.
	OrgId string `protobuf:"bytes,1,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty"`
	// identity account type to identify with the account is for a machine or user
	IdentityAccountType identityaccounttype.IdentityAccountType `protobuf:"varint,2,opt,name=identity_account_type,json=identityAccountType,proto3,enum=cloud.planton.apis.iam.v1.identityaccount.enums.identityaccounttype.IdentityAccountType" json:"identity_account_type,omitempty"`
	// email of the identity account.
	// (ignored for create) this value is assigned by planton backend.
	// the format for email for machine account is <machine-account-name>@<owner-id>.iam.pmachineaccount.com
	// <machine-account-name> is entered by user at the time of creation of machine accounts
	// (update not allowed)
	// email in case user account is based on the email used to signup to the application.
	Email string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	// first name of the identity account this is more relevant in case of user accounts.
	// first name might be blank in case of machine accounts
	FirstName string `protobuf:"bytes,4,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	// last name of the identity account this is more relevant in case of user accounts.
	// last name might be blank in case of machine accounts
	LastName string `protobuf:"bytes,5,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	// (read-only) url of the identity account on idp.
	// this is constructed when a client requests for the account object.
	// the format of the url for machine accounts is https://manage.auth0.com/dashboard/us/<auth0-tenant>/applications/<client-id>/settings
	// ex: https://manage.auth0.com/dashboard/us/planton-pcs-dev/applications/pKfND0v14bJRhk8XOFGVFY15H9gNHIov/settings
	// the format of the url for user accounts is https://manage.auth0.com/dashboard/us/<auth0-tenant>/users/<client-id>/settings
	// ex: https://manage.auth0.com/dashboard/us/planton-pcs-dev/users/pKfND0v14bJRhk8XOFGVFY15H9gNHIov/settings
	IdpUrl string `protobuf:"bytes,6,opt,name=idp_url,json=idpUrl,proto3" json:"idp_url,omitempty"`
	// url of the profile picture for the user on auth0
	PictureUrl string `protobuf:"bytes,7,opt,name=picture_url,json=pictureUrl,proto3" json:"picture_url,omitempty"`
}

func (x *IdentityAccountSpec) Reset() {
	*x = IdentityAccountSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_iam_v1_identityaccount_model_state_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IdentityAccountSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IdentityAccountSpec) ProtoMessage() {}

func (x *IdentityAccountSpec) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_iam_v1_identityaccount_model_state_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IdentityAccountSpec.ProtoReflect.Descriptor instead.
func (*IdentityAccountSpec) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_iam_v1_identityaccount_model_state_proto_rawDescGZIP(), []int{1}
}

func (x *IdentityAccountSpec) GetOrgId() string {
	if x != nil {
		return x.OrgId
	}
	return ""
}

func (x *IdentityAccountSpec) GetIdentityAccountType() identityaccounttype.IdentityAccountType {
	if x != nil {
		return x.IdentityAccountType
	}
	return identityaccounttype.IdentityAccountType(0)
}

func (x *IdentityAccountSpec) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *IdentityAccountSpec) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *IdentityAccountSpec) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *IdentityAccountSpec) GetIdpUrl() string {
	if x != nil {
		return x.IdpUrl
	}
	return ""
}

func (x *IdentityAccountSpec) GetPictureUrl() string {
	if x != nil {
		return x.PictureUrl
	}
	return ""
}

// identity-account spec
type IdentityAccountStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// lifecycle of the resource
	Lifecycle *model.ApiResourceLifecycle `protobuf:"bytes,99,opt,name=lifecycle,proto3" json:"lifecycle,omitempty"`
	// audit-info
	Audit *model.ApiResourceAudit `protobuf:"bytes,98,opt,name=audit,proto3" json:"audit,omitempty"`
}

func (x *IdentityAccountStatus) Reset() {
	*x = IdentityAccountStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_iam_v1_identityaccount_model_state_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IdentityAccountStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IdentityAccountStatus) ProtoMessage() {}

func (x *IdentityAccountStatus) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_iam_v1_identityaccount_model_state_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IdentityAccountStatus.ProtoReflect.Descriptor instead.
func (*IdentityAccountStatus) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_iam_v1_identityaccount_model_state_proto_rawDescGZIP(), []int{2}
}

func (x *IdentityAccountStatus) GetLifecycle() *model.ApiResourceLifecycle {
	if x != nil {
		return x.Lifecycle
	}
	return nil
}

func (x *IdentityAccountStatus) GetAudit() *model.ApiResourceAudit {
	if x != nil {
		return x.Audit
	}
	return nil
}

// https://auth0.com/docs/customize/log-streams/custom-log-streams
type Auth0Webhook struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Date        string `protobuf:"bytes,1,opt,name=date,proto3" json:"date,omitempty"`
	Type        string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	ClientId    string `protobuf:"bytes,4,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	ClientName  string `protobuf:"bytes,5,opt,name=client_name,json=clientName,proto3" json:"client_name,omitempty"`
	Ip          string `protobuf:"bytes,6,opt,name=ip,proto3" json:"ip,omitempty"`
	UserAgent   string `protobuf:"bytes,7,opt,name=user_agent,json=userAgent,proto3" json:"user_agent,omitempty"`
	UserId      string `protobuf:"bytes,8,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *Auth0Webhook) Reset() {
	*x = Auth0Webhook{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_iam_v1_identityaccount_model_state_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Auth0Webhook) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Auth0Webhook) ProtoMessage() {}

func (x *Auth0Webhook) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_iam_v1_identityaccount_model_state_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Auth0Webhook.ProtoReflect.Descriptor instead.
func (*Auth0Webhook) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_iam_v1_identityaccount_model_state_proto_rawDescGZIP(), []int{3}
}

func (x *Auth0Webhook) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

func (x *Auth0Webhook) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Auth0Webhook) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Auth0Webhook) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *Auth0Webhook) GetClientName() string {
	if x != nil {
		return x.ClientName
	}
	return ""
}

func (x *Auth0Webhook) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *Auth0Webhook) GetUserAgent() string {
	if x != nil {
		return x.UserAgent
	}
	return ""
}

func (x *Auth0Webhook) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

var File_cloud_planton_apis_iam_v1_identityaccount_model_state_proto protoreflect.FileDescriptor

var file_cloud_planton_apis_iam_v1_identityaccount_model_state_proto_rawDesc = []byte{
	0x0a, 0x3b, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x2f, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x1a, 0x1b,
	0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x38, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x67, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61,
	0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x73, 0x2f, 0x61, 0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x6f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x70,
	0x69, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x6b,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70,
	0x69, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61,
	0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x5f, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x6f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x5f, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f,
	0x69, 0x61, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x69, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x74, 0x79, 0x70, 0x65,
	0x2f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe5, 0x02, 0x0a,
	0x0f, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x70, 0x69, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x70, 0x69, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x5d, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x41, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x41, 0x70, 0x69, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x12, 0x58, 0x0a, 0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x44, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74,
	0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x69,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x53, 0x70, 0x65, 0x63, 0x52, 0x04, 0x73, 0x70, 0x65, 0x63, 0x12, 0x5e,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x46,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x3a, 0x04,
	0x88, 0xa6, 0x1d, 0x14, 0x22, 0xf2, 0x02, 0x0a, 0x13, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x53, 0x70, 0x65, 0x63, 0x12, 0x22, 0x0a, 0x06,
	0x6f, 0x72, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0b, 0xba, 0x48,
	0x04, 0x72, 0x02, 0x10, 0x01, 0xd0, 0xb8, 0x18, 0x01, 0x52, 0x05, 0x6f, 0x72, 0x67, 0x49, 0x64,
	0x12, 0x98, 0x01, 0x0a, 0x15, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x58, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e,
	0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x69, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x65, 0x6e, 0x75,
	0x6d, 0x73, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x42, 0x0a, 0xba, 0x48, 0x03, 0xc8,
	0x01, 0x01, 0xd0, 0xb8, 0x18, 0x01, 0x52, 0x13, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0xc8, 0xb8, 0x18, 0x01,
	0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73, 0x74,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72,
	0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x07, 0x69, 0x64, 0x70, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0xc8, 0xb8, 0x18, 0x01, 0x52, 0x06, 0x69, 0x64, 0x70, 0x55,
	0x72, 0x6c, 0x12, 0x25, 0x0a, 0x0b, 0x70, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x5f, 0x75, 0x72,
	0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0xc8, 0xb8, 0x18, 0x01, 0x52, 0x0a, 0x70,
	0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x55, 0x72, 0x6c, 0x22, 0xcf, 0x01, 0x0a, 0x15, 0x49, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x60, 0x0a, 0x09, 0x6c, 0x69, 0x66, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65,
	0x18, 0x63, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x42, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70,
	0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x41, 0x70, 0x69, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x4c, 0x69, 0x66, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x52, 0x09, 0x6c, 0x69, 0x66, 0x65,
	0x63, 0x79, 0x63, 0x6c, 0x65, 0x12, 0x54, 0x0a, 0x05, 0x61, 0x75, 0x64, 0x69, 0x74, 0x18, 0x62,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x3e, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61,
	0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x73, 0x2e, 0x61, 0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x41, 0x70, 0x69, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x41,
	0x75, 0x64, 0x69, 0x74, 0x52, 0x05, 0x61, 0x75, 0x64, 0x69, 0x74, 0x22, 0xde, 0x01, 0x0a, 0x0c,
	0x41, 0x75, 0x74, 0x68, 0x30, 0x57, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x12, 0x12, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x70, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x61, 0x67, 0x65,
	0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72, 0x41, 0x67,
	0x65, 0x6e, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x42, 0x92, 0x03, 0x0a,
	0x3d, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x67, 0x65, 0x6e, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x42, 0x0a,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x5f, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2d, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x7a, 0x7a, 0x67, 0x6f, 0x2f, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73,
	0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0xa2, 0x02, 0x07,
	0x43, 0x50, 0x41, 0x49, 0x56, 0x49, 0x4d, 0xaa, 0x02, 0x2f, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x41, 0x70, 0x69, 0x73, 0x2e, 0x49, 0x61, 0x6d,
	0x2e, 0x56, 0x31, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0xca, 0x02, 0x2f, 0x43, 0x6c, 0x6f, 0x75,
	0x64, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70, 0x69, 0x73, 0x5c, 0x49,
	0x61, 0x6d, 0x5c, 0x56, 0x31, 0x5c, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5c, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0xe2, 0x02, 0x3b, 0x43, 0x6c,
	0x6f, 0x75, 0x64, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70, 0x69, 0x73,
	0x5c, 0x49, 0x61, 0x6d, 0x5c, 0x56, 0x31, 0x5c, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5c, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x5c, 0x47, 0x50,
	0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x35, 0x43, 0x6c, 0x6f, 0x75,
	0x64, 0x3a, 0x3a, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x3a, 0x3a, 0x41, 0x70, 0x69, 0x73,
	0x3a, 0x3a, 0x49, 0x61, 0x6d, 0x3a, 0x3a, 0x56, 0x31, 0x3a, 0x3a, 0x49, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x3a, 0x3a, 0x4d, 0x6f, 0x64, 0x65,
	0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cloud_planton_apis_iam_v1_identityaccount_model_state_proto_rawDescOnce sync.Once
	file_cloud_planton_apis_iam_v1_identityaccount_model_state_proto_rawDescData = file_cloud_planton_apis_iam_v1_identityaccount_model_state_proto_rawDesc
)

func file_cloud_planton_apis_iam_v1_identityaccount_model_state_proto_rawDescGZIP() []byte {
	file_cloud_planton_apis_iam_v1_identityaccount_model_state_proto_rawDescOnce.Do(func() {
		file_cloud_planton_apis_iam_v1_identityaccount_model_state_proto_rawDescData = protoimpl.X.CompressGZIP(file_cloud_planton_apis_iam_v1_identityaccount_model_state_proto_rawDescData)
	})
	return file_cloud_planton_apis_iam_v1_identityaccount_model_state_proto_rawDescData
}

var file_cloud_planton_apis_iam_v1_identityaccount_model_state_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_cloud_planton_apis_iam_v1_identityaccount_model_state_proto_goTypes = []interface{}{
	(*IdentityAccount)(nil),                      // 0: cloud.planton.apis.iam.v1.identityaccount.model.IdentityAccount
	(*IdentityAccountSpec)(nil),                  // 1: cloud.planton.apis.iam.v1.identityaccount.model.IdentityAccountSpec
	(*IdentityAccountStatus)(nil),                // 2: cloud.planton.apis.iam.v1.identityaccount.model.IdentityAccountStatus
	(*Auth0Webhook)(nil),                         // 3: cloud.planton.apis.iam.v1.identityaccount.model.Auth0Webhook
	(*model.ApiResourceMetadata)(nil),            // 4: cloud.planton.apis.commons.apiresource.model.ApiResourceMetadata
	(identityaccounttype.IdentityAccountType)(0), // 5: cloud.planton.apis.iam.v1.identityaccount.enums.identityaccounttype.IdentityAccountType
	(*model.ApiResourceLifecycle)(nil),           // 6: cloud.planton.apis.commons.apiresource.model.ApiResourceLifecycle
	(*model.ApiResourceAudit)(nil),               // 7: cloud.planton.apis.commons.apiresource.model.ApiResourceAudit
}
var file_cloud_planton_apis_iam_v1_identityaccount_model_state_proto_depIdxs = []int32{
	4, // 0: cloud.planton.apis.iam.v1.identityaccount.model.IdentityAccount.metadata:type_name -> cloud.planton.apis.commons.apiresource.model.ApiResourceMetadata
	1, // 1: cloud.planton.apis.iam.v1.identityaccount.model.IdentityAccount.spec:type_name -> cloud.planton.apis.iam.v1.identityaccount.model.IdentityAccountSpec
	2, // 2: cloud.planton.apis.iam.v1.identityaccount.model.IdentityAccount.status:type_name -> cloud.planton.apis.iam.v1.identityaccount.model.IdentityAccountStatus
	5, // 3: cloud.planton.apis.iam.v1.identityaccount.model.IdentityAccountSpec.identity_account_type:type_name -> cloud.planton.apis.iam.v1.identityaccount.enums.identityaccounttype.IdentityAccountType
	6, // 4: cloud.planton.apis.iam.v1.identityaccount.model.IdentityAccountStatus.lifecycle:type_name -> cloud.planton.apis.commons.apiresource.model.ApiResourceLifecycle
	7, // 5: cloud.planton.apis.iam.v1.identityaccount.model.IdentityAccountStatus.audit:type_name -> cloud.planton.apis.commons.apiresource.model.ApiResourceAudit
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_cloud_planton_apis_iam_v1_identityaccount_model_state_proto_init() }
func file_cloud_planton_apis_iam_v1_identityaccount_model_state_proto_init() {
	if File_cloud_planton_apis_iam_v1_identityaccount_model_state_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cloud_planton_apis_iam_v1_identityaccount_model_state_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IdentityAccount); i {
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
		file_cloud_planton_apis_iam_v1_identityaccount_model_state_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IdentityAccountSpec); i {
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
		file_cloud_planton_apis_iam_v1_identityaccount_model_state_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IdentityAccountStatus); i {
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
		file_cloud_planton_apis_iam_v1_identityaccount_model_state_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Auth0Webhook); i {
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
			RawDescriptor: file_cloud_planton_apis_iam_v1_identityaccount_model_state_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cloud_planton_apis_iam_v1_identityaccount_model_state_proto_goTypes,
		DependencyIndexes: file_cloud_planton_apis_iam_v1_identityaccount_model_state_proto_depIdxs,
		MessageInfos:      file_cloud_planton_apis_iam_v1_identityaccount_model_state_proto_msgTypes,
	}.Build()
	File_cloud_planton_apis_iam_v1_identityaccount_model_state_proto = out.File
	file_cloud_planton_apis_iam_v1_identityaccount_model_state_proto_rawDesc = nil
	file_cloud_planton_apis_iam_v1_identityaccount_model_state_proto_goTypes = nil
	file_cloud_planton_apis_iam_v1_identityaccount_model_state_proto_depIdxs = nil
}
