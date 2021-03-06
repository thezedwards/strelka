// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/strelka/strelka.proto

package strelka

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

type Request struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Client               string   `protobuf:"bytes,2,opt,name=client,proto3" json:"client,omitempty"`
	Source               string   `protobuf:"bytes,3,opt,name=source,proto3" json:"source,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_bbe6981b6c359f03, []int{0}
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

func (m *Request) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Request) GetClient() string {
	if m != nil {
		return m.Client
	}
	return ""
}

func (m *Request) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

type Attributes struct {
	Filename             string            `protobuf:"bytes,1,opt,name=filename,proto3" json:"filename,omitempty"`
	Metadata             map[string]string `protobuf:"bytes,2,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Attributes) Reset()         { *m = Attributes{} }
func (m *Attributes) String() string { return proto.CompactTextString(m) }
func (*Attributes) ProtoMessage()    {}
func (*Attributes) Descriptor() ([]byte, []int) {
	return fileDescriptor_bbe6981b6c359f03, []int{1}
}

func (m *Attributes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Attributes.Unmarshal(m, b)
}
func (m *Attributes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Attributes.Marshal(b, m, deterministic)
}
func (m *Attributes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Attributes.Merge(m, src)
}
func (m *Attributes) XXX_Size() int {
	return xxx_messageInfo_Attributes.Size(m)
}
func (m *Attributes) XXX_DiscardUnknown() {
	xxx_messageInfo_Attributes.DiscardUnknown(m)
}

var xxx_messageInfo_Attributes proto.InternalMessageInfo

func (m *Attributes) GetFilename() string {
	if m != nil {
		return m.Filename
	}
	return ""
}

func (m *Attributes) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

type ScanFileRequest struct {
	Data                 []byte      `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Request              *Request    `protobuf:"bytes,2,opt,name=request,proto3" json:"request,omitempty"`
	Attributes           *Attributes `protobuf:"bytes,3,opt,name=attributes,proto3" json:"attributes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *ScanFileRequest) Reset()         { *m = ScanFileRequest{} }
func (m *ScanFileRequest) String() string { return proto.CompactTextString(m) }
func (*ScanFileRequest) ProtoMessage()    {}
func (*ScanFileRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bbe6981b6c359f03, []int{2}
}

func (m *ScanFileRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ScanFileRequest.Unmarshal(m, b)
}
func (m *ScanFileRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ScanFileRequest.Marshal(b, m, deterministic)
}
func (m *ScanFileRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScanFileRequest.Merge(m, src)
}
func (m *ScanFileRequest) XXX_Size() int {
	return xxx_messageInfo_ScanFileRequest.Size(m)
}
func (m *ScanFileRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ScanFileRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ScanFileRequest proto.InternalMessageInfo

func (m *ScanFileRequest) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *ScanFileRequest) GetRequest() *Request {
	if m != nil {
		return m.Request
	}
	return nil
}

func (m *ScanFileRequest) GetAttributes() *Attributes {
	if m != nil {
		return m.Attributes
	}
	return nil
}

type ScanHttpRequest struct {
	Url                  string      `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	Request              *Request    `protobuf:"bytes,2,opt,name=request,proto3" json:"request,omitempty"`
	Attributes           *Attributes `protobuf:"bytes,3,opt,name=attributes,proto3" json:"attributes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *ScanHttpRequest) Reset()         { *m = ScanHttpRequest{} }
func (m *ScanHttpRequest) String() string { return proto.CompactTextString(m) }
func (*ScanHttpRequest) ProtoMessage()    {}
func (*ScanHttpRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bbe6981b6c359f03, []int{3}
}

func (m *ScanHttpRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ScanHttpRequest.Unmarshal(m, b)
}
func (m *ScanHttpRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ScanHttpRequest.Marshal(b, m, deterministic)
}
func (m *ScanHttpRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScanHttpRequest.Merge(m, src)
}
func (m *ScanHttpRequest) XXX_Size() int {
	return xxx_messageInfo_ScanHttpRequest.Size(m)
}
func (m *ScanHttpRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ScanHttpRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ScanHttpRequest proto.InternalMessageInfo

func (m *ScanHttpRequest) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *ScanHttpRequest) GetRequest() *Request {
	if m != nil {
		return m.Request
	}
	return nil
}

func (m *ScanHttpRequest) GetAttributes() *Attributes {
	if m != nil {
		return m.Attributes
	}
	return nil
}

type ScanResponse struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Event                string   `protobuf:"bytes,2,opt,name=event,proto3" json:"event,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ScanResponse) Reset()         { *m = ScanResponse{} }
func (m *ScanResponse) String() string { return proto.CompactTextString(m) }
func (*ScanResponse) ProtoMessage()    {}
func (*ScanResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_bbe6981b6c359f03, []int{4}
}

func (m *ScanResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ScanResponse.Unmarshal(m, b)
}
func (m *ScanResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ScanResponse.Marshal(b, m, deterministic)
}
func (m *ScanResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScanResponse.Merge(m, src)
}
func (m *ScanResponse) XXX_Size() int {
	return xxx_messageInfo_ScanResponse.Size(m)
}
func (m *ScanResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ScanResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ScanResponse proto.InternalMessageInfo

func (m *ScanResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ScanResponse) GetEvent() string {
	if m != nil {
		return m.Event
	}
	return ""
}

func init() {
	proto.RegisterType((*Request)(nil), "Request")
	proto.RegisterType((*Attributes)(nil), "Attributes")
	proto.RegisterMapType((map[string]string)(nil), "Attributes.MetadataEntry")
	proto.RegisterType((*ScanFileRequest)(nil), "ScanFileRequest")
	proto.RegisterType((*ScanHttpRequest)(nil), "ScanHttpRequest")
	proto.RegisterType((*ScanResponse)(nil), "ScanResponse")
}

func init() { proto.RegisterFile("api/strelka/strelka.proto", fileDescriptor_bbe6981b6c359f03) }

var fileDescriptor_bbe6981b6c359f03 = []byte{
	// 332 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x52, 0xcd, 0x4e, 0xf3, 0x30,
	0x10, 0xfc, 0x92, 0x7c, 0x6d, 0xc3, 0xa6, 0x85, 0xca, 0xaa, 0x50, 0xda, 0x53, 0x95, 0x53, 0x25,
	0xa4, 0x00, 0x01, 0x24, 0x04, 0xe2, 0xc0, 0x81, 0x0a, 0x0e, 0x5c, 0xcc, 0x13, 0xb8, 0xcd, 0x22,
	0x59, 0x75, 0x9d, 0xe0, 0x6c, 0x2a, 0xf5, 0x5d, 0x78, 0x58, 0x94, 0x1f, 0xb7, 0x05, 0xae, 0x9c,
	0xe2, 0xd9, 0xb1, 0x67, 0x66, 0x77, 0x03, 0x63, 0x91, 0xcb, 0xf3, 0x82, 0x0c, 0xaa, 0x95, 0xb0,
	0xdf, 0x38, 0x37, 0x19, 0x65, 0xd1, 0x0b, 0xf4, 0x38, 0x7e, 0x94, 0x58, 0x10, 0x3b, 0x06, 0x57,
	0xa6, 0xa1, 0x33, 0x75, 0x66, 0x47, 0xdc, 0x95, 0x29, 0x3b, 0x85, 0xee, 0x52, 0x49, 0xd4, 0x14,
	0xba, 0x75, 0xad, 0x45, 0x55, 0xbd, 0xc8, 0x4a, 0xb3, 0xc4, 0xd0, 0x6b, 0xea, 0x0d, 0x8a, 0x3e,
	0x1d, 0x80, 0x47, 0x22, 0x23, 0x17, 0x25, 0x61, 0xc1, 0x26, 0xe0, 0xbf, 0x4b, 0x85, 0x5a, 0xac,
	0xb1, 0x15, 0xdd, 0x61, 0x76, 0x03, 0xfe, 0x1a, 0x49, 0xa4, 0x82, 0x44, 0xe8, 0x4e, 0xbd, 0x59,
	0x90, 0x8c, 0xe3, 0xfd, 0xd3, 0xf8, 0xb5, 0xe5, 0x9e, 0x34, 0x99, 0x2d, 0xdf, 0x5d, 0x9d, 0xdc,
	0xc3, 0xe0, 0x1b, 0xc5, 0x86, 0xe0, 0xad, 0x70, 0xdb, 0xca, 0x57, 0x47, 0x36, 0x82, 0xce, 0x46,
	0xa8, 0x12, 0xdb, 0xcc, 0x0d, 0xb8, 0x73, 0x6f, 0x9d, 0x68, 0x03, 0x27, 0x6f, 0x4b, 0xa1, 0xe7,
	0x52, 0xa1, 0xed, 0x98, 0xc1, 0xff, 0x3a, 0x42, 0xf5, 0xbe, 0xcf, 0xeb, 0x33, 0x8b, 0xa0, 0x67,
	0x1a, 0xba, 0x96, 0x08, 0x12, 0x3f, 0x6e, 0xaf, 0x73, 0x4b, 0xb0, 0x33, 0x00, 0xb1, 0x4b, 0x5b,
	0x4f, 0x21, 0x48, 0x82, 0x83, 0x06, 0xf8, 0x01, 0x1d, 0x51, 0xe3, 0xfb, 0x4c, 0x94, 0x5b, 0xdf,
	0x21, 0x78, 0xa5, 0x51, 0x36, 0x76, 0x69, 0xd4, 0xdf, 0xbb, 0x5e, 0x43, 0xbf, 0x72, 0xe5, 0x58,
	0xe4, 0x99, 0x2e, 0xf0, 0xd7, 0x72, 0x47, 0xd0, 0xc1, 0xcd, 0x7e, 0xb7, 0x0d, 0x48, 0x1e, 0xc0,
	0x9f, 0x9b, 0x4c, 0x13, 0xea, 0x94, 0x5d, 0x82, 0x6f, 0xe7, 0xc5, 0x86, 0xf1, 0x8f, 0xd1, 0x4d,
	0x06, 0xf1, 0xa1, 0x7c, 0xf4, 0x6f, 0xe6, 0x5c, 0x38, 0x8b, 0x6e, 0xfd, 0x4f, 0x5d, 0x7d, 0x05,
	0x00, 0x00, 0xff, 0xff, 0xdd, 0x00, 0xeb, 0x9a, 0x70, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// FrontendClient is the client API for Frontend service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FrontendClient interface {
	ScanFile(ctx context.Context, opts ...grpc.CallOption) (Frontend_ScanFileClient, error)
}

type frontendClient struct {
	cc *grpc.ClientConn
}

func NewFrontendClient(cc *grpc.ClientConn) FrontendClient {
	return &frontendClient{cc}
}

func (c *frontendClient) ScanFile(ctx context.Context, opts ...grpc.CallOption) (Frontend_ScanFileClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Frontend_serviceDesc.Streams[0], "/Frontend/ScanFile", opts...)
	if err != nil {
		return nil, err
	}
	x := &frontendScanFileClient{stream}
	return x, nil
}

type Frontend_ScanFileClient interface {
	Send(*ScanFileRequest) error
	Recv() (*ScanResponse, error)
	grpc.ClientStream
}

type frontendScanFileClient struct {
	grpc.ClientStream
}

func (x *frontendScanFileClient) Send(m *ScanFileRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *frontendScanFileClient) Recv() (*ScanResponse, error) {
	m := new(ScanResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// FrontendServer is the server API for Frontend service.
type FrontendServer interface {
	ScanFile(Frontend_ScanFileServer) error
}

func RegisterFrontendServer(s *grpc.Server, srv FrontendServer) {
	s.RegisterService(&_Frontend_serviceDesc, srv)
}

func _Frontend_ScanFile_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(FrontendServer).ScanFile(&frontendScanFileServer{stream})
}

type Frontend_ScanFileServer interface {
	Send(*ScanResponse) error
	Recv() (*ScanFileRequest, error)
	grpc.ServerStream
}

type frontendScanFileServer struct {
	grpc.ServerStream
}

func (x *frontendScanFileServer) Send(m *ScanResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *frontendScanFileServer) Recv() (*ScanFileRequest, error) {
	m := new(ScanFileRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Frontend_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Frontend",
	HandlerType: (*FrontendServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ScanFile",
			Handler:       _Frontend_ScanFile_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "api/strelka/strelka.proto",
}
