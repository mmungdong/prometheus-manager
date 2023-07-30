// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: strategy/v1/crud.proto

package v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	api "prometheus-manager/api"
	strategy "prometheus-manager/api/strategy"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateRuleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rule    *strategy.Rule `protobuf:"bytes,1,opt,name=rule,proto3" json:"rule,omitempty"`
	GroupId uint32         `protobuf:"varint,2,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
}

func (x *CreateRuleRequest) Reset() {
	*x = CreateRuleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strategy_v1_crud_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRuleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRuleRequest) ProtoMessage() {}

func (x *CreateRuleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_strategy_v1_crud_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRuleRequest.ProtoReflect.Descriptor instead.
func (*CreateRuleRequest) Descriptor() ([]byte, []int) {
	return file_strategy_v1_crud_proto_rawDescGZIP(), []int{0}
}

func (x *CreateRuleRequest) GetRule() *strategy.Rule {
	if x != nil {
		return x.Rule
	}
	return nil
}

func (x *CreateRuleRequest) GetGroupId() uint32 {
	if x != nil {
		return x.GroupId
	}
	return 0
}

type CreateRuleReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 响应内容
	Response *api.Response `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *CreateRuleReply) Reset() {
	*x = CreateRuleReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strategy_v1_crud_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRuleReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRuleReply) ProtoMessage() {}

func (x *CreateRuleReply) ProtoReflect() protoreflect.Message {
	mi := &file_strategy_v1_crud_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRuleReply.ProtoReflect.Descriptor instead.
func (*CreateRuleReply) Descriptor() ([]byte, []int) {
	return file_strategy_v1_crud_proto_rawDescGZIP(), []int{1}
}

func (x *CreateRuleReply) GetResponse() *api.Response {
	if x != nil {
		return x.Response
	}
	return nil
}

type UpdateRuleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string         `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Rule *strategy.Rule `protobuf:"bytes,2,opt,name=rule,proto3" json:"rule,omitempty"`
}

func (x *UpdateRuleRequest) Reset() {
	*x = UpdateRuleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strategy_v1_crud_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRuleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRuleRequest) ProtoMessage() {}

func (x *UpdateRuleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_strategy_v1_crud_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRuleRequest.ProtoReflect.Descriptor instead.
func (*UpdateRuleRequest) Descriptor() ([]byte, []int) {
	return file_strategy_v1_crud_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateRuleRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateRuleRequest) GetRule() *strategy.Rule {
	if x != nil {
		return x.Rule
	}
	return nil
}

type UpdateRuleReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 响应内容
	Response *api.Response `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *UpdateRuleReply) Reset() {
	*x = UpdateRuleReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strategy_v1_crud_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRuleReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRuleReply) ProtoMessage() {}

func (x *UpdateRuleReply) ProtoReflect() protoreflect.Message {
	mi := &file_strategy_v1_crud_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRuleReply.ProtoReflect.Descriptor instead.
func (*UpdateRuleReply) Descriptor() ([]byte, []int) {
	return file_strategy_v1_crud_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateRuleReply) GetResponse() *api.Response {
	if x != nil {
		return x.Response
	}
	return nil
}

type DeleteRuleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteRuleRequest) Reset() {
	*x = DeleteRuleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strategy_v1_crud_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRuleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRuleRequest) ProtoMessage() {}

func (x *DeleteRuleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_strategy_v1_crud_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRuleRequest.ProtoReflect.Descriptor instead.
func (*DeleteRuleRequest) Descriptor() ([]byte, []int) {
	return file_strategy_v1_crud_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteRuleRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteRuleReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 响应内容
	Response *api.Response `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *DeleteRuleReply) Reset() {
	*x = DeleteRuleReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strategy_v1_crud_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRuleReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRuleReply) ProtoMessage() {}

func (x *DeleteRuleReply) ProtoReflect() protoreflect.Message {
	mi := &file_strategy_v1_crud_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRuleReply.ProtoReflect.Descriptor instead.
func (*DeleteRuleReply) Descriptor() ([]byte, []int) {
	return file_strategy_v1_crud_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteRuleReply) GetResponse() *api.Response {
	if x != nil {
		return x.Response
	}
	return nil
}

type GetRuleDetailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetRuleDetailRequest) Reset() {
	*x = GetRuleDetailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strategy_v1_crud_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRuleDetailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRuleDetailRequest) ProtoMessage() {}

func (x *GetRuleDetailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_strategy_v1_crud_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRuleDetailRequest.ProtoReflect.Descriptor instead.
func (*GetRuleDetailRequest) Descriptor() ([]byte, []int) {
	return file_strategy_v1_crud_proto_rawDescGZIP(), []int{6}
}

func (x *GetRuleDetailRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetRuleDetailReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 响应内容
	Response *api.Response  `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
	Rule     *strategy.Rule `protobuf:"bytes,2,opt,name=rule,proto3" json:"rule,omitempty"`
}

func (x *GetRuleDetailReply) Reset() {
	*x = GetRuleDetailReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strategy_v1_crud_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRuleDetailReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRuleDetailReply) ProtoMessage() {}

func (x *GetRuleDetailReply) ProtoReflect() protoreflect.Message {
	mi := &file_strategy_v1_crud_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRuleDetailReply.ProtoReflect.Descriptor instead.
func (*GetRuleDetailReply) Descriptor() ([]byte, []int) {
	return file_strategy_v1_crud_proto_rawDescGZIP(), []int{7}
}

func (x *GetRuleDetailReply) GetResponse() *api.Response {
	if x != nil {
		return x.Response
	}
	return nil
}

func (x *GetRuleDetailReply) GetRule() *strategy.Rule {
	if x != nil {
		return x.Rule
	}
	return nil
}

type StrategiesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeId string `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
}

func (x *StrategiesRequest) Reset() {
	*x = StrategiesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strategy_v1_crud_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StrategiesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StrategiesRequest) ProtoMessage() {}

func (x *StrategiesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_strategy_v1_crud_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StrategiesRequest.ProtoReflect.Descriptor instead.
func (*StrategiesRequest) Descriptor() ([]byte, []int) {
	return file_strategy_v1_crud_proto_rawDescGZIP(), []int{8}
}

func (x *StrategiesRequest) GetNodeId() string {
	if x != nil {
		return x.NodeId
	}
	return ""
}

type StrategiesReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 响应内容
	Response    *api.Response           `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
	StrategyDir []*strategy.StrategyDir `protobuf:"bytes,2,rep,name=strategy_dir,json=strategyDir,proto3" json:"strategy_dir,omitempty"`
}

func (x *StrategiesReply) Reset() {
	*x = StrategiesReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strategy_v1_crud_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StrategiesReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StrategiesReply) ProtoMessage() {}

func (x *StrategiesReply) ProtoReflect() protoreflect.Message {
	mi := &file_strategy_v1_crud_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StrategiesReply.ProtoReflect.Descriptor instead.
func (*StrategiesReply) Descriptor() ([]byte, []int) {
	return file_strategy_v1_crud_proto_rawDescGZIP(), []int{9}
}

func (x *StrategiesReply) GetResponse() *api.Response {
	if x != nil {
		return x.Response
	}
	return nil
}

func (x *StrategiesReply) GetStrategyDir() []*strategy.StrategyDir {
	if x != nil {
		return x.StrategyDir
	}
	return nil
}

var File_strategy_v1_crud_proto protoreflect.FileDescriptor

var file_strategy_v1_crud_proto_rawDesc = []byte{
	0x0a, 0x16, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x72,
	0x75, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74,
	0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x0a, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x73, 0x74,
	0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x2f, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5f, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52,
	0x75, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x04, 0x72, 0x75,
	0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73,
	0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x2e, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x04, 0x72, 0x75,
	0x6c, 0x65, 0x12, 0x22, 0x0a, 0x08, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0d, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x2a, 0x02, 0x20, 0x00, 0x52, 0x07, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x22, 0x3c, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x52, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x29, 0x0a, 0x08, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x4b, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x75,
	0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x26, 0x0a, 0x04, 0x72, 0x75, 0x6c,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74,
	0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x2e, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x04, 0x72, 0x75, 0x6c,
	0x65, 0x22, 0x3c, 0x0a, 0x0f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x75, 0x6c, 0x65, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x29, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x23, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x22, 0x3c, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x75,
	0x6c, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x29, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x26, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x52, 0x75, 0x6c, 0x65, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x67, 0x0a, 0x12, 0x47, 0x65,
	0x74, 0x52, 0x75, 0x6c, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x29, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x04, 0x72,
	0x75, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x2e, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x04, 0x72,
	0x75, 0x6c, 0x65, 0x22, 0x35, 0x0a, 0x11, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x69, 0x65,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x07, 0x6e, 0x6f, 0x64, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02,
	0x10, 0x01, 0x52, 0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x22, 0x7a, 0x0a, 0x0f, 0x53, 0x74,
	0x72, 0x61, 0x74, 0x65, 0x67, 0x69, 0x65, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x29, 0x0a,
	0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x08,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x0c, 0x73, 0x74, 0x72, 0x61,
	0x74, 0x65, 0x67, 0x79, 0x5f, 0x64, 0x69, 0x72, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x2e, 0x53, 0x74,
	0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x44, 0x69, 0x72, 0x52, 0x0b, 0x73, 0x74, 0x72, 0x61, 0x74,
	0x65, 0x67, 0x79, 0x44, 0x69, 0x72, 0x32, 0xc6, 0x04, 0x0a, 0x04, 0x43, 0x72, 0x75, 0x64, 0x12,
	0x6b, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x75, 0x6c, 0x65, 0x12, 0x22, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x22, 0x0c, 0x2f, 0x72, 0x75,
	0x6c, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x64, 0x64, 0x3a, 0x01, 0x2a, 0x12, 0x71, 0x0a, 0x0a,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x75, 0x6c, 0x65, 0x12, 0x22, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x2e, 0x76, 0x31,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x1a, 0x12, 0x2f, 0x72, 0x75, 0x6c, 0x65, 0x2f,
	0x76, 0x31, 0x2f, 0x65, 0x64, 0x69, 0x74, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x3a, 0x01, 0x2a, 0x12,
	0x70, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x75, 0x6c, 0x65, 0x12, 0x22, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x2e, 0x76, 0x31, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79,
	0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x2a, 0x14, 0x2f, 0x72, 0x75,
	0x6c, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x2f, 0x7b, 0x69, 0x64,
	0x7d, 0x12, 0x76, 0x0a, 0x0a, 0x52, 0x75, 0x6c, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12,
	0x25, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x75, 0x6c, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x72,
	0x61, 0x74, 0x65, 0x67, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x75, 0x6c, 0x65,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1c, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x16, 0x12, 0x14, 0x2f, 0x72, 0x75, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x74, 0x0a, 0x0a, 0x53, 0x74, 0x72,
	0x61, 0x74, 0x65, 0x67, 0x69, 0x65, 0x73, 0x12, 0x22, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74,
	0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65,
	0x67, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74,
	0x72, 0x61, 0x74, 0x65, 0x67, 0x69, 0x65, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x20, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x12, 0x18, 0x2f, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x69,
	0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x7b, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x69, 0x64, 0x7d, 0x42,
	0x3a, 0x0a, 0x0f, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x2e,
	0x76, 0x31, 0x50, 0x01, 0x5a, 0x25, 0x70, 0x72, 0x6f, 0x6d, 0x65, 0x74, 0x68, 0x65, 0x75, 0x73,
	0x2d, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x74, 0x72,
	0x61, 0x74, 0x65, 0x67, 0x79, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_strategy_v1_crud_proto_rawDescOnce sync.Once
	file_strategy_v1_crud_proto_rawDescData = file_strategy_v1_crud_proto_rawDesc
)

func file_strategy_v1_crud_proto_rawDescGZIP() []byte {
	file_strategy_v1_crud_proto_rawDescOnce.Do(func() {
		file_strategy_v1_crud_proto_rawDescData = protoimpl.X.CompressGZIP(file_strategy_v1_crud_proto_rawDescData)
	})
	return file_strategy_v1_crud_proto_rawDescData
}

var file_strategy_v1_crud_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_strategy_v1_crud_proto_goTypes = []interface{}{
	(*CreateRuleRequest)(nil),    // 0: api.strategy.v1.CreateRuleRequest
	(*CreateRuleReply)(nil),      // 1: api.strategy.v1.CreateRuleReply
	(*UpdateRuleRequest)(nil),    // 2: api.strategy.v1.UpdateRuleRequest
	(*UpdateRuleReply)(nil),      // 3: api.strategy.v1.UpdateRuleReply
	(*DeleteRuleRequest)(nil),    // 4: api.strategy.v1.DeleteRuleRequest
	(*DeleteRuleReply)(nil),      // 5: api.strategy.v1.DeleteRuleReply
	(*GetRuleDetailRequest)(nil), // 6: api.strategy.v1.GetRuleDetailRequest
	(*GetRuleDetailReply)(nil),   // 7: api.strategy.v1.GetRuleDetailReply
	(*StrategiesRequest)(nil),    // 8: api.strategy.v1.StrategiesRequest
	(*StrategiesReply)(nil),      // 9: api.strategy.v1.StrategiesReply
	(*strategy.Rule)(nil),        // 10: api.strategy.Rule
	(*api.Response)(nil),         // 11: api.Response
	(*strategy.StrategyDir)(nil), // 12: api.strategy.StrategyDir
}
var file_strategy_v1_crud_proto_depIdxs = []int32{
	10, // 0: api.strategy.v1.CreateRuleRequest.rule:type_name -> api.strategy.Rule
	11, // 1: api.strategy.v1.CreateRuleReply.response:type_name -> api.Response
	10, // 2: api.strategy.v1.UpdateRuleRequest.rule:type_name -> api.strategy.Rule
	11, // 3: api.strategy.v1.UpdateRuleReply.response:type_name -> api.Response
	11, // 4: api.strategy.v1.DeleteRuleReply.response:type_name -> api.Response
	11, // 5: api.strategy.v1.GetRuleDetailReply.response:type_name -> api.Response
	10, // 6: api.strategy.v1.GetRuleDetailReply.rule:type_name -> api.strategy.Rule
	11, // 7: api.strategy.v1.StrategiesReply.response:type_name -> api.Response
	12, // 8: api.strategy.v1.StrategiesReply.strategy_dir:type_name -> api.strategy.StrategyDir
	0,  // 9: api.strategy.v1.Crud.CreateRule:input_type -> api.strategy.v1.CreateRuleRequest
	2,  // 10: api.strategy.v1.Crud.UpdateRule:input_type -> api.strategy.v1.UpdateRuleRequest
	4,  // 11: api.strategy.v1.Crud.DeleteRule:input_type -> api.strategy.v1.DeleteRuleRequest
	6,  // 12: api.strategy.v1.Crud.RuleDetail:input_type -> api.strategy.v1.GetRuleDetailRequest
	8,  // 13: api.strategy.v1.Crud.Strategies:input_type -> api.strategy.v1.StrategiesRequest
	1,  // 14: api.strategy.v1.Crud.CreateRule:output_type -> api.strategy.v1.CreateRuleReply
	3,  // 15: api.strategy.v1.Crud.UpdateRule:output_type -> api.strategy.v1.UpdateRuleReply
	5,  // 16: api.strategy.v1.Crud.DeleteRule:output_type -> api.strategy.v1.DeleteRuleReply
	7,  // 17: api.strategy.v1.Crud.RuleDetail:output_type -> api.strategy.v1.GetRuleDetailReply
	9,  // 18: api.strategy.v1.Crud.Strategies:output_type -> api.strategy.v1.StrategiesReply
	14, // [14:19] is the sub-list for method output_type
	9,  // [9:14] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_strategy_v1_crud_proto_init() }
func file_strategy_v1_crud_proto_init() {
	if File_strategy_v1_crud_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_strategy_v1_crud_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRuleRequest); i {
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
		file_strategy_v1_crud_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRuleReply); i {
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
		file_strategy_v1_crud_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateRuleRequest); i {
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
		file_strategy_v1_crud_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateRuleReply); i {
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
		file_strategy_v1_crud_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteRuleRequest); i {
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
		file_strategy_v1_crud_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteRuleReply); i {
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
		file_strategy_v1_crud_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRuleDetailRequest); i {
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
		file_strategy_v1_crud_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRuleDetailReply); i {
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
		file_strategy_v1_crud_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StrategiesRequest); i {
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
		file_strategy_v1_crud_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StrategiesReply); i {
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
			RawDescriptor: file_strategy_v1_crud_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_strategy_v1_crud_proto_goTypes,
		DependencyIndexes: file_strategy_v1_crud_proto_depIdxs,
		MessageInfos:      file_strategy_v1_crud_proto_msgTypes,
	}.Build()
	File_strategy_v1_crud_proto = out.File
	file_strategy_v1_crud_proto_rawDesc = nil
	file_strategy_v1_crud_proto_goTypes = nil
	file_strategy_v1_crud_proto_depIdxs = nil
}
