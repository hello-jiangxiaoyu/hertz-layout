// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.9
// source: oauth.proto

package oauth

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	_ "hertz/demo/biz/model/api"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type NoReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *NoReq) Reset() {
	*x = NoReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oauth_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NoReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NoReq) ProtoMessage() {}

func (x *NoReq) ProtoReflect() protoreflect.Message {
	mi := &file_oauth_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NoReq.ProtoReflect.Descriptor instead.
func (*NoReq) Descriptor() ([]byte, []int) {
	return file_oauth_proto_rawDescGZIP(), []int{0}
}

type CommonResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int64      `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty" form:"code" query:"code"`
	Msg  string     `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty" form:"msg" query:"msg"`
	Data *anypb.Any `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty" form:"data" query:"data"`
}

func (x *CommonResp) Reset() {
	*x = CommonResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oauth_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommonResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommonResp) ProtoMessage() {}

func (x *CommonResp) ProtoReflect() protoreflect.Message {
	mi := &file_oauth_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommonResp.ProtoReflect.Descriptor instead.
func (*CommonResp) Descriptor() ([]byte, []int) {
	return file_oauth_proto_rawDescGZIP(), []int{1}
}

func (x *CommonResp) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *CommonResp) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *CommonResp) GetData() *anypb.Any {
	if x != nil {
		return x.Data
	}
	return nil
}

// 登录
type LoginReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty" form:"username" vd:"(len($) > 2 && len($) < 100)"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty" form:"username" vd:"(len($) > 6 && len($) < 100)"`
}

func (x *LoginReq) Reset() {
	*x = LoginReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oauth_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginReq) ProtoMessage() {}

func (x *LoginReq) ProtoReflect() protoreflect.Message {
	mi := &file_oauth_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginReq.ProtoReflect.Descriptor instead.
func (*LoginReq) Descriptor() ([]byte, []int) {
	return file_oauth_proto_rawDescGZIP(), []int{2}
}

func (x *LoginReq) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *LoginReq) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type TypeReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty" path:"type" vd:"(len($) > 2 && len($) < 100)"`
}

func (x *TypeReq) Reset() {
	*x = TypeReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oauth_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TypeReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TypeReq) ProtoMessage() {}

func (x *TypeReq) ProtoReflect() protoreflect.Message {
	mi := &file_oauth_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TypeReq.ProtoReflect.Descriptor instead.
func (*TypeReq) Descriptor() ([]byte, []int) {
	return file_oauth_proto_rawDescGZIP(), []int{3}
}

func (x *TypeReq) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DisplayName string `protobuf:"bytes,1,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty" form:"display_name" query:"display_name"`
	Gender      string `protobuf:"bytes,2,opt,name=gender,proto3" json:"gender,omitempty" form:"gender" query:"gender"`
	Avatar      string `protobuf:"bytes,3,opt,name=avatar,proto3" json:"avatar,omitempty" form:"avatar" query:"avatar"`
	Type        int32  `protobuf:"varint,4,opt,name=type,proto3" json:"type,omitempty" form:"type" query:"type"`
	IsOnline    bool   `protobuf:"varint,5,opt,name=is_online,json=isOnline,proto3" json:"is_online,omitempty" form:"is_online" query:"is_online"`
	IsDisabled  bool   `protobuf:"varint,6,opt,name=is_disabled,json=isDisabled,proto3" json:"is_disabled,omitempty" form:"is_disabled" query:"is_disabled"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oauth_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_oauth_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_oauth_proto_rawDescGZIP(), []int{4}
}

func (x *User) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *User) GetGender() string {
	if x != nil {
		return x.Gender
	}
	return ""
}

func (x *User) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *User) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *User) GetIsOnline() bool {
	if x != nil {
		return x.IsOnline
	}
	return false
}

func (x *User) GetIsDisabled() bool {
	if x != nil {
		return x.IsDisabled
	}
	return false
}

// 授权
type TokenReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientID     string `protobuf:"bytes,1,opt,name=clientID,proto3" json:"clientID,omitempty" query:"client_id" vd:"len($) < 100"`
	ClientSecret string `protobuf:"bytes,2,opt,name=clientSecret,proto3" json:"clientSecret,omitempty" query:"client_secret" vd:"len($) < 100"`
	GrantType    string `protobuf:"bytes,3,opt,name=grantType,proto3" json:"grantType,omitempty" query:"grant_type" vd:"len($) < 100"`
}

func (x *TokenReq) Reset() {
	*x = TokenReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oauth_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokenReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenReq) ProtoMessage() {}

func (x *TokenReq) ProtoReflect() protoreflect.Message {
	mi := &file_oauth_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenReq.ProtoReflect.Descriptor instead.
func (*TokenReq) Descriptor() ([]byte, []int) {
	return file_oauth_proto_rawDescGZIP(), []int{5}
}

func (x *TokenReq) GetClientID() string {
	if x != nil {
		return x.ClientID
	}
	return ""
}

func (x *TokenReq) GetClientSecret() string {
	if x != nil {
		return x.ClientSecret
	}
	return ""
}

func (x *TokenReq) GetGrantType() string {
	if x != nil {
		return x.GrantType
	}
	return ""
}

type JWKSKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Use string `protobuf:"bytes,1,opt,name=use,proto3" form:"use" json:"use,omitempty"`
	Kty string `protobuf:"bytes,2,opt,name=kty,proto3" form:"kty" json:"kty,omitempty"`
	Kid string `protobuf:"bytes,3,opt,name=kid,proto3" form:"kid" json:"kid,omitempty"`
	Alg string `protobuf:"bytes,4,opt,name=alg,proto3" form:"alg" json:"alg,omitempty"`
	N   string `protobuf:"bytes,5,opt,name=n,proto3" form:"e" json:"e,omitempty"`
	E   string `protobuf:"bytes,6,opt,name=e,proto3" form:"n" json:"n,omitempty"`
}

func (x *JWKSKey) Reset() {
	*x = JWKSKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oauth_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JWKSKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JWKSKey) ProtoMessage() {}

func (x *JWKSKey) ProtoReflect() protoreflect.Message {
	mi := &file_oauth_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JWKSKey.ProtoReflect.Descriptor instead.
func (*JWKSKey) Descriptor() ([]byte, []int) {
	return file_oauth_proto_rawDescGZIP(), []int{6}
}

func (x *JWKSKey) GetUse() string {
	if x != nil {
		return x.Use
	}
	return ""
}

func (x *JWKSKey) GetKty() string {
	if x != nil {
		return x.Kty
	}
	return ""
}

func (x *JWKSKey) GetKid() string {
	if x != nil {
		return x.Kid
	}
	return ""
}

func (x *JWKSKey) GetAlg() string {
	if x != nil {
		return x.Alg
	}
	return ""
}

func (x *JWKSKey) GetN() string {
	if x != nil {
		return x.N
	}
	return ""
}

func (x *JWKSKey) GetE() string {
	if x != nil {
		return x.E
	}
	return ""
}

type JWKSResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Keys []*JWKSKey `protobuf:"bytes,1,rep,name=keys,proto3" json:"keys,omitempty" form:"keys" query:"keys"`
}

func (x *JWKSResp) Reset() {
	*x = JWKSResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oauth_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JWKSResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JWKSResp) ProtoMessage() {}

func (x *JWKSResp) ProtoReflect() protoreflect.Message {
	mi := &file_oauth_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JWKSResp.ProtoReflect.Descriptor instead.
func (*JWKSResp) Descriptor() ([]byte, []int) {
	return file_oauth_proto_rawDescGZIP(), []int{7}
}

func (x *JWKSResp) GetKeys() []*JWKSKey {
	if x != nil {
		return x.Keys
	}
	return nil
}

var File_oauth_proto protoreflect.FileDescriptor

var file_oauth_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6f,
	0x61, 0x75, 0x74, 0x68, 0x1a, 0x09, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x07, 0x0a, 0x05, 0x4e, 0x6f,
	0x52, 0x65, 0x71, 0x22, 0x5c, 0x0a, 0x0a, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x28, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x22, 0x9e, 0x01, 0x0a, 0x08, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x48,
	0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x2c, 0xda, 0xbb, 0x18, 0x1c, 0x28, 0x6c, 0x65, 0x6e, 0x28, 0x24, 0x29, 0x20, 0x3e, 0x20,
	0x32, 0x20, 0x26, 0x26, 0x20, 0x6c, 0x65, 0x6e, 0x28, 0x24, 0x29, 0x20, 0x3c, 0x20, 0x31, 0x30,
	0x30, 0x29, 0xe2, 0xbb, 0x18, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x52, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x48, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x2c, 0xda, 0xbb, 0x18, 0x1c,
	0x28, 0x6c, 0x65, 0x6e, 0x28, 0x24, 0x29, 0x20, 0x3e, 0x20, 0x36, 0x20, 0x26, 0x26, 0x20, 0x6c,
	0x65, 0x6e, 0x28, 0x24, 0x29, 0x20, 0x3c, 0x20, 0x31, 0x30, 0x30, 0x29, 0xe2, 0xbb, 0x18, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x22, 0x47, 0x0a, 0x07, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x71, 0x12, 0x3c, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x28, 0xd2, 0xbb, 0x18,
	0x04, 0x74, 0x79, 0x70, 0x65, 0xda, 0xbb, 0x18, 0x1c, 0x28, 0x6c, 0x65, 0x6e, 0x28, 0x24, 0x29,
	0x20, 0x3e, 0x20, 0x32, 0x20, 0x26, 0x26, 0x20, 0x6c, 0x65, 0x6e, 0x28, 0x24, 0x29, 0x20, 0x3c,
	0x20, 0x31, 0x30, 0x30, 0x29, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0xab, 0x01, 0x0a, 0x04,
	0x55, 0x73, 0x65, 0x72, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70,
	0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12,
	0x16, 0x0a, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x69,
	0x73, 0x5f, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08,
	0x69, 0x73, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x73, 0x5f, 0x64,
	0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x69,
	0x73, 0x44, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x22, 0xca, 0x01, 0x0a, 0x08, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x39, 0x0a, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1d, 0xb2, 0xbb, 0x18, 0x09, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0xda, 0xbb, 0x18, 0x0c, 0x6c, 0x65, 0x6e, 0x28, 0x24,
	0x29, 0x20, 0x3c, 0x20, 0x31, 0x30, 0x30, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49,
	0x44, 0x12, 0x45, 0x0a, 0x0c, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x63, 0x72, 0x65,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x21, 0xb2, 0xbb, 0x18, 0x0d, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0xda, 0xbb, 0x18, 0x0c, 0x6c, 0x65,
	0x6e, 0x28, 0x24, 0x29, 0x20, 0x3c, 0x20, 0x31, 0x30, 0x30, 0x52, 0x0c, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x3c, 0x0a, 0x09, 0x67, 0x72, 0x61, 0x6e,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1e, 0xb2, 0xbb, 0x18,
	0x0a, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0xda, 0xbb, 0x18, 0x0c, 0x6c,
	0x65, 0x6e, 0x28, 0x24, 0x29, 0x20, 0x3c, 0x20, 0x31, 0x30, 0x30, 0x52, 0x09, 0x67, 0x72, 0x61,
	0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x22, 0x9f, 0x01, 0x0a, 0x07, 0x4a, 0x57, 0x4b, 0x53, 0x4b,
	0x65, 0x79, 0x12, 0x19, 0x0a, 0x03, 0x75, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x07, 0xca, 0xbb, 0x18, 0x03, 0x75, 0x73, 0x65, 0x52, 0x03, 0x75, 0x73, 0x65, 0x12, 0x19, 0x0a,
	0x03, 0x6b, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xca, 0xbb, 0x18, 0x03,
	0x6b, 0x74, 0x79, 0x52, 0x03, 0x6b, 0x74, 0x79, 0x12, 0x19, 0x0a, 0x03, 0x6b, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xca, 0xbb, 0x18, 0x03, 0x6b, 0x69, 0x64, 0x52, 0x03,
	0x6b, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x03, 0x61, 0x6c, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x07, 0xca, 0xbb, 0x18, 0x03, 0x61, 0x6c, 0x67, 0x52, 0x03, 0x61, 0x6c, 0x67, 0x12, 0x13,
	0x0a, 0x01, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x05, 0xca, 0xbb, 0x18, 0x01, 0x65,
	0x52, 0x01, 0x6e, 0x12, 0x13, 0x0a, 0x01, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x42, 0x05,
	0xca, 0xbb, 0x18, 0x01, 0x6e, 0x52, 0x01, 0x65, 0x22, 0x2e, 0x0a, 0x08, 0x4a, 0x57, 0x4b, 0x53,
	0x52, 0x65, 0x73, 0x70, 0x12, 0x22, 0x0a, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x4a, 0x57, 0x4b, 0x53, 0x4b,
	0x65, 0x79, 0x52, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x32, 0xf8, 0x02, 0x0a, 0x05, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x12, 0x4b, 0x0a, 0x0d, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x50, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x12, 0x0f, 0x2e, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x4c, 0x6f, 0x67, 0x69,
	0x6e, 0x52, 0x65, 0x71, 0x1a, 0x11, 0x2e, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x43, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x22, 0x16, 0xd2, 0xc1, 0x18, 0x12, 0x2f, 0x76, 0x31,
	0x2f, 0x69, 0x6d, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x32, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x12,
	0x50, 0x0a, 0x0d, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72,
	0x12, 0x0e, 0x2e, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x71,
	0x1a, 0x11, 0x2e, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x22, 0x1c, 0xd2, 0xc1, 0x18, 0x18, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6d, 0x2f,
	0x61, 0x75, 0x74, 0x68, 0x32, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2f, 0x3a, 0x74, 0x79, 0x70,
	0x65, 0x12, 0x44, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12,
	0x0c, 0x2e, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x4e, 0x6f, 0x52, 0x65, 0x71, 0x1a, 0x11, 0x2e,
	0x6f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x22, 0x15, 0xca, 0xc1, 0x18, 0x11, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6d, 0x2f, 0x6d, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x46, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x0b, 0x2e, 0x6f, 0x61, 0x75, 0x74, 0x68,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x1a, 0x11, 0x2e, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x43, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x22, 0x15, 0xda, 0xc1, 0x18, 0x11, 0x2f, 0x76,
	0x31, 0x2f, 0x69, 0x6d, 0x2f, 0x6d, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12,
	0x42, 0x0a, 0x06, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x12, 0x0c, 0x2e, 0x6f, 0x61, 0x75, 0x74,
	0x68, 0x2e, 0x4e, 0x6f, 0x52, 0x65, 0x71, 0x1a, 0x11, 0x2e, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x2e,
	0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x22, 0x17, 0xd2, 0xc1, 0x18, 0x13,
	0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6d, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x32, 0x2f, 0x6c, 0x6f, 0x67,
	0x6f, 0x75, 0x74, 0x32, 0xa1, 0x01, 0x0a, 0x05, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x12, 0x49, 0x0a,
	0x08, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x0f, 0x2e, 0x6f, 0x61, 0x75, 0x74,
	0x68, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x11, 0x2e, 0x6f, 0x61, 0x75,
	0x74, 0x68, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x22, 0x19, 0xd2,
	0xc1, 0x18, 0x15, 0x2f, 0x76, 0x31, 0x2f, 0x68, 0x65, 0x72, 0x74, 0x7a, 0x2f, 0x61, 0x75, 0x74,
	0x68, 0x32, 0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x4d, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x4a,
	0x57, 0x4b, 0x73, 0x12, 0x0c, 0x2e, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x4e, 0x6f, 0x52, 0x65,
	0x71, 0x1a, 0x0f, 0x2e, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x4a, 0x57, 0x4b, 0x53, 0x52, 0x65,
	0x73, 0x70, 0x22, 0x23, 0xca, 0xc1, 0x18, 0x1f, 0x2f, 0x76, 0x31, 0x2f, 0x68, 0x65, 0x72, 0x74,
	0x7a, 0x2f, 0x2e, 0x77, 0x65, 0x6c, 0x6c, 0x2d, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x2f, 0x6a, 0x77,
	0x6b, 0x73, 0x2e, 0x6a, 0x73, 0x6f, 0x6e, 0x42, 0x22, 0x5a, 0x20, 0x68, 0x65, 0x72, 0x74, 0x7a,
	0x2f, 0x64, 0x65, 0x6d, 0x6f, 0x2f, 0x62, 0x69, 0x7a, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f,
	0x68, 0x65, 0x72, 0x74, 0x7a, 0x2f, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_oauth_proto_rawDescOnce sync.Once
	file_oauth_proto_rawDescData = file_oauth_proto_rawDesc
)

func file_oauth_proto_rawDescGZIP() []byte {
	file_oauth_proto_rawDescOnce.Do(func() {
		file_oauth_proto_rawDescData = protoimpl.X.CompressGZIP(file_oauth_proto_rawDescData)
	})
	return file_oauth_proto_rawDescData
}

var file_oauth_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_oauth_proto_goTypes = []interface{}{
	(*NoReq)(nil),      // 0: oauth.NoReq
	(*CommonResp)(nil), // 1: oauth.CommonResp
	(*LoginReq)(nil),   // 2: oauth.LoginReq
	(*TypeReq)(nil),    // 3: oauth.TypeReq
	(*User)(nil),       // 4: oauth.User
	(*TokenReq)(nil),   // 5: oauth.TokenReq
	(*JWKSKey)(nil),    // 6: oauth.JWKSKey
	(*JWKSResp)(nil),   // 7: oauth.JWKSResp
	(*anypb.Any)(nil),  // 8: google.protobuf.Any
}
var file_oauth_proto_depIdxs = []int32{
	8, // 0: oauth.CommonResp.data:type_name -> google.protobuf.Any
	6, // 1: oauth.JWKSResp.keys:type_name -> oauth.JWKSKey
	2, // 2: oauth.Login.LoginPassword:input_type -> oauth.LoginReq
	3, // 3: oauth.Login.LoginProvider:input_type -> oauth.TypeReq
	0, // 4: oauth.Login.GetProfile:input_type -> oauth.NoReq
	4, // 5: oauth.Login.UpdateProfile:input_type -> oauth.User
	0, // 6: oauth.Login.Logout:input_type -> oauth.NoReq
	5, // 7: oauth.OAuth.GetToken:input_type -> oauth.TokenReq
	0, // 8: oauth.OAuth.GetJWKs:input_type -> oauth.NoReq
	1, // 9: oauth.Login.LoginPassword:output_type -> oauth.CommonResp
	1, // 10: oauth.Login.LoginProvider:output_type -> oauth.CommonResp
	1, // 11: oauth.Login.GetProfile:output_type -> oauth.CommonResp
	1, // 12: oauth.Login.UpdateProfile:output_type -> oauth.CommonResp
	1, // 13: oauth.Login.Logout:output_type -> oauth.CommonResp
	1, // 14: oauth.OAuth.GetToken:output_type -> oauth.CommonResp
	7, // 15: oauth.OAuth.GetJWKs:output_type -> oauth.JWKSResp
	9, // [9:16] is the sub-list for method output_type
	2, // [2:9] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_oauth_proto_init() }
func file_oauth_proto_init() {
	if File_oauth_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_oauth_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NoReq); i {
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
		file_oauth_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommonResp); i {
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
		file_oauth_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginReq); i {
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
		file_oauth_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TypeReq); i {
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
		file_oauth_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_oauth_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TokenReq); i {
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
		file_oauth_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JWKSKey); i {
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
		file_oauth_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JWKSResp); i {
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
			RawDescriptor: file_oauth_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_oauth_proto_goTypes,
		DependencyIndexes: file_oauth_proto_depIdxs,
		MessageInfos:      file_oauth_proto_msgTypes,
	}.Build()
	File_oauth_proto = out.File
	file_oauth_proto_rawDesc = nil
	file_oauth_proto_goTypes = nil
	file_oauth_proto_depIdxs = nil
}