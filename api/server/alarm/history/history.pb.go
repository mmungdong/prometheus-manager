// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: server/alarm/history/history.proto

package history

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	api "prometheus-manager/api"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// 获取报警历史详情请求参数
type GetHistoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 报警历史ID, 0 < id
	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetHistoryRequest) Reset() {
	*x = GetHistoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_alarm_history_history_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetHistoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetHistoryRequest) ProtoMessage() {}

func (x *GetHistoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_server_alarm_history_history_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetHistoryRequest.ProtoReflect.Descriptor instead.
func (*GetHistoryRequest) Descriptor() ([]byte, []int) {
	return file_server_alarm_history_history_proto_rawDescGZIP(), []int{0}
}

func (x *GetHistoryRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

// 获取报警历史详情响应参数
type GetHistoryReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 报警历史详情
	AlarmHistory *api.AlarmHistoryV1 `protobuf:"bytes,1,opt,name=alarmHistory,proto3" json:"alarmHistory,omitempty"`
}

func (x *GetHistoryReply) Reset() {
	*x = GetHistoryReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_alarm_history_history_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetHistoryReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetHistoryReply) ProtoMessage() {}

func (x *GetHistoryReply) ProtoReflect() protoreflect.Message {
	mi := &file_server_alarm_history_history_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetHistoryReply.ProtoReflect.Descriptor instead.
func (*GetHistoryReply) Descriptor() ([]byte, []int) {
	return file_server_alarm_history_history_proto_rawDescGZIP(), []int{1}
}

func (x *GetHistoryReply) GetAlarmHistory() *api.AlarmHistoryV1 {
	if x != nil {
		return x.AlarmHistory
	}
	return nil
}

// 获取报警历史列表请求参数
type ListHistoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 分页参数
	Page *api.PageRequest `protobuf:"bytes,1,opt,name=page,proto3" json:"page,omitempty"`
	// 关键字
	Keyword string `protobuf:"bytes,2,opt,name=keyword,proto3" json:"keyword,omitempty"`
	// 报警状态, 对应PromDictV1 id
	Status api.Status `protobuf:"varint,3,opt,name=status,proto3,enum=api.Status" json:"status,omitempty"`
	// 报警页面, 对应AlarmPageV1 id
	AlarmPages []uint32 `protobuf:"varint,4,rep,packed,name=alarmPages,proto3" json:"alarmPages,omitempty"`
	// 开始时间, unix时间戳
	FiringStartAt int64 `protobuf:"varint,5,opt,name=firingStartAt,proto3" json:"firingStartAt,omitempty"`
	// 结束时间, unix时间戳
	FiringEndAt int64 `protobuf:"varint,6,opt,name=firingEndAt,proto3" json:"firingEndAt,omitempty"`
	// 报警等级
	AlarmLevelIds []uint32 `protobuf:"varint,7,rep,packed,name=alarmLevelIds,proto3" json:"alarmLevelIds,omitempty"`
	// 告警策略
	StrategyIds []uint32 `protobuf:"varint,8,rep,packed,name=strategyIds,proto3" json:"strategyIds,omitempty"`
	// 恢复开始时间
	ResolvedStartAt int64 `protobuf:"varint,9,opt,name=resolvedStartAt,proto3" json:"resolvedStartAt,omitempty"`
	// 恢复结束时间
	ResolvedEndAt int64 `protobuf:"varint,10,opt,name=resolvedEndAt,proto3" json:"resolvedEndAt,omitempty"`
	// 持续时间
	Duration int64 `protobuf:"varint,11,opt,name=duration,proto3" json:"duration,omitempty"`
}

func (x *ListHistoryRequest) Reset() {
	*x = ListHistoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_alarm_history_history_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListHistoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListHistoryRequest) ProtoMessage() {}

func (x *ListHistoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_server_alarm_history_history_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListHistoryRequest.ProtoReflect.Descriptor instead.
func (*ListHistoryRequest) Descriptor() ([]byte, []int) {
	return file_server_alarm_history_history_proto_rawDescGZIP(), []int{2}
}

func (x *ListHistoryRequest) GetPage() *api.PageRequest {
	if x != nil {
		return x.Page
	}
	return nil
}

func (x *ListHistoryRequest) GetKeyword() string {
	if x != nil {
		return x.Keyword
	}
	return ""
}

func (x *ListHistoryRequest) GetStatus() api.Status {
	if x != nil {
		return x.Status
	}
	return api.Status(0)
}

func (x *ListHistoryRequest) GetAlarmPages() []uint32 {
	if x != nil {
		return x.AlarmPages
	}
	return nil
}

func (x *ListHistoryRequest) GetFiringStartAt() int64 {
	if x != nil {
		return x.FiringStartAt
	}
	return 0
}

func (x *ListHistoryRequest) GetFiringEndAt() int64 {
	if x != nil {
		return x.FiringEndAt
	}
	return 0
}

func (x *ListHistoryRequest) GetAlarmLevelIds() []uint32 {
	if x != nil {
		return x.AlarmLevelIds
	}
	return nil
}

func (x *ListHistoryRequest) GetStrategyIds() []uint32 {
	if x != nil {
		return x.StrategyIds
	}
	return nil
}

func (x *ListHistoryRequest) GetResolvedStartAt() int64 {
	if x != nil {
		return x.ResolvedStartAt
	}
	return 0
}

func (x *ListHistoryRequest) GetResolvedEndAt() int64 {
	if x != nil {
		return x.ResolvedEndAt
	}
	return 0
}

func (x *ListHistoryRequest) GetDuration() int64 {
	if x != nil {
		return x.Duration
	}
	return 0
}

// 获取报警历史列表响应参数
type ListHistoryReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 分页参数
	Page *api.PageReply `protobuf:"bytes,1,opt,name=page,proto3" json:"page,omitempty"`
	// 报警历史列表
	List []*api.AlarmHistoryV1 `protobuf:"bytes,2,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *ListHistoryReply) Reset() {
	*x = ListHistoryReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_alarm_history_history_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListHistoryReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListHistoryReply) ProtoMessage() {}

func (x *ListHistoryReply) ProtoReflect() protoreflect.Message {
	mi := &file_server_alarm_history_history_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListHistoryReply.ProtoReflect.Descriptor instead.
func (*ListHistoryReply) Descriptor() ([]byte, []int) {
	return file_server_alarm_history_history_proto_rawDescGZIP(), []int{3}
}

func (x *ListHistoryReply) GetPage() *api.PageReply {
	if x != nil {
		return x.Page
	}
	return nil
}

func (x *ListHistoryReply) GetList() []*api.AlarmHistoryV1 {
	if x != nil {
		return x.List
	}
	return nil
}

var File_server_alarm_history_history_proto protoreflect.FileDescriptor

var file_server_alarm_history_history_proto_rawDesc = []byte{
	0x0a, 0x22, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x61, 0x6c, 0x61, 0x72, 0x6d, 0x2f, 0x68,
	0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x2f, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2e, 0x61, 0x6c, 0x61, 0x72, 0x6d, 0x2e, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0a, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x0b, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2c,
	0x0a, 0x11, 0x47, 0x65, 0x74, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x42,
	0x07, 0xfa, 0x42, 0x04, 0x2a, 0x02, 0x20, 0x00, 0x52, 0x02, 0x69, 0x64, 0x22, 0x4a, 0x0a, 0x0f,
	0x47, 0x65, 0x74, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x37, 0x0a, 0x0c, 0x61, 0x6c, 0x61, 0x72, 0x6d, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x41, 0x6c, 0x61, 0x72,
	0x6d, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x56, 0x31, 0x52, 0x0c, 0x61, 0x6c, 0x61, 0x72,
	0x6d, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x22, 0xb2, 0x03, 0x0a, 0x12, 0x4c, 0x69, 0x73,
	0x74, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x2e, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42,
	0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x2d, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x82, 0x01, 0x02, 0x10, 0x01,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x61, 0x6c, 0x61, 0x72,
	0x6d, 0x50, 0x61, 0x67, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0a, 0x61, 0x6c,
	0x61, 0x72, 0x6d, 0x50, 0x61, 0x67, 0x65, 0x73, 0x12, 0x24, 0x0a, 0x0d, 0x66, 0x69, 0x72, 0x69,
	0x6e, 0x67, 0x53, 0x74, 0x61, 0x72, 0x74, 0x41, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0d, 0x66, 0x69, 0x72, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x61, 0x72, 0x74, 0x41, 0x74, 0x12, 0x20,
	0x0a, 0x0b, 0x66, 0x69, 0x72, 0x69, 0x6e, 0x67, 0x45, 0x6e, 0x64, 0x41, 0x74, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0b, 0x66, 0x69, 0x72, 0x69, 0x6e, 0x67, 0x45, 0x6e, 0x64, 0x41, 0x74,
	0x12, 0x24, 0x0a, 0x0d, 0x61, 0x6c, 0x61, 0x72, 0x6d, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x49, 0x64,
	0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0d, 0x61, 0x6c, 0x61, 0x72, 0x6d, 0x4c, 0x65,
	0x76, 0x65, 0x6c, 0x49, 0x64, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65,
	0x67, 0x79, 0x49, 0x64, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0b, 0x73, 0x74, 0x72,
	0x61, 0x74, 0x65, 0x67, 0x79, 0x49, 0x64, 0x73, 0x12, 0x28, 0x0a, 0x0f, 0x72, 0x65, 0x73, 0x6f,
	0x6c, 0x76, 0x65, 0x64, 0x53, 0x74, 0x61, 0x72, 0x74, 0x41, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0f, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x64, 0x53, 0x74, 0x61, 0x72, 0x74,
	0x41, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x64, 0x45, 0x6e,
	0x64, 0x41, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x6c,
	0x76, 0x65, 0x64, 0x45, 0x6e, 0x64, 0x41, 0x74, 0x12, 0x23, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x22,
	0x02, 0x28, 0x00, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x5f, 0x0a,
	0x10, 0x4c, 0x69, 0x73, 0x74, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x22, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x52,
	0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x27, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x41, 0x6c, 0x61, 0x72, 0x6d, 0x48,
	0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x56, 0x31, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x32, 0xa7,
	0x02, 0x0a, 0x07, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x8a, 0x01, 0x0a, 0x0a, 0x47,
	0x65, 0x74, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x2b, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x61, 0x6c, 0x61, 0x72, 0x6d, 0x2e, 0x68, 0x69, 0x73,
	0x74, 0x6f, 0x72, 0x79, 0x2e, 0x47, 0x65, 0x74, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2e, 0x61, 0x6c, 0x61, 0x72, 0x6d, 0x2e, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72,
	0x79, 0x2e, 0x47, 0x65, 0x74, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x24, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x22, 0x19, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x31, 0x2f, 0x61, 0x6c, 0x61, 0x72, 0x6d, 0x2f, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79,
	0x2f, 0x67, 0x65, 0x74, 0x3a, 0x01, 0x2a, 0x12, 0x8e, 0x01, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74,
	0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x2c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2e, 0x61, 0x6c, 0x61, 0x72, 0x6d, 0x2e, 0x68, 0x69, 0x73, 0x74, 0x6f,
	0x72, 0x79, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2e, 0x61, 0x6c, 0x61, 0x72, 0x6d, 0x2e, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x25, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1f, 0x22, 0x1a, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x31, 0x2f, 0x61, 0x6c, 0x61, 0x72, 0x6d, 0x2f, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79,
	0x2f, 0x6c, 0x69, 0x73, 0x74, 0x3a, 0x01, 0x2a, 0x42, 0x51, 0x0a, 0x18, 0x61, 0x70, 0x69, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x61, 0x6c, 0x61, 0x72, 0x6d, 0x2e, 0x68, 0x69, 0x73,
	0x74, 0x6f, 0x72, 0x79, 0x50, 0x01, 0x5a, 0x33, 0x70, 0x72, 0x6f, 0x6d, 0x65, 0x74, 0x68, 0x65,
	0x75, 0x73, 0x2d, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x61, 0x6c, 0x61, 0x72, 0x6d, 0x2f, 0x68, 0x69, 0x73, 0x74,
	0x6f, 0x72, 0x79, 0x3b, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_server_alarm_history_history_proto_rawDescOnce sync.Once
	file_server_alarm_history_history_proto_rawDescData = file_server_alarm_history_history_proto_rawDesc
)

func file_server_alarm_history_history_proto_rawDescGZIP() []byte {
	file_server_alarm_history_history_proto_rawDescOnce.Do(func() {
		file_server_alarm_history_history_proto_rawDescData = protoimpl.X.CompressGZIP(file_server_alarm_history_history_proto_rawDescData)
	})
	return file_server_alarm_history_history_proto_rawDescData
}

var file_server_alarm_history_history_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_server_alarm_history_history_proto_goTypes = []interface{}{
	(*GetHistoryRequest)(nil),  // 0: api.server.alarm.history.GetHistoryRequest
	(*GetHistoryReply)(nil),    // 1: api.server.alarm.history.GetHistoryReply
	(*ListHistoryRequest)(nil), // 2: api.server.alarm.history.ListHistoryRequest
	(*ListHistoryReply)(nil),   // 3: api.server.alarm.history.ListHistoryReply
	(*api.AlarmHistoryV1)(nil), // 4: api.AlarmHistoryV1
	(*api.PageRequest)(nil),    // 5: api.PageRequest
	(api.Status)(0),            // 6: api.Status
	(*api.PageReply)(nil),      // 7: api.PageReply
}
var file_server_alarm_history_history_proto_depIdxs = []int32{
	4, // 0: api.server.alarm.history.GetHistoryReply.alarmHistory:type_name -> api.AlarmHistoryV1
	5, // 1: api.server.alarm.history.ListHistoryRequest.page:type_name -> api.PageRequest
	6, // 2: api.server.alarm.history.ListHistoryRequest.status:type_name -> api.Status
	7, // 3: api.server.alarm.history.ListHistoryReply.page:type_name -> api.PageReply
	4, // 4: api.server.alarm.history.ListHistoryReply.list:type_name -> api.AlarmHistoryV1
	0, // 5: api.server.alarm.history.History.GetHistory:input_type -> api.server.alarm.history.GetHistoryRequest
	2, // 6: api.server.alarm.history.History.ListHistory:input_type -> api.server.alarm.history.ListHistoryRequest
	1, // 7: api.server.alarm.history.History.GetHistory:output_type -> api.server.alarm.history.GetHistoryReply
	3, // 8: api.server.alarm.history.History.ListHistory:output_type -> api.server.alarm.history.ListHistoryReply
	7, // [7:9] is the sub-list for method output_type
	5, // [5:7] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_server_alarm_history_history_proto_init() }
func file_server_alarm_history_history_proto_init() {
	if File_server_alarm_history_history_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_server_alarm_history_history_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetHistoryRequest); i {
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
		file_server_alarm_history_history_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetHistoryReply); i {
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
		file_server_alarm_history_history_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListHistoryRequest); i {
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
		file_server_alarm_history_history_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListHistoryReply); i {
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
			RawDescriptor: file_server_alarm_history_history_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_server_alarm_history_history_proto_goTypes,
		DependencyIndexes: file_server_alarm_history_history_proto_depIdxs,
		MessageInfos:      file_server_alarm_history_history_proto_msgTypes,
	}.Build()
	File_server_alarm_history_history_proto = out.File
	file_server_alarm_history_history_proto_rawDesc = nil
	file_server_alarm_history_history_proto_goTypes = nil
	file_server_alarm_history_history_proto_depIdxs = nil
}
