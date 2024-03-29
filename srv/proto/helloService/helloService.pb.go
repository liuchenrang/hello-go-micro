// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/helloService/helloService.proto

package go_micro_srv_helloService

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Message struct {
	Say                  string   `protobuf:"bytes,1,opt,name=say,proto3" json:"say,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_28efb5bd2b35a9b7, []int{0}
}

func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetSay() string {
	if m != nil {
		return m.Say
	}
	return ""
}

type Request struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_28efb5bd2b35a9b7, []int{1}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Response struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_28efb5bd2b35a9b7, []int{2}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type StreamingRequest struct {
	Count                int64    `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamingRequest) Reset()         { *m = StreamingRequest{} }
func (m *StreamingRequest) String() string { return proto.CompactTextString(m) }
func (*StreamingRequest) ProtoMessage()    {}
func (*StreamingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_28efb5bd2b35a9b7, []int{3}
}

func (m *StreamingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamingRequest.Unmarshal(m, b)
}
func (m *StreamingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamingRequest.Marshal(b, m, deterministic)
}
func (m *StreamingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamingRequest.Merge(m, src)
}
func (m *StreamingRequest) XXX_Size() int {
	return xxx_messageInfo_StreamingRequest.Size(m)
}
func (m *StreamingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StreamingRequest proto.InternalMessageInfo

func (m *StreamingRequest) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

type StreamingResponse struct {
	Count                int64    `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamingResponse) Reset()         { *m = StreamingResponse{} }
func (m *StreamingResponse) String() string { return proto.CompactTextString(m) }
func (*StreamingResponse) ProtoMessage()    {}
func (*StreamingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_28efb5bd2b35a9b7, []int{4}
}

func (m *StreamingResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamingResponse.Unmarshal(m, b)
}
func (m *StreamingResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamingResponse.Marshal(b, m, deterministic)
}
func (m *StreamingResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamingResponse.Merge(m, src)
}
func (m *StreamingResponse) XXX_Size() int {
	return xxx_messageInfo_StreamingResponse.Size(m)
}
func (m *StreamingResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamingResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StreamingResponse proto.InternalMessageInfo

func (m *StreamingResponse) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

type Ping struct {
	Stroke               int64    `protobuf:"varint,1,opt,name=stroke,proto3" json:"stroke,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ping) Reset()         { *m = Ping{} }
func (m *Ping) String() string { return proto.CompactTextString(m) }
func (*Ping) ProtoMessage()    {}
func (*Ping) Descriptor() ([]byte, []int) {
	return fileDescriptor_28efb5bd2b35a9b7, []int{5}
}

func (m *Ping) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ping.Unmarshal(m, b)
}
func (m *Ping) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ping.Marshal(b, m, deterministic)
}
func (m *Ping) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ping.Merge(m, src)
}
func (m *Ping) XXX_Size() int {
	return xxx_messageInfo_Ping.Size(m)
}
func (m *Ping) XXX_DiscardUnknown() {
	xxx_messageInfo_Ping.DiscardUnknown(m)
}

var xxx_messageInfo_Ping proto.InternalMessageInfo

func (m *Ping) GetStroke() int64 {
	if m != nil {
		return m.Stroke
	}
	return 0
}

type Pong struct {
	Stroke               int64    `protobuf:"varint,1,opt,name=stroke,proto3" json:"stroke,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Pong) Reset()         { *m = Pong{} }
func (m *Pong) String() string { return proto.CompactTextString(m) }
func (*Pong) ProtoMessage()    {}
func (*Pong) Descriptor() ([]byte, []int) {
	return fileDescriptor_28efb5bd2b35a9b7, []int{6}
}

func (m *Pong) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Pong.Unmarshal(m, b)
}
func (m *Pong) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Pong.Marshal(b, m, deterministic)
}
func (m *Pong) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pong.Merge(m, src)
}
func (m *Pong) XXX_Size() int {
	return xxx_messageInfo_Pong.Size(m)
}
func (m *Pong) XXX_DiscardUnknown() {
	xxx_messageInfo_Pong.DiscardUnknown(m)
}

var xxx_messageInfo_Pong proto.InternalMessageInfo

func (m *Pong) GetStroke() int64 {
	if m != nil {
		return m.Stroke
	}
	return 0
}

func init() {
	proto.RegisterType((*Message)(nil), "go.micro.srv.helloService.Message")
	proto.RegisterType((*Request)(nil), "go.micro.srv.helloService.Request")
	proto.RegisterType((*Response)(nil), "go.micro.srv.helloService.Response")
	proto.RegisterType((*StreamingRequest)(nil), "go.micro.srv.helloService.StreamingRequest")
	proto.RegisterType((*StreamingResponse)(nil), "go.micro.srv.helloService.StreamingResponse")
	proto.RegisterType((*Ping)(nil), "go.micro.srv.helloService.Ping")
	proto.RegisterType((*Pong)(nil), "go.micro.srv.helloService.Pong")
}

func init() {
	proto.RegisterFile("proto/helloService/helloService.proto", fileDescriptor_28efb5bd2b35a9b7)
}

var fileDescriptor_28efb5bd2b35a9b7 = []byte{
	// 280 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xdf, 0x4a, 0xf3, 0x40,
	0x10, 0xc5, 0x9b, 0xaf, 0xf9, 0xd2, 0x3a, 0x78, 0x51, 0x17, 0x11, 0x8d, 0x7f, 0x59, 0x11, 0x22,
	0xca, 0x5a, 0xf4, 0x11, 0xbc, 0xf1, 0x46, 0xd0, 0xf4, 0x09, 0x62, 0x18, 0xd6, 0x60, 0xb2, 0x53,
	0x77, 0xd2, 0x82, 0x2f, 0xe3, 0xb3, 0x4a, 0x36, 0x1b, 0x88, 0x62, 0x83, 0x77, 0x33, 0x9c, 0xdf,
	0x99, 0x9d, 0x39, 0x2c, 0x5c, 0x2c, 0x2d, 0xd5, 0x74, 0xf3, 0x8a, 0x65, 0x49, 0x0b, 0xb4, 0xeb,
	0x22, 0xc7, 0x6f, 0x8d, 0x72, 0xba, 0x38, 0xd0, 0xa4, 0xaa, 0x22, 0xb7, 0xa4, 0xd8, 0xae, 0x55,
	0x1f, 0x90, 0x87, 0x30, 0x79, 0x44, 0xe6, 0x4c, 0xa3, 0x98, 0xc1, 0x98, 0xb3, 0x8f, 0xfd, 0xe0,
	0x2c, 0x48, 0xb6, 0xd2, 0xa6, 0x94, 0xc7, 0x30, 0x49, 0xf1, 0x7d, 0x85, 0x5c, 0x0b, 0x01, 0xa1,
	0xc9, 0x2a, 0xf4, 0xaa, 0xab, 0xe5, 0x11, 0x4c, 0x53, 0xe4, 0x25, 0x19, 0x76, 0xe6, 0x8a, 0x75,
	0x67, 0xae, 0x58, 0xcb, 0x04, 0x66, 0x8b, 0xda, 0x62, 0x56, 0x15, 0x46, 0x77, 0x53, 0x76, 0xe1,
	0x7f, 0x4e, 0x2b, 0x53, 0x3b, 0x6e, 0x9c, 0xb6, 0x8d, 0xbc, 0x84, 0x9d, 0x1e, 0xe9, 0x07, 0xfe,
	0x8e, 0x9e, 0x40, 0xf8, 0x54, 0x18, 0x2d, 0xf6, 0x20, 0xe2, 0xda, 0xd2, 0x1b, 0x7a, 0xd9, 0x77,
	0x4e, 0xa7, 0xcd, 0xfa, 0xed, 0xe7, 0x3f, 0xd8, 0x7e, 0xe8, 0xdd, 0x2f, 0x9e, 0x21, 0xbc, 0xcf,
	0xca, 0x52, 0x48, 0xb5, 0x31, 0x23, 0xe5, 0xb7, 0x8f, 0xcf, 0x07, 0x99, 0x76, 0x6f, 0x39, 0x12,
	0x1a, 0xa2, 0xf6, 0x1c, 0x71, 0x35, 0x60, 0xf8, 0x99, 0x4d, 0x7c, 0xfd, 0x37, 0xb8, 0x7b, 0x66,
	0x1e, 0x88, 0x14, 0xa6, 0x4d, 0x18, 0xee, 0xe0, 0xd3, 0x01, 0x77, 0x03, 0xc5, 0x83, 0x00, 0x19,
	0x2d, 0x47, 0x49, 0x30, 0x0f, 0x5e, 0x22, 0xf7, 0x63, 0xee, 0xbe, 0x02, 0x00, 0x00, 0xff, 0xff,
	0x69, 0xed, 0xec, 0xc0, 0x5a, 0x02, 0x00, 0x00,
}
