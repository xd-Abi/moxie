// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.23.0--rc1
// source: pkg/proto/health/health.proto

package health

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

type Record struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message   string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Timestamp int64  `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Healthy   bool   `protobuf:"varint,3,opt,name=healthy,proto3" json:"healthy,omitempty"`
}

func (x *Record) Reset() {
	*x = Record{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_health_health_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Record) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Record) ProtoMessage() {}

func (x *Record) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_health_health_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Record.ProtoReflect.Descriptor instead.
func (*Record) Descriptor() ([]byte, []int) {
	return file_pkg_proto_health_health_proto_rawDescGZIP(), []int{0}
}

func (x *Record) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Record) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *Record) GetHealthy() bool {
	if x != nil {
		return x.Healthy
	}
	return false
}

type Checks struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Auth    *Record `protobuf:"bytes,1,opt,name=auth,proto3" json:"auth,omitempty"`
	Profile *Record `protobuf:"bytes,2,opt,name=profile,proto3" json:"profile,omitempty"`
}

func (x *Checks) Reset() {
	*x = Checks{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_health_health_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Checks) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Checks) ProtoMessage() {}

func (x *Checks) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_health_health_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Checks.ProtoReflect.Descriptor instead.
func (*Checks) Descriptor() ([]byte, []int) {
	return file_pkg_proto_health_health_proto_rawDescGZIP(), []int{1}
}

func (x *Checks) GetAuth() *Record {
	if x != nil {
		return x.Auth
	}
	return nil
}

func (x *Checks) GetProfile() *Record {
	if x != nil {
		return x.Profile
	}
	return nil
}

type CheckHealthRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CheckHealthRequest) Reset() {
	*x = CheckHealthRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_health_health_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckHealthRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckHealthRequest) ProtoMessage() {}

func (x *CheckHealthRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_health_health_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckHealthRequest.ProtoReflect.Descriptor instead.
func (*CheckHealthRequest) Descriptor() ([]byte, []int) {
	return file_pkg_proto_health_health_proto_rawDescGZIP(), []int{2}
}

type CheckHealthResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Checks    *Checks `protobuf:"bytes,1,opt,name=checks,proto3" json:"checks,omitempty"`
	Timestamp int64   `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *CheckHealthResponse) Reset() {
	*x = CheckHealthResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_health_health_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckHealthResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckHealthResponse) ProtoMessage() {}

func (x *CheckHealthResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_health_health_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckHealthResponse.ProtoReflect.Descriptor instead.
func (*CheckHealthResponse) Descriptor() ([]byte, []int) {
	return file_pkg_proto_health_health_proto_rawDescGZIP(), []int{3}
}

func (x *CheckHealthResponse) GetChecks() *Checks {
	if x != nil {
		return x.Checks
	}
	return nil
}

func (x *CheckHealthResponse) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

var File_pkg_proto_health_health_proto protoreflect.FileDescriptor

var file_pkg_proto_health_health_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x68, 0x65, 0x61, 0x6c,
	0x74, 0x68, 0x2f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x06, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x22, 0x5a, 0x0a, 0x06, 0x52, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x68, 0x65, 0x61,
	0x6c, 0x74, 0x68, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x68, 0x65, 0x61, 0x6c,
	0x74, 0x68, 0x79, 0x22, 0x56, 0x0a, 0x06, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x73, 0x12, 0x22, 0x0a,
	0x04, 0x61, 0x75, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x68, 0x65,
	0x61, 0x6c, 0x74, 0x68, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x04, 0x61, 0x75, 0x74,
	0x68, 0x12, 0x28, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x2e, 0x52, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x22, 0x14, 0x0a, 0x12, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x22, 0x5b, 0x0a, 0x13, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x06, 0x63, 0x68, 0x65, 0x63,
	0x6b, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x68, 0x65, 0x61, 0x6c, 0x74,
	0x68, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x73, 0x52, 0x06, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x73,
	0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x32, 0x59,
	0x0a, 0x0d, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x48, 0x0a, 0x0b, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x12, 0x1a,
	0x2e, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x48, 0x65, 0x61,
	0x6c, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x68, 0x65, 0x61,
	0x6c, 0x74, 0x68, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x2a, 0x5a, 0x28, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x78, 0x64, 0x2d, 0x41, 0x62, 0x69, 0x2f, 0x6d,
	0x6f, 0x78, 0x69, 0x65, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x68,
	0x65, 0x61, 0x6c, 0x74, 0x68, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_proto_health_health_proto_rawDescOnce sync.Once
	file_pkg_proto_health_health_proto_rawDescData = file_pkg_proto_health_health_proto_rawDesc
)

func file_pkg_proto_health_health_proto_rawDescGZIP() []byte {
	file_pkg_proto_health_health_proto_rawDescOnce.Do(func() {
		file_pkg_proto_health_health_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_proto_health_health_proto_rawDescData)
	})
	return file_pkg_proto_health_health_proto_rawDescData
}

var file_pkg_proto_health_health_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_pkg_proto_health_health_proto_goTypes = []interface{}{
	(*Record)(nil),              // 0: health.Record
	(*Checks)(nil),              // 1: health.Checks
	(*CheckHealthRequest)(nil),  // 2: health.CheckHealthRequest
	(*CheckHealthResponse)(nil), // 3: health.CheckHealthResponse
}
var file_pkg_proto_health_health_proto_depIdxs = []int32{
	0, // 0: health.Checks.auth:type_name -> health.Record
	0, // 1: health.Checks.profile:type_name -> health.Record
	1, // 2: health.CheckHealthResponse.checks:type_name -> health.Checks
	2, // 3: health.HealthService.CheckHealth:input_type -> health.CheckHealthRequest
	3, // 4: health.HealthService.CheckHealth:output_type -> health.CheckHealthResponse
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_pkg_proto_health_health_proto_init() }
func file_pkg_proto_health_health_proto_init() {
	if File_pkg_proto_health_health_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_proto_health_health_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Record); i {
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
		file_pkg_proto_health_health_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Checks); i {
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
		file_pkg_proto_health_health_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckHealthRequest); i {
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
		file_pkg_proto_health_health_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckHealthResponse); i {
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
			RawDescriptor: file_pkg_proto_health_health_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_proto_health_health_proto_goTypes,
		DependencyIndexes: file_pkg_proto_health_health_proto_depIdxs,
		MessageInfos:      file_pkg_proto_health_health_proto_msgTypes,
	}.Build()
	File_pkg_proto_health_health_proto = out.File
	file_pkg_proto_health_health_proto_rawDesc = nil
	file_pkg_proto_health_health_proto_goTypes = nil
	file_pkg_proto_health_health_proto_depIdxs = nil
}
