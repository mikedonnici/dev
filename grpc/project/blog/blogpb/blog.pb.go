// Code generated by protoc-gen-go. DO NOT EDIT.
// source: blog.proto

package blogpb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Post struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	AuthorId             string   `protobuf:"bytes,2,opt,name=author_id,json=authorId,proto3" json:"author_id,omitempty"`
	Title                string   `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Content              string   `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Post) Reset()         { *m = Post{} }
func (m *Post) String() string { return proto.CompactTextString(m) }
func (*Post) ProtoMessage()    {}
func (*Post) Descriptor() ([]byte, []int) {
	return fileDescriptor_6745b25902462fb1, []int{0}
}

func (m *Post) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Post.Unmarshal(m, b)
}
func (m *Post) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Post.Marshal(b, m, deterministic)
}
func (m *Post) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Post.Merge(m, src)
}
func (m *Post) XXX_Size() int {
	return xxx_messageInfo_Post.Size(m)
}
func (m *Post) XXX_DiscardUnknown() {
	xxx_messageInfo_Post.DiscardUnknown(m)
}

var xxx_messageInfo_Post proto.InternalMessageInfo

func (m *Post) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Post) GetAuthorId() string {
	if m != nil {
		return m.AuthorId
	}
	return ""
}

func (m *Post) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Post) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

type CreatePostRequest struct {
	Post                 *Post    `protobuf:"bytes,1,opt,name=post,proto3" json:"post,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreatePostRequest) Reset()         { *m = CreatePostRequest{} }
func (m *CreatePostRequest) String() string { return proto.CompactTextString(m) }
func (*CreatePostRequest) ProtoMessage()    {}
func (*CreatePostRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6745b25902462fb1, []int{1}
}

func (m *CreatePostRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreatePostRequest.Unmarshal(m, b)
}
func (m *CreatePostRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreatePostRequest.Marshal(b, m, deterministic)
}
func (m *CreatePostRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreatePostRequest.Merge(m, src)
}
func (m *CreatePostRequest) XXX_Size() int {
	return xxx_messageInfo_CreatePostRequest.Size(m)
}
func (m *CreatePostRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreatePostRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreatePostRequest proto.InternalMessageInfo

func (m *CreatePostRequest) GetPost() *Post {
	if m != nil {
		return m.Post
	}
	return nil
}

type CreatePostResponse struct {
	Post                 *Post    `protobuf:"bytes,1,opt,name=post,proto3" json:"post,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreatePostResponse) Reset()         { *m = CreatePostResponse{} }
func (m *CreatePostResponse) String() string { return proto.CompactTextString(m) }
func (*CreatePostResponse) ProtoMessage()    {}
func (*CreatePostResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6745b25902462fb1, []int{2}
}

func (m *CreatePostResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreatePostResponse.Unmarshal(m, b)
}
func (m *CreatePostResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreatePostResponse.Marshal(b, m, deterministic)
}
func (m *CreatePostResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreatePostResponse.Merge(m, src)
}
func (m *CreatePostResponse) XXX_Size() int {
	return xxx_messageInfo_CreatePostResponse.Size(m)
}
func (m *CreatePostResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreatePostResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreatePostResponse proto.InternalMessageInfo

func (m *CreatePostResponse) GetPost() *Post {
	if m != nil {
		return m.Post
	}
	return nil
}

func init() {
	proto.RegisterType((*Post)(nil), "blog.Post")
	proto.RegisterType((*CreatePostRequest)(nil), "blog.CreatePostRequest")
	proto.RegisterType((*CreatePostResponse)(nil), "blog.CreatePostResponse")
}

func init() { proto.RegisterFile("blog.proto", fileDescriptor_6745b25902462fb1) }

var fileDescriptor_6745b25902462fb1 = []byte{
	// 217 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0x41, 0x4b, 0x03, 0x31,
	0x10, 0x85, 0xed, 0x1a, 0x6b, 0x3b, 0x05, 0xc1, 0x41, 0x30, 0x28, 0x88, 0xe4, 0xe4, 0xa9, 0x87,
	0xd6, 0x3f, 0x60, 0x3d, 0x79, 0x2b, 0xf5, 0xe6, 0x45, 0xba, 0x9b, 0x61, 0x0d, 0x2c, 0x3b, 0x31,
	0x99, 0xf5, 0xf7, 0xcb, 0x4e, 0x10, 0x85, 0x3d, 0x78, 0xcb, 0x7c, 0x6f, 0x5e, 0xde, 0x63, 0x00,
	0xea, 0x8e, 0xdb, 0x75, 0x4c, 0x2c, 0x8c, 0x66, 0x7c, 0xbb, 0x06, 0xcc, 0x9e, 0xb3, 0xe0, 0x05,
	0x54, 0xc1, 0xdb, 0xd9, 0xfd, 0xec, 0x61, 0x79, 0xa8, 0x82, 0xc7, 0x5b, 0x58, 0x1e, 0x07, 0xf9,
	0xe0, 0xf4, 0x1e, 0xbc, 0xad, 0x14, 0x2f, 0x0a, 0x78, 0xf1, 0x78, 0x05, 0x67, 0x12, 0xa4, 0x23,
	0x7b, 0xaa, 0x42, 0x19, 0xd0, 0xc2, 0x79, 0xc3, 0xbd, 0x50, 0x2f, 0xd6, 0x28, 0xff, 0x19, 0xdd,
	0x16, 0x2e, 0x9f, 0x13, 0x1d, 0x85, 0xc6, 0xa8, 0x03, 0x7d, 0x0e, 0x94, 0x05, 0xef, 0xc0, 0x44,
	0xce, 0xa2, 0x99, 0xab, 0x0d, 0xac, 0xb5, 0x9a, 0x2e, 0x28, 0x77, 0x8f, 0x80, 0x7f, 0x4d, 0x39,
	0x72, 0x9f, 0xe9, 0x3f, 0xd7, 0x66, 0x0f, 0xab, 0x5d, 0xc7, 0xed, 0x2b, 0xa5, 0xaf, 0xd0, 0x10,
	0x3e, 0x01, 0xfc, 0x7e, 0x82, 0xd7, 0x65, 0x7d, 0xd2, 0xe5, 0xc6, 0x4e, 0x85, 0x92, 0xe7, 0x4e,
	0x76, 0x8b, 0xb7, 0xf9, 0x28, 0xc6, 0xba, 0x9e, 0xeb, 0xe1, 0xb6, 0xdf, 0x01, 0x00, 0x00, 0xff,
	0xff, 0x8b, 0xba, 0x83, 0x80, 0x46, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BlogServiceClient is the client API for BlogService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BlogServiceClient interface {
	CreatePost(ctx context.Context, in *CreatePostRequest, opts ...grpc.CallOption) (*CreatePostResponse, error)
}

type blogServiceClient struct {
	cc *grpc.ClientConn
}

func NewBlogServiceClient(cc *grpc.ClientConn) BlogServiceClient {
	return &blogServiceClient{cc}
}

func (c *blogServiceClient) CreatePost(ctx context.Context, in *CreatePostRequest, opts ...grpc.CallOption) (*CreatePostResponse, error) {
	out := new(CreatePostResponse)
	err := c.cc.Invoke(ctx, "/blog.BlogService/CreatePost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BlogServiceServer is the server API for BlogService service.
type BlogServiceServer interface {
	CreatePost(context.Context, *CreatePostRequest) (*CreatePostResponse, error)
}

// UnimplementedBlogServiceServer can be embedded to have forward compatible implementations.
type UnimplementedBlogServiceServer struct {
}

func (*UnimplementedBlogServiceServer) CreatePost(ctx context.Context, req *CreatePostRequest) (*CreatePostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePost not implemented")
}

func RegisterBlogServiceServer(s *grpc.Server, srv BlogServiceServer) {
	s.RegisterService(&_BlogService_serviceDesc, srv)
}

func _BlogService_CreatePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogServiceServer).CreatePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blog.BlogService/CreatePost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogServiceServer).CreatePost(ctx, req.(*CreatePostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _BlogService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "blog.BlogService",
	HandlerType: (*BlogServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePost",
			Handler:    _BlogService_CreatePost_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "blog.proto",
}
