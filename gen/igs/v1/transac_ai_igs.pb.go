// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: igs/v1/transac_ai_igs.proto

package igs_v1

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

// Request parameters for GenerateInsights
type GenerateInsightsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReqId                   string `protobuf:"bytes,1,opt,name=req_id,json=reqId,proto3" json:"req_id,omitempty"`
	ClientId                string `protobuf:"bytes,2,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	PromptId                int32  `protobuf:"varint,3,opt,name=prompt_id,json=promptId,proto3" json:"prompt_id,omitempty"`
	RecordsSourceId         string `protobuf:"bytes,4,opt,name=records_source_id,json=recordsSourceId,proto3" json:"records_source_id,omitempty"`
	PromptTemplatesSourceId string `protobuf:"bytes,5,opt,name=prompt_templates_source_id,json=promptTemplatesSourceId,proto3" json:"prompt_templates_source_id,omitempty"`
	FromTime                string `protobuf:"bytes,6,opt,name=from_time,json=fromTime,proto3" json:"from_time,omitempty"`
	ToTime                  string `protobuf:"bytes,7,opt,name=to_time,json=toTime,proto3" json:"to_time,omitempty"`
}

func (x *GenerateInsightsRequest) Reset() {
	*x = GenerateInsightsRequest{}
	mi := &file_igs_v1_transac_ai_igs_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GenerateInsightsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateInsightsRequest) ProtoMessage() {}

func (x *GenerateInsightsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_igs_v1_transac_ai_igs_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateInsightsRequest.ProtoReflect.Descriptor instead.
func (*GenerateInsightsRequest) Descriptor() ([]byte, []int) {
	return file_igs_v1_transac_ai_igs_proto_rawDescGZIP(), []int{0}
}

func (x *GenerateInsightsRequest) GetReqId() string {
	if x != nil {
		return x.ReqId
	}
	return ""
}

func (x *GenerateInsightsRequest) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *GenerateInsightsRequest) GetPromptId() int32 {
	if x != nil {
		return x.PromptId
	}
	return 0
}

func (x *GenerateInsightsRequest) GetRecordsSourceId() string {
	if x != nil {
		return x.RecordsSourceId
	}
	return ""
}

func (x *GenerateInsightsRequest) GetPromptTemplatesSourceId() string {
	if x != nil {
		return x.PromptTemplatesSourceId
	}
	return ""
}

func (x *GenerateInsightsRequest) GetFromTime() string {
	if x != nil {
		return x.FromTime
	}
	return ""
}

func (x *GenerateInsightsRequest) GetToTime() string {
	if x != nil {
		return x.ToTime
	}
	return ""
}

type GenerateInsightsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Received bool `protobuf:"varint,1,opt,name=received,proto3" json:"received,omitempty"`
}

func (x *GenerateInsightsResponse) Reset() {
	*x = GenerateInsightsResponse{}
	mi := &file_igs_v1_transac_ai_igs_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GenerateInsightsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateInsightsResponse) ProtoMessage() {}

func (x *GenerateInsightsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_igs_v1_transac_ai_igs_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateInsightsResponse.ProtoReflect.Descriptor instead.
func (*GenerateInsightsResponse) Descriptor() ([]byte, []int) {
	return file_igs_v1_transac_ai_igs_proto_rawDescGZIP(), []int{1}
}

func (x *GenerateInsightsResponse) GetReceived() bool {
	if x != nil {
		return x.Received
	}
	return false
}

var File_igs_v1_transac_ai_igs_proto protoreflect.FileDescriptor

var file_igs_v1_transac_ai_igs_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x69, 0x67, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x5f, 0x61, 0x69, 0x5f, 0x69, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x69,
	0x67, 0x73, 0x2e, 0x76, 0x31, 0x22, 0x89, 0x02, 0x0a, 0x17, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x74, 0x65, 0x49, 0x6e, 0x73, 0x69, 0x67, 0x68, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x15, 0x0a, 0x06, 0x72, 0x65, 0x71, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x72, 0x65, 0x71, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74,
	0x49, 0x64, 0x12, 0x2a, 0x0a, 0x11, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x5f, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x72,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x12, 0x3b,
	0x0a, 0x1a, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x5f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74,
	0x65, 0x73, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x17, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61,
	0x74, 0x65, 0x73, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x66,
	0x72, 0x6f, 0x6d, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x66, 0x72, 0x6f, 0x6d, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x6f, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x6f, 0x54, 0x69, 0x6d,
	0x65, 0x22, 0x36, 0x0a, 0x18, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x73,
	0x69, 0x67, 0x68, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x08, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x32, 0x63, 0x0a, 0x0a, 0x49, 0x47, 0x53,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x55, 0x0a, 0x10, 0x47, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x74, 0x65, 0x49, 0x6e, 0x73, 0x69, 0x67, 0x68, 0x74, 0x73, 0x12, 0x1f, 0x2e, 0x69, 0x67,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x73,
	0x69, 0x67, 0x68, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x69,
	0x67, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x49, 0x6e,
	0x73, 0x69, 0x67, 0x68, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x21,
	0x5a, 0x1f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x61, 0x69, 0x2d, 0x77, 0x6d, 0x73, 0x2f,
	0x67, 0x65, 0x6e, 0x2f, 0x69, 0x67, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x69, 0x67, 0x73, 0x5f, 0x76,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_igs_v1_transac_ai_igs_proto_rawDescOnce sync.Once
	file_igs_v1_transac_ai_igs_proto_rawDescData = file_igs_v1_transac_ai_igs_proto_rawDesc
)

func file_igs_v1_transac_ai_igs_proto_rawDescGZIP() []byte {
	file_igs_v1_transac_ai_igs_proto_rawDescOnce.Do(func() {
		file_igs_v1_transac_ai_igs_proto_rawDescData = protoimpl.X.CompressGZIP(file_igs_v1_transac_ai_igs_proto_rawDescData)
	})
	return file_igs_v1_transac_ai_igs_proto_rawDescData
}

var file_igs_v1_transac_ai_igs_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_igs_v1_transac_ai_igs_proto_goTypes = []any{
	(*GenerateInsightsRequest)(nil),  // 0: igs.v1.GenerateInsightsRequest
	(*GenerateInsightsResponse)(nil), // 1: igs.v1.GenerateInsightsResponse
}
var file_igs_v1_transac_ai_igs_proto_depIdxs = []int32{
	0, // 0: igs.v1.IGSService.GenerateInsights:input_type -> igs.v1.GenerateInsightsRequest
	1, // 1: igs.v1.IGSService.GenerateInsights:output_type -> igs.v1.GenerateInsightsResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_igs_v1_transac_ai_igs_proto_init() }
func file_igs_v1_transac_ai_igs_proto_init() {
	if File_igs_v1_transac_ai_igs_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_igs_v1_transac_ai_igs_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_igs_v1_transac_ai_igs_proto_goTypes,
		DependencyIndexes: file_igs_v1_transac_ai_igs_proto_depIdxs,
		MessageInfos:      file_igs_v1_transac_ai_igs_proto_msgTypes,
	}.Build()
	File_igs_v1_transac_ai_igs_proto = out.File
	file_igs_v1_transac_ai_igs_proto_rawDesc = nil
	file_igs_v1_transac_ai_igs_proto_goTypes = nil
	file_igs_v1_transac_ai_igs_proto_depIdxs = nil
}
