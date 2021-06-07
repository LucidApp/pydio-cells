// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: jobs.proto

/*
Package jobs is a generated protocol buffer package.

It is generated from these files:
	jobs.proto

It has these top-level messages:
	NodesSelector
	IdmSelector
	UsersSelector
	DataSourceSelector
	TriggerFilterQuery
	TriggerFilter
	ActionOutputFilter
	ContextMetaFilter
	ContextMetaSingleQuery
	Schedule
	Action
	Job
	JobParameter
	JobChangeEvent
	TaskChangeEvent
	PutJobRequest
	PutJobResponse
	GetJobRequest
	GetJobResponse
	DeleteJobRequest
	DeleteJobResponse
	ListJobsRequest
	ListJobsResponse
	ListTasksRequest
	ListTasksResponse
	PutTaskRequest
	PutTaskResponse
	DeleteTasksRequest
	DeleteTasksResponse
	DetectStuckTasksRequest
	DetectStuckTasksResponse
	Task
	CtrlCommand
	CtrlCommandResponse
	ActionLog
	JobTriggerEvent
	ActionOutput
	ActionOutputSingleQuery
	ActionMessage
*/
package jobs

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/golang/protobuf/ptypes/any"
import _ "github.com/pydio/cells/common/service/proto"
import _ "github.com/pydio/cells/common/proto/tree"
import _ "github.com/pydio/cells/common/proto/idm"
import _ "github.com/pydio/cells/common/proto/activity"
import _ "github.com/pydio/cells/common/proto/object"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "context"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for JobService service

type JobServiceClient interface {
	PutJob(ctx context.Context, in *PutJobRequest, opts ...client.CallOption) (*PutJobResponse, error)
	GetJob(ctx context.Context, in *GetJobRequest, opts ...client.CallOption) (*GetJobResponse, error)
	DeleteJob(ctx context.Context, in *DeleteJobRequest, opts ...client.CallOption) (*DeleteJobResponse, error)
	ListJobs(ctx context.Context, in *ListJobsRequest, opts ...client.CallOption) (JobService_ListJobsClient, error)
	PutTask(ctx context.Context, in *PutTaskRequest, opts ...client.CallOption) (*PutTaskResponse, error)
	PutTaskStream(ctx context.Context, opts ...client.CallOption) (JobService_PutTaskStreamClient, error)
	ListTasks(ctx context.Context, in *ListTasksRequest, opts ...client.CallOption) (JobService_ListTasksClient, error)
	DeleteTasks(ctx context.Context, in *DeleteTasksRequest, opts ...client.CallOption) (*DeleteTasksResponse, error)
	DetectStuckTasks(ctx context.Context, in *DetectStuckTasksRequest, opts ...client.CallOption) (*DetectStuckTasksResponse, error)
}

type jobServiceClient struct {
	c           client.Client
	serviceName string
}

func NewJobServiceClient(serviceName string, c client.Client) JobServiceClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "jobs"
	}
	return &jobServiceClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *jobServiceClient) PutJob(ctx context.Context, in *PutJobRequest, opts ...client.CallOption) (*PutJobResponse, error) {
	req := c.c.NewRequest(c.serviceName, "JobService.PutJob", in)
	out := new(PutJobResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobServiceClient) GetJob(ctx context.Context, in *GetJobRequest, opts ...client.CallOption) (*GetJobResponse, error) {
	req := c.c.NewRequest(c.serviceName, "JobService.GetJob", in)
	out := new(GetJobResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobServiceClient) DeleteJob(ctx context.Context, in *DeleteJobRequest, opts ...client.CallOption) (*DeleteJobResponse, error) {
	req := c.c.NewRequest(c.serviceName, "JobService.DeleteJob", in)
	out := new(DeleteJobResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobServiceClient) ListJobs(ctx context.Context, in *ListJobsRequest, opts ...client.CallOption) (JobService_ListJobsClient, error) {
	req := c.c.NewRequest(c.serviceName, "JobService.ListJobs", &ListJobsRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	if err := stream.Send(in); err != nil {
		return nil, err
	}
	return &jobServiceListJobsClient{stream}, nil
}

type JobService_ListJobsClient interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Recv() (*ListJobsResponse, error)
}

type jobServiceListJobsClient struct {
	stream client.Streamer
}

func (x *jobServiceListJobsClient) Close() error {
	return x.stream.Close()
}

func (x *jobServiceListJobsClient) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *jobServiceListJobsClient) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *jobServiceListJobsClient) Recv() (*ListJobsResponse, error) {
	m := new(ListJobsResponse)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (c *jobServiceClient) PutTask(ctx context.Context, in *PutTaskRequest, opts ...client.CallOption) (*PutTaskResponse, error) {
	req := c.c.NewRequest(c.serviceName, "JobService.PutTask", in)
	out := new(PutTaskResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobServiceClient) PutTaskStream(ctx context.Context, opts ...client.CallOption) (JobService_PutTaskStreamClient, error) {
	req := c.c.NewRequest(c.serviceName, "JobService.PutTaskStream", &PutTaskRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	return &jobServicePutTaskStreamClient{stream}, nil
}

type JobService_PutTaskStreamClient interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*PutTaskRequest) error
	Recv() (*PutTaskResponse, error)
}

type jobServicePutTaskStreamClient struct {
	stream client.Streamer
}

func (x *jobServicePutTaskStreamClient) Close() error {
	return x.stream.Close()
}

func (x *jobServicePutTaskStreamClient) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *jobServicePutTaskStreamClient) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *jobServicePutTaskStreamClient) Send(m *PutTaskRequest) error {
	return x.stream.Send(m)
}

func (x *jobServicePutTaskStreamClient) Recv() (*PutTaskResponse, error) {
	m := new(PutTaskResponse)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (c *jobServiceClient) ListTasks(ctx context.Context, in *ListTasksRequest, opts ...client.CallOption) (JobService_ListTasksClient, error) {
	req := c.c.NewRequest(c.serviceName, "JobService.ListTasks", &ListTasksRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	if err := stream.Send(in); err != nil {
		return nil, err
	}
	return &jobServiceListTasksClient{stream}, nil
}

type JobService_ListTasksClient interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Recv() (*ListTasksResponse, error)
}

type jobServiceListTasksClient struct {
	stream client.Streamer
}

func (x *jobServiceListTasksClient) Close() error {
	return x.stream.Close()
}

func (x *jobServiceListTasksClient) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *jobServiceListTasksClient) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *jobServiceListTasksClient) Recv() (*ListTasksResponse, error) {
	m := new(ListTasksResponse)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (c *jobServiceClient) DeleteTasks(ctx context.Context, in *DeleteTasksRequest, opts ...client.CallOption) (*DeleteTasksResponse, error) {
	req := c.c.NewRequest(c.serviceName, "JobService.DeleteTasks", in)
	out := new(DeleteTasksResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobServiceClient) DetectStuckTasks(ctx context.Context, in *DetectStuckTasksRequest, opts ...client.CallOption) (*DetectStuckTasksResponse, error) {
	req := c.c.NewRequest(c.serviceName, "JobService.DetectStuckTasks", in)
	out := new(DetectStuckTasksResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for JobService service

type JobServiceHandler interface {
	PutJob(context.Context, *PutJobRequest, *PutJobResponse) error
	GetJob(context.Context, *GetJobRequest, *GetJobResponse) error
	DeleteJob(context.Context, *DeleteJobRequest, *DeleteJobResponse) error
	ListJobs(context.Context, *ListJobsRequest, JobService_ListJobsStream) error
	PutTask(context.Context, *PutTaskRequest, *PutTaskResponse) error
	PutTaskStream(context.Context, JobService_PutTaskStreamStream) error
	ListTasks(context.Context, *ListTasksRequest, JobService_ListTasksStream) error
	DeleteTasks(context.Context, *DeleteTasksRequest, *DeleteTasksResponse) error
	DetectStuckTasks(context.Context, *DetectStuckTasksRequest, *DetectStuckTasksResponse) error
}

func RegisterJobServiceHandler(s server.Server, hdlr JobServiceHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&JobService{hdlr}, opts...))
}

type JobService struct {
	JobServiceHandler
}

func (h *JobService) PutJob(ctx context.Context, in *PutJobRequest, out *PutJobResponse) error {
	return h.JobServiceHandler.PutJob(ctx, in, out)
}

func (h *JobService) GetJob(ctx context.Context, in *GetJobRequest, out *GetJobResponse) error {
	return h.JobServiceHandler.GetJob(ctx, in, out)
}

func (h *JobService) DeleteJob(ctx context.Context, in *DeleteJobRequest, out *DeleteJobResponse) error {
	return h.JobServiceHandler.DeleteJob(ctx, in, out)
}

func (h *JobService) ListJobs(ctx context.Context, stream server.Streamer) error {
	m := new(ListJobsRequest)
	if err := stream.Recv(m); err != nil {
		return err
	}
	return h.JobServiceHandler.ListJobs(ctx, m, &jobServiceListJobsStream{stream})
}

type JobService_ListJobsStream interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*ListJobsResponse) error
}

type jobServiceListJobsStream struct {
	stream server.Streamer
}

func (x *jobServiceListJobsStream) Close() error {
	return x.stream.Close()
}

func (x *jobServiceListJobsStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *jobServiceListJobsStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *jobServiceListJobsStream) Send(m *ListJobsResponse) error {
	return x.stream.Send(m)
}

func (h *JobService) PutTask(ctx context.Context, in *PutTaskRequest, out *PutTaskResponse) error {
	return h.JobServiceHandler.PutTask(ctx, in, out)
}

func (h *JobService) PutTaskStream(ctx context.Context, stream server.Streamer) error {
	return h.JobServiceHandler.PutTaskStream(ctx, &jobServicePutTaskStreamStream{stream})
}

type JobService_PutTaskStreamStream interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*PutTaskResponse) error
	Recv() (*PutTaskRequest, error)
}

type jobServicePutTaskStreamStream struct {
	stream server.Streamer
}

func (x *jobServicePutTaskStreamStream) Close() error {
	return x.stream.Close()
}

func (x *jobServicePutTaskStreamStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *jobServicePutTaskStreamStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *jobServicePutTaskStreamStream) Send(m *PutTaskResponse) error {
	return x.stream.Send(m)
}

func (x *jobServicePutTaskStreamStream) Recv() (*PutTaskRequest, error) {
	m := new(PutTaskRequest)
	if err := x.stream.Recv(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (h *JobService) ListTasks(ctx context.Context, stream server.Streamer) error {
	m := new(ListTasksRequest)
	if err := stream.Recv(m); err != nil {
		return err
	}
	return h.JobServiceHandler.ListTasks(ctx, m, &jobServiceListTasksStream{stream})
}

type JobService_ListTasksStream interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*ListTasksResponse) error
}

type jobServiceListTasksStream struct {
	stream server.Streamer
}

func (x *jobServiceListTasksStream) Close() error {
	return x.stream.Close()
}

func (x *jobServiceListTasksStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *jobServiceListTasksStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *jobServiceListTasksStream) Send(m *ListTasksResponse) error {
	return x.stream.Send(m)
}

func (h *JobService) DeleteTasks(ctx context.Context, in *DeleteTasksRequest, out *DeleteTasksResponse) error {
	return h.JobServiceHandler.DeleteTasks(ctx, in, out)
}

func (h *JobService) DetectStuckTasks(ctx context.Context, in *DetectStuckTasksRequest, out *DetectStuckTasksResponse) error {
	return h.JobServiceHandler.DetectStuckTasks(ctx, in, out)
}

// Client API for TaskService service

type TaskServiceClient interface {
	Control(ctx context.Context, in *CtrlCommand, opts ...client.CallOption) (*CtrlCommandResponse, error)
}

type taskServiceClient struct {
	c           client.Client
	serviceName string
}

func NewTaskServiceClient(serviceName string, c client.Client) TaskServiceClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "jobs"
	}
	return &taskServiceClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *taskServiceClient) Control(ctx context.Context, in *CtrlCommand, opts ...client.CallOption) (*CtrlCommandResponse, error) {
	req := c.c.NewRequest(c.serviceName, "TaskService.Control", in)
	out := new(CtrlCommandResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for TaskService service

type TaskServiceHandler interface {
	Control(context.Context, *CtrlCommand, *CtrlCommandResponse) error
}

func RegisterTaskServiceHandler(s server.Server, hdlr TaskServiceHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&TaskService{hdlr}, opts...))
}

type TaskService struct {
	TaskServiceHandler
}

func (h *TaskService) Control(ctx context.Context, in *CtrlCommand, out *CtrlCommandResponse) error {
	return h.TaskServiceHandler.Control(ctx, in, out)
}
