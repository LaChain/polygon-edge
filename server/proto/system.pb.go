// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.0
// source: server/proto/system.proto

package proto

import (
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type GetTraceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number uint64 `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
}

func (x *GetTraceRequest) Reset() {
	*x = GetTraceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_proto_system_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTraceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTraceRequest) ProtoMessage() {}

func (x *GetTraceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_system_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTraceRequest.ProtoReflect.Descriptor instead.
func (*GetTraceRequest) Descriptor() ([]byte, []int) {
	return file_server_proto_system_proto_rawDescGZIP(), []int{0}
}

func (x *GetTraceRequest) GetNumber() uint64 {
	if x != nil {
		return x.Number
	}
	return 0
}

type GetTraceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Trace []byte `protobuf:"bytes,1,opt,name=trace,proto3" json:"trace,omitempty"`
}

func (x *GetTraceResponse) Reset() {
	*x = GetTraceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_proto_system_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTraceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTraceResponse) ProtoMessage() {}

func (x *GetTraceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_system_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTraceResponse.ProtoReflect.Descriptor instead.
func (*GetTraceResponse) Descriptor() ([]byte, []int) {
	return file_server_proto_system_proto_rawDescGZIP(), []int{1}
}

func (x *GetTraceResponse) GetTrace() []byte {
	if x != nil {
		return x.Trace
	}
	return nil
}

type BlockchainEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Added   []*BlockchainEvent_Header `protobuf:"bytes,1,rep,name=added,proto3" json:"added,omitempty"`
	Removed []*BlockchainEvent_Header `protobuf:"bytes,2,rep,name=removed,proto3" json:"removed,omitempty"`
}

func (x *BlockchainEvent) Reset() {
	*x = BlockchainEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_proto_system_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockchainEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockchainEvent) ProtoMessage() {}

func (x *BlockchainEvent) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_system_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockchainEvent.ProtoReflect.Descriptor instead.
func (*BlockchainEvent) Descriptor() ([]byte, []int) {
	return file_server_proto_system_proto_rawDescGZIP(), []int{2}
}

func (x *BlockchainEvent) GetAdded() []*BlockchainEvent_Header {
	if x != nil {
		return x.Added
	}
	return nil
}

func (x *BlockchainEvent) GetRemoved() []*BlockchainEvent_Header {
	if x != nil {
		return x.Removed
	}
	return nil
}

type ServerStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Network int64               `protobuf:"varint,1,opt,name=network,proto3" json:"network,omitempty"`
	Genesis string              `protobuf:"bytes,2,opt,name=genesis,proto3" json:"genesis,omitempty"`
	Current *ServerStatus_Block `protobuf:"bytes,3,opt,name=current,proto3" json:"current,omitempty"`
	P2PAddr string              `protobuf:"bytes,4,opt,name=p2pAddr,proto3" json:"p2pAddr,omitempty"`
}

func (x *ServerStatus) Reset() {
	*x = ServerStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_proto_system_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerStatus) ProtoMessage() {}

func (x *ServerStatus) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_system_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerStatus.ProtoReflect.Descriptor instead.
func (*ServerStatus) Descriptor() ([]byte, []int) {
	return file_server_proto_system_proto_rawDescGZIP(), []int{3}
}

func (x *ServerStatus) GetNetwork() int64 {
	if x != nil {
		return x.Network
	}
	return 0
}

func (x *ServerStatus) GetGenesis() string {
	if x != nil {
		return x.Genesis
	}
	return ""
}

func (x *ServerStatus) GetCurrent() *ServerStatus_Block {
	if x != nil {
		return x.Current
	}
	return nil
}

func (x *ServerStatus) GetP2PAddr() string {
	if x != nil {
		return x.P2PAddr
	}
	return ""
}

type Peer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Protocols []string `protobuf:"bytes,2,rep,name=protocols,proto3" json:"protocols,omitempty"`
	Addrs     []string `protobuf:"bytes,3,rep,name=addrs,proto3" json:"addrs,omitempty"`
}

func (x *Peer) Reset() {
	*x = Peer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_proto_system_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Peer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Peer) ProtoMessage() {}

func (x *Peer) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_system_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Peer.ProtoReflect.Descriptor instead.
func (*Peer) Descriptor() ([]byte, []int) {
	return file_server_proto_system_proto_rawDescGZIP(), []int{4}
}

func (x *Peer) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Peer) GetProtocols() []string {
	if x != nil {
		return x.Protocols
	}
	return nil
}

func (x *Peer) GetAddrs() []string {
	if x != nil {
		return x.Addrs
	}
	return nil
}

type PeersAddRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *PeersAddRequest) Reset() {
	*x = PeersAddRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_proto_system_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PeersAddRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PeersAddRequest) ProtoMessage() {}

func (x *PeersAddRequest) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_system_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PeersAddRequest.ProtoReflect.Descriptor instead.
func (*PeersAddRequest) Descriptor() ([]byte, []int) {
	return file_server_proto_system_proto_rawDescGZIP(), []int{5}
}

func (x *PeersAddRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type PeersAddResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *PeersAddResponse) Reset() {
	*x = PeersAddResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_proto_system_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PeersAddResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PeersAddResponse) ProtoMessage() {}

func (x *PeersAddResponse) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_system_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PeersAddResponse.ProtoReflect.Descriptor instead.
func (*PeersAddResponse) Descriptor() ([]byte, []int) {
	return file_server_proto_system_proto_rawDescGZIP(), []int{6}
}

func (x *PeersAddResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type PeersStatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *PeersStatusRequest) Reset() {
	*x = PeersStatusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_proto_system_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PeersStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PeersStatusRequest) ProtoMessage() {}

func (x *PeersStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_system_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PeersStatusRequest.ProtoReflect.Descriptor instead.
func (*PeersStatusRequest) Descriptor() ([]byte, []int) {
	return file_server_proto_system_proto_rawDescGZIP(), []int{7}
}

func (x *PeersStatusRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type PeersListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Peers []*Peer `protobuf:"bytes,1,rep,name=peers,proto3" json:"peers,omitempty"`
}

func (x *PeersListResponse) Reset() {
	*x = PeersListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_proto_system_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PeersListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PeersListResponse) ProtoMessage() {}

func (x *PeersListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_system_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PeersListResponse.ProtoReflect.Descriptor instead.
func (*PeersListResponse) Descriptor() ([]byte, []int) {
	return file_server_proto_system_proto_rawDescGZIP(), []int{8}
}

func (x *PeersListResponse) GetPeers() []*Peer {
	if x != nil {
		return x.Peers
	}
	return nil
}

type BlockByNumberRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number uint64 `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
}

func (x *BlockByNumberRequest) Reset() {
	*x = BlockByNumberRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_proto_system_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockByNumberRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockByNumberRequest) ProtoMessage() {}

func (x *BlockByNumberRequest) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_system_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockByNumberRequest.ProtoReflect.Descriptor instead.
func (*BlockByNumberRequest) Descriptor() ([]byte, []int) {
	return file_server_proto_system_proto_rawDescGZIP(), []int{9}
}

func (x *BlockByNumberRequest) GetNumber() uint64 {
	if x != nil {
		return x.Number
	}
	return 0
}

type BlockResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *BlockResponse) Reset() {
	*x = BlockResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_proto_system_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockResponse) ProtoMessage() {}

func (x *BlockResponse) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_system_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockResponse.ProtoReflect.Descriptor instead.
func (*BlockResponse) Descriptor() ([]byte, []int) {
	return file_server_proto_system_proto_rawDescGZIP(), []int{10}
}

func (x *BlockResponse) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type ExportRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	From uint64 `protobuf:"varint,1,opt,name=from,proto3" json:"from,omitempty"`
	To   uint64 `protobuf:"varint,2,opt,name=to,proto3" json:"to,omitempty"`
}

func (x *ExportRequest) Reset() {
	*x = ExportRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_proto_system_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExportRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExportRequest) ProtoMessage() {}

func (x *ExportRequest) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_system_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExportRequest.ProtoReflect.Descriptor instead.
func (*ExportRequest) Descriptor() ([]byte, []int) {
	return file_server_proto_system_proto_rawDescGZIP(), []int{11}
}

func (x *ExportRequest) GetFrom() uint64 {
	if x != nil {
		return x.From
	}
	return 0
}

func (x *ExportRequest) GetTo() uint64 {
	if x != nil {
		return x.To
	}
	return 0
}

type ExportEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	From uint64 `protobuf:"varint,1,opt,name=from,proto3" json:"from,omitempty"`
	// null when zero
	To     uint64 `protobuf:"varint,2,opt,name=to,proto3" json:"to,omitempty"`
	Latest uint64 `protobuf:"varint,3,opt,name=latest,proto3" json:"latest,omitempty"`
	Data   []byte `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *ExportEvent) Reset() {
	*x = ExportEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_proto_system_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExportEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExportEvent) ProtoMessage() {}

func (x *ExportEvent) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_system_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExportEvent.ProtoReflect.Descriptor instead.
func (*ExportEvent) Descriptor() ([]byte, []int) {
	return file_server_proto_system_proto_rawDescGZIP(), []int{12}
}

func (x *ExportEvent) GetFrom() uint64 {
	if x != nil {
		return x.From
	}
	return 0
}

func (x *ExportEvent) GetTo() uint64 {
	if x != nil {
		return x.To
	}
	return 0
}

func (x *ExportEvent) GetLatest() uint64 {
	if x != nil {
		return x.Latest
	}
	return 0
}

func (x *ExportEvent) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type BlockchainEvent_Header struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number int64  `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	Hash   string `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
}

func (x *BlockchainEvent_Header) Reset() {
	*x = BlockchainEvent_Header{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_proto_system_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockchainEvent_Header) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockchainEvent_Header) ProtoMessage() {}

func (x *BlockchainEvent_Header) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_system_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockchainEvent_Header.ProtoReflect.Descriptor instead.
func (*BlockchainEvent_Header) Descriptor() ([]byte, []int) {
	return file_server_proto_system_proto_rawDescGZIP(), []int{2, 0}
}

func (x *BlockchainEvent_Header) GetNumber() int64 {
	if x != nil {
		return x.Number
	}
	return 0
}

func (x *BlockchainEvent_Header) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

type ServerStatus_Block struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number int64  `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	Hash   string `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
}

func (x *ServerStatus_Block) Reset() {
	*x = ServerStatus_Block{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_proto_system_proto_msgTypes[14]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerStatus_Block) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerStatus_Block) ProtoMessage() {}

func (x *ServerStatus_Block) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_system_proto_msgTypes[14]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerStatus_Block.ProtoReflect.Descriptor instead.
func (*ServerStatus_Block) Descriptor() ([]byte, []int) {
	return file_server_proto_system_proto_rawDescGZIP(), []int{3, 0}
}

func (x *ServerStatus_Block) GetNumber() int64 {
	if x != nil {
		return x.Number
	}
	return 0
}

func (x *ServerStatus_Block) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

var File_server_proto_system_proto protoreflect.FileDescriptor

var file_server_proto_system_proto_rawDesc = []byte{
	0x0a, 0x19, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x76, 0x31, 0x1a,
	0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x29, 0x0a, 0x0f,
	0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x28, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x54, 0x72,
	0x61, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x72, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x74, 0x72, 0x61, 0x63,
	0x65, 0x22, 0xaf, 0x01, 0x0a, 0x0f, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x30, 0x0a, 0x05, 0x61, 0x64, 0x64, 0x65, 0x64, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x63,
	0x68, 0x61, 0x69, 0x6e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x52, 0x05, 0x61, 0x64, 0x64, 0x65, 0x64, 0x12, 0x34, 0x0a, 0x07, 0x72, 0x65, 0x6d, 0x6f, 0x76,
	0x65, 0x64, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x6c,
	0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x48, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x52, 0x07, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x1a, 0x34, 0x0a,
	0x06, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12,
	0x12, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68,
	0x61, 0x73, 0x68, 0x22, 0xc3, 0x01, 0x0a, 0x0c, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12, 0x18,
	0x0a, 0x07, 0x67, 0x65, 0x6e, 0x65, 0x73, 0x69, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x67, 0x65, 0x6e, 0x65, 0x73, 0x69, 0x73, 0x12, 0x30, 0x0a, 0x07, 0x63, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x76, 0x31, 0x2e, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x42, 0x6c, 0x6f, 0x63,
	0x6b, 0x52, 0x07, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x32,
	0x70, 0x41, 0x64, 0x64, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x32, 0x70,
	0x41, 0x64, 0x64, 0x72, 0x1a, 0x33, 0x0a, 0x05, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x16, 0x0a,
	0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x22, 0x4a, 0x0a, 0x04, 0x50, 0x65, 0x65,
	0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x73, 0x12,
	0x14, 0x0a, 0x05, 0x61, 0x64, 0x64, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05,
	0x61, 0x64, 0x64, 0x72, 0x73, 0x22, 0x21, 0x0a, 0x0f, 0x50, 0x65, 0x65, 0x72, 0x73, 0x41, 0x64,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x2c, 0x0a, 0x10, 0x50, 0x65, 0x65, 0x72,
	0x73, 0x41, 0x64, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x24, 0x0a, 0x12, 0x50, 0x65, 0x65, 0x72, 0x73, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x33, 0x0a, 0x11,
	0x50, 0x65, 0x65, 0x72, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x1e, 0x0a, 0x05, 0x70, 0x65, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x08, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x52, 0x05, 0x70, 0x65, 0x65, 0x72,
	0x73, 0x22, 0x2e, 0x0a, 0x14, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x42, 0x79, 0x4e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x22, 0x23, 0x0a, 0x0d, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x33, 0x0a, 0x0d, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x74,
	0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x74, 0x6f, 0x22, 0x5d, 0x0a, 0x0b, 0x45,
	0x78, 0x70, 0x6f, 0x72, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72,
	0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x0e,
	0x0a, 0x02, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x16,
	0x0a, 0x06, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06,
	0x6c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x32, 0xc4, 0x03, 0x0a, 0x06, 0x53,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0x12, 0x35, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x10, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x35, 0x0a, 0x08,
	0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x63, 0x65, 0x12, 0x13, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x54, 0x72, 0x61, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x08, 0x50, 0x65, 0x65, 0x72, 0x73, 0x41, 0x64, 0x64, 0x12,
	0x13, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x73, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x73, 0x41,
	0x64, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x09, 0x50, 0x65,
	0x65, 0x72, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a,
	0x15, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x0b, 0x50, 0x65, 0x65, 0x72, 0x73, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x73,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x08, 0x2e,
	0x76, 0x31, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x12, 0x3a, 0x0a, 0x09, 0x53, 0x75, 0x62, 0x73, 0x63,
	0x72, 0x69, 0x62, 0x65, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x13, 0x2e, 0x76,
	0x31, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x30, 0x01, 0x12, 0x3c, 0x0a, 0x0d, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x42, 0x79, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x12, 0x18, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x42,
	0x79, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11,
	0x2e, 0x76, 0x31, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x2e, 0x0a, 0x06, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x11, 0x2e, 0x76, 0x31,
	0x2e, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f,
	0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x30,
	0x01, 0x42, 0x0f, 0x5a, 0x0d, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_server_proto_system_proto_rawDescOnce sync.Once
	file_server_proto_system_proto_rawDescData = file_server_proto_system_proto_rawDesc
)

func file_server_proto_system_proto_rawDescGZIP() []byte {
	file_server_proto_system_proto_rawDescOnce.Do(func() {
		file_server_proto_system_proto_rawDescData = protoimpl.X.CompressGZIP(file_server_proto_system_proto_rawDescData)
	})
	return file_server_proto_system_proto_rawDescData
}

var file_server_proto_system_proto_msgTypes = make([]protoimpl.MessageInfo, 15)
var file_server_proto_system_proto_goTypes = []interface{}{
	(*GetTraceRequest)(nil),        // 0: v1.GetTraceRequest
	(*GetTraceResponse)(nil),       // 1: v1.GetTraceResponse
	(*BlockchainEvent)(nil),        // 2: v1.BlockchainEvent
	(*ServerStatus)(nil),           // 3: v1.ServerStatus
	(*Peer)(nil),                   // 4: v1.Peer
	(*PeersAddRequest)(nil),        // 5: v1.PeersAddRequest
	(*PeersAddResponse)(nil),       // 6: v1.PeersAddResponse
	(*PeersStatusRequest)(nil),     // 7: v1.PeersStatusRequest
	(*PeersListResponse)(nil),      // 8: v1.PeersListResponse
	(*BlockByNumberRequest)(nil),   // 9: v1.BlockByNumberRequest
	(*BlockResponse)(nil),          // 10: v1.BlockResponse
	(*ExportRequest)(nil),          // 11: v1.ExportRequest
	(*ExportEvent)(nil),            // 12: v1.ExportEvent
	(*BlockchainEvent_Header)(nil), // 13: v1.BlockchainEvent.Header
	(*ServerStatus_Block)(nil),     // 14: v1.ServerStatus.Block
	(*empty.Empty)(nil),            // 15: google.protobuf.Empty
}
var file_server_proto_system_proto_depIdxs = []int32{
	13, // 0: v1.BlockchainEvent.added:type_name -> v1.BlockchainEvent.Header
	13, // 1: v1.BlockchainEvent.removed:type_name -> v1.BlockchainEvent.Header
	14, // 2: v1.ServerStatus.current:type_name -> v1.ServerStatus.Block
	4,  // 3: v1.PeersListResponse.peers:type_name -> v1.Peer
	15, // 4: v1.System.GetStatus:input_type -> google.protobuf.Empty
	0,  // 5: v1.System.GetTrace:input_type -> v1.GetTraceRequest
	5,  // 6: v1.System.PeersAdd:input_type -> v1.PeersAddRequest
	15, // 7: v1.System.PeersList:input_type -> google.protobuf.Empty
	7,  // 8: v1.System.PeersStatus:input_type -> v1.PeersStatusRequest
	15, // 9: v1.System.Subscribe:input_type -> google.protobuf.Empty
	9,  // 10: v1.System.BlockByNumber:input_type -> v1.BlockByNumberRequest
	11, // 11: v1.System.Export:input_type -> v1.ExportRequest
	3,  // 12: v1.System.GetStatus:output_type -> v1.ServerStatus
	1,  // 13: v1.System.GetTrace:output_type -> v1.GetTraceResponse
	6,  // 14: v1.System.PeersAdd:output_type -> v1.PeersAddResponse
	8,  // 15: v1.System.PeersList:output_type -> v1.PeersListResponse
	4,  // 16: v1.System.PeersStatus:output_type -> v1.Peer
	2,  // 17: v1.System.Subscribe:output_type -> v1.BlockchainEvent
	10, // 18: v1.System.BlockByNumber:output_type -> v1.BlockResponse
	12, // 19: v1.System.Export:output_type -> v1.ExportEvent
	12, // [12:20] is the sub-list for method output_type
	4,  // [4:12] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_server_proto_system_proto_init() }
func file_server_proto_system_proto_init() {
	if File_server_proto_system_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_server_proto_system_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTraceRequest); i {
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
		file_server_proto_system_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTraceResponse); i {
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
		file_server_proto_system_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockchainEvent); i {
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
		file_server_proto_system_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerStatus); i {
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
		file_server_proto_system_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Peer); i {
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
		file_server_proto_system_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PeersAddRequest); i {
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
		file_server_proto_system_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PeersAddResponse); i {
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
		file_server_proto_system_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PeersStatusRequest); i {
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
		file_server_proto_system_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PeersListResponse); i {
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
		file_server_proto_system_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockByNumberRequest); i {
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
		file_server_proto_system_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockResponse); i {
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
		file_server_proto_system_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExportRequest); i {
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
		file_server_proto_system_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExportEvent); i {
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
		file_server_proto_system_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockchainEvent_Header); i {
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
		file_server_proto_system_proto_msgTypes[14].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerStatus_Block); i {
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
			RawDescriptor: file_server_proto_system_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   15,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_server_proto_system_proto_goTypes,
		DependencyIndexes: file_server_proto_system_proto_depIdxs,
		MessageInfos:      file_server_proto_system_proto_msgTypes,
	}.Build()
	File_server_proto_system_proto = out.File
	file_server_proto_system_proto_rawDesc = nil
	file_server_proto_system_proto_goTypes = nil
	file_server_proto_system_proto_depIdxs = nil
}
