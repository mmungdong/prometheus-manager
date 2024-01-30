// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: ping/ping.proto

package ping

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// 健康检查请求
type PingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 请求者名称, 非必填
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *PingRequest) Reset() {
	*x = PingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ping_ping_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingRequest) ProtoMessage() {}

func (x *PingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ping_ping_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingRequest.ProtoReflect.Descriptor instead.
func (*PingRequest) Descriptor() ([]byte, []int) {
	return file_ping_ping_proto_rawDescGZIP(), []int{0}
}

func (x *PingRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// 健康检查响应
type PingReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 响应者名称
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// 版本号
	Version string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	// 命名空间
	Namespace string `protobuf:"bytes,3,opt,name=namespace,proto3" json:"namespace,omitempty"`
	// 元数据
	Metadata map[string]string `protobuf:"bytes,4,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *PingReply) Reset() {
	*x = PingReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ping_ping_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingReply) ProtoMessage() {}

func (x *PingReply) ProtoReflect() protoreflect.Message {
	mi := &file_ping_ping_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingReply.ProtoReflect.Descriptor instead.
func (*PingReply) Descriptor() ([]byte, []int) {
	return file_ping_ping_proto_rawDescGZIP(), []int{1}
}

func (x *PingReply) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PingReply) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *PingReply) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *PingReply) GetMetadata() map[string]string {
	if x != nil {
		return x.Metadata
	}
	return nil
}

var File_ping_ping_proto protoreflect.FileDescriptor

var file_ping_ping_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x70, 0x69, 0x6e, 0x67, 0x2f, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x03, 0x61, 0x70, 0x69, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x21, 0x0a, 0x0b, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0xce, 0x01, 0x0a, 0x09, 0x50, 0x69, 0x6e, 0x67,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x12, 0x38, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x3b, 0x0a, 0x0d, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x32, 0x40, 0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67,
	0x12, 0x38, 0x0a, 0x05, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x10, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x0d, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x07, 0x12, 0x05, 0x2f, 0x70, 0x69, 0x6e, 0x67, 0x42, 0x2a, 0x0a, 0x04, 0x70, 0x69,
	0x6e, 0x67, 0x50, 0x01, 0x5a, 0x20, 0x70, 0x72, 0x6f, 0x6d, 0x65, 0x74, 0x68, 0x65, 0x75, 0x73,
	0x2d, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x69, 0x6e,
	0x67, 0x3b, 0x70, 0x69, 0x6e, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ping_ping_proto_rawDescOnce sync.Once
	file_ping_ping_proto_rawDescData = file_ping_ping_proto_rawDesc
)

func file_ping_ping_proto_rawDescGZIP() []byte {
	file_ping_ping_proto_rawDescOnce.Do(func() {
		file_ping_ping_proto_rawDescData = protoimpl.X.CompressGZIP(file_ping_ping_proto_rawDescData)
	})
	return file_ping_ping_proto_rawDescData
}

var file_ping_ping_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_ping_ping_proto_goTypes = []interface{}{
	(*PingRequest)(nil), // 0: api.PingRequest
	(*PingReply)(nil),   // 1: api.PingReply
	nil,                 // 2: api.PingReply.MetadataEntry
}
var file_ping_ping_proto_depIdxs = []int32{
	2, // 0: api.PingReply.metadata:type_name -> api.PingReply.MetadataEntry
	0, // 1: api.Ping.Check:input_type -> api.PingRequest
	1, // 2: api.Ping.Check:output_type -> api.PingReply
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_ping_ping_proto_init() }
func file_ping_ping_proto_init() {
	if File_ping_ping_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ping_ping_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingRequest); i {
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
		file_ping_ping_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingReply); i {
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
			RawDescriptor: file_ping_ping_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ping_ping_proto_goTypes,
		DependencyIndexes: file_ping_ping_proto_depIdxs,
		MessageInfos:      file_ping_ping_proto_msgTypes,
	}.Build()
	File_ping_ping_proto = out.File
	file_ping_ping_proto_rawDesc = nil
	file_ping_ping_proto_goTypes = nil
	file_ping_ping_proto_depIdxs = nil
}
