// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.23.0--rc1
// source: pkg/proto/reset-password/reset-password.proto

package reset_password

import (
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

type ResetPasswordRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *ResetPasswordRequest) Reset() {
	*x = ResetPasswordRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_reset_password_reset_password_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResetPasswordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResetPasswordRequest) ProtoMessage() {}

func (x *ResetPasswordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_reset_password_reset_password_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResetPasswordRequest.ProtoReflect.Descriptor instead.
func (*ResetPasswordRequest) Descriptor() ([]byte, []int) {
	return file_pkg_proto_reset_password_reset_password_proto_rawDescGZIP(), []int{0}
}

func (x *ResetPasswordRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type ResetPasswordResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PasswordReset bool `protobuf:"varint,1,opt,name=password_reset,json=passwordReset,proto3" json:"password_reset,omitempty"`
}

func (x *ResetPasswordResponse) Reset() {
	*x = ResetPasswordResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_reset_password_reset_password_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResetPasswordResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResetPasswordResponse) ProtoMessage() {}

func (x *ResetPasswordResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_reset_password_reset_password_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResetPasswordResponse.ProtoReflect.Descriptor instead.
func (*ResetPasswordResponse) Descriptor() ([]byte, []int) {
	return file_pkg_proto_reset_password_reset_password_proto_rawDescGZIP(), []int{1}
}

func (x *ResetPasswordResponse) GetPasswordReset() bool {
	if x != nil {
		return x.PasswordReset
	}
	return false
}

var File_pkg_proto_reset_password_reset_password_proto protoreflect.FileDescriptor

var file_pkg_proto_reset_password_reset_password_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x65, 0x73, 0x65,
	0x74, 0x2d, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x2f, 0x72, 0x65, 0x73, 0x65, 0x74,
	0x2d, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0d, 0x72, 0x65, 0x73, 0x65, 0x74, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x2c,
	0x0a, 0x14, 0x52, 0x65, 0x73, 0x65, 0x74, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x3e, 0x0a, 0x15,
	0x52, 0x65, 0x73, 0x65, 0x74, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x5f, 0x72, 0x65, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x70,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x65, 0x74, 0x32, 0x72, 0x0a, 0x14,
	0x52, 0x65, 0x73, 0x65, 0x74, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x5a, 0x0a, 0x0d, 0x52, 0x65, 0x73, 0x65, 0x74, 0x50, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x23, 0x2e, 0x72, 0x65, 0x73, 0x65, 0x74, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x2e, 0x52, 0x65, 0x73, 0x65, 0x74, 0x50, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x72, 0x65, 0x73,
	0x65, 0x74, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x2e, 0x52, 0x65, 0x73, 0x65, 0x74,
	0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x35, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x78,
	0x64, 0x2d, 0x41, 0x62, 0x69, 0x2f, 0x6d, 0x6f, 0x78, 0x69, 0x65, 0x2d, 0x67, 0x6f, 0x2f, 0x70,
	0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x65, 0x73, 0x65, 0x74, 0x2d, 0x70,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_proto_reset_password_reset_password_proto_rawDescOnce sync.Once
	file_pkg_proto_reset_password_reset_password_proto_rawDescData = file_pkg_proto_reset_password_reset_password_proto_rawDesc
)

func file_pkg_proto_reset_password_reset_password_proto_rawDescGZIP() []byte {
	file_pkg_proto_reset_password_reset_password_proto_rawDescOnce.Do(func() {
		file_pkg_proto_reset_password_reset_password_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_proto_reset_password_reset_password_proto_rawDescData)
	})
	return file_pkg_proto_reset_password_reset_password_proto_rawDescData
}

var file_pkg_proto_reset_password_reset_password_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pkg_proto_reset_password_reset_password_proto_goTypes = []interface{}{
	(*ResetPasswordRequest)(nil),  // 0: resetpassword.ResetPasswordRequest
	(*ResetPasswordResponse)(nil), // 1: resetpassword.ResetPasswordResponse
}
var file_pkg_proto_reset_password_reset_password_proto_depIdxs = []int32{
	0, // 0: resetpassword.ResetPasswordService.ResetPassword:input_type -> resetpassword.ResetPasswordRequest
	1, // 1: resetpassword.ResetPasswordService.ResetPassword:output_type -> resetpassword.ResetPasswordResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pkg_proto_reset_password_reset_password_proto_init() }
func file_pkg_proto_reset_password_reset_password_proto_init() {
	if File_pkg_proto_reset_password_reset_password_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_proto_reset_password_reset_password_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResetPasswordRequest); i {
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
		file_pkg_proto_reset_password_reset_password_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResetPasswordResponse); i {
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
			RawDescriptor: file_pkg_proto_reset_password_reset_password_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_proto_reset_password_reset_password_proto_goTypes,
		DependencyIndexes: file_pkg_proto_reset_password_reset_password_proto_depIdxs,
		MessageInfos:      file_pkg_proto_reset_password_reset_password_proto_msgTypes,
	}.Build()
	File_pkg_proto_reset_password_reset_password_proto = out.File
	file_pkg_proto_reset_password_reset_password_proto_rawDesc = nil
	file_pkg_proto_reset_password_reset_password_proto_goTypes = nil
	file_pkg_proto_reset_password_reset_password_proto_depIdxs = nil
}
