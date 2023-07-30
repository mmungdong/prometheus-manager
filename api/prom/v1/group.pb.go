// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: prom/v1/group.proto

package v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	api "prometheus-manager/api"
	prom "prometheus-manager/api/prom"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Group *prom.GroupItem `protobuf:"bytes,1,opt,name=group,proto3" json:"group,omitempty"`
}

func (x *CreateGroupRequest) Reset() {
	*x = CreateGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prom_v1_group_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateGroupRequest) ProtoMessage() {}

func (x *CreateGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_prom_v1_group_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateGroupRequest.ProtoReflect.Descriptor instead.
func (*CreateGroupRequest) Descriptor() ([]byte, []int) {
	return file_prom_v1_group_proto_rawDescGZIP(), []int{0}
}

func (x *CreateGroupRequest) GetGroup() *prom.GroupItem {
	if x != nil {
		return x.Group
	}
	return nil
}

type CreateGroupReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response *api.Response `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *CreateGroupReply) Reset() {
	*x = CreateGroupReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prom_v1_group_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateGroupReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateGroupReply) ProtoMessage() {}

func (x *CreateGroupReply) ProtoReflect() protoreflect.Message {
	mi := &file_prom_v1_group_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateGroupReply.ProtoReflect.Descriptor instead.
func (*CreateGroupReply) Descriptor() ([]byte, []int) {
	return file_prom_v1_group_proto_rawDescGZIP(), []int{1}
}

func (x *CreateGroupReply) GetResponse() *api.Response {
	if x != nil {
		return x.Response
	}
	return nil
}

type UpdateGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    uint32          `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Group *prom.GroupItem `protobuf:"bytes,2,opt,name=group,proto3" json:"group,omitempty"`
}

func (x *UpdateGroupRequest) Reset() {
	*x = UpdateGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prom_v1_group_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateGroupRequest) ProtoMessage() {}

func (x *UpdateGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_prom_v1_group_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateGroupRequest.ProtoReflect.Descriptor instead.
func (*UpdateGroupRequest) Descriptor() ([]byte, []int) {
	return file_prom_v1_group_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateGroupRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateGroupRequest) GetGroup() *prom.GroupItem {
	if x != nil {
		return x.Group
	}
	return nil
}

type UpdateGroupReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response *api.Response `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *UpdateGroupReply) Reset() {
	*x = UpdateGroupReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prom_v1_group_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateGroupReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateGroupReply) ProtoMessage() {}

func (x *UpdateGroupReply) ProtoReflect() protoreflect.Message {
	mi := &file_prom_v1_group_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateGroupReply.ProtoReflect.Descriptor instead.
func (*UpdateGroupReply) Descriptor() ([]byte, []int) {
	return file_prom_v1_group_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateGroupReply) GetResponse() *api.Response {
	if x != nil {
		return x.Response
	}
	return nil
}

type DeleteGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteGroupRequest) Reset() {
	*x = DeleteGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prom_v1_group_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteGroupRequest) ProtoMessage() {}

func (x *DeleteGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_prom_v1_group_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteGroupRequest.ProtoReflect.Descriptor instead.
func (*DeleteGroupRequest) Descriptor() ([]byte, []int) {
	return file_prom_v1_group_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteGroupRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteGroupReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response *api.Response `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *DeleteGroupReply) Reset() {
	*x = DeleteGroupReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prom_v1_group_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteGroupReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteGroupReply) ProtoMessage() {}

func (x *DeleteGroupReply) ProtoReflect() protoreflect.Message {
	mi := &file_prom_v1_group_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteGroupReply.ProtoReflect.Descriptor instead.
func (*DeleteGroupReply) Descriptor() ([]byte, []int) {
	return file_prom_v1_group_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteGroupReply) GetResponse() *api.Response {
	if x != nil {
		return x.Response
	}
	return nil
}

type GetGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetGroupRequest) Reset() {
	*x = GetGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prom_v1_group_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGroupRequest) ProtoMessage() {}

func (x *GetGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_prom_v1_group_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGroupRequest.ProtoReflect.Descriptor instead.
func (*GetGroupRequest) Descriptor() ([]byte, []int) {
	return file_prom_v1_group_proto_rawDescGZIP(), []int{6}
}

func (x *GetGroupRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetGroupReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response *api.Response   `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
	Group    *prom.GroupItem `protobuf:"bytes,2,opt,name=group,proto3" json:"group,omitempty"`
}

func (x *GetGroupReply) Reset() {
	*x = GetGroupReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prom_v1_group_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGroupReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGroupReply) ProtoMessage() {}

func (x *GetGroupReply) ProtoReflect() protoreflect.Message {
	mi := &file_prom_v1_group_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGroupReply.ProtoReflect.Descriptor instead.
func (*GetGroupReply) Descriptor() ([]byte, []int) {
	return file_prom_v1_group_proto_rawDescGZIP(), []int{7}
}

func (x *GetGroupReply) GetResponse() *api.Response {
	if x != nil {
		return x.Response
	}
	return nil
}

func (x *GetGroupReply) GetGroup() *prom.GroupItem {
	if x != nil {
		return x.Group
	}
	return nil
}

type ListGroupRequestParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Keyword string `protobuf:"bytes,1,opt,name=keyword,proto3" json:"keyword,omitempty"`
}

func (x *ListGroupRequestParams) Reset() {
	*x = ListGroupRequestParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prom_v1_group_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListGroupRequestParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListGroupRequestParams) ProtoMessage() {}

func (x *ListGroupRequestParams) ProtoReflect() protoreflect.Message {
	mi := &file_prom_v1_group_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListGroupRequestParams.ProtoReflect.Descriptor instead.
func (*ListGroupRequestParams) Descriptor() ([]byte, []int) {
	return file_prom_v1_group_proto_rawDescGZIP(), []int{8}
}

func (x *ListGroupRequestParams) GetKeyword() string {
	if x != nil {
		return x.Keyword
	}
	return ""
}

type ListGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page   *api.PageRequest        `protobuf:"bytes,1,opt,name=page,proto3" json:"page,omitempty"`
	Params *ListGroupRequestParams `protobuf:"bytes,2,opt,name=params,proto3" json:"params,omitempty"`
}

func (x *ListGroupRequest) Reset() {
	*x = ListGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prom_v1_group_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListGroupRequest) ProtoMessage() {}

func (x *ListGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_prom_v1_group_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListGroupRequest.ProtoReflect.Descriptor instead.
func (*ListGroupRequest) Descriptor() ([]byte, []int) {
	return file_prom_v1_group_proto_rawDescGZIP(), []int{9}
}

func (x *ListGroupRequest) GetPage() *api.PageRequest {
	if x != nil {
		return x.Page
	}
	return nil
}

func (x *ListGroupRequest) GetParams() *ListGroupRequestParams {
	if x != nil {
		return x.Params
	}
	return nil
}

type ListGroupReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response *api.Response     `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
	Page     *api.PageReply    `protobuf:"bytes,2,opt,name=page,proto3" json:"page,omitempty"`
	List     []*prom.GroupItem `protobuf:"bytes,3,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *ListGroupReply) Reset() {
	*x = ListGroupReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prom_v1_group_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListGroupReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListGroupReply) ProtoMessage() {}

func (x *ListGroupReply) ProtoReflect() protoreflect.Message {
	mi := &file_prom_v1_group_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListGroupReply.ProtoReflect.Descriptor instead.
func (*ListGroupReply) Descriptor() ([]byte, []int) {
	return file_prom_v1_group_proto_rawDescGZIP(), []int{10}
}

func (x *ListGroupReply) GetResponse() *api.Response {
	if x != nil {
		return x.Response
	}
	return nil
}

func (x *ListGroupReply) GetPage() *api.PageReply {
	if x != nil {
		return x.Page
	}
	return nil
}

func (x *ListGroupReply) GetList() []*prom.GroupItem {
	if x != nil {
		return x.List
	}
	return nil
}

var File_prom_v1_group_proto protoreflect.FileDescriptor

var file_prom_v1_group_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x6d, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0a, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x0f, 0x70, 0x72, 0x6f, 0x6d, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x49, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x33, 0x0a, 0x05, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70,
	0x72, 0x6f, 0x6d, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x74, 0x65, 0x6d, 0x42, 0x08, 0xfa,
	0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x22, 0x3d,
	0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x29, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x62, 0x0a,
	0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x42,
	0x07, 0xfa, 0x42, 0x04, 0x2a, 0x02, 0x20, 0x00, 0x52, 0x02, 0x69, 0x64, 0x12, 0x33, 0x0a, 0x05,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x70, 0x72, 0x6f, 0x6d, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x74, 0x65, 0x6d,
	0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x05, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x22, 0x3d, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x29, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x2d, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x2a, 0x02, 0x20, 0x00, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x3d, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x29, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2a,
	0x0a, 0x0f, 0x47, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x17, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x07, 0xfa,
	0x42, 0x04, 0x2a, 0x02, 0x20, 0x00, 0x52, 0x02, 0x69, 0x64, 0x22, 0x65, 0x0a, 0x0d, 0x47, 0x65,
	0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x29, 0x0a, 0x08, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x08, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x6d,
	0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x22, 0x3b, 0x0a, 0x16, 0x4c, 0x69, 0x73, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x21, 0x0a, 0x07, 0x6b,
	0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42,
	0x04, 0x72, 0x02, 0x18, 0x40, 0x52, 0x07, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x72,
	0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x24, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x10, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x38, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x61,
	0x6d, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70,
	0x72, 0x6f, 0x6d, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61,
	0x6d, 0x73, 0x22, 0x88, 0x01, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x29, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x22, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x52, 0x04,
	0x70, 0x61, 0x67, 0x65, 0x12, 0x27, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x13, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x6d, 0x2e, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x32, 0x91, 0x04,
	0x0a, 0x05, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x66, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f,
	0x6d, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x6d, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x22, 0x12, 0x2f, 0x70, 0x72, 0x6f, 0x6d, 0x2f,
	0x76, 0x31, 0x2f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2f, 0x61, 0x64, 0x64, 0x3a, 0x01, 0x2a, 0x12,
	0x6c, 0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x1c,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x6d, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x6d, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x23, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d,
	0x1a, 0x18, 0x2f, 0x70, 0x72, 0x6f, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x2f, 0x65, 0x64, 0x69, 0x74, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x3a, 0x01, 0x2a, 0x12, 0x6b, 0x0a,
	0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x1c, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x6d, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x70, 0x72, 0x6f, 0x6d, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x22, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x2a, 0x1a,
	0x2f, 0x70, 0x72, 0x6f, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2f, 0x64,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x62, 0x0a, 0x08, 0x47, 0x65,
	0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x19, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f,
	0x6d, 0x2e, 0x47, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x17, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x6d, 0x2e, 0x47, 0x65, 0x74,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x22, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x1c, 0x12, 0x1a, 0x2f, 0x70, 0x72, 0x6f, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x2f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x61,
	0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x1a, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x70, 0x72, 0x6f, 0x6d, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72,
	0x6f, 0x6d, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x1e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x22, 0x13, 0x2f, 0x70, 0x72, 0x6f, 0x6d,
	0x2f, 0x76, 0x31, 0x2f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x3a, 0x01,
	0x2a, 0x42, 0x32, 0x0a, 0x0b, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x6d, 0x2e, 0x76, 0x31,
	0x50, 0x01, 0x5a, 0x21, 0x70, 0x72, 0x6f, 0x6d, 0x65, 0x74, 0x68, 0x65, 0x75, 0x73, 0x2d, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x6d, 0x2f,
	0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_prom_v1_group_proto_rawDescOnce sync.Once
	file_prom_v1_group_proto_rawDescData = file_prom_v1_group_proto_rawDesc
)

func file_prom_v1_group_proto_rawDescGZIP() []byte {
	file_prom_v1_group_proto_rawDescOnce.Do(func() {
		file_prom_v1_group_proto_rawDescData = protoimpl.X.CompressGZIP(file_prom_v1_group_proto_rawDescData)
	})
	return file_prom_v1_group_proto_rawDescData
}

var file_prom_v1_group_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_prom_v1_group_proto_goTypes = []interface{}{
	(*CreateGroupRequest)(nil),     // 0: api.prom.CreateGroupRequest
	(*CreateGroupReply)(nil),       // 1: api.prom.CreateGroupReply
	(*UpdateGroupRequest)(nil),     // 2: api.prom.UpdateGroupRequest
	(*UpdateGroupReply)(nil),       // 3: api.prom.UpdateGroupReply
	(*DeleteGroupRequest)(nil),     // 4: api.prom.DeleteGroupRequest
	(*DeleteGroupReply)(nil),       // 5: api.prom.DeleteGroupReply
	(*GetGroupRequest)(nil),        // 6: api.prom.GetGroupRequest
	(*GetGroupReply)(nil),          // 7: api.prom.GetGroupReply
	(*ListGroupRequestParams)(nil), // 8: api.prom.ListGroupRequestParams
	(*ListGroupRequest)(nil),       // 9: api.prom.ListGroupRequest
	(*ListGroupReply)(nil),         // 10: api.prom.ListGroupReply
	(*prom.GroupItem)(nil),         // 11: api.prom.GroupItem
	(*api.Response)(nil),           // 12: api.Response
	(*api.PageRequest)(nil),        // 13: api.PageRequest
	(*api.PageReply)(nil),          // 14: api.PageReply
}
var file_prom_v1_group_proto_depIdxs = []int32{
	11, // 0: api.prom.CreateGroupRequest.group:type_name -> api.prom.GroupItem
	12, // 1: api.prom.CreateGroupReply.response:type_name -> api.Response
	11, // 2: api.prom.UpdateGroupRequest.group:type_name -> api.prom.GroupItem
	12, // 3: api.prom.UpdateGroupReply.response:type_name -> api.Response
	12, // 4: api.prom.DeleteGroupReply.response:type_name -> api.Response
	12, // 5: api.prom.GetGroupReply.response:type_name -> api.Response
	11, // 6: api.prom.GetGroupReply.group:type_name -> api.prom.GroupItem
	13, // 7: api.prom.ListGroupRequest.page:type_name -> api.PageRequest
	8,  // 8: api.prom.ListGroupRequest.params:type_name -> api.prom.ListGroupRequestParams
	12, // 9: api.prom.ListGroupReply.response:type_name -> api.Response
	14, // 10: api.prom.ListGroupReply.page:type_name -> api.PageReply
	11, // 11: api.prom.ListGroupReply.list:type_name -> api.prom.GroupItem
	0,  // 12: api.prom.Group.CreateGroup:input_type -> api.prom.CreateGroupRequest
	2,  // 13: api.prom.Group.UpdateGroup:input_type -> api.prom.UpdateGroupRequest
	4,  // 14: api.prom.Group.DeleteGroup:input_type -> api.prom.DeleteGroupRequest
	6,  // 15: api.prom.Group.GetGroup:input_type -> api.prom.GetGroupRequest
	9,  // 16: api.prom.Group.ListGroup:input_type -> api.prom.ListGroupRequest
	1,  // 17: api.prom.Group.CreateGroup:output_type -> api.prom.CreateGroupReply
	3,  // 18: api.prom.Group.UpdateGroup:output_type -> api.prom.UpdateGroupReply
	5,  // 19: api.prom.Group.DeleteGroup:output_type -> api.prom.DeleteGroupReply
	7,  // 20: api.prom.Group.GetGroup:output_type -> api.prom.GetGroupReply
	10, // 21: api.prom.Group.ListGroup:output_type -> api.prom.ListGroupReply
	17, // [17:22] is the sub-list for method output_type
	12, // [12:17] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_prom_v1_group_proto_init() }
func file_prom_v1_group_proto_init() {
	if File_prom_v1_group_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_prom_v1_group_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateGroupRequest); i {
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
		file_prom_v1_group_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateGroupReply); i {
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
		file_prom_v1_group_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateGroupRequest); i {
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
		file_prom_v1_group_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateGroupReply); i {
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
		file_prom_v1_group_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteGroupRequest); i {
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
		file_prom_v1_group_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteGroupReply); i {
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
		file_prom_v1_group_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGroupRequest); i {
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
		file_prom_v1_group_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGroupReply); i {
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
		file_prom_v1_group_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListGroupRequestParams); i {
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
		file_prom_v1_group_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListGroupRequest); i {
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
		file_prom_v1_group_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListGroupReply); i {
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
			RawDescriptor: file_prom_v1_group_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_prom_v1_group_proto_goTypes,
		DependencyIndexes: file_prom_v1_group_proto_depIdxs,
		MessageInfos:      file_prom_v1_group_proto_msgTypes,
	}.Build()
	File_prom_v1_group_proto = out.File
	file_prom_v1_group_proto_rawDesc = nil
	file_prom_v1_group_proto_goTypes = nil
	file_prom_v1_group_proto_depIdxs = nil
}
