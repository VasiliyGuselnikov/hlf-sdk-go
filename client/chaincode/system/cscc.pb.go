// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.7.1
// source: cscc.proto

package system

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	common "github.com/hyperledger/fabric-protos-go/common"
	peer "github.com/hyperledger/fabric-protos-go/peer"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type JoinChainRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Channel      string        `protobuf:"bytes,1,opt,name=channel,proto3" json:"channel,omitempty"`
	GenesisBlock *common.Block `protobuf:"bytes,2,opt,name=genesis_block,json=genesisBlock,proto3" json:"genesis_block,omitempty"`
}

func (x *JoinChainRequest) Reset() {
	*x = JoinChainRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cscc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JoinChainRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinChainRequest) ProtoMessage() {}

func (x *JoinChainRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cscc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinChainRequest.ProtoReflect.Descriptor instead.
func (*JoinChainRequest) Descriptor() ([]byte, []int) {
	return file_cscc_proto_rawDescGZIP(), []int{0}
}

func (x *JoinChainRequest) GetChannel() string {
	if x != nil {
		return x.Channel
	}
	return ""
}

func (x *JoinChainRequest) GetGenesisBlock() *common.Block {
	if x != nil {
		return x.GenesisBlock
	}
	return nil
}

type GetConfigBlockRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Channel string `protobuf:"bytes,1,opt,name=channel,proto3" json:"channel,omitempty"`
}

func (x *GetConfigBlockRequest) Reset() {
	*x = GetConfigBlockRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cscc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetConfigBlockRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetConfigBlockRequest) ProtoMessage() {}

func (x *GetConfigBlockRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cscc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetConfigBlockRequest.ProtoReflect.Descriptor instead.
func (*GetConfigBlockRequest) Descriptor() ([]byte, []int) {
	return file_cscc_proto_rawDescGZIP(), []int{1}
}

func (x *GetConfigBlockRequest) GetChannel() string {
	if x != nil {
		return x.Channel
	}
	return ""
}

type GetChannelConfigRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Channel string `protobuf:"bytes,1,opt,name=channel,proto3" json:"channel,omitempty"`
}

func (x *GetChannelConfigRequest) Reset() {
	*x = GetChannelConfigRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cscc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetChannelConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetChannelConfigRequest) ProtoMessage() {}

func (x *GetChannelConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cscc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetChannelConfigRequest.ProtoReflect.Descriptor instead.
func (*GetChannelConfigRequest) Descriptor() ([]byte, []int) {
	return file_cscc_proto_rawDescGZIP(), []int{2}
}

func (x *GetChannelConfigRequest) GetChannel() string {
	if x != nil {
		return x.Channel
	}
	return ""
}

var File_cscc_proto protoreflect.FileDescriptor

var file_cscc_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x63, 0x73, 0x63, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x68, 0x6c,
	0x66, 0x73, 0x64, 0x6b, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x63, 0x68, 0x61, 0x69,
	0x6e, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x1a, 0x13, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x15, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x74, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x70, 0x65, 0x65, 0x72, 0x2f, 0x71,
	0x75, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x60, 0x0a, 0x10, 0x4a, 0x6f, 0x69, 0x6e, 0x43, 0x68, 0x61,
	0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x68, 0x61,
	0x6e, 0x6e, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x12, 0x32, 0x0a, 0x0d, 0x67, 0x65, 0x6e, 0x65, 0x73, 0x69, 0x73, 0x5f, 0x62,
	0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x0c, 0x67, 0x65, 0x6e, 0x65, 0x73,
	0x69, 0x73, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x22, 0x31, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x18, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x22, 0x33, 0x0a, 0x17, 0x47, 0x65,
	0x74, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x32,
	0xd8, 0x03, 0x0a, 0x0b, 0x43, 0x53, 0x43, 0x43, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x71, 0x0a, 0x09, 0x4a, 0x6f, 0x69, 0x6e, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x12, 0x30, 0x2e, 0x68,
	0x6c, 0x66, 0x73, 0x64, 0x6b, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x63, 0x68, 0x61,
	0x69, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x4a, 0x6f,
	0x69, 0x6e, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x22, 0x0f,
	0x2f, 0x63, 0x73, 0x63, 0x63, 0x2f, 0x6a, 0x6f, 0x69, 0x6e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x3a,
	0x01, 0x2a, 0x12, 0x59, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x14, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x12,
	0x0c, 0x2f, 0x63, 0x73, 0x63, 0x63, 0x2f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x73, 0x12, 0x76, 0x0a,
	0x0e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12,
	0x35, 0x2e, 0x68, 0x6c, 0x66, 0x73, 0x64, 0x6b, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e,
	0x63, 0x68, 0x61, 0x69, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d,
	0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x22, 0x1e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x12, 0x16, 0x2f,
	0x63, 0x73, 0x63, 0x63, 0x2f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x73, 0x2f, 0x7b, 0x63, 0x68, 0x61,
	0x6e, 0x6e, 0x65, 0x6c, 0x7d, 0x12, 0x82, 0x01, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x43, 0x68, 0x61,
	0x6e, 0x6e, 0x65, 0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x37, 0x2e, 0x68, 0x6c, 0x66,
	0x73, 0x64, 0x6b, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e,
	0x63, 0x6f, 0x64, 0x65, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x47, 0x65, 0x74, 0x43,
	0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x22, 0x25, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1f, 0x12, 0x1d, 0x2f, 0x63, 0x73,
	0x63, 0x63, 0x2f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x73, 0x2f, 0x7b, 0x63, 0x68, 0x61, 0x6e, 0x6e,
	0x65, 0x6c, 0x7d, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42, 0x20, 0x5a, 0x1e, 0x68, 0x6c,
	0x66, 0x73, 0x64, 0x6b, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x63, 0x68, 0x61, 0x69,
	0x6e, 0x63, 0x6f, 0x64, 0x65, 0x2f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cscc_proto_rawDescOnce sync.Once
	file_cscc_proto_rawDescData = file_cscc_proto_rawDesc
)

func file_cscc_proto_rawDescGZIP() []byte {
	file_cscc_proto_rawDescOnce.Do(func() {
		file_cscc_proto_rawDescData = protoimpl.X.CompressGZIP(file_cscc_proto_rawDescData)
	})
	return file_cscc_proto_rawDescData
}

var file_cscc_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_cscc_proto_goTypes = []interface{}{
	(*JoinChainRequest)(nil),          // 0: hlfsdk.client.chaincode.system.JoinChainRequest
	(*GetConfigBlockRequest)(nil),     // 1: hlfsdk.client.chaincode.system.GetConfigBlockRequest
	(*GetChannelConfigRequest)(nil),   // 2: hlfsdk.client.chaincode.system.GetChannelConfigRequest
	(*common.Block)(nil),              // 3: common.Block
	(*empty.Empty)(nil),               // 4: google.protobuf.Empty
	(*peer.ChannelQueryResponse)(nil), // 5: protos.ChannelQueryResponse
	(*common.Config)(nil),             // 6: common.Config
}
var file_cscc_proto_depIdxs = []int32{
	3, // 0: hlfsdk.client.chaincode.system.JoinChainRequest.genesis_block:type_name -> common.Block
	0, // 1: hlfsdk.client.chaincode.system.CSCCService.JoinChain:input_type -> hlfsdk.client.chaincode.system.JoinChainRequest
	4, // 2: hlfsdk.client.chaincode.system.CSCCService.GetChannels:input_type -> google.protobuf.Empty
	1, // 3: hlfsdk.client.chaincode.system.CSCCService.GetConfigBlock:input_type -> hlfsdk.client.chaincode.system.GetConfigBlockRequest
	2, // 4: hlfsdk.client.chaincode.system.CSCCService.GetChannelConfig:input_type -> hlfsdk.client.chaincode.system.GetChannelConfigRequest
	4, // 5: hlfsdk.client.chaincode.system.CSCCService.JoinChain:output_type -> google.protobuf.Empty
	5, // 6: hlfsdk.client.chaincode.system.CSCCService.GetChannels:output_type -> protos.ChannelQueryResponse
	3, // 7: hlfsdk.client.chaincode.system.CSCCService.GetConfigBlock:output_type -> common.Block
	6, // 8: hlfsdk.client.chaincode.system.CSCCService.GetChannelConfig:output_type -> common.Config
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_cscc_proto_init() }
func file_cscc_proto_init() {
	if File_cscc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cscc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JoinChainRequest); i {
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
		file_cscc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetConfigBlockRequest); i {
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
		file_cscc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetChannelConfigRequest); i {
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
			RawDescriptor: file_cscc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cscc_proto_goTypes,
		DependencyIndexes: file_cscc_proto_depIdxs,
		MessageInfos:      file_cscc_proto_msgTypes,
	}.Build()
	File_cscc_proto = out.File
	file_cscc_proto_rawDesc = nil
	file_cscc_proto_goTypes = nil
	file_cscc_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CSCCServiceClient is the client API for CSCCService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CSCCServiceClient interface {
	// GetChainInfo allows joining channel using presented genesis block
	JoinChain(ctx context.Context, in *JoinChainRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	GetChannels(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*peer.ChannelQueryResponse, error)
	// GetConfigBlock returns genesis block of channel
	GetConfigBlock(ctx context.Context, in *GetConfigBlockRequest, opts ...grpc.CallOption) (*common.Block, error)
	// GetChannelConfig returns channel configuration
	GetChannelConfig(ctx context.Context, in *GetChannelConfigRequest, opts ...grpc.CallOption) (*common.Config, error)
}

type cSCCServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCSCCServiceClient(cc grpc.ClientConnInterface) CSCCServiceClient {
	return &cSCCServiceClient{cc}
}

func (c *cSCCServiceClient) JoinChain(ctx context.Context, in *JoinChainRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/hlfsdk.client.chaincode.system.CSCCService/JoinChain", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cSCCServiceClient) GetChannels(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*peer.ChannelQueryResponse, error) {
	out := new(peer.ChannelQueryResponse)
	err := c.cc.Invoke(ctx, "/hlfsdk.client.chaincode.system.CSCCService/GetChannels", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cSCCServiceClient) GetConfigBlock(ctx context.Context, in *GetConfigBlockRequest, opts ...grpc.CallOption) (*common.Block, error) {
	out := new(common.Block)
	err := c.cc.Invoke(ctx, "/hlfsdk.client.chaincode.system.CSCCService/GetConfigBlock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cSCCServiceClient) GetChannelConfig(ctx context.Context, in *GetChannelConfigRequest, opts ...grpc.CallOption) (*common.Config, error) {
	out := new(common.Config)
	err := c.cc.Invoke(ctx, "/hlfsdk.client.chaincode.system.CSCCService/GetChannelConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CSCCServiceServer is the server API for CSCCService service.
type CSCCServiceServer interface {
	// GetChainInfo allows joining channel using presented genesis block
	JoinChain(context.Context, *JoinChainRequest) (*empty.Empty, error)
	GetChannels(context.Context, *empty.Empty) (*peer.ChannelQueryResponse, error)
	// GetConfigBlock returns genesis block of channel
	GetConfigBlock(context.Context, *GetConfigBlockRequest) (*common.Block, error)
	// GetChannelConfig returns channel configuration
	GetChannelConfig(context.Context, *GetChannelConfigRequest) (*common.Config, error)
}

// UnimplementedCSCCServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCSCCServiceServer struct {
}

func (*UnimplementedCSCCServiceServer) JoinChain(context.Context, *JoinChainRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JoinChain not implemented")
}
func (*UnimplementedCSCCServiceServer) GetChannels(context.Context, *empty.Empty) (*peer.ChannelQueryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChannels not implemented")
}
func (*UnimplementedCSCCServiceServer) GetConfigBlock(context.Context, *GetConfigBlockRequest) (*common.Block, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConfigBlock not implemented")
}
func (*UnimplementedCSCCServiceServer) GetChannelConfig(context.Context, *GetChannelConfigRequest) (*common.Config, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChannelConfig not implemented")
}

func RegisterCSCCServiceServer(s *grpc.Server, srv CSCCServiceServer) {
	s.RegisterService(&_CSCCService_serviceDesc, srv)
}

func _CSCCService_JoinChain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JoinChainRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CSCCServiceServer).JoinChain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hlfsdk.client.chaincode.system.CSCCService/JoinChain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CSCCServiceServer).JoinChain(ctx, req.(*JoinChainRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CSCCService_GetChannels_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CSCCServiceServer).GetChannels(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hlfsdk.client.chaincode.system.CSCCService/GetChannels",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CSCCServiceServer).GetChannels(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _CSCCService_GetConfigBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetConfigBlockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CSCCServiceServer).GetConfigBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hlfsdk.client.chaincode.system.CSCCService/GetConfigBlock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CSCCServiceServer).GetConfigBlock(ctx, req.(*GetConfigBlockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CSCCService_GetChannelConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetChannelConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CSCCServiceServer).GetChannelConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hlfsdk.client.chaincode.system.CSCCService/GetChannelConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CSCCServiceServer).GetChannelConfig(ctx, req.(*GetChannelConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CSCCService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "hlfsdk.client.chaincode.system.CSCCService",
	HandlerType: (*CSCCServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "JoinChain",
			Handler:    _CSCCService_JoinChain_Handler,
		},
		{
			MethodName: "GetChannels",
			Handler:    _CSCCService_GetChannels_Handler,
		},
		{
			MethodName: "GetConfigBlock",
			Handler:    _CSCCService_GetConfigBlock_Handler,
		},
		{
			MethodName: "GetChannelConfig",
			Handler:    _CSCCService_GetChannelConfig_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cscc.proto",
}
