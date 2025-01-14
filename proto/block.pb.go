// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.15.8
// source: block.proto

package proto

import (
	common "github.com/hyperledger/fabric-protos-go/common"
	kvrwset "github.com/hyperledger/fabric-protos-go/ledger/rwset/kvrwset"
	msp "github.com/hyperledger/fabric-protos-go/msp"
	peer "github.com/hyperledger/fabric-protos-go/peer"
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

type Block struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header            *common.BlockHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Envelopes         []*Envelope         `protobuf:"bytes,2,rep,name=envelopes,proto3" json:"envelopes,omitempty"`
	OrdererSignatures []*OrdererSignature `protobuf:"bytes,3,rep,name=orderer_signatures,json=ordererSignatures,proto3" json:"orderer_signatures,omitempty"`
}

func (x *Block) Reset() {
	*x = Block{}
	if protoimpl.UnsafeEnabled {
		mi := &file_block_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Block) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Block) ProtoMessage() {}

func (x *Block) ProtoReflect() protoreflect.Message {
	mi := &file_block_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Block.ProtoReflect.Descriptor instead.
func (*Block) Descriptor() ([]byte, []int) {
	return file_block_proto_rawDescGZIP(), []int{0}
}

func (x *Block) GetHeader() *common.BlockHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *Block) GetEnvelopes() []*Envelope {
	if x != nil {
		return x.Envelopes
	}
	return nil
}

func (x *Block) GetOrdererSignatures() []*OrdererSignature {
	if x != nil {
		return x.OrdererSignatures
	}
	return nil
}

type Envelope struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Signature      []byte                `protobuf:"bytes,1,opt,name=signature,proto3" json:"signature,omitempty"`
	ChannelHeader  *common.ChannelHeader `protobuf:"bytes,2,opt,name=channel_header,json=channelHeader,proto3" json:"channel_header,omitempty"`
	ValidationCode peer.TxValidationCode `protobuf:"varint,3,opt,name=validation_code,json=validationCode,proto3,enum=protos.TxValidationCode" json:"validation_code,omitempty"`
	Transaction    *Transaction          `protobuf:"bytes,4,opt,name=transaction,proto3" json:"transaction,omitempty"`
	ChannelConfig  *ChannelConfig        `protobuf:"bytes,5,opt,name=channel_config,json=channelConfig,proto3" json:"channel_config,omitempty"`
}

func (x *Envelope) Reset() {
	*x = Envelope{}
	if protoimpl.UnsafeEnabled {
		mi := &file_block_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Envelope) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Envelope) ProtoMessage() {}

func (x *Envelope) ProtoReflect() protoreflect.Message {
	mi := &file_block_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Envelope.ProtoReflect.Descriptor instead.
func (*Envelope) Descriptor() ([]byte, []int) {
	return file_block_proto_rawDescGZIP(), []int{1}
}

func (x *Envelope) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

func (x *Envelope) GetChannelHeader() *common.ChannelHeader {
	if x != nil {
		return x.ChannelHeader
	}
	return nil
}

func (x *Envelope) GetValidationCode() peer.TxValidationCode {
	if x != nil {
		return x.ValidationCode
	}
	return peer.TxValidationCode(0)
}

func (x *Envelope) GetTransaction() *Transaction {
	if x != nil {
		return x.Transaction
	}
	return nil
}

func (x *Envelope) GetChannelConfig() *ChannelConfig {
	if x != nil {
		return x.ChannelConfig
	}
	return nil
}

type Transaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Actions         []*TransactionAction    `protobuf:"bytes,1,rep,name=actions,proto3" json:"actions,omitempty"`
	CreatorIdentity *msp.SerializedIdentity `protobuf:"bytes,2,opt,name=creator_identity,json=creatorIdentity,proto3" json:"creator_identity,omitempty"`
}

func (x *Transaction) Reset() {
	*x = Transaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_block_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Transaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Transaction) ProtoMessage() {}

func (x *Transaction) ProtoReflect() protoreflect.Message {
	mi := &file_block_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Transaction.ProtoReflect.Descriptor instead.
func (*Transaction) Descriptor() ([]byte, []int) {
	return file_block_proto_rawDescGZIP(), []int{2}
}

func (x *Transaction) GetActions() []*TransactionAction {
	if x != nil {
		return x.Actions
	}
	return nil
}

func (x *Transaction) GetCreatorIdentity() *msp.SerializedIdentity {
	if x != nil {
		return x.CreatorIdentity
	}
	return nil
}

type TransactionAction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Event                   *peer.ChaincodeEvent          `protobuf:"bytes,1,opt,name=event,proto3" json:"event,omitempty"`
	Endorsers               []*msp.SerializedIdentity     `protobuf:"bytes,2,rep,name=endorsers,proto3" json:"endorsers,omitempty"`
	ReadWriteSets           []*kvrwset.KVRWSet            `protobuf:"bytes,3,rep,name=read_write_sets,json=readWriteSets,proto3" json:"read_write_sets,omitempty"`
	ChaincodeInvocationSpec *peer.ChaincodeInvocationSpec `protobuf:"bytes,4,opt,name=chaincode_invocation_spec,json=chaincodeInvocationSpec,proto3" json:"chaincode_invocation_spec,omitempty"`
	CreatorIdentity         *msp.SerializedIdentity       `protobuf:"bytes,5,opt,name=creator_identity,json=creatorIdentity,proto3" json:"creator_identity,omitempty"`
	Payload                 []byte                        `protobuf:"bytes,6,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *TransactionAction) Reset() {
	*x = TransactionAction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_block_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionAction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionAction) ProtoMessage() {}

func (x *TransactionAction) ProtoReflect() protoreflect.Message {
	mi := &file_block_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionAction.ProtoReflect.Descriptor instead.
func (*TransactionAction) Descriptor() ([]byte, []int) {
	return file_block_proto_rawDescGZIP(), []int{3}
}

func (x *TransactionAction) GetEvent() *peer.ChaincodeEvent {
	if x != nil {
		return x.Event
	}
	return nil
}

func (x *TransactionAction) GetEndorsers() []*msp.SerializedIdentity {
	if x != nil {
		return x.Endorsers
	}
	return nil
}

func (x *TransactionAction) GetReadWriteSets() []*kvrwset.KVRWSet {
	if x != nil {
		return x.ReadWriteSets
	}
	return nil
}

func (x *TransactionAction) GetChaincodeInvocationSpec() *peer.ChaincodeInvocationSpec {
	if x != nil {
		return x.ChaincodeInvocationSpec
	}
	return nil
}

func (x *TransactionAction) GetCreatorIdentity() *msp.SerializedIdentity {
	if x != nil {
		return x.CreatorIdentity
	}
	return nil
}

func (x *TransactionAction) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

type OrdererSignature struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Identity  *msp.SerializedIdentity `protobuf:"bytes,1,opt,name=identity,proto3" json:"identity,omitempty"`
	Signature []byte                  `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (x *OrdererSignature) Reset() {
	*x = OrdererSignature{}
	if protoimpl.UnsafeEnabled {
		mi := &file_block_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrdererSignature) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrdererSignature) ProtoMessage() {}

func (x *OrdererSignature) ProtoReflect() protoreflect.Message {
	mi := &file_block_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrdererSignature.ProtoReflect.Descriptor instead.
func (*OrdererSignature) Descriptor() ([]byte, []int) {
	return file_block_proto_rawDescGZIP(), []int{4}
}

func (x *OrdererSignature) GetIdentity() *msp.SerializedIdentity {
	if x != nil {
		return x.Identity
	}
	return nil
}

func (x *OrdererSignature) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

var File_block_proto protoreflect.FileDescriptor

var file_block_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x68,
	0x6c, 0x66, 0x73, 0x64, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x63, 0x68, 0x61,
	0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2f, 0x72, 0x77, 0x73, 0x65,
	0x74, 0x2f, 0x6b, 0x76, 0x72, 0x77, 0x73, 0x65, 0x74, 0x2f, 0x6b, 0x76, 0x5f, 0x72, 0x77, 0x73,
	0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x6d, 0x73, 0x70, 0x2f, 0x69, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14,
	0x70, 0x65, 0x65, 0x72, 0x2f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x70, 0x65, 0x65, 0x72, 0x2f, 0x63, 0x68, 0x61, 0x69, 0x6e,
	0x63, 0x6f, 0x64, 0x65, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x16, 0x70, 0x65, 0x65, 0x72, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb9, 0x01, 0x0a, 0x05, 0x42, 0x6c, 0x6f,
	0x63, 0x6b, 0x12, 0x2b, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x13, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x42, 0x6c, 0x6f, 0x63,
	0x6b, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12,
	0x34, 0x0a, 0x09, 0x65, 0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x16, 0x2e, 0x68, 0x6c, 0x66, 0x73, 0x64, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x45, 0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x52, 0x09, 0x65, 0x6e, 0x76, 0x65,
	0x6c, 0x6f, 0x70, 0x65, 0x73, 0x12, 0x4d, 0x0a, 0x12, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x65, 0x72,
	0x5f, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1e, 0x2e, 0x68, 0x6c, 0x66, 0x73, 0x64, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x65, 0x72, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x52, 0x11, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x65, 0x72, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x73, 0x22, 0xaa, 0x02, 0x0a, 0x08, 0x45, 0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70,
	0x65, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12,
	0x3c, 0x0a, 0x0e, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x0d,
	0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x41, 0x0a,
	0x0f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e,
	0x54, 0x78, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65,
	0x52, 0x0e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65,
	0x12, 0x3b, 0x0a, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x68, 0x6c, 0x66, 0x73, 0x64, 0x6b, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x42, 0x0a,
	0x0e, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x68, 0x6c, 0x66, 0x73, 0x64, 0x6b, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x52, 0x0d, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x22, 0x8c, 0x01, 0x0a, 0x0b, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x39, 0x0a, 0x07, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x68, 0x6c, 0x66, 0x73, 0x64, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x07, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x42, 0x0a, 0x10,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x53, 0x65, 0x72,
	0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52,
	0x0f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x22, 0xed, 0x02, 0x0a, 0x11, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2c, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x43,
	0x68, 0x61, 0x69, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x05, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x12, 0x35, 0x0a, 0x09, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x73, 0x65, 0x72,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x53, 0x65,
	0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x52, 0x09, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x73, 0x65, 0x72, 0x73, 0x12, 0x38, 0x0a, 0x0f, 0x72,
	0x65, 0x61, 0x64, 0x5f, 0x77, 0x72, 0x69, 0x74, 0x65, 0x5f, 0x73, 0x65, 0x74, 0x73, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6b, 0x76, 0x72, 0x77, 0x73, 0x65, 0x74, 0x2e, 0x4b,
	0x56, 0x52, 0x57, 0x53, 0x65, 0x74, 0x52, 0x0d, 0x72, 0x65, 0x61, 0x64, 0x57, 0x72, 0x69, 0x74,
	0x65, 0x53, 0x65, 0x74, 0x73, 0x12, 0x5b, 0x0a, 0x19, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x63, 0x6f,
	0x64, 0x65, 0x5f, 0x69, 0x6e, 0x76, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x70,
	0x65, 0x63, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2e, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x76, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x70, 0x65, 0x63, 0x52, 0x17, 0x63, 0x68, 0x61, 0x69, 0x6e,
	0x63, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x76, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x70,
	0x65, 0x63, 0x12, 0x42, 0x0a, 0x10, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x69, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6d,
	0x73, 0x70, 0x2e, 0x53, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x49, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x0f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64,
	0x22, 0x65, 0x0a, 0x10, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x65, 0x72, 0x53, 0x69, 0x67, 0x6e, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x12, 0x33, 0x0a, 0x08, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x53, 0x65, 0x72,
	0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52,
	0x08, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67,
	0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x73, 0x69,
	0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x42, 0x28, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x74, 0x6f, 0x6d, 0x79, 0x7a, 0x65, 0x2d, 0x72, 0x75,
	0x2f, 0x68, 0x6c, 0x66, 0x2d, 0x73, 0x64, 0x6b, 0x2d, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_block_proto_rawDescOnce sync.Once
	file_block_proto_rawDescData = file_block_proto_rawDesc
)

func file_block_proto_rawDescGZIP() []byte {
	file_block_proto_rawDescOnce.Do(func() {
		file_block_proto_rawDescData = protoimpl.X.CompressGZIP(file_block_proto_rawDescData)
	})
	return file_block_proto_rawDescData
}

var file_block_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_block_proto_goTypes = []interface{}{
	(*Block)(nil),                        // 0: hlfsdk.proto.Block
	(*Envelope)(nil),                     // 1: hlfsdk.proto.Envelope
	(*Transaction)(nil),                  // 2: hlfsdk.proto.Transaction
	(*TransactionAction)(nil),            // 3: hlfsdk.proto.TransactionAction
	(*OrdererSignature)(nil),             // 4: hlfsdk.proto.OrdererSignature
	(*common.BlockHeader)(nil),           // 5: common.BlockHeader
	(*common.ChannelHeader)(nil),         // 6: common.ChannelHeader
	(peer.TxValidationCode)(0),           // 7: protos.TxValidationCode
	(*ChannelConfig)(nil),                // 8: hlfsdk.proto.ChannelConfig
	(*msp.SerializedIdentity)(nil),       // 9: msp.SerializedIdentity
	(*peer.ChaincodeEvent)(nil),          // 10: protos.ChaincodeEvent
	(*kvrwset.KVRWSet)(nil),              // 11: kvrwset.KVRWSet
	(*peer.ChaincodeInvocationSpec)(nil), // 12: protos.ChaincodeInvocationSpec
}
var file_block_proto_depIdxs = []int32{
	5,  // 0: hlfsdk.proto.Block.header:type_name -> common.BlockHeader
	1,  // 1: hlfsdk.proto.Block.envelopes:type_name -> hlfsdk.proto.Envelope
	4,  // 2: hlfsdk.proto.Block.orderer_signatures:type_name -> hlfsdk.proto.OrdererSignature
	6,  // 3: hlfsdk.proto.Envelope.channel_header:type_name -> common.ChannelHeader
	7,  // 4: hlfsdk.proto.Envelope.validation_code:type_name -> protos.TxValidationCode
	2,  // 5: hlfsdk.proto.Envelope.transaction:type_name -> hlfsdk.proto.Transaction
	8,  // 6: hlfsdk.proto.Envelope.channel_config:type_name -> hlfsdk.proto.ChannelConfig
	3,  // 7: hlfsdk.proto.Transaction.actions:type_name -> hlfsdk.proto.TransactionAction
	9,  // 8: hlfsdk.proto.Transaction.creator_identity:type_name -> msp.SerializedIdentity
	10, // 9: hlfsdk.proto.TransactionAction.event:type_name -> protos.ChaincodeEvent
	9,  // 10: hlfsdk.proto.TransactionAction.endorsers:type_name -> msp.SerializedIdentity
	11, // 11: hlfsdk.proto.TransactionAction.read_write_sets:type_name -> kvrwset.KVRWSet
	12, // 12: hlfsdk.proto.TransactionAction.chaincode_invocation_spec:type_name -> protos.ChaincodeInvocationSpec
	9,  // 13: hlfsdk.proto.TransactionAction.creator_identity:type_name -> msp.SerializedIdentity
	9,  // 14: hlfsdk.proto.OrdererSignature.identity:type_name -> msp.SerializedIdentity
	15, // [15:15] is the sub-list for method output_type
	15, // [15:15] is the sub-list for method input_type
	15, // [15:15] is the sub-list for extension type_name
	15, // [15:15] is the sub-list for extension extendee
	0,  // [0:15] is the sub-list for field type_name
}

func init() { file_block_proto_init() }
func file_block_proto_init() {
	if File_block_proto != nil {
		return
	}
	file_chan_config_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_block_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Block); i {
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
		file_block_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Envelope); i {
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
		file_block_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Transaction); i {
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
		file_block_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactionAction); i {
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
		file_block_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrdererSignature); i {
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
			RawDescriptor: file_block_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_block_proto_goTypes,
		DependencyIndexes: file_block_proto_depIdxs,
		MessageInfos:      file_block_proto_msgTypes,
	}.Build()
	File_block_proto = out.File
	file_block_proto_rawDesc = nil
	file_block_proto_goTypes = nil
	file_block_proto_depIdxs = nil
}
