// Code generated by protoc-gen-go. DO NOT EDIT.
// source: message.proto

package p2ppb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type BroadcastMsg struct {
	ChainId              uint32               `protobuf:"varint,1,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	MsgType              uint32               `protobuf:"varint,2,opt,name=msg_type,json=msgType,proto3" json:"msg_type,omitempty"`
	MsgBody              []byte               `protobuf:"bytes,3,opt,name=msg_body,json=msgBody,proto3" json:"msg_body,omitempty"`
	PeerId               string               `protobuf:"bytes,4,opt,name=peer_id,json=peerId,proto3" json:"peer_id,omitempty"`
	Timestamp            *timestamp.Timestamp `protobuf:"bytes,5,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *BroadcastMsg) Reset()         { *m = BroadcastMsg{} }
func (m *BroadcastMsg) String() string { return proto.CompactTextString(m) }
func (*BroadcastMsg) ProtoMessage()    {}
func (*BroadcastMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_message_d969953dc02c4436, []int{0}
}
func (m *BroadcastMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BroadcastMsg.Unmarshal(m, b)
}
func (m *BroadcastMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BroadcastMsg.Marshal(b, m, deterministic)
}
func (dst *BroadcastMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BroadcastMsg.Merge(dst, src)
}
func (m *BroadcastMsg) XXX_Size() int {
	return xxx_messageInfo_BroadcastMsg.Size(m)
}
func (m *BroadcastMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_BroadcastMsg.DiscardUnknown(m)
}

var xxx_messageInfo_BroadcastMsg proto.InternalMessageInfo

func (m *BroadcastMsg) GetChainId() uint32 {
	if m != nil {
		return m.ChainId
	}
	return 0
}

func (m *BroadcastMsg) GetMsgType() uint32 {
	if m != nil {
		return m.MsgType
	}
	return 0
}

func (m *BroadcastMsg) GetMsgBody() []byte {
	if m != nil {
		return m.MsgBody
	}
	return nil
}

func (m *BroadcastMsg) GetPeerId() string {
	if m != nil {
		return m.PeerId
	}
	return ""
}

func (m *BroadcastMsg) GetTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

type UnicastMsg struct {
	ChainId              uint32               `protobuf:"varint,1,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	Addr                 string               `protobuf:"bytes,2,opt,name=addr,proto3" json:"addr,omitempty"`
	MsgType              uint32               `protobuf:"varint,3,opt,name=msg_type,json=msgType,proto3" json:"msg_type,omitempty"`
	MsgBody              []byte               `protobuf:"bytes,4,opt,name=msg_body,json=msgBody,proto3" json:"msg_body,omitempty"`
	PeerId               string               `protobuf:"bytes,5,opt,name=peer_id,json=peerId,proto3" json:"peer_id,omitempty"`
	Timestamp            *timestamp.Timestamp `protobuf:"bytes,6,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *UnicastMsg) Reset()         { *m = UnicastMsg{} }
func (m *UnicastMsg) String() string { return proto.CompactTextString(m) }
func (*UnicastMsg) ProtoMessage()    {}
func (*UnicastMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_message_d969953dc02c4436, []int{1}
}
func (m *UnicastMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UnicastMsg.Unmarshal(m, b)
}
func (m *UnicastMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UnicastMsg.Marshal(b, m, deterministic)
}
func (dst *UnicastMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UnicastMsg.Merge(dst, src)
}
func (m *UnicastMsg) XXX_Size() int {
	return xxx_messageInfo_UnicastMsg.Size(m)
}
func (m *UnicastMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_UnicastMsg.DiscardUnknown(m)
}

var xxx_messageInfo_UnicastMsg proto.InternalMessageInfo

func (m *UnicastMsg) GetChainId() uint32 {
	if m != nil {
		return m.ChainId
	}
	return 0
}

func (m *UnicastMsg) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

func (m *UnicastMsg) GetMsgType() uint32 {
	if m != nil {
		return m.MsgType
	}
	return 0
}

func (m *UnicastMsg) GetMsgBody() []byte {
	if m != nil {
		return m.MsgBody
	}
	return nil
}

func (m *UnicastMsg) GetPeerId() string {
	if m != nil {
		return m.PeerId
	}
	return ""
}

func (m *UnicastMsg) GetTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func init() {
	proto.RegisterType((*BroadcastMsg)(nil), "p2ppb.BroadcastMsg")
	proto.RegisterType((*UnicastMsg)(nil), "p2ppb.UnicastMsg")
}

func init() { proto.RegisterFile("message.proto", fileDescriptor_message_d969953dc02c4436) }

var fileDescriptor_message_d969953dc02c4436 = []byte{
	// 242 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x8f, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x86, 0x65, 0x9a, 0xa4, 0xe4, 0x68, 0x17, 0x2f, 0x84, 0x2e, 0x44, 0x9d, 0x32, 0xa5, 0x52,
	0x59, 0x98, 0xbb, 0x75, 0x60, 0x89, 0xca, 0x5c, 0x39, 0xbd, 0xc3, 0x44, 0xc2, 0xf1, 0x29, 0x36,
	0x43, 0x5e, 0x8b, 0x67, 0xe0, 0xc1, 0x50, 0x1c, 0x05, 0xc4, 0x10, 0x09, 0xb6, 0x5c, 0xbe, 0x5f,
	0xd6, 0xf7, 0xc1, 0xda, 0x90, 0x73, 0x4a, 0x53, 0xc9, 0x9d, 0xf5, 0x56, 0xc6, 0xbc, 0x67, 0xae,
	0x37, 0xf7, 0xda, 0x5a, 0xfd, 0x46, 0xbb, 0xf0, 0xb3, 0x7e, 0x7f, 0xd9, 0xf9, 0xc6, 0x90, 0xf3,
	0xca, 0xf0, 0xb8, 0xdb, 0x7e, 0x08, 0x58, 0x1d, 0x3a, 0xab, 0xf0, 0xa2, 0x9c, 0x7f, 0x72, 0x5a,
	0xde, 0xc1, 0xf5, 0xe5, 0x55, 0x35, 0xed, 0xb9, 0xc1, 0x4c, 0xe4, 0xa2, 0x58, 0x57, 0xcb, 0x70,
	0x1f, 0x71, 0x40, 0xc6, 0xe9, 0xb3, 0xef, 0x99, 0xb2, 0xab, 0x11, 0x19, 0xa7, 0x4f, 0x3d, 0xd3,
	0x84, 0x6a, 0x8b, 0x7d, 0xb6, 0xc8, 0x45, 0xb1, 0x0a, 0xe8, 0x60, 0xb1, 0x97, 0xb7, 0xb0, 0x64,
	0xa2, 0x6e, 0x78, 0x2f, 0xca, 0x45, 0x91, 0x56, 0xc9, 0x70, 0x1e, 0x51, 0x3e, 0x42, 0xfa, 0x6d,
	0x93, 0xc5, 0xb9, 0x28, 0x6e, 0xf6, 0x9b, 0x72, 0xf4, 0x2d, 0x27, 0xdf, 0xf2, 0x34, 0x2d, 0xaa,
	0x9f, 0xf1, 0xf6, 0x53, 0x00, 0x3c, 0xb7, 0xcd, 0x1f, 0x94, 0x25, 0x44, 0x0a, 0xb1, 0x0b, 0xba,
	0x69, 0x15, 0xbe, 0x7f, 0x65, 0x2c, 0xe6, 0x33, 0xa2, 0xd9, 0x8c, 0x78, 0x3e, 0x23, 0xf9, 0x47,
	0x46, 0x9d, 0x04, 0xfc, 0xf0, 0x15, 0x00, 0x00, 0xff, 0xff, 0x9c, 0xdc, 0x49, 0x95, 0xbb, 0x01,
	0x00, 0x00,
}