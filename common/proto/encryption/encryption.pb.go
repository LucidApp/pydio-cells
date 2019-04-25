// Code generated by protoc-gen-go. DO NOT EDIT.
// source: encryption.proto

/*
Package encryption is a generated protocol buffer package.

It is generated from these files:
	encryption.proto

It has these top-level messages:
	Export
	Import
	KeyInfo
	Key
	AddKeyRequest
	AddKeyResponse
	GetKeyRequest
	GetKeyResponse
	AdminListKeysRequest
	AdminListKeysResponse
	AdminDeleteKeyRequest
	AdminDeleteKeyResponse
	AdminExportKeyRequest
	AdminExportKeyResponse
	AdminImportKeyRequest
	AdminImportKeyResponse
	AdminCreateKeyRequest
	AdminCreateKeyResponse
	NodeKey
	Node
	NodeInfo
	Block
	GetNodeInfoRequest
	GetNodeInfoResponse
	SetNodeInfoRequest
	SetNodeInfoResponse
	DeleteNodeRequest
	DeleteNodeResponse
	DeleteNodeKeyRequest
	DeleteNodeKeyResponse
	DeleteNodeSharedKeyRequest
	DeleteNodeSharedKeyResponse
*/
package encryption

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Export struct {
	By   string `protobuf:"bytes,1,opt,name=By" json:"By,omitempty"`
	Date int32  `protobuf:"varint,2,opt,name=Date" json:"Date,omitempty"`
}

func (m *Export) Reset()                    { *m = Export{} }
func (m *Export) String() string            { return proto.CompactTextString(m) }
func (*Export) ProtoMessage()               {}
func (*Export) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Export) GetBy() string {
	if m != nil {
		return m.By
	}
	return ""
}

func (m *Export) GetDate() int32 {
	if m != nil {
		return m.Date
	}
	return 0
}

type Import struct {
	By   string `protobuf:"bytes,1,opt,name=By" json:"By,omitempty"`
	Date int32  `protobuf:"varint,3,opt,name=Date" json:"Date,omitempty"`
}

func (m *Import) Reset()                    { *m = Import{} }
func (m *Import) String() string            { return proto.CompactTextString(m) }
func (*Import) ProtoMessage()               {}
func (*Import) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Import) GetBy() string {
	if m != nil {
		return m.By
	}
	return ""
}

func (m *Import) GetDate() int32 {
	if m != nil {
		return m.Date
	}
	return 0
}

type KeyInfo struct {
	Exports []*Export `protobuf:"bytes,1,rep,name=Exports" json:"Exports,omitempty"`
	Imports []*Import `protobuf:"bytes,2,rep,name=Imports" json:"Imports,omitempty"`
}

func (m *KeyInfo) Reset()                    { *m = KeyInfo{} }
func (m *KeyInfo) String() string            { return proto.CompactTextString(m) }
func (*KeyInfo) ProtoMessage()               {}
func (*KeyInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *KeyInfo) GetExports() []*Export {
	if m != nil {
		return m.Exports
	}
	return nil
}

func (m *KeyInfo) GetImports() []*Import {
	if m != nil {
		return m.Imports
	}
	return nil
}

type Key struct {
	Owner        string   `protobuf:"bytes,1,opt,name=Owner" json:"Owner,omitempty"`
	ID           string   `protobuf:"bytes,2,opt,name=ID" json:"ID,omitempty"`
	Label        string   `protobuf:"bytes,3,opt,name=Label" json:"Label,omitempty"`
	Content      string   `protobuf:"bytes,4,opt,name=Content" json:"Content,omitempty"`
	CreationDate int32    `protobuf:"varint,5,opt,name=CreationDate" json:"CreationDate,omitempty"`
	Info         *KeyInfo `protobuf:"bytes,6,opt,name=Info" json:"Info,omitempty"`
}

func (m *Key) Reset()                    { *m = Key{} }
func (m *Key) String() string            { return proto.CompactTextString(m) }
func (*Key) ProtoMessage()               {}
func (*Key) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Key) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *Key) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *Key) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

func (m *Key) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *Key) GetCreationDate() int32 {
	if m != nil {
		return m.CreationDate
	}
	return 0
}

func (m *Key) GetInfo() *KeyInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

type AddKeyRequest struct {
	Key         *Key   `protobuf:"bytes,1,opt,name=Key" json:"Key,omitempty"`
	StrPassword string `protobuf:"bytes,2,opt,name=StrPassword" json:"StrPassword,omitempty"`
}

func (m *AddKeyRequest) Reset()                    { *m = AddKeyRequest{} }
func (m *AddKeyRequest) String() string            { return proto.CompactTextString(m) }
func (*AddKeyRequest) ProtoMessage()               {}
func (*AddKeyRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *AddKeyRequest) GetKey() *Key {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *AddKeyRequest) GetStrPassword() string {
	if m != nil {
		return m.StrPassword
	}
	return ""
}

type AddKeyResponse struct {
	Success bool `protobuf:"varint,1,opt,name=Success" json:"Success,omitempty"`
}

func (m *AddKeyResponse) Reset()                    { *m = AddKeyResponse{} }
func (m *AddKeyResponse) String() string            { return proto.CompactTextString(m) }
func (*AddKeyResponse) ProtoMessage()               {}
func (*AddKeyResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *AddKeyResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

type GetKeyRequest struct {
	Owner       string `protobuf:"bytes,1,opt,name=Owner" json:"Owner,omitempty"`
	KeyID       string `protobuf:"bytes,2,opt,name=KeyID" json:"KeyID,omitempty"`
	StrPassword string `protobuf:"bytes,3,opt,name=StrPassword" json:"StrPassword,omitempty"`
}

func (m *GetKeyRequest) Reset()                    { *m = GetKeyRequest{} }
func (m *GetKeyRequest) String() string            { return proto.CompactTextString(m) }
func (*GetKeyRequest) ProtoMessage()               {}
func (*GetKeyRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *GetKeyRequest) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *GetKeyRequest) GetKeyID() string {
	if m != nil {
		return m.KeyID
	}
	return ""
}

func (m *GetKeyRequest) GetStrPassword() string {
	if m != nil {
		return m.StrPassword
	}
	return ""
}

type GetKeyResponse struct {
	Key *Key `protobuf:"bytes,1,opt,name=Key" json:"Key,omitempty"`
}

func (m *GetKeyResponse) Reset()                    { *m = GetKeyResponse{} }
func (m *GetKeyResponse) String() string            { return proto.CompactTextString(m) }
func (*GetKeyResponse) ProtoMessage()               {}
func (*GetKeyResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *GetKeyResponse) GetKey() *Key {
	if m != nil {
		return m.Key
	}
	return nil
}

type AdminListKeysRequest struct {
}

func (m *AdminListKeysRequest) Reset()                    { *m = AdminListKeysRequest{} }
func (m *AdminListKeysRequest) String() string            { return proto.CompactTextString(m) }
func (*AdminListKeysRequest) ProtoMessage()               {}
func (*AdminListKeysRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

type AdminListKeysResponse struct {
	Keys []*Key `protobuf:"bytes,1,rep,name=Keys" json:"Keys,omitempty"`
}

func (m *AdminListKeysResponse) Reset()                    { *m = AdminListKeysResponse{} }
func (m *AdminListKeysResponse) String() string            { return proto.CompactTextString(m) }
func (*AdminListKeysResponse) ProtoMessage()               {}
func (*AdminListKeysResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *AdminListKeysResponse) GetKeys() []*Key {
	if m != nil {
		return m.Keys
	}
	return nil
}

type AdminDeleteKeyRequest struct {
	KeyID string `protobuf:"bytes,1,opt,name=KeyID" json:"KeyID,omitempty"`
}

func (m *AdminDeleteKeyRequest) Reset()                    { *m = AdminDeleteKeyRequest{} }
func (m *AdminDeleteKeyRequest) String() string            { return proto.CompactTextString(m) }
func (*AdminDeleteKeyRequest) ProtoMessage()               {}
func (*AdminDeleteKeyRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *AdminDeleteKeyRequest) GetKeyID() string {
	if m != nil {
		return m.KeyID
	}
	return ""
}

type AdminDeleteKeyResponse struct {
	Success bool `protobuf:"varint,1,opt,name=Success" json:"Success,omitempty"`
}

func (m *AdminDeleteKeyResponse) Reset()                    { *m = AdminDeleteKeyResponse{} }
func (m *AdminDeleteKeyResponse) String() string            { return proto.CompactTextString(m) }
func (*AdminDeleteKeyResponse) ProtoMessage()               {}
func (*AdminDeleteKeyResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *AdminDeleteKeyResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

type AdminExportKeyRequest struct {
	KeyID       string `protobuf:"bytes,1,opt,name=KeyID" json:"KeyID,omitempty"`
	StrPassword string `protobuf:"bytes,2,opt,name=StrPassword" json:"StrPassword,omitempty"`
}

func (m *AdminExportKeyRequest) Reset()                    { *m = AdminExportKeyRequest{} }
func (m *AdminExportKeyRequest) String() string            { return proto.CompactTextString(m) }
func (*AdminExportKeyRequest) ProtoMessage()               {}
func (*AdminExportKeyRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *AdminExportKeyRequest) GetKeyID() string {
	if m != nil {
		return m.KeyID
	}
	return ""
}

func (m *AdminExportKeyRequest) GetStrPassword() string {
	if m != nil {
		return m.StrPassword
	}
	return ""
}

type AdminExportKeyResponse struct {
	Key *Key `protobuf:"bytes,1,opt,name=Key" json:"Key,omitempty"`
}

func (m *AdminExportKeyResponse) Reset()                    { *m = AdminExportKeyResponse{} }
func (m *AdminExportKeyResponse) String() string            { return proto.CompactTextString(m) }
func (*AdminExportKeyResponse) ProtoMessage()               {}
func (*AdminExportKeyResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *AdminExportKeyResponse) GetKey() *Key {
	if m != nil {
		return m.Key
	}
	return nil
}

type AdminImportKeyRequest struct {
	Key         *Key   `protobuf:"bytes,1,opt,name=Key" json:"Key,omitempty"`
	StrPassword string `protobuf:"bytes,2,opt,name=StrPassword" json:"StrPassword,omitempty"`
	Override    bool   `protobuf:"varint,3,opt,name=Override" json:"Override,omitempty"`
}

func (m *AdminImportKeyRequest) Reset()                    { *m = AdminImportKeyRequest{} }
func (m *AdminImportKeyRequest) String() string            { return proto.CompactTextString(m) }
func (*AdminImportKeyRequest) ProtoMessage()               {}
func (*AdminImportKeyRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

func (m *AdminImportKeyRequest) GetKey() *Key {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *AdminImportKeyRequest) GetStrPassword() string {
	if m != nil {
		return m.StrPassword
	}
	return ""
}

func (m *AdminImportKeyRequest) GetOverride() bool {
	if m != nil {
		return m.Override
	}
	return false
}

type AdminImportKeyResponse struct {
	Success bool `protobuf:"varint,1,opt,name=Success" json:"Success,omitempty"`
}

func (m *AdminImportKeyResponse) Reset()                    { *m = AdminImportKeyResponse{} }
func (m *AdminImportKeyResponse) String() string            { return proto.CompactTextString(m) }
func (*AdminImportKeyResponse) ProtoMessage()               {}
func (*AdminImportKeyResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{15} }

func (m *AdminImportKeyResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

type AdminCreateKeyRequest struct {
	KeyID string `protobuf:"bytes,1,opt,name=KeyID" json:"KeyID,omitempty"`
	Label string `protobuf:"bytes,2,opt,name=Label" json:"Label,omitempty"`
}

func (m *AdminCreateKeyRequest) Reset()                    { *m = AdminCreateKeyRequest{} }
func (m *AdminCreateKeyRequest) String() string            { return proto.CompactTextString(m) }
func (*AdminCreateKeyRequest) ProtoMessage()               {}
func (*AdminCreateKeyRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{16} }

func (m *AdminCreateKeyRequest) GetKeyID() string {
	if m != nil {
		return m.KeyID
	}
	return ""
}

func (m *AdminCreateKeyRequest) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

type AdminCreateKeyResponse struct {
	Success bool `protobuf:"varint,1,opt,name=Success" json:"Success,omitempty"`
}

func (m *AdminCreateKeyResponse) Reset()                    { *m = AdminCreateKeyResponse{} }
func (m *AdminCreateKeyResponse) String() string            { return proto.CompactTextString(m) }
func (*AdminCreateKeyResponse) ProtoMessage()               {}
func (*AdminCreateKeyResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{17} }

func (m *AdminCreateKeyResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

type NodeKey struct {
	NodeId  string `protobuf:"bytes,1,opt,name=NodeId" json:"NodeId,omitempty"`
	UserId  string `protobuf:"bytes,2,opt,name=UserId" json:"UserId,omitempty"`
	OwnerId string `protobuf:"bytes,3,opt,name=OwnerId" json:"OwnerId,omitempty"`
	KeyData []byte `protobuf:"bytes,6,opt,name=KeyData,proto3" json:"KeyData,omitempty"`
}

func (m *NodeKey) Reset()                    { *m = NodeKey{} }
func (m *NodeKey) String() string            { return proto.CompactTextString(m) }
func (*NodeKey) ProtoMessage()               {}
func (*NodeKey) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{18} }

func (m *NodeKey) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

func (m *NodeKey) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *NodeKey) GetOwnerId() string {
	if m != nil {
		return m.OwnerId
	}
	return ""
}

func (m *NodeKey) GetKeyData() []byte {
	if m != nil {
		return m.KeyData
	}
	return nil
}

type Node struct {
	NodeId string `protobuf:"bytes,1,opt,name=NodeId" json:"NodeId,omitempty"`
	Legacy bool   `protobuf:"varint,2,opt,name=Legacy" json:"Legacy,omitempty"`
}

func (m *Node) Reset()                    { *m = Node{} }
func (m *Node) String() string            { return proto.CompactTextString(m) }
func (*Node) ProtoMessage()               {}
func (*Node) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{19} }

func (m *Node) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

func (m *Node) GetLegacy() bool {
	if m != nil {
		return m.Legacy
	}
	return false
}

type NodeInfo struct {
	Node    *Node    `protobuf:"bytes,1,opt,name=Node" json:"Node,omitempty"`
	NodeKey *NodeKey `protobuf:"bytes,2,opt,name=NodeKey" json:"NodeKey,omitempty"`
	Block   *Block   `protobuf:"bytes,3,opt,name=Block" json:"Block,omitempty"`
}

func (m *NodeInfo) Reset()                    { *m = NodeInfo{} }
func (m *NodeInfo) String() string            { return proto.CompactTextString(m) }
func (*NodeInfo) ProtoMessage()               {}
func (*NodeInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{20} }

func (m *NodeInfo) GetNode() *Node {
	if m != nil {
		return m.Node
	}
	return nil
}

func (m *NodeInfo) GetNodeKey() *NodeKey {
	if m != nil {
		return m.NodeKey
	}
	return nil
}

func (m *NodeInfo) GetBlock() *Block {
	if m != nil {
		return m.Block
	}
	return nil
}

type Block struct {
	OwnerId    string `protobuf:"bytes,1,opt,name=OwnerId" json:"OwnerId,omitempty"`
	PartId     uint32 `protobuf:"varint,2,opt,name=PartId" json:"PartId,omitempty"`
	Position   uint32 `protobuf:"varint,3,opt,name=Position" json:"Position,omitempty"`
	HeaderSize uint32 `protobuf:"varint,4,opt,name=HeaderSize" json:"HeaderSize,omitempty"`
	BlockSize  uint32 `protobuf:"varint,5,opt,name=BlockSize" json:"BlockSize,omitempty"`
	Nonce      []byte `protobuf:"bytes,6,opt,name=Nonce,proto3" json:"Nonce,omitempty"`
}

func (m *Block) Reset()                    { *m = Block{} }
func (m *Block) String() string            { return proto.CompactTextString(m) }
func (*Block) ProtoMessage()               {}
func (*Block) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{21} }

func (m *Block) GetOwnerId() string {
	if m != nil {
		return m.OwnerId
	}
	return ""
}

func (m *Block) GetPartId() uint32 {
	if m != nil {
		return m.PartId
	}
	return 0
}

func (m *Block) GetPosition() uint32 {
	if m != nil {
		return m.Position
	}
	return 0
}

func (m *Block) GetHeaderSize() uint32 {
	if m != nil {
		return m.HeaderSize
	}
	return 0
}

func (m *Block) GetBlockSize() uint32 {
	if m != nil {
		return m.BlockSize
	}
	return 0
}

func (m *Block) GetNonce() []byte {
	if m != nil {
		return m.Nonce
	}
	return nil
}

type GetNodeInfoRequest struct {
	UserId      string `protobuf:"bytes,1,opt,name=UserId" json:"UserId,omitempty"`
	NodeId      string `protobuf:"bytes,2,opt,name=NodeId" json:"NodeId,omitempty"`
	WithRange   bool   `protobuf:"varint,3,opt,name=WithRange" json:"WithRange,omitempty"`
	PlainOffset int64  `protobuf:"varint,4,opt,name=PlainOffset" json:"PlainOffset,omitempty"`
	PlainLength int64  `protobuf:"varint,5,opt,name=PlainLength" json:"PlainLength,omitempty"`
}

func (m *GetNodeInfoRequest) Reset()                    { *m = GetNodeInfoRequest{} }
func (m *GetNodeInfoRequest) String() string            { return proto.CompactTextString(m) }
func (*GetNodeInfoRequest) ProtoMessage()               {}
func (*GetNodeInfoRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{22} }

func (m *GetNodeInfoRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *GetNodeInfoRequest) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

func (m *GetNodeInfoRequest) GetWithRange() bool {
	if m != nil {
		return m.WithRange
	}
	return false
}

func (m *GetNodeInfoRequest) GetPlainOffset() int64 {
	if m != nil {
		return m.PlainOffset
	}
	return 0
}

func (m *GetNodeInfoRequest) GetPlainLength() int64 {
	if m != nil {
		return m.PlainLength
	}
	return 0
}

type GetNodeInfoResponse struct {
	NodeInfo                   *NodeInfo `protobuf:"bytes,1,opt,name=NodeInfo" json:"NodeInfo,omitempty"`
	HeadSKippedPlainBytesCount int64     `protobuf:"varint,2,opt,name=HeadSKippedPlainBytesCount" json:"HeadSKippedPlainBytesCount,omitempty"`
	WithRange                  bool      `protobuf:"varint,3,opt,name=WithRange" json:"WithRange,omitempty"`
	EncryptedOffset            int64     `protobuf:"varint,4,opt,name=EncryptedOffset" json:"EncryptedOffset,omitempty"`
	EncryptedCount             int64     `protobuf:"varint,5,opt,name=EncryptedCount" json:"EncryptedCount,omitempty"`
}

func (m *GetNodeInfoResponse) Reset()                    { *m = GetNodeInfoResponse{} }
func (m *GetNodeInfoResponse) String() string            { return proto.CompactTextString(m) }
func (*GetNodeInfoResponse) ProtoMessage()               {}
func (*GetNodeInfoResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{23} }

func (m *GetNodeInfoResponse) GetNodeInfo() *NodeInfo {
	if m != nil {
		return m.NodeInfo
	}
	return nil
}

func (m *GetNodeInfoResponse) GetHeadSKippedPlainBytesCount() int64 {
	if m != nil {
		return m.HeadSKippedPlainBytesCount
	}
	return 0
}

func (m *GetNodeInfoResponse) GetWithRange() bool {
	if m != nil {
		return m.WithRange
	}
	return false
}

func (m *GetNodeInfoResponse) GetEncryptedOffset() int64 {
	if m != nil {
		return m.EncryptedOffset
	}
	return 0
}

func (m *GetNodeInfoResponse) GetEncryptedCount() int64 {
	if m != nil {
		return m.EncryptedCount
	}
	return 0
}

type SetNodeInfoRequest struct {
	NodeId  string   `protobuf:"bytes,1,opt,name=NodeId" json:"NodeId,omitempty"`
	NodeKey *NodeKey `protobuf:"bytes,2,opt,name=NodeKey" json:"NodeKey,omitempty"`
	Block   *Block   `protobuf:"bytes,3,opt,name=Block" json:"Block,omitempty"`
}

func (m *SetNodeInfoRequest) Reset()                    { *m = SetNodeInfoRequest{} }
func (m *SetNodeInfoRequest) String() string            { return proto.CompactTextString(m) }
func (*SetNodeInfoRequest) ProtoMessage()               {}
func (*SetNodeInfoRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{24} }

func (m *SetNodeInfoRequest) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

func (m *SetNodeInfoRequest) GetNodeKey() *NodeKey {
	if m != nil {
		return m.NodeKey
	}
	return nil
}

func (m *SetNodeInfoRequest) GetBlock() *Block {
	if m != nil {
		return m.Block
	}
	return nil
}

type SetNodeInfoResponse struct {
}

func (m *SetNodeInfoResponse) Reset()                    { *m = SetNodeInfoResponse{} }
func (m *SetNodeInfoResponse) String() string            { return proto.CompactTextString(m) }
func (*SetNodeInfoResponse) ProtoMessage()               {}
func (*SetNodeInfoResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{25} }

type DeleteNodeRequest struct {
	NodeId string `protobuf:"bytes,1,opt,name=NodeId" json:"NodeId,omitempty"`
}

func (m *DeleteNodeRequest) Reset()                    { *m = DeleteNodeRequest{} }
func (m *DeleteNodeRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteNodeRequest) ProtoMessage()               {}
func (*DeleteNodeRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{26} }

func (m *DeleteNodeRequest) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

type DeleteNodeResponse struct {
}

func (m *DeleteNodeResponse) Reset()                    { *m = DeleteNodeResponse{} }
func (m *DeleteNodeResponse) String() string            { return proto.CompactTextString(m) }
func (*DeleteNodeResponse) ProtoMessage()               {}
func (*DeleteNodeResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{27} }

type DeleteNodeKeyRequest struct {
	UserId string `protobuf:"bytes,1,opt,name=UserId" json:"UserId,omitempty"`
	NodeId string `protobuf:"bytes,2,opt,name=NodeId" json:"NodeId,omitempty"`
}

func (m *DeleteNodeKeyRequest) Reset()                    { *m = DeleteNodeKeyRequest{} }
func (m *DeleteNodeKeyRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteNodeKeyRequest) ProtoMessage()               {}
func (*DeleteNodeKeyRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{28} }

func (m *DeleteNodeKeyRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *DeleteNodeKeyRequest) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

type DeleteNodeKeyResponse struct {
}

func (m *DeleteNodeKeyResponse) Reset()                    { *m = DeleteNodeKeyResponse{} }
func (m *DeleteNodeKeyResponse) String() string            { return proto.CompactTextString(m) }
func (*DeleteNodeKeyResponse) ProtoMessage()               {}
func (*DeleteNodeKeyResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{29} }

type DeleteNodeSharedKeyRequest struct {
	UserId  string `protobuf:"bytes,1,opt,name=UserId" json:"UserId,omitempty"`
	OwnerId string `protobuf:"bytes,2,opt,name=OwnerId" json:"OwnerId,omitempty"`
	NodeId  string `protobuf:"bytes,3,opt,name=NodeId" json:"NodeId,omitempty"`
}

func (m *DeleteNodeSharedKeyRequest) Reset()                    { *m = DeleteNodeSharedKeyRequest{} }
func (m *DeleteNodeSharedKeyRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteNodeSharedKeyRequest) ProtoMessage()               {}
func (*DeleteNodeSharedKeyRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{30} }

func (m *DeleteNodeSharedKeyRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *DeleteNodeSharedKeyRequest) GetOwnerId() string {
	if m != nil {
		return m.OwnerId
	}
	return ""
}

func (m *DeleteNodeSharedKeyRequest) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

type DeleteNodeSharedKeyResponse struct {
}

func (m *DeleteNodeSharedKeyResponse) Reset()                    { *m = DeleteNodeSharedKeyResponse{} }
func (m *DeleteNodeSharedKeyResponse) String() string            { return proto.CompactTextString(m) }
func (*DeleteNodeSharedKeyResponse) ProtoMessage()               {}
func (*DeleteNodeSharedKeyResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{31} }

func init() {
	proto.RegisterType((*Export)(nil), "encryption.Export")
	proto.RegisterType((*Import)(nil), "encryption.Import")
	proto.RegisterType((*KeyInfo)(nil), "encryption.KeyInfo")
	proto.RegisterType((*Key)(nil), "encryption.Key")
	proto.RegisterType((*AddKeyRequest)(nil), "encryption.AddKeyRequest")
	proto.RegisterType((*AddKeyResponse)(nil), "encryption.AddKeyResponse")
	proto.RegisterType((*GetKeyRequest)(nil), "encryption.GetKeyRequest")
	proto.RegisterType((*GetKeyResponse)(nil), "encryption.GetKeyResponse")
	proto.RegisterType((*AdminListKeysRequest)(nil), "encryption.AdminListKeysRequest")
	proto.RegisterType((*AdminListKeysResponse)(nil), "encryption.AdminListKeysResponse")
	proto.RegisterType((*AdminDeleteKeyRequest)(nil), "encryption.AdminDeleteKeyRequest")
	proto.RegisterType((*AdminDeleteKeyResponse)(nil), "encryption.AdminDeleteKeyResponse")
	proto.RegisterType((*AdminExportKeyRequest)(nil), "encryption.AdminExportKeyRequest")
	proto.RegisterType((*AdminExportKeyResponse)(nil), "encryption.AdminExportKeyResponse")
	proto.RegisterType((*AdminImportKeyRequest)(nil), "encryption.AdminImportKeyRequest")
	proto.RegisterType((*AdminImportKeyResponse)(nil), "encryption.AdminImportKeyResponse")
	proto.RegisterType((*AdminCreateKeyRequest)(nil), "encryption.AdminCreateKeyRequest")
	proto.RegisterType((*AdminCreateKeyResponse)(nil), "encryption.AdminCreateKeyResponse")
	proto.RegisterType((*NodeKey)(nil), "encryption.NodeKey")
	proto.RegisterType((*Node)(nil), "encryption.Node")
	proto.RegisterType((*NodeInfo)(nil), "encryption.NodeInfo")
	proto.RegisterType((*Block)(nil), "encryption.Block")
	proto.RegisterType((*GetNodeInfoRequest)(nil), "encryption.GetNodeInfoRequest")
	proto.RegisterType((*GetNodeInfoResponse)(nil), "encryption.GetNodeInfoResponse")
	proto.RegisterType((*SetNodeInfoRequest)(nil), "encryption.SetNodeInfoRequest")
	proto.RegisterType((*SetNodeInfoResponse)(nil), "encryption.SetNodeInfoResponse")
	proto.RegisterType((*DeleteNodeRequest)(nil), "encryption.DeleteNodeRequest")
	proto.RegisterType((*DeleteNodeResponse)(nil), "encryption.DeleteNodeResponse")
	proto.RegisterType((*DeleteNodeKeyRequest)(nil), "encryption.DeleteNodeKeyRequest")
	proto.RegisterType((*DeleteNodeKeyResponse)(nil), "encryption.DeleteNodeKeyResponse")
	proto.RegisterType((*DeleteNodeSharedKeyRequest)(nil), "encryption.DeleteNodeSharedKeyRequest")
	proto.RegisterType((*DeleteNodeSharedKeyResponse)(nil), "encryption.DeleteNodeSharedKeyResponse")
}

func init() { proto.RegisterFile("encryption.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 1101 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x57, 0x4b, 0x6f, 0x1b, 0x55,
	0x14, 0xce, 0xf8, 0x95, 0xf8, 0x38, 0x76, 0xdb, 0x1b, 0x27, 0x98, 0xa1, 0x0d, 0xee, 0x05, 0x35,
	0x16, 0xb4, 0x15, 0x72, 0x25, 0x36, 0x20, 0xa4, 0x24, 0x2e, 0xc5, 0x72, 0x48, 0xac, 0x3b, 0x05,
	0xc4, 0x82, 0xc5, 0xd4, 0x73, 0x9c, 0x58, 0x24, 0x33, 0x66, 0x66, 0x42, 0x3b, 0xac, 0x59, 0xf0,
	0x4b, 0x10, 0xfc, 0x08, 0xfe, 0x07, 0xff, 0x84, 0x2d, 0xba, 0x8f, 0x99, 0xb9, 0xf3, 0x88, 0x1d,
	0x24, 0xba, 0xf3, 0x79, 0xdc, 0xef, 0xbc, 0xcf, 0x1c, 0xc3, 0x5d, 0x74, 0x67, 0x7e, 0xb4, 0x0c,
	0x17, 0x9e, 0xfb, 0x74, 0xe9, 0x7b, 0xa1, 0x47, 0x20, 0xe5, 0xd0, 0xc7, 0xd0, 0x78, 0xfe, 0x66,
	0xe9, 0xf9, 0x21, 0xe9, 0x40, 0xe5, 0x28, 0xea, 0x19, 0x7d, 0x63, 0xd0, 0x64, 0x95, 0xa3, 0x88,
	0x10, 0xa8, 0x8d, 0xec, 0x10, 0x7b, 0x95, 0xbe, 0x31, 0xa8, 0x33, 0xf1, 0x9b, 0x6b, 0x8f, 0xaf,
	0x56, 0x6a, 0x57, 0x35, 0x6d, 0x84, 0xcd, 0x09, 0x46, 0x63, 0x77, 0xee, 0x91, 0xc7, 0xb0, 0x29,
	0xcd, 0x04, 0x3d, 0xa3, 0x5f, 0x1d, 0xb4, 0x86, 0xe4, 0xa9, 0xe6, 0x96, 0x14, 0xb1, 0x58, 0x85,
	0x6b, 0x4b, 0x33, 0x41, 0xaf, 0x52, 0xd4, 0x96, 0x22, 0x16, 0xab, 0xd0, 0x3f, 0x0c, 0xa8, 0x4e,
	0x30, 0x22, 0x5d, 0xa8, 0x9f, 0xbd, 0x76, 0xd1, 0x57, 0x5e, 0x49, 0x82, 0x3b, 0x3a, 0x1e, 0x89,
	0x20, 0x9a, 0xac, 0x32, 0x1e, 0x71, 0xad, 0x13, 0xfb, 0x15, 0x5e, 0x0a, 0x4f, 0x9b, 0x4c, 0x12,
	0xa4, 0x07, 0x9b, 0xc7, 0x9e, 0x1b, 0xa2, 0x1b, 0xf6, 0x6a, 0x82, 0x1f, 0x93, 0x84, 0xc2, 0xf6,
	0xb1, 0x8f, 0x36, 0xb7, 0x2c, 0x02, 0xac, 0x8b, 0x00, 0x33, 0x3c, 0x72, 0x00, 0x35, 0x1e, 0x65,
	0xaf, 0xd1, 0x37, 0x06, 0xad, 0xe1, 0x8e, 0xee, 0xac, 0x4a, 0x00, 0x13, 0x0a, 0xf4, 0x25, 0xb4,
	0x0f, 0x1d, 0x67, 0x82, 0x11, 0xc3, 0x9f, 0xae, 0x31, 0x08, 0xc9, 0x43, 0xe1, 0xba, 0xf0, 0xb8,
	0x35, 0xbc, 0x93, 0x7b, 0xc8, 0x44, 0x58, 0x7d, 0x68, 0x59, 0xa1, 0x3f, 0xb5, 0x83, 0xe0, 0xb5,
	0xe7, 0x3b, 0x2a, 0x12, 0x9d, 0x45, 0x3f, 0x82, 0x4e, 0x8c, 0x1a, 0x2c, 0x3d, 0x37, 0x40, 0x1e,
	0x8e, 0x75, 0x3d, 0x9b, 0x61, 0x10, 0x08, 0xe8, 0x2d, 0x16, 0x93, 0xf4, 0x07, 0x68, 0xbf, 0xc0,
	0x50, 0xf3, 0xa0, 0x3c, 0x6b, 0x5d, 0xa8, 0x73, 0xcf, 0xe3, 0xc4, 0x49, 0x22, 0xef, 0x4a, 0xb5,
	0xe8, 0xca, 0x33, 0xe8, 0xc4, 0xf0, 0xca, 0x95, 0xf5, 0x11, 0xd2, 0x3d, 0xe8, 0x1e, 0x3a, 0x57,
	0x0b, 0xf7, 0x64, 0x11, 0xf0, 0xa7, 0x81, 0x72, 0x8d, 0x7e, 0x0e, 0xbb, 0x39, 0xbe, 0xc2, 0xfc,
	0x00, 0x6a, 0x9c, 0x56, 0xad, 0x54, 0x00, 0x15, 0x42, 0xfa, 0x44, 0xbd, 0x1e, 0xe1, 0x25, 0x86,
	0x98, 0x8d, 0x58, 0xc6, 0x66, 0x68, 0xb1, 0xd1, 0x21, 0xec, 0xe5, 0xd5, 0xd7, 0x26, 0xf3, 0x4c,
	0x99, 0x90, 0x7d, 0xbb, 0xce, 0xc4, 0x2d, 0x2a, 0xf9, 0x99, 0x72, 0x42, 0x03, 0xbc, 0x7d, 0x1a,
	0xdf, 0x28, 0x6f, 0xe4, 0x5c, 0xfc, 0xcf, 0x4d, 0x46, 0x4c, 0xd8, 0x3a, 0xfb, 0x19, 0x7d, 0x7f,
	0xe1, 0xc8, 0x21, 0xdf, 0x62, 0x09, 0x9d, 0xe4, 0x4e, 0xb3, 0xbc, 0x36, 0x77, 0xc7, 0xca, 0x5b,
	0x31, 0x48, 0x6b, 0xcb, 0x93, 0x8e, 0x6d, 0x45, 0x1b, 0xdb, 0xc4, 0xb0, 0x06, 0xb2, 0xd6, 0xf0,
	0x15, 0x6c, 0x9e, 0x7a, 0x0e, 0x57, 0x26, 0x7b, 0xd0, 0xe0, 0x3f, 0xc7, 0x8e, 0xb2, 0xa5, 0x28,
	0xce, 0xff, 0x26, 0x40, 0x7f, 0x1c, 0x27, 0x42, 0x51, 0x1c, 0x54, 0x8c, 0xc7, 0x38, 0xee, 0xfd,
	0x98, 0xe4, 0x92, 0x09, 0x46, 0x23, 0x3b, 0xb4, 0xc5, 0x12, 0xd8, 0x66, 0x31, 0x49, 0x3f, 0x85,
	0x1a, 0x47, 0x5d, 0x65, 0xeb, 0x04, 0xcf, 0xed, 0x59, 0x24, 0x6c, 0x6d, 0x31, 0x45, 0xd1, 0xdf,
	0x0c, 0xd8, 0x12, 0x2a, 0x7c, 0x7d, 0x7e, 0x28, 0x41, 0x54, 0x09, 0xef, 0xea, 0x25, 0xe4, 0x7c,
	0x26, 0x4d, 0x3c, 0x49, 0x22, 0x13, 0x58, 0xb9, 0x4d, 0xa4, 0x44, 0x2c, 0x89, 0xfe, 0x00, 0xea,
	0x47, 0x97, 0xde, 0xec, 0x47, 0x11, 0x4b, 0x6b, 0x78, 0x4f, 0x57, 0x16, 0x02, 0x26, 0xe5, 0xf4,
	0x4f, 0x43, 0x69, 0xea, 0x09, 0x30, 0xb2, 0x09, 0xd8, 0x83, 0xc6, 0xd4, 0xf6, 0x43, 0x95, 0xb2,
	0x36, 0x53, 0x14, 0x6f, 0x9b, 0xa9, 0x17, 0x2c, 0x38, 0xa8, 0xb0, 0xd3, 0x66, 0x09, 0x4d, 0xf6,
	0x01, 0xbe, 0x42, 0xdb, 0x41, 0xdf, 0x5a, 0xfc, 0x82, 0x62, 0xef, 0xb6, 0x99, 0xc6, 0x21, 0xf7,
	0xa1, 0x29, 0xcc, 0x0a, 0x71, 0x5d, 0x88, 0x53, 0x06, 0xef, 0x88, 0x53, 0xcf, 0x9d, 0xa1, 0x4a,
	0xb8, 0x24, 0xe8, 0xef, 0x06, 0x90, 0x17, 0x18, 0xc6, 0x99, 0x8b, 0x9b, 0x2a, 0xad, 0xa8, 0x91,
	0xa9, 0x68, 0x5a, 0x95, 0x4a, 0xa6, 0x2a, 0xf7, 0xa1, 0xf9, 0xdd, 0x22, 0xbc, 0x60, 0xb6, 0x7b,
	0x1e, 0xb7, 0x7b, 0xca, 0xe0, 0xd3, 0x32, 0xbd, 0xb4, 0x17, 0xee, 0xd9, 0x7c, 0x1e, 0xa0, 0xfc,
	0x62, 0x54, 0x99, 0xce, 0x4a, 0x34, 0x4e, 0xd0, 0x3d, 0x0f, 0x2f, 0x84, 0xf3, 0xb1, 0x86, 0x64,
	0xd1, 0x7f, 0x0c, 0xd8, 0xc9, 0x38, 0xaa, 0x1a, 0xf7, 0x93, 0xb4, 0xec, 0xaa, 0xdc, 0xdd, 0x7c,
	0x15, 0x85, 0x7e, 0xda, 0x1c, 0x5f, 0x80, 0xc9, 0x93, 0x66, 0x4d, 0x16, 0xcb, 0x25, 0x3a, 0xc2,
	0xc6, 0x51, 0x14, 0x62, 0x70, 0xec, 0x5d, 0xbb, 0xa1, 0x88, 0xab, 0xca, 0x56, 0x68, 0xac, 0x89,
	0x75, 0x00, 0x77, 0x9e, 0x4b, 0xf3, 0xe8, 0x64, 0xe2, 0xcd, 0xb3, 0xc9, 0x23, 0xe8, 0x24, 0x2c,
	0x69, 0x5b, 0x86, 0x9d, 0xe3, 0xd2, 0x5f, 0x0d, 0x20, 0x56, 0x69, 0x89, 0x4a, 0x07, 0xe4, 0x6d,
	0x75, 0xf5, 0x2e, 0xec, 0x58, 0xc5, 0xfc, 0xd3, 0x8f, 0xe1, 0x9e, 0xfc, 0x04, 0x88, 0xc1, 0x5a,
	0xed, 0x1b, 0xed, 0x02, 0xd1, 0x95, 0x15, 0xc4, 0x97, 0xd0, 0x4d, 0xb9, 0xda, 0x66, 0xfb, 0x8f,
	0x4d, 0x48, 0xdf, 0x81, 0xdd, 0x1c, 0x8e, 0x32, 0x30, 0x07, 0x33, 0x15, 0x58, 0x17, 0xb6, 0x8f,
	0xce, 0x2d, 0xcc, 0x68, 0xc3, 0x5b, 0x29, 0x0c, 0xaf, 0x72, 0xa0, 0x9a, 0x71, 0xe0, 0x01, 0xbc,
	0x57, 0x6a, 0x47, 0xba, 0x31, 0xfc, 0xbb, 0x06, 0xdb, 0x1c, 0x7b, 0x82, 0x91, 0x15, 0x7a, 0x3e,
	0x92, 0x43, 0x68, 0xc8, 0x43, 0x84, 0xbc, 0xab, 0xa7, 0x3d, 0x73, 0xf2, 0x98, 0x66, 0x99, 0x48,
	0x05, 0xb6, 0xc1, 0x21, 0xe4, 0x01, 0x91, 0x85, 0xc8, 0xdc, 0x2c, 0x59, 0x88, 0xec, 0xbd, 0x41,
	0x37, 0xc8, 0xb7, 0xfc, 0xc8, 0xd2, 0xce, 0x06, 0xd2, 0xcf, 0x5a, 0x2c, 0x5e, 0x1a, 0xe6, 0xc3,
	0x15, 0x1a, 0x09, 0xee, 0xf7, 0xfc, 0xcc, 0xd2, 0x3f, 0x36, 0xa4, 0xf8, 0x2c, 0xff, 0x35, 0x33,
	0xe9, 0x2a, 0x95, 0x02, 0x74, 0x72, 0x7c, 0x94, 0x40, 0xe7, 0xef, 0x98, 0x12, 0xe8, 0xc2, 0xed,
	0xa2, 0x41, 0x27, 0x27, 0x45, 0x09, 0x74, 0xfe, 0x7e, 0x29, 0x81, 0x2e, 0x5c, 0x24, 0x1a, 0x74,
	0xf2, 0xd9, 0x2f, 0x81, 0xce, 0x1f, 0x23, 0x25, 0xd0, 0x85, 0xab, 0x81, 0x6e, 0x0c, 0xff, 0xaa,
	0x42, 0x47, 0x75, 0xfd, 0xd7, 0xb6, 0x6b, 0x9f, 0xa3, 0x4f, 0x4e, 0xa1, 0xa5, 0xed, 0x4b, 0xb2,
	0x9f, 0xeb, 0x81, 0xdc, 0x3a, 0x31, 0xdf, 0xbf, 0x51, 0xae, 0x16, 0xed, 0x14, 0x5a, 0xd6, 0x4d,
	0x78, 0xd6, 0x1a, 0xbc, 0x92, 0xc5, 0x31, 0x30, 0xc8, 0x04, 0x20, 0x1d, 0x17, 0xf2, 0x40, 0x7f,
	0x50, 0x58, 0x29, 0xe6, 0xfe, 0x4d, 0x62, 0xe5, 0xde, 0x4b, 0x68, 0x67, 0x86, 0x3f, 0xdb, 0xc5,
	0x65, 0xfb, 0x25, 0xdb, 0xc5, 0xa5, 0x9b, 0x83, 0xcc, 0x61, 0xa7, 0x64, 0xa2, 0xc9, 0xa3, 0xf2,
	0x97, 0xf9, 0xd5, 0x62, 0x1e, 0xac, 0xd5, 0x93, 0x76, 0x5e, 0x35, 0xc4, 0x3f, 0xcd, 0x67, 0xff,
	0x06, 0x00, 0x00, 0xff, 0xff, 0xd7, 0x76, 0x6d, 0xa1, 0x7d, 0x0e, 0x00, 0x00,
}
