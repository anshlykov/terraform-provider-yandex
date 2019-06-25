// Code generated by protoc-gen-go. DO NOT EDIT.
// source: yandex/cloud/mdb/mongodb/v1/backup_service.proto

package mongodb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/yandex-cloud/go-genproto/yandex/cloud/validation"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type GetBackupRequest struct {
	// ID of the backup to return information about.
	// To get the backup ID, use a [ClusterService.ListBackups] request.
	BackupId             string   `protobuf:"bytes,1,opt,name=backup_id,json=backupId,proto3" json:"backup_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetBackupRequest) Reset()         { *m = GetBackupRequest{} }
func (m *GetBackupRequest) String() string { return proto.CompactTextString(m) }
func (*GetBackupRequest) ProtoMessage()    {}
func (*GetBackupRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_11029ddc87f407e2, []int{0}
}

func (m *GetBackupRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBackupRequest.Unmarshal(m, b)
}
func (m *GetBackupRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBackupRequest.Marshal(b, m, deterministic)
}
func (m *GetBackupRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBackupRequest.Merge(m, src)
}
func (m *GetBackupRequest) XXX_Size() int {
	return xxx_messageInfo_GetBackupRequest.Size(m)
}
func (m *GetBackupRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBackupRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetBackupRequest proto.InternalMessageInfo

func (m *GetBackupRequest) GetBackupId() string {
	if m != nil {
		return m.BackupId
	}
	return ""
}

type ListBackupsRequest struct {
	// ID of the folder to list backups in.
	// To get the folder ID, use a [yandex.cloud.resourcemanager.v1.FolderService.List] request.
	FolderId string `protobuf:"bytes,1,opt,name=folder_id,json=folderId,proto3" json:"folder_id,omitempty"`
	PageSize int64  `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// Page token. To get the next page of results, set [page_token] to the
	// [ListBackupsResponse.next_page_token] returned by a previous list request.
	PageToken            string   `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListBackupsRequest) Reset()         { *m = ListBackupsRequest{} }
func (m *ListBackupsRequest) String() string { return proto.CompactTextString(m) }
func (*ListBackupsRequest) ProtoMessage()    {}
func (*ListBackupsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_11029ddc87f407e2, []int{1}
}

func (m *ListBackupsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListBackupsRequest.Unmarshal(m, b)
}
func (m *ListBackupsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListBackupsRequest.Marshal(b, m, deterministic)
}
func (m *ListBackupsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListBackupsRequest.Merge(m, src)
}
func (m *ListBackupsRequest) XXX_Size() int {
	return xxx_messageInfo_ListBackupsRequest.Size(m)
}
func (m *ListBackupsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListBackupsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListBackupsRequest proto.InternalMessageInfo

func (m *ListBackupsRequest) GetFolderId() string {
	if m != nil {
		return m.FolderId
	}
	return ""
}

func (m *ListBackupsRequest) GetPageSize() int64 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListBackupsRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

type ListBackupsResponse struct {
	// List of Backup resources.
	Backups []*Backup `protobuf:"bytes,1,rep,name=backups,proto3" json:"backups,omitempty"`
	// This token allows you to get the next page of results for list requests. If the number of results
	// is larger than [ListBackupsRequest.page_size], use the [next_page_token] as the value
	// for the [ListBackupsRequest.page_token] parameter in the next list request. Each subsequent
	// list request will have its own [next_page_token] to continue paging through the results.
	NextPageToken        string   `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListBackupsResponse) Reset()         { *m = ListBackupsResponse{} }
func (m *ListBackupsResponse) String() string { return proto.CompactTextString(m) }
func (*ListBackupsResponse) ProtoMessage()    {}
func (*ListBackupsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_11029ddc87f407e2, []int{2}
}

func (m *ListBackupsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListBackupsResponse.Unmarshal(m, b)
}
func (m *ListBackupsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListBackupsResponse.Marshal(b, m, deterministic)
}
func (m *ListBackupsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListBackupsResponse.Merge(m, src)
}
func (m *ListBackupsResponse) XXX_Size() int {
	return xxx_messageInfo_ListBackupsResponse.Size(m)
}
func (m *ListBackupsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListBackupsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListBackupsResponse proto.InternalMessageInfo

func (m *ListBackupsResponse) GetBackups() []*Backup {
	if m != nil {
		return m.Backups
	}
	return nil
}

func (m *ListBackupsResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

func init() {
	proto.RegisterType((*GetBackupRequest)(nil), "yandex.cloud.mdb.mongodb.v1.GetBackupRequest")
	proto.RegisterType((*ListBackupsRequest)(nil), "yandex.cloud.mdb.mongodb.v1.ListBackupsRequest")
	proto.RegisterType((*ListBackupsResponse)(nil), "yandex.cloud.mdb.mongodb.v1.ListBackupsResponse")
}

func init() {
	proto.RegisterFile("yandex/cloud/mdb/mongodb/v1/backup_service.proto", fileDescriptor_11029ddc87f407e2)
}

var fileDescriptor_11029ddc87f407e2 = []byte{
	// 459 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x31, 0x6f, 0xd3, 0x40,
	0x14, 0xc7, 0xe5, 0x24, 0x94, 0xf8, 0xa0, 0x02, 0x1d, 0x4b, 0x94, 0x52, 0x29, 0xb8, 0x12, 0x75,
	0x87, 0xf8, 0xec, 0xa2, 0x4e, 0x34, 0x4b, 0x96, 0x28, 0x12, 0x03, 0x72, 0x99, 0x58, 0xa2, 0x73,
	0xee, 0x71, 0x9c, 0x1a, 0xdf, 0x99, 0xdc, 0xc5, 0x2a, 0x05, 0x16, 0xc6, 0x0c, 0x0c, 0xf0, 0x39,
	0xf8, 0x1c, 0xed, 0xce, 0x57, 0x60, 0xe0, 0x33, 0x30, 0x21, 0xdf, 0x39, 0x40, 0xa9, 0x64, 0xba,
	0x9d, 0xee, 0xff, 0x7e, 0xef, 0xfd, 0xf5, 0x7f, 0x0f, 0xc5, 0x6f, 0xa9, 0x64, 0x70, 0x46, 0xe6,
	0x0b, 0xb5, 0x62, 0x24, 0x67, 0x19, 0xc9, 0x95, 0xe4, 0x8a, 0x65, 0xa4, 0x4c, 0x48, 0x46, 0xe7,
	0xa7, 0xab, 0x62, 0xa6, 0x61, 0x59, 0x8a, 0x39, 0x44, 0xc5, 0x52, 0x19, 0x85, 0x77, 0x1c, 0x11,
	0x59, 0x22, 0xca, 0x59, 0x16, 0xd5, 0x44, 0x54, 0x26, 0xfd, 0x87, 0x5c, 0x29, 0xbe, 0x00, 0x42,
	0x0b, 0x41, 0xa8, 0x94, 0xca, 0x50, 0x23, 0x94, 0xd4, 0x0e, 0xed, 0xef, 0x5e, 0x19, 0x56, 0xd2,
	0x85, 0x60, 0x56, 0xaf, 0xe5, 0xf0, 0xff, 0x5e, 0x5c, 0x65, 0x70, 0x84, 0xee, 0x4f, 0xc0, 0x8c,
	0xed, 0x57, 0x0a, 0x6f, 0x56, 0xa0, 0x0d, 0x7e, 0x84, 0xfc, 0xda, 0xaf, 0x60, 0x3d, 0x6f, 0xe0,
	0x85, 0xfe, 0xb8, 0xf3, 0xe3, 0x22, 0xf1, 0xd2, 0xae, 0xfb, 0x9e, 0xb2, 0xe0, 0xb3, 0x87, 0xf0,
	0x33, 0xa1, 0x6b, 0x50, 0x6f, 0xc8, 0x03, 0xe4, 0xbf, 0x52, 0x0b, 0x06, 0xcb, 0x3f, 0xe4, 0xdd,
	0x8a, 0x5c, 0x5f, 0x26, 0x9d, 0xe3, 0xd1, 0x51, 0x9c, 0x76, 0x9d, 0x3c, 0x65, 0x78, 0x1f, 0xf9,
	0x05, 0xe5, 0x30, 0xd3, 0xe2, 0x1c, 0x7a, 0xad, 0x81, 0x17, 0xb6, 0xc7, 0xe8, 0xe7, 0x45, 0xb2,
	0x75, 0x3c, 0x4a, 0xe2, 0x38, 0x4e, 0xbb, 0x95, 0x78, 0x22, 0xce, 0x01, 0x87, 0x08, 0xd9, 0x42,
	0xa3, 0x4e, 0x41, 0xf6, 0xda, 0xb6, 0xa9, 0xbf, 0xbe, 0x4c, 0x6e, 0xd9, 0xca, 0xd4, 0x76, 0x79,
	0x51, 0x69, 0xc1, 0x7b, 0xf4, 0xe0, 0x8a, 0x27, 0x5d, 0x28, 0xa9, 0x01, 0x8f, 0xd0, 0x6d, 0xe7,
	0x5b, 0xf7, 0xbc, 0x41, 0x3b, 0xbc, 0x73, 0xb8, 0x17, 0x35, 0x04, 0x1f, 0xd5, 0x59, 0x6c, 0x18,
	0xfc, 0x18, 0xdd, 0x93, 0x70, 0x66, 0x66, 0x7f, 0x99, 0xa8, 0xec, 0xfa, 0xe9, 0x76, 0xf5, 0xfd,
	0x7c, 0x33, 0xfd, 0xf0, 0x6b, 0x0b, 0x6d, 0x3b, 0xf6, 0xc4, 0x6d, 0x19, 0xaf, 0x3d, 0xd4, 0x9e,
	0x80, 0xc1, 0xc3, 0xc6, 0x79, 0xff, 0xc6, 0xdf, 0xbf, 0x89, 0xbd, 0x80, 0x7c, 0xfc, 0xf6, 0xfd,
	0x4b, 0xeb, 0x00, 0xef, 0x93, 0x9c, 0x4a, 0xca, 0x81, 0x0d, 0xaf, 0x6d, 0x58, 0x93, 0x77, 0xbf,
	0xd7, 0xf8, 0x01, 0x7f, 0xf2, 0x50, 0xa7, 0x4a, 0x07, 0x93, 0xc6, 0xf6, 0xd7, 0x97, 0xda, 0x8f,
	0x6f, 0x0e, 0xb8, 0xc4, 0x83, 0x3d, 0x6b, 0x6e, 0x17, 0xef, 0x34, 0x98, 0x1b, 0x4f, 0x5f, 0x4e,
	0xb8, 0x30, 0xaf, 0x57, 0x59, 0x34, 0x57, 0x39, 0x71, 0x23, 0x86, 0xee, 0x60, 0xb9, 0x1a, 0x72,
	0x90, 0xf6, 0x40, 0x49, 0xc3, 0x25, 0x3f, 0xad, 0x9f, 0xd9, 0x96, 0x2d, 0x7d, 0xf2, 0x2b, 0x00,
	0x00, 0xff, 0xff, 0x1f, 0x32, 0xcb, 0x4a, 0x83, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BackupServiceClient is the client API for BackupService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BackupServiceClient interface {
	// Returns the specified MongoDB Backup resource.
	//
	// To get the list of available MongoDB Backup resources, make a [List] request.
	Get(ctx context.Context, in *GetBackupRequest, opts ...grpc.CallOption) (*Backup, error)
	// Retrieves the list of Backup resources available for the specified folder.
	List(ctx context.Context, in *ListBackupsRequest, opts ...grpc.CallOption) (*ListBackupsResponse, error)
}

type backupServiceClient struct {
	cc *grpc.ClientConn
}

func NewBackupServiceClient(cc *grpc.ClientConn) BackupServiceClient {
	return &backupServiceClient{cc}
}

func (c *backupServiceClient) Get(ctx context.Context, in *GetBackupRequest, opts ...grpc.CallOption) (*Backup, error) {
	out := new(Backup)
	err := c.cc.Invoke(ctx, "/yandex.cloud.mdb.mongodb.v1.BackupService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *backupServiceClient) List(ctx context.Context, in *ListBackupsRequest, opts ...grpc.CallOption) (*ListBackupsResponse, error) {
	out := new(ListBackupsResponse)
	err := c.cc.Invoke(ctx, "/yandex.cloud.mdb.mongodb.v1.BackupService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BackupServiceServer is the server API for BackupService service.
type BackupServiceServer interface {
	// Returns the specified MongoDB Backup resource.
	//
	// To get the list of available MongoDB Backup resources, make a [List] request.
	Get(context.Context, *GetBackupRequest) (*Backup, error)
	// Retrieves the list of Backup resources available for the specified folder.
	List(context.Context, *ListBackupsRequest) (*ListBackupsResponse, error)
}

// UnimplementedBackupServiceServer can be embedded to have forward compatible implementations.
type UnimplementedBackupServiceServer struct {
}

func (*UnimplementedBackupServiceServer) Get(ctx context.Context, req *GetBackupRequest) (*Backup, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedBackupServiceServer) List(ctx context.Context, req *ListBackupsRequest) (*ListBackupsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}

func RegisterBackupServiceServer(s *grpc.Server, srv BackupServiceServer) {
	s.RegisterService(&_BackupService_serviceDesc, srv)
}

func _BackupService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBackupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackupServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.mdb.mongodb.v1.BackupService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackupServiceServer).Get(ctx, req.(*GetBackupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BackupService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListBackupsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackupServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.mdb.mongodb.v1.BackupService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackupServiceServer).List(ctx, req.(*ListBackupsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _BackupService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "yandex.cloud.mdb.mongodb.v1.BackupService",
	HandlerType: (*BackupServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _BackupService_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _BackupService_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "yandex/cloud/mdb/mongodb/v1/backup_service.proto",
}
