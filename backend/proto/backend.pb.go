// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v3.19.6
// source: backend.proto

package proto

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

type BackendResponse_State int32

const (
	BackendResponse_UNINITIALIZED BackendResponse_State = 0
	BackendResponse_BUSY          BackendResponse_State = 1
	BackendResponse_READY         BackendResponse_State = 2
	BackendResponse_ERROR         BackendResponse_State = -1
)

// Enum value maps for BackendResponse_State.
var (
	BackendResponse_State_name = map[int32]string{
		0:  "UNINITIALIZED",
		1:  "BUSY",
		2:  "READY",
		-1: "ERROR",
	}
	BackendResponse_State_value = map[string]int32{
		"UNINITIALIZED": 0,
		"BUSY":          1,
		"READY":         2,
		"ERROR":         -1,
	}
)

func (x BackendResponse_State) Enum() *BackendResponse_State {
	p := new(BackendResponse_State)
	*p = x
	return p
}

func (x BackendResponse_State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BackendResponse_State) Descriptor() protoreflect.EnumDescriptor {
	return file_backend_proto_enumTypes[0].Descriptor()
}

func (BackendResponse_State) Type() protoreflect.EnumType {
	return &file_backend_proto_enumTypes[0]
}

func (x BackendResponse_State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BackendResponse_State.Descriptor instead.
func (BackendResponse_State) EnumDescriptor() ([]byte, []int) {
	return file_backend_proto_rawDescGZIP(), []int{4, 0}
}

type BackendOptions struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Port          string                 `protobuf:"bytes,1,opt,name=Port,proto3" json:"Port,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BackendOptions) Reset() {
	*x = BackendOptions{}
	mi := &file_backend_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BackendOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BackendOptions) ProtoMessage() {}

func (x *BackendOptions) ProtoReflect() protoreflect.Message {
	mi := &file_backend_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BackendOptions.ProtoReflect.Descriptor instead.
func (*BackendOptions) Descriptor() ([]byte, []int) {
	return file_backend_proto_rawDescGZIP(), []int{0}
}

func (x *BackendOptions) GetPort() string {
	if x != nil {
		return x.Port
	}
	return ""
}

type Reply struct {
	state                  protoimpl.MessageState `protogen:"open.v1"`
	Message                []byte                 `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Tokens                 int32                  `protobuf:"varint,2,opt,name=tokens,proto3" json:"tokens,omitempty"`
	PromptTokens           int32                  `protobuf:"varint,3,opt,name=prompt_tokens,json=promptTokens,proto3" json:"prompt_tokens,omitempty"`
	TimingPromptProcessing float64                `protobuf:"fixed64,4,opt,name=timing_prompt_processing,json=timingPromptProcessing,proto3" json:"timing_prompt_processing,omitempty"`
	TimingTokenGeneration  float64                `protobuf:"fixed64,5,opt,name=timing_token_generation,json=timingTokenGeneration,proto3" json:"timing_token_generation,omitempty"`
	unknownFields          protoimpl.UnknownFields
	sizeCache              protoimpl.SizeCache
}

func (x *Reply) Reset() {
	*x = Reply{}
	mi := &file_backend_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Reply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Reply) ProtoMessage() {}

func (x *Reply) ProtoReflect() protoreflect.Message {
	mi := &file_backend_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Reply.ProtoReflect.Descriptor instead.
func (*Reply) Descriptor() ([]byte, []int) {
	return file_backend_proto_rawDescGZIP(), []int{1}
}

func (x *Reply) GetMessage() []byte {
	if x != nil {
		return x.Message
	}
	return nil
}

func (x *Reply) GetTokens() int32 {
	if x != nil {
		return x.Tokens
	}
	return 0
}

func (x *Reply) GetPromptTokens() int32 {
	if x != nil {
		return x.PromptTokens
	}
	return 0
}

func (x *Reply) GetTimingPromptProcessing() float64 {
	if x != nil {
		return x.TimingPromptProcessing
	}
	return 0
}

func (x *Reply) GetTimingTokenGeneration() float64 {
	if x != nil {
		return x.TimingTokenGeneration
	}
	return 0
}

type ChatOptions struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Prompt        string                 `protobuf:"bytes,1,opt,name=Prompt,proto3" json:"Prompt,omitempty"`
	Model         string                 `protobuf:"bytes,2,opt,name=Model,proto3" json:"Model,omitempty"`
	Predict       int32                  `protobuf:"varint,3,opt,name=Predict,proto3" json:"Predict,omitempty"`
	RepeatPenalty float32                `protobuf:"fixed32,4,opt,name=RepeatPenalty,proto3" json:"RepeatPenalty,omitempty"`
	Backends      string                 `protobuf:"bytes,5,opt,name=Backends,proto3" json:"Backends,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ChatOptions) Reset() {
	*x = ChatOptions{}
	mi := &file_backend_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ChatOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatOptions) ProtoMessage() {}

func (x *ChatOptions) ProtoReflect() protoreflect.Message {
	mi := &file_backend_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatOptions.ProtoReflect.Descriptor instead.
func (*ChatOptions) Descriptor() ([]byte, []int) {
	return file_backend_proto_rawDescGZIP(), []int{2}
}

func (x *ChatOptions) GetPrompt() string {
	if x != nil {
		return x.Prompt
	}
	return ""
}

func (x *ChatOptions) GetModel() string {
	if x != nil {
		return x.Model
	}
	return ""
}

func (x *ChatOptions) GetPredict() int32 {
	if x != nil {
		return x.Predict
	}
	return 0
}

func (x *ChatOptions) GetRepeatPenalty() float32 {
	if x != nil {
		return x.RepeatPenalty
	}
	return 0
}

func (x *ChatOptions) GetBackends() string {
	if x != nil {
		return x.Backends
	}
	return ""
}

type MemoryUsageData struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Total         uint64                 `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Breakdown     map[string]uint64      `protobuf:"bytes,2,rep,name=breakdown,proto3" json:"breakdown,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MemoryUsageData) Reset() {
	*x = MemoryUsageData{}
	mi := &file_backend_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MemoryUsageData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemoryUsageData) ProtoMessage() {}

func (x *MemoryUsageData) ProtoReflect() protoreflect.Message {
	mi := &file_backend_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemoryUsageData.ProtoReflect.Descriptor instead.
func (*MemoryUsageData) Descriptor() ([]byte, []int) {
	return file_backend_proto_rawDescGZIP(), []int{3}
}

func (x *MemoryUsageData) GetTotal() uint64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *MemoryUsageData) GetBreakdown() map[string]uint64 {
	if x != nil {
		return x.Breakdown
	}
	return nil
}

type BackendResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	State         BackendResponse_State  `protobuf:"varint,1,opt,name=state,proto3,enum=BackendResponse_State" json:"state,omitempty"`
	Memory        *MemoryUsageData       `protobuf:"bytes,2,opt,name=memory,proto3" json:"memory,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BackendResponse) Reset() {
	*x = BackendResponse{}
	mi := &file_backend_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BackendResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BackendResponse) ProtoMessage() {}

func (x *BackendResponse) ProtoReflect() protoreflect.Message {
	mi := &file_backend_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BackendResponse.ProtoReflect.Descriptor instead.
func (*BackendResponse) Descriptor() ([]byte, []int) {
	return file_backend_proto_rawDescGZIP(), []int{4}
}

func (x *BackendResponse) GetState() BackendResponse_State {
	if x != nil {
		return x.State
	}
	return BackendResponse_UNINITIALIZED
}

func (x *BackendResponse) GetMemory() *MemoryUsageData {
	if x != nil {
		return x.Memory
	}
	return nil
}

var File_backend_proto protoreflect.FileDescriptor

var file_backend_proto_rawDesc = string([]byte{
	0x0a, 0x0d, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x24, 0x0a, 0x0e, 0x42, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x12, 0x12, 0x0a, 0x04, 0x50, 0x6f, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x50, 0x6f, 0x72, 0x74, 0x22, 0xd0, 0x01, 0x0a, 0x05, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x73, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x5f, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x12, 0x38, 0x0a, 0x18, 0x74, 0x69, 0x6d, 0x69, 0x6e, 0x67,
	0x5f, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69,
	0x6e, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x16, 0x74, 0x69, 0x6d, 0x69, 0x6e, 0x67,
	0x50, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67,
	0x12, 0x36, 0x0a, 0x17, 0x74, 0x69, 0x6d, 0x69, 0x6e, 0x67, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x5f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x15, 0x74, 0x69, 0x6d, 0x69, 0x6e, 0x67, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x47, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x97, 0x01, 0x0a, 0x0b, 0x43, 0x68, 0x61,
	0x74, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x50, 0x72, 0x6f, 0x6d,
	0x70, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x50, 0x72, 0x6f, 0x6d, 0x70, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74,
	0x12, 0x24, 0x0a, 0x0d, 0x52, 0x65, 0x70, 0x65, 0x61, 0x74, 0x50, 0x65, 0x6e, 0x61, 0x6c, 0x74,
	0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0d, 0x52, 0x65, 0x70, 0x65, 0x61, 0x74, 0x50,
	0x65, 0x6e, 0x61, 0x6c, 0x74, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x42, 0x61, 0x63, 0x6b, 0x65, 0x6e,
	0x64, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x42, 0x61, 0x63, 0x6b, 0x65, 0x6e,
	0x64, 0x73, 0x22, 0xa4, 0x01, 0x0a, 0x0f, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x55, 0x73, 0x61,
	0x67, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x3d, 0x0a, 0x09,
	0x62, 0x72, 0x65, 0x61, 0x6b, 0x64, 0x6f, 0x77, 0x6e, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1f, 0x2e, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x55, 0x73, 0x61, 0x67, 0x65, 0x44, 0x61, 0x74,
	0x61, 0x2e, 0x42, 0x72, 0x65, 0x61, 0x6b, 0x64, 0x6f, 0x77, 0x6e, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x09, 0x62, 0x72, 0x65, 0x61, 0x6b, 0x64, 0x6f, 0x77, 0x6e, 0x1a, 0x3c, 0x0a, 0x0e, 0x42,
	0x72, 0x65, 0x61, 0x6b, 0x64, 0x6f, 0x77, 0x6e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xae, 0x01, 0x0a, 0x0f, 0x42, 0x61,
	0x63, 0x6b, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a,
	0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x42,
	0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x28, 0x0a, 0x06, 0x6d,
	0x65, 0x6d, 0x6f, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x4d, 0x65,
	0x6d, 0x6f, 0x72, 0x79, 0x55, 0x73, 0x61, 0x67, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x06, 0x6d,
	0x65, 0x6d, 0x6f, 0x72, 0x79, 0x22, 0x43, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x11,
	0x0a, 0x0d, 0x55, 0x4e, 0x49, 0x4e, 0x49, 0x54, 0x49, 0x41, 0x4c, 0x49, 0x5a, 0x45, 0x44, 0x10,
	0x00, 0x12, 0x08, 0x0a, 0x04, 0x42, 0x55, 0x53, 0x59, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x52,
	0x45, 0x41, 0x44, 0x59, 0x10, 0x02, 0x12, 0x12, 0x0a, 0x05, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x01, 0x32, 0x66, 0x0a, 0x07, 0x42, 0x61,
	0x63, 0x6b, 0x65, 0x6e, 0x64, 0x12, 0x26, 0x0a, 0x0a, 0x43, 0x68, 0x61, 0x74, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x12, 0x0c, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x1a, 0x06, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x30, 0x01, 0x12, 0x33, 0x0a,
	0x0c, 0x53, 0x74, 0x61, 0x72, 0x74, 0x42, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x12, 0x0f, 0x2e,
	0x42, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x10,
	0x2e, 0x42, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0x33, 0x5a, 0x31, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x61, 0x74, 0x6f, 0x76, 0x61, 0x6c,
	0x2f, 0x61, 0x69, 0x72, 0x74, 0x48, 0x69, 0x76, 0x65, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e,
	0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_backend_proto_rawDescOnce sync.Once
	file_backend_proto_rawDescData []byte
)

func file_backend_proto_rawDescGZIP() []byte {
	file_backend_proto_rawDescOnce.Do(func() {
		file_backend_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_backend_proto_rawDesc), len(file_backend_proto_rawDesc)))
	})
	return file_backend_proto_rawDescData
}

var file_backend_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_backend_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_backend_proto_goTypes = []any{
	(BackendResponse_State)(0), // 0: BackendResponse.State
	(*BackendOptions)(nil),     // 1: BackendOptions
	(*Reply)(nil),              // 2: Reply
	(*ChatOptions)(nil),        // 3: ChatOptions
	(*MemoryUsageData)(nil),    // 4: MemoryUsageData
	(*BackendResponse)(nil),    // 5: BackendResponse
	nil,                        // 6: MemoryUsageData.BreakdownEntry
}
var file_backend_proto_depIdxs = []int32{
	6, // 0: MemoryUsageData.breakdown:type_name -> MemoryUsageData.BreakdownEntry
	0, // 1: BackendResponse.state:type_name -> BackendResponse.State
	4, // 2: BackendResponse.memory:type_name -> MemoryUsageData
	3, // 3: Backend.ChatStream:input_type -> ChatOptions
	1, // 4: Backend.StartBackend:input_type -> BackendOptions
	2, // 5: Backend.ChatStream:output_type -> Reply
	5, // 6: Backend.StartBackend:output_type -> BackendResponse
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_backend_proto_init() }
func file_backend_proto_init() {
	if File_backend_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_backend_proto_rawDesc), len(file_backend_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_backend_proto_goTypes,
		DependencyIndexes: file_backend_proto_depIdxs,
		EnumInfos:         file_backend_proto_enumTypes,
		MessageInfos:      file_backend_proto_msgTypes,
	}.Build()
	File_backend_proto = out.File
	file_backend_proto_goTypes = nil
	file_backend_proto_depIdxs = nil
}
