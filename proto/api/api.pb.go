// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/api.proto

package api

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

type Endpoint struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Host                 []string `protobuf:"bytes,2,rep,name=host,proto3" json:"host,omitempty"`
	Path                 []string `protobuf:"bytes,3,rep,name=path,proto3" json:"path,omitempty"`
	Method               []string `protobuf:"bytes,4,rep,name=method,proto3" json:"method,omitempty"`
	Stream               bool     `protobuf:"varint,5,opt,name=stream,proto3" json:"stream,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Endpoint) Reset()         { *m = Endpoint{} }
func (m *Endpoint) String() string { return proto.CompactTextString(m) }
func (*Endpoint) ProtoMessage()    {}
func (*Endpoint) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b40cafcd4234784, []int{0}
}

func (m *Endpoint) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Endpoint.Unmarshal(m, b)
}
func (m *Endpoint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Endpoint.Marshal(b, m, deterministic)
}
func (m *Endpoint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Endpoint.Merge(m, src)
}
func (m *Endpoint) XXX_Size() int {
	return xxx_messageInfo_Endpoint.Size(m)
}
func (m *Endpoint) XXX_DiscardUnknown() {
	xxx_messageInfo_Endpoint.DiscardUnknown(m)
}

var xxx_messageInfo_Endpoint proto.InternalMessageInfo

func (m *Endpoint) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Endpoint) GetHost() []string {
	if m != nil {
		return m.Host
	}
	return nil
}

func (m *Endpoint) GetPath() []string {
	if m != nil {
		return m.Path
	}
	return nil
}

func (m *Endpoint) GetMethod() []string {
	if m != nil {
		return m.Method
	}
	return nil
}

func (m *Endpoint) GetStream() bool {
	if m != nil {
		return m.Stream
	}
	return false
}

type EmptyResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EmptyResponse) Reset()         { *m = EmptyResponse{} }
func (m *EmptyResponse) String() string { return proto.CompactTextString(m) }
func (*EmptyResponse) ProtoMessage()    {}
func (*EmptyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b40cafcd4234784, []int{1}
}

func (m *EmptyResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EmptyResponse.Unmarshal(m, b)
}
func (m *EmptyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EmptyResponse.Marshal(b, m, deterministic)
}
func (m *EmptyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmptyResponse.Merge(m, src)
}
func (m *EmptyResponse) XXX_Size() int {
	return xxx_messageInfo_EmptyResponse.Size(m)
}
func (m *EmptyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EmptyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EmptyResponse proto.InternalMessageInfo

type Pair struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Values               []string `protobuf:"bytes,2,rep,name=values,proto3" json:"values,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Pair) Reset()         { *m = Pair{} }
func (m *Pair) String() string { return proto.CompactTextString(m) }
func (*Pair) ProtoMessage()    {}
func (*Pair) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b40cafcd4234784, []int{2}
}

func (m *Pair) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Pair.Unmarshal(m, b)
}
func (m *Pair) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Pair.Marshal(b, m, deterministic)
}
func (m *Pair) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pair.Merge(m, src)
}
func (m *Pair) XXX_Size() int {
	return xxx_messageInfo_Pair.Size(m)
}
func (m *Pair) XXX_DiscardUnknown() {
	xxx_messageInfo_Pair.DiscardUnknown(m)
}

var xxx_messageInfo_Pair proto.InternalMessageInfo

func (m *Pair) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Pair) GetValues() []string {
	if m != nil {
		return m.Values
	}
	return nil
}

type Request struct {
	Method               string           `protobuf:"bytes,1,opt,name=method,proto3" json:"method,omitempty"`
	Path                 string           `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	Header               map[string]*Pair `protobuf:"bytes,3,rep,name=header,proto3" json:"header,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Get                  map[string]*Pair `protobuf:"bytes,4,rep,name=get,proto3" json:"get,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Post                 map[string]*Pair `protobuf:"bytes,5,rep,name=post,proto3" json:"post,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Body                 string           `protobuf:"bytes,6,opt,name=body,proto3" json:"body,omitempty"`
	Url                  string           `protobuf:"bytes,7,opt,name=url,proto3" json:"url,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b40cafcd4234784, []int{3}
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

func (m *Request) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *Request) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *Request) GetHeader() map[string]*Pair {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Request) GetGet() map[string]*Pair {
	if m != nil {
		return m.Get
	}
	return nil
}

func (m *Request) GetPost() map[string]*Pair {
	if m != nil {
		return m.Post
	}
	return nil
}

func (m *Request) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func (m *Request) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

type Response struct {
	StatusCode           int32            `protobuf:"varint,1,opt,name=statusCode,proto3" json:"statusCode,omitempty"`
	Header               map[string]*Pair `protobuf:"bytes,2,rep,name=header,proto3" json:"header,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Body                 string           `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b40cafcd4234784, []int{4}
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

func (m *Response) GetStatusCode() int32 {
	if m != nil {
		return m.StatusCode
	}
	return 0
}

func (m *Response) GetHeader() map[string]*Pair {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Response) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

// A HTTP event as RPC
// Forwarded by the event handler
type Event struct {
	// e.g login
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// uuid
	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// unix timestamp of event
	Timestamp int64 `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// event headers
	Header map[string]*Pair `protobuf:"bytes,4,rep,name=header,proto3" json:"header,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// the event data
	Data                 string   `protobuf:"bytes,5,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()    {}
func (*Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b40cafcd4234784, []int{5}
}

func (m *Event) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Event.Unmarshal(m, b)
}
func (m *Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Event.Marshal(b, m, deterministic)
}
func (m *Event) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event.Merge(m, src)
}
func (m *Event) XXX_Size() int {
	return xxx_messageInfo_Event.Size(m)
}
func (m *Event) XXX_DiscardUnknown() {
	xxx_messageInfo_Event.DiscardUnknown(m)
}

var xxx_messageInfo_Event proto.InternalMessageInfo

func (m *Event) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Event) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Event) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *Event) GetHeader() map[string]*Pair {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Event) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func init() {
	proto.RegisterType((*Endpoint)(nil), "api.Endpoint")
	proto.RegisterType((*EmptyResponse)(nil), "api.EmptyResponse")
	proto.RegisterType((*Pair)(nil), "api.Pair")
	proto.RegisterType((*Request)(nil), "api.Request")
	proto.RegisterMapType((map[string]*Pair)(nil), "api.Request.GetEntry")
	proto.RegisterMapType((map[string]*Pair)(nil), "api.Request.HeaderEntry")
	proto.RegisterMapType((map[string]*Pair)(nil), "api.Request.PostEntry")
	proto.RegisterType((*Response)(nil), "api.Response")
	proto.RegisterMapType((map[string]*Pair)(nil), "api.Response.HeaderEntry")
	proto.RegisterType((*Event)(nil), "api.Event")
	proto.RegisterMapType((map[string]*Pair)(nil), "api.Event.HeaderEntry")
}

func init() { proto.RegisterFile("api/api.proto", fileDescriptor_1b40cafcd4234784) }

var fileDescriptor_1b40cafcd4234784 = []byte{
	// 501 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0x4d, 0x6f, 0x13, 0x31,
	0x10, 0x65, 0xbf, 0xd2, 0xec, 0x44, 0x01, 0x64, 0x89, 0x6a, 0x89, 0x10, 0x44, 0x7b, 0x69, 0xe1,
	0xb0, 0x69, 0xd3, 0x0b, 0x82, 0x53, 0x4b, 0x23, 0x38, 0x56, 0x7b, 0xe4, 0xe6, 0x74, 0xad, 0xac,
	0x45, 0x76, 0x6d, 0x6c, 0x6f, 0xa4, 0xfc, 0x2a, 0x8e, 0xfc, 0x10, 0xfe, 0x10, 0xf2, 0xd8, 0x9b,
	0x26, 0x52, 0x7b, 0x68, 0xd5, 0xcb, 0x6a, 0xe6, 0xed, 0x78, 0xfc, 0xde, 0x1b, 0xdb, 0x30, 0xa6,
	0x92, 0xcf, 0xa8, 0xe4, 0x85, 0x54, 0xc2, 0x08, 0x12, 0x51, 0xc9, 0xf3, 0x0d, 0x0c, 0x17, 0x6d,
	0x25, 0x05, 0x6f, 0x0d, 0x21, 0x10, 0xb7, 0xb4, 0x61, 0x59, 0x30, 0x0d, 0x4e, 0xd3, 0x12, 0x63,
	0x8b, 0xd5, 0x42, 0x9b, 0x2c, 0x9c, 0x46, 0x16, 0xb3, 0xb1, 0xc5, 0x24, 0x35, 0x75, 0x16, 0x39,
	0xcc, 0xc6, 0xe4, 0x18, 0x06, 0x0d, 0x33, 0xb5, 0xa8, 0xb2, 0x18, 0x51, 0x9f, 0x59, 0x5c, 0x1b,
	0xc5, 0x68, 0x93, 0x25, 0xd3, 0xe0, 0x74, 0x58, 0xfa, 0x2c, 0x7f, 0x05, 0xe3, 0x45, 0x23, 0xcd,
	0xb6, 0x64, 0x5a, 0x8a, 0x56, 0xb3, 0xfc, 0x0c, 0xe2, 0x1b, 0xca, 0x15, 0x79, 0x0d, 0xd1, 0x2f,
	0xb6, 0xf5, 0x1c, 0x6c, 0x68, 0x5b, 0x6c, 0xe8, 0xba, 0x63, 0xda, 0x93, 0xf0, 0x59, 0xfe, 0x27,
	0x82, 0xa3, 0x92, 0xfd, 0xee, 0x98, 0x36, 0x7b, 0xdb, 0xbb, 0x85, 0xfd, 0xf6, 0x3d, 0xd5, 0xd0,
	0x49, 0x42, 0xaa, 0x67, 0x30, 0xa8, 0x19, 0xad, 0x98, 0x42, 0x01, 0xa3, 0x79, 0x56, 0x58, 0x4f,
	0x7c, 0xa7, 0xe2, 0x07, 0xfe, 0x5a, 0xb4, 0x46, 0x6d, 0x4b, 0x5f, 0x47, 0x4e, 0x20, 0x5a, 0x31,
	0x83, 0xca, 0x46, 0xf3, 0x37, 0x07, 0xe5, 0xdf, 0x99, 0x71, 0xb5, 0xb6, 0x82, 0x7c, 0x82, 0x58,
	0x5a, 0xb7, 0x12, 0xac, 0x3c, 0x3e, 0xa8, 0xbc, 0x11, 0xda, 0x97, 0x62, 0x8d, 0xa5, 0xb6, 0x14,
	0xd5, 0x36, 0x1b, 0x38, 0x6a, 0x36, 0xb6, 0xe2, 0x3b, 0xb5, 0xce, 0x8e, 0x9c, 0xf8, 0x4e, 0xad,
	0x27, 0xd7, 0x30, 0xda, 0x63, 0x74, 0x8f, 0x3b, 0x1f, 0x20, 0x41, 0x3f, 0x50, 0xe2, 0x68, 0x9e,
	0xe2, 0x9e, 0xd6, 0xc9, 0xd2, 0xe1, 0x5f, 0xc2, 0xcf, 0xc1, 0xe4, 0x12, 0x86, 0x3d, 0xd1, 0xa7,
	0xb6, 0xb8, 0x82, 0x74, 0xa7, 0xe0, 0x89, 0x3d, 0xf2, 0xbf, 0x01, 0x0c, 0xfb, 0x81, 0x93, 0xf7,
	0x00, 0xda, 0x50, 0xd3, 0xe9, 0x6f, 0xa2, 0x72, 0x67, 0x2e, 0x29, 0xf7, 0x10, 0x72, 0xbe, 0x1b,
	0x53, 0x88, 0x6e, 0xbe, 0xf5, 0x6e, 0xba, 0xe5, 0xf7, 0xce, 0xa9, 0xb7, 0x34, 0xba, 0xb3, 0xf4,
	0x79, 0x0c, 0xcc, 0xff, 0x05, 0x90, 0x2c, 0x36, 0xec, 0x81, 0x4b, 0xf2, 0x12, 0x42, 0x5e, 0xf9,
	0x33, 0x16, 0xf2, 0x8a, 0xbc, 0x83, 0xd4, 0xf0, 0x86, 0x69, 0x43, 0x1b, 0x89, 0x64, 0xa2, 0xf2,
	0x0e, 0x20, 0xc5, 0x4e, 0x58, 0xbc, 0x77, 0x4c, 0xb0, 0xfb, 0x43, 0xaa, 0x2a, 0x6a, 0x28, 0x5e,
	0xa0, 0xb4, 0xc4, 0xf8, 0x79, 0x54, 0xcd, 0x39, 0x44, 0x97, 0x92, 0x93, 0x99, 0x9d, 0xca, 0x8a,
	0x6b, 0xc3, 0x14, 0x19, 0x3b, 0x32, 0xfe, 0x49, 0x98, 0x10, 0x97, 0x1e, 0xdc, 0xd4, 0x17, 0xe4,
	0x1c, 0xe0, 0x9a, 0xa9, 0xc7, 0x2c, 0xb9, 0xfa, 0xf8, 0xf3, 0x64, 0xc5, 0x4d, 0xdd, 0x2d, 0x8b,
	0x5b, 0xd1, 0xcc, 0x1a, 0x7e, 0xab, 0x84, 0xff, 0x6e, 0x2e, 0x66, 0xf8, 0x1e, 0xd9, 0x97, 0xe9,
	0x2b, 0x95, 0x7c, 0x39, 0xc0, 0xf4, 0xe2, 0x7f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xaa, 0x47, 0x03,
	0x6e, 0xaf, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ApiClient is the client API for Api service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ApiClient interface {
	Register(ctx context.Context, in *Endpoint, opts ...grpc.CallOption) (*EmptyResponse, error)
	Deregister(ctx context.Context, in *Endpoint, opts ...grpc.CallOption) (*EmptyResponse, error)
}

type apiClient struct {
	cc *grpc.ClientConn
}

func NewApiClient(cc *grpc.ClientConn) ApiClient {
	return &apiClient{cc}
}

func (c *apiClient) Register(ctx context.Context, in *Endpoint, opts ...grpc.CallOption) (*EmptyResponse, error) {
	out := new(EmptyResponse)
	err := c.cc.Invoke(ctx, "/api.Api/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) Deregister(ctx context.Context, in *Endpoint, opts ...grpc.CallOption) (*EmptyResponse, error) {
	out := new(EmptyResponse)
	err := c.cc.Invoke(ctx, "/api.Api/Deregister", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ApiServer is the server API for Api service.
type ApiServer interface {
	Register(context.Context, *Endpoint) (*EmptyResponse, error)
	Deregister(context.Context, *Endpoint) (*EmptyResponse, error)
}

func RegisterApiServer(s *grpc.Server, srv ApiServer) {
	s.RegisterService(&_Api_serviceDesc, srv)
}

func _Api_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Endpoint)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Api/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).Register(ctx, req.(*Endpoint))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_Deregister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Endpoint)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).Deregister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Api/Deregister",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).Deregister(ctx, req.(*Endpoint))
	}
	return interceptor(ctx, in, info, handler)
}

var _Api_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.Api",
	HandlerType: (*ApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _Api_Register_Handler,
		},
		{
			MethodName: "Deregister",
			Handler:    _Api_Deregister_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/api.proto",
}
