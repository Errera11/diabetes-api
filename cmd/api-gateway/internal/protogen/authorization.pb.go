// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: authorization.proto

package diabetesproto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type SignupRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Username      string                 `protobuf:"bytes,1,opt,name=Username,proto3" json:"Username,omitempty"`
	Password      string                 `protobuf:"bytes,2,opt,name=Password,proto3" json:"Password,omitempty"`
	Email         string                 `protobuf:"bytes,3,opt,name=Email,proto3" json:"Email,omitempty"`
	Image         *string                `protobuf:"bytes,4,opt,name=Image,proto3,oneof" json:"Image,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SignupRequest) Reset() {
	*x = SignupRequest{}
	mi := &file_authorization_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SignupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignupRequest) ProtoMessage() {}

func (x *SignupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_authorization_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignupRequest.ProtoReflect.Descriptor instead.
func (*SignupRequest) Descriptor() ([]byte, []int) {
	return file_authorization_proto_rawDescGZIP(), []int{0}
}

func (x *SignupRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *SignupRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *SignupRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *SignupRequest) GetImage() string {
	if x != nil && x.Image != nil {
		return *x.Image
	}
	return ""
}

type SigninRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Email         string                 `protobuf:"bytes,1,opt,name=Email,proto3" json:"Email,omitempty"`
	Password      string                 `protobuf:"bytes,2,opt,name=Password,proto3" json:"Password,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SigninRequest) Reset() {
	*x = SigninRequest{}
	mi := &file_authorization_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SigninRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SigninRequest) ProtoMessage() {}

func (x *SigninRequest) ProtoReflect() protoreflect.Message {
	mi := &file_authorization_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SigninRequest.ProtoReflect.Descriptor instead.
func (*SigninRequest) Descriptor() ([]byte, []int) {
	return file_authorization_proto_rawDescGZIP(), []int{1}
}

func (x *SigninRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *SigninRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type SigninResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        int32                  `protobuf:"varint,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
	Token         string                 `protobuf:"bytes,2,opt,name=Token,proto3" json:"Token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SigninResponse) Reset() {
	*x = SigninResponse{}
	mi := &file_authorization_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SigninResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SigninResponse) ProtoMessage() {}

func (x *SigninResponse) ProtoReflect() protoreflect.Message {
	mi := &file_authorization_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SigninResponse.ProtoReflect.Descriptor instead.
func (*SigninResponse) Descriptor() ([]byte, []int) {
	return file_authorization_proto_rawDescGZIP(), []int{2}
}

func (x *SigninResponse) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *SigninResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type SignupResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        int32                  `protobuf:"varint,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
	Token         string                 `protobuf:"bytes,2,opt,name=Token,proto3" json:"Token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SignupResponse) Reset() {
	*x = SignupResponse{}
	mi := &file_authorization_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SignupResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignupResponse) ProtoMessage() {}

func (x *SignupResponse) ProtoReflect() protoreflect.Message {
	mi := &file_authorization_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignupResponse.ProtoReflect.Descriptor instead.
func (*SignupResponse) Descriptor() ([]byte, []int) {
	return file_authorization_proto_rawDescGZIP(), []int{3}
}

func (x *SignupResponse) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *SignupResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type LogoutRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Token         string                 `protobuf:"bytes,1,opt,name=Token,proto3" json:"Token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LogoutRequest) Reset() {
	*x = LogoutRequest{}
	mi := &file_authorization_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LogoutRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogoutRequest) ProtoMessage() {}

func (x *LogoutRequest) ProtoReflect() protoreflect.Message {
	mi := &file_authorization_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogoutRequest.ProtoReflect.Descriptor instead.
func (*LogoutRequest) Descriptor() ([]byte, []int) {
	return file_authorization_proto_rawDescGZIP(), []int{4}
}

func (x *LogoutRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type LogoutResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Message       string                 `protobuf:"bytes,1,opt,name=Message,proto3" json:"Message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LogoutResponse) Reset() {
	*x = LogoutResponse{}
	mi := &file_authorization_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LogoutResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogoutResponse) ProtoMessage() {}

func (x *LogoutResponse) ProtoReflect() protoreflect.Message {
	mi := &file_authorization_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogoutResponse.ProtoReflect.Descriptor instead.
func (*LogoutResponse) Descriptor() ([]byte, []int) {
	return file_authorization_proto_rawDescGZIP(), []int{5}
}

func (x *LogoutResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type AuthRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Token         string                 `protobuf:"bytes,1,opt,name=Token,proto3" json:"Token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AuthRequest) Reset() {
	*x = AuthRequest{}
	mi := &file_authorization_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AuthRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthRequest) ProtoMessage() {}

func (x *AuthRequest) ProtoReflect() protoreflect.Message {
	mi := &file_authorization_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthRequest.ProtoReflect.Descriptor instead.
func (*AuthRequest) Descriptor() ([]byte, []int) {
	return file_authorization_proto_rawDescGZIP(), []int{6}
}

func (x *AuthRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type AuthResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int32                  `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Username      string                 `protobuf:"bytes,2,opt,name=Username,proto3" json:"Username,omitempty"`
	Email         string                 `protobuf:"bytes,4,opt,name=Email,proto3" json:"Email,omitempty"`
	Image         string                 `protobuf:"bytes,5,opt,name=Image,proto3" json:"Image,omitempty"`
	CreatedAt     string                 `protobuf:"bytes,6,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AuthResponse) Reset() {
	*x = AuthResponse{}
	mi := &file_authorization_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AuthResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthResponse) ProtoMessage() {}

func (x *AuthResponse) ProtoReflect() protoreflect.Message {
	mi := &file_authorization_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthResponse.ProtoReflect.Descriptor instead.
func (*AuthResponse) Descriptor() ([]byte, []int) {
	return file_authorization_proto_rawDescGZIP(), []int{7}
}

func (x *AuthResponse) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *AuthResponse) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *AuthResponse) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *AuthResponse) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *AuthResponse) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

var File_authorization_proto protoreflect.FileDescriptor

const file_authorization_proto_rawDesc = "" +
	"\n" +
	"\x13authorization.proto\x12\rdiabetesproto\"\x82\x01\n" +
	"\rSignupRequest\x12\x1a\n" +
	"\bUsername\x18\x01 \x01(\tR\bUsername\x12\x1a\n" +
	"\bPassword\x18\x02 \x01(\tR\bPassword\x12\x14\n" +
	"\x05Email\x18\x03 \x01(\tR\x05Email\x12\x19\n" +
	"\x05Image\x18\x04 \x01(\tH\x00R\x05Image\x88\x01\x01B\b\n" +
	"\x06_Image\"A\n" +
	"\rSigninRequest\x12\x14\n" +
	"\x05Email\x18\x01 \x01(\tR\x05Email\x12\x1a\n" +
	"\bPassword\x18\x02 \x01(\tR\bPassword\">\n" +
	"\x0eSigninResponse\x12\x16\n" +
	"\x06UserId\x18\x01 \x01(\x05R\x06UserId\x12\x14\n" +
	"\x05Token\x18\x02 \x01(\tR\x05Token\">\n" +
	"\x0eSignupResponse\x12\x16\n" +
	"\x06UserId\x18\x01 \x01(\x05R\x06UserId\x12\x14\n" +
	"\x05Token\x18\x02 \x01(\tR\x05Token\"%\n" +
	"\rLogoutRequest\x12\x14\n" +
	"\x05Token\x18\x01 \x01(\tR\x05Token\"*\n" +
	"\x0eLogoutResponse\x12\x18\n" +
	"\aMessage\x18\x01 \x01(\tR\aMessage\"#\n" +
	"\vAuthRequest\x12\x14\n" +
	"\x05Token\x18\x01 \x01(\tR\x05Token\"\x84\x01\n" +
	"\fAuthResponse\x12\x0e\n" +
	"\x02Id\x18\x01 \x01(\x05R\x02Id\x12\x1a\n" +
	"\bUsername\x18\x02 \x01(\tR\bUsername\x12\x14\n" +
	"\x05Email\x18\x04 \x01(\tR\x05Email\x12\x14\n" +
	"\x05Image\x18\x05 \x01(\tR\x05Image\x12\x1c\n" +
	"\tCreatedAt\x18\x06 \x01(\tR\tCreatedAt2\xab\x02\n" +
	"\vAuthService\x12G\n" +
	"\x06Signin\x12\x1c.diabetesproto.SigninRequest\x1a\x1d.diabetesproto.SigninResponse\"\x00\x12G\n" +
	"\x06Signup\x12\x1c.diabetesproto.SignupRequest\x1a\x1d.diabetesproto.SignupResponse\"\x00\x12G\n" +
	"\x06Logout\x12\x1c.diabetesproto.LogoutRequest\x1a\x1d.diabetesproto.LogoutResponse\"\x00\x12A\n" +
	"\x04Auth\x12\x1a.diabetesproto.AuthRequest\x1a\x1b.diabetesproto.AuthResponse\"\x00B#Z!github.com/Errera11/diabetesprotob\x06proto3"

var (
	file_authorization_proto_rawDescOnce sync.Once
	file_authorization_proto_rawDescData []byte
)

func file_authorization_proto_rawDescGZIP() []byte {
	file_authorization_proto_rawDescOnce.Do(func() {
		file_authorization_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_authorization_proto_rawDesc), len(file_authorization_proto_rawDesc)))
	})
	return file_authorization_proto_rawDescData
}

var file_authorization_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_authorization_proto_goTypes = []any{
	(*SignupRequest)(nil),  // 0: diabetesproto.SignupRequest
	(*SigninRequest)(nil),  // 1: diabetesproto.SigninRequest
	(*SigninResponse)(nil), // 2: diabetesproto.SigninResponse
	(*SignupResponse)(nil), // 3: diabetesproto.SignupResponse
	(*LogoutRequest)(nil),  // 4: diabetesproto.LogoutRequest
	(*LogoutResponse)(nil), // 5: diabetesproto.LogoutResponse
	(*AuthRequest)(nil),    // 6: diabetesproto.AuthRequest
	(*AuthResponse)(nil),   // 7: diabetesproto.AuthResponse
}
var file_authorization_proto_depIdxs = []int32{
	1, // 0: diabetesproto.AuthService.Signin:input_type -> diabetesproto.SigninRequest
	0, // 1: diabetesproto.AuthService.Signup:input_type -> diabetesproto.SignupRequest
	4, // 2: diabetesproto.AuthService.Logout:input_type -> diabetesproto.LogoutRequest
	6, // 3: diabetesproto.AuthService.Auth:input_type -> diabetesproto.AuthRequest
	2, // 4: diabetesproto.AuthService.Signin:output_type -> diabetesproto.SigninResponse
	3, // 5: diabetesproto.AuthService.Signup:output_type -> diabetesproto.SignupResponse
	5, // 6: diabetesproto.AuthService.Logout:output_type -> diabetesproto.LogoutResponse
	7, // 7: diabetesproto.AuthService.Auth:output_type -> diabetesproto.AuthResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_authorization_proto_init() }
func file_authorization_proto_init() {
	if File_authorization_proto != nil {
		return
	}
	file_authorization_proto_msgTypes[0].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_authorization_proto_rawDesc), len(file_authorization_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_authorization_proto_goTypes,
		DependencyIndexes: file_authorization_proto_depIdxs,
		MessageInfos:      file_authorization_proto_msgTypes,
	}.Build()
	File_authorization_proto = out.File
	file_authorization_proto_goTypes = nil
	file_authorization_proto_depIdxs = nil
}
