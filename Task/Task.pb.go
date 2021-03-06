// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.1
// source: Task.proto

package Task

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Task struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CompId        int32 `protobuf:"varint,1,opt,name=comp_id,json=compId,proto3" json:"comp_id,omitempty"`
	SeqNo         int32 `protobuf:"varint,2,opt,name=seq_no,json=seqNo,proto3" json:"seq_no,omitempty"`
	TimeToExecute int32 `protobuf:"varint,3,opt,name=time_to_execute,json=timeToExecute,proto3" json:"time_to_execute,omitempty"`
}

func (x *Task) Reset() {
	*x = Task{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Task_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Task) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Task) ProtoMessage() {}

func (x *Task) ProtoReflect() protoreflect.Message {
	mi := &file_Task_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Task.ProtoReflect.Descriptor instead.
func (*Task) Descriptor() ([]byte, []int) {
	return file_Task_proto_rawDescGZIP(), []int{0}
}

func (x *Task) GetCompId() int32 {
	if x != nil {
		return x.CompId
	}
	return 0
}

func (x *Task) GetSeqNo() int32 {
	if x != nil {
		return x.SeqNo
	}
	return 0
}

func (x *Task) GetTimeToExecute() int32 {
	if x != nil {
		return x.TimeToExecute
	}
	return 0
}

type TaskAddAck struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *TaskAddAck) Reset() {
	*x = TaskAddAck{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Task_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskAddAck) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskAddAck) ProtoMessage() {}

func (x *TaskAddAck) ProtoReflect() protoreflect.Message {
	mi := &file_Task_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskAddAck.ProtoReflect.Descriptor instead.
func (*TaskAddAck) Descriptor() ([]byte, []int) {
	return file_Task_proto_rawDescGZIP(), []int{1}
}

func (x *TaskAddAck) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type TaskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CompId   int32 `protobuf:"varint,1,opt,name=comp_id,json=compId,proto3" json:"comp_id,omitempty"`
	NumTasks int32 `protobuf:"varint,2,opt,name=num_tasks,json=numTasks,proto3" json:"num_tasks,omitempty"`
}

func (x *TaskRequest) Reset() {
	*x = TaskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Task_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskRequest) ProtoMessage() {}

func (x *TaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_Task_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskRequest.ProtoReflect.Descriptor instead.
func (*TaskRequest) Descriptor() ([]byte, []int) {
	return file_Task_proto_rawDescGZIP(), []int{2}
}

func (x *TaskRequest) GetCompId() int32 {
	if x != nil {
		return x.CompId
	}
	return 0
}

func (x *TaskRequest) GetNumTasks() int32 {
	if x != nil {
		return x.NumTasks
	}
	return 0
}

type TaskResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count uint32  `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	Tasks []*Task `protobuf:"bytes,2,rep,name=tasks,proto3" json:"tasks,omitempty"`
}

func (x *TaskResponse) Reset() {
	*x = TaskResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Task_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskResponse) ProtoMessage() {}

func (x *TaskResponse) ProtoReflect() protoreflect.Message {
	mi := &file_Task_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskResponse.ProtoReflect.Descriptor instead.
func (*TaskResponse) Descriptor() ([]byte, []int) {
	return file_Task_proto_rawDescGZIP(), []int{3}
}

func (x *TaskResponse) GetCount() uint32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *TaskResponse) GetTasks() []*Task {
	if x != nil {
		return x.Tasks
	}
	return nil
}

var File_Task_proto protoreflect.FileDescriptor

var file_Task_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x54, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x54, 0x61,
	0x73, 0x6b, 0x22, 0x5e, 0x0a, 0x04, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x17, 0x0a, 0x07, 0x63, 0x6f,
	0x6d, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x63, 0x6f, 0x6d,
	0x70, 0x49, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x73, 0x65, 0x71, 0x5f, 0x6e, 0x6f, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x73, 0x65, 0x71, 0x4e, 0x6f, 0x12, 0x26, 0x0a, 0x0f, 0x74, 0x69,
	0x6d, 0x65, 0x5f, 0x74, 0x6f, 0x5f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0d, 0x74, 0x69, 0x6d, 0x65, 0x54, 0x6f, 0x45, 0x78, 0x65, 0x63, 0x75,
	0x74, 0x65, 0x22, 0x24, 0x0a, 0x0a, 0x54, 0x61, 0x73, 0x6b, 0x41, 0x64, 0x64, 0x41, 0x63, 0x6b,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x43, 0x0a, 0x0b, 0x54, 0x61, 0x73, 0x6b,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x70, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x63, 0x6f, 0x6d, 0x70, 0x49, 0x64,
	0x12, 0x1b, 0x0a, 0x09, 0x6e, 0x75, 0x6d, 0x5f, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x08, 0x6e, 0x75, 0x6d, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x22, 0x46, 0x0a,
	0x0c, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x20, 0x0a, 0x05, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x05,
	0x74, 0x61, 0x73, 0x6b, 0x73, 0x32, 0xa2, 0x01, 0x0a, 0x0b, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x29, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x54, 0x61, 0x73, 0x6b,
	0x12, 0x0a, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x1a, 0x10, 0x2e, 0x54,
	0x61, 0x73, 0x6b, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x41, 0x64, 0x64, 0x41, 0x63, 0x6b, 0x28, 0x01,
	0x12, 0x32, 0x0a, 0x0d, 0x67, 0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x12, 0x11, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x0a, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x2e, 0x54, 0x61, 0x73, 0x6b,
	0x28, 0x01, 0x30, 0x01, 0x12, 0x34, 0x0a, 0x0b, 0x67, 0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x50,
	0x61, 0x67, 0x65, 0x12, 0x11, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x2e, 0x54, 0x61,
	0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x31, 0x5a, 0x2f, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x53, 0x61, 0x75, 0x72, 0x61, 0x76, 0x73,
	0x6b, 0x76, 0x30, 0x37, 0x2f, 0x62, 0x65, 0x6e, 0x63, 0x68, 0x6d, 0x61, 0x72, 0x6b, 0x5f, 0x67,
	0x72, 0x70, 0x63, 0x5f, 0x74, 0x61, 0x73, 0x6b, 0x2f, 0x54, 0x61, 0x73, 0x6b, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Task_proto_rawDescOnce sync.Once
	file_Task_proto_rawDescData = file_Task_proto_rawDesc
)

func file_Task_proto_rawDescGZIP() []byte {
	file_Task_proto_rawDescOnce.Do(func() {
		file_Task_proto_rawDescData = protoimpl.X.CompressGZIP(file_Task_proto_rawDescData)
	})
	return file_Task_proto_rawDescData
}

var file_Task_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_Task_proto_goTypes = []interface{}{
	(*Task)(nil),         // 0: Task.Task
	(*TaskAddAck)(nil),   // 1: Task.TaskAddAck
	(*TaskRequest)(nil),  // 2: Task.TaskRequest
	(*TaskResponse)(nil), // 3: Task.TaskResponse
}
var file_Task_proto_depIdxs = []int32{
	0, // 0: Task.TaskResponse.tasks:type_name -> Task.Task
	0, // 1: Task.TaskService.addTask:input_type -> Task.Task
	2, // 2: Task.TaskService.getTaskStream:input_type -> Task.TaskRequest
	2, // 3: Task.TaskService.getTaskPage:input_type -> Task.TaskRequest
	1, // 4: Task.TaskService.addTask:output_type -> Task.TaskAddAck
	0, // 5: Task.TaskService.getTaskStream:output_type -> Task.Task
	3, // 6: Task.TaskService.getTaskPage:output_type -> Task.TaskResponse
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_Task_proto_init() }
func file_Task_proto_init() {
	if File_Task_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Task_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Task); i {
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
		file_Task_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskAddAck); i {
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
		file_Task_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskRequest); i {
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
		file_Task_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskResponse); i {
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
			RawDescriptor: file_Task_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_Task_proto_goTypes,
		DependencyIndexes: file_Task_proto_depIdxs,
		MessageInfos:      file_Task_proto_msgTypes,
	}.Build()
	File_Task_proto = out.File
	file_Task_proto_rawDesc = nil
	file_Task_proto_goTypes = nil
	file_Task_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// TaskServiceClient is the client API for TaskService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TaskServiceClient interface {
	AddTask(ctx context.Context, opts ...grpc.CallOption) (TaskService_AddTaskClient, error)
	GetTaskStream(ctx context.Context, opts ...grpc.CallOption) (TaskService_GetTaskStreamClient, error)
	GetTaskPage(ctx context.Context, in *TaskRequest, opts ...grpc.CallOption) (*TaskResponse, error)
}

type taskServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTaskServiceClient(cc grpc.ClientConnInterface) TaskServiceClient {
	return &taskServiceClient{cc}
}

func (c *taskServiceClient) AddTask(ctx context.Context, opts ...grpc.CallOption) (TaskService_AddTaskClient, error) {
	stream, err := c.cc.NewStream(ctx, &_TaskService_serviceDesc.Streams[0], "/Task.TaskService/addTask", opts...)
	if err != nil {
		return nil, err
	}
	x := &taskServiceAddTaskClient{stream}
	return x, nil
}

type TaskService_AddTaskClient interface {
	Send(*Task) error
	CloseAndRecv() (*TaskAddAck, error)
	grpc.ClientStream
}

type taskServiceAddTaskClient struct {
	grpc.ClientStream
}

func (x *taskServiceAddTaskClient) Send(m *Task) error {
	return x.ClientStream.SendMsg(m)
}

func (x *taskServiceAddTaskClient) CloseAndRecv() (*TaskAddAck, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(TaskAddAck)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *taskServiceClient) GetTaskStream(ctx context.Context, opts ...grpc.CallOption) (TaskService_GetTaskStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_TaskService_serviceDesc.Streams[1], "/Task.TaskService/getTaskStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &taskServiceGetTaskStreamClient{stream}
	return x, nil
}

type TaskService_GetTaskStreamClient interface {
	Send(*TaskRequest) error
	Recv() (*Task, error)
	grpc.ClientStream
}

type taskServiceGetTaskStreamClient struct {
	grpc.ClientStream
}

func (x *taskServiceGetTaskStreamClient) Send(m *TaskRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *taskServiceGetTaskStreamClient) Recv() (*Task, error) {
	m := new(Task)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *taskServiceClient) GetTaskPage(ctx context.Context, in *TaskRequest, opts ...grpc.CallOption) (*TaskResponse, error) {
	out := new(TaskResponse)
	err := c.cc.Invoke(ctx, "/Task.TaskService/getTaskPage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TaskServiceServer is the server API for TaskService service.
type TaskServiceServer interface {
	AddTask(TaskService_AddTaskServer) error
	GetTaskStream(TaskService_GetTaskStreamServer) error
	GetTaskPage(context.Context, *TaskRequest) (*TaskResponse, error)
}

// UnimplementedTaskServiceServer can be embedded to have forward compatible implementations.
type UnimplementedTaskServiceServer struct {
}

func (*UnimplementedTaskServiceServer) AddTask(TaskService_AddTaskServer) error {
	return status.Errorf(codes.Unimplemented, "method AddTask not implemented")
}
func (*UnimplementedTaskServiceServer) GetTaskStream(TaskService_GetTaskStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method GetTaskStream not implemented")
}
func (*UnimplementedTaskServiceServer) GetTaskPage(context.Context, *TaskRequest) (*TaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTaskPage not implemented")
}

func RegisterTaskServiceServer(s *grpc.Server, srv TaskServiceServer) {
	s.RegisterService(&_TaskService_serviceDesc, srv)
}

func _TaskService_AddTask_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TaskServiceServer).AddTask(&taskServiceAddTaskServer{stream})
}

type TaskService_AddTaskServer interface {
	SendAndClose(*TaskAddAck) error
	Recv() (*Task, error)
	grpc.ServerStream
}

type taskServiceAddTaskServer struct {
	grpc.ServerStream
}

func (x *taskServiceAddTaskServer) SendAndClose(m *TaskAddAck) error {
	return x.ServerStream.SendMsg(m)
}

func (x *taskServiceAddTaskServer) Recv() (*Task, error) {
	m := new(Task)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _TaskService_GetTaskStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TaskServiceServer).GetTaskStream(&taskServiceGetTaskStreamServer{stream})
}

type TaskService_GetTaskStreamServer interface {
	Send(*Task) error
	Recv() (*TaskRequest, error)
	grpc.ServerStream
}

type taskServiceGetTaskStreamServer struct {
	grpc.ServerStream
}

func (x *taskServiceGetTaskStreamServer) Send(m *Task) error {
	return x.ServerStream.SendMsg(m)
}

func (x *taskServiceGetTaskStreamServer) Recv() (*TaskRequest, error) {
	m := new(TaskRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _TaskService_GetTaskPage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).GetTaskPage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Task.TaskService/GetTaskPage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).GetTaskPage(ctx, req.(*TaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TaskService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Task.TaskService",
	HandlerType: (*TaskServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getTaskPage",
			Handler:    _TaskService_GetTaskPage_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "addTask",
			Handler:       _TaskService_AddTask_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "getTaskStream",
			Handler:       _TaskService_GetTaskStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "Task.proto",
}
