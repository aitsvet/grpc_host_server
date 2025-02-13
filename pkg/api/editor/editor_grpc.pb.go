// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: api/editor/editor.proto

package editor

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	QuizService_CreateQuiz_FullMethodName         = "/editor.QuizService/CreateQuiz"
	QuizService_GetAllQuizzes_FullMethodName      = "/editor.QuizService/GetAllQuizzes"
	QuizService_AddQuestionToQuiz_FullMethodName  = "/editor.QuizService/AddQuestionToQuiz"
	QuizService_GetQuizQuestions_FullMethodName   = "/editor.QuizService/GetQuizQuestions"
	QuizService_EditQuestionInQuiz_FullMethodName = "/editor.QuizService/EditQuestionInQuiz"
)

// QuizServiceClient is the client API for QuizService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QuizServiceClient interface {
	CreateQuiz(ctx context.Context, in *CreateQuizRequest, opts ...grpc.CallOption) (*CreateQuizResponse, error)
	GetAllQuizzes(ctx context.Context, in *GetAllQuizzesRequest, opts ...grpc.CallOption) (*GetAllQuizzesResponse, error)
	AddQuestionToQuiz(ctx context.Context, in *AddQuestionToQuizRequest, opts ...grpc.CallOption) (*AddQuestionToQuizResponse, error)
	GetQuizQuestions(ctx context.Context, in *GetQuizQuestionsRequest, opts ...grpc.CallOption) (*GetQuizQuestionsResponse, error)
	EditQuestionInQuiz(ctx context.Context, in *EditQuestionInQuizRequest, opts ...grpc.CallOption) (*EditQuestionInQuizResponse, error)
}

type quizServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewQuizServiceClient(cc grpc.ClientConnInterface) QuizServiceClient {
	return &quizServiceClient{cc}
}

func (c *quizServiceClient) CreateQuiz(ctx context.Context, in *CreateQuizRequest, opts ...grpc.CallOption) (*CreateQuizResponse, error) {
	out := new(CreateQuizResponse)
	err := c.cc.Invoke(ctx, QuizService_CreateQuiz_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *quizServiceClient) GetAllQuizzes(ctx context.Context, in *GetAllQuizzesRequest, opts ...grpc.CallOption) (*GetAllQuizzesResponse, error) {
	out := new(GetAllQuizzesResponse)
	err := c.cc.Invoke(ctx, QuizService_GetAllQuizzes_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *quizServiceClient) AddQuestionToQuiz(ctx context.Context, in *AddQuestionToQuizRequest, opts ...grpc.CallOption) (*AddQuestionToQuizResponse, error) {
	out := new(AddQuestionToQuizResponse)
	err := c.cc.Invoke(ctx, QuizService_AddQuestionToQuiz_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *quizServiceClient) GetQuizQuestions(ctx context.Context, in *GetQuizQuestionsRequest, opts ...grpc.CallOption) (*GetQuizQuestionsResponse, error) {
	out := new(GetQuizQuestionsResponse)
	err := c.cc.Invoke(ctx, QuizService_GetQuizQuestions_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *quizServiceClient) EditQuestionInQuiz(ctx context.Context, in *EditQuestionInQuizRequest, opts ...grpc.CallOption) (*EditQuestionInQuizResponse, error) {
	out := new(EditQuestionInQuizResponse)
	err := c.cc.Invoke(ctx, QuizService_EditQuestionInQuiz_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QuizServiceServer is the server API for QuizService service.
// All implementations must embed UnimplementedQuizServiceServer
// for forward compatibility
type QuizServiceServer interface {
	CreateQuiz(context.Context, *CreateQuizRequest) (*CreateQuizResponse, error)
	GetAllQuizzes(context.Context, *GetAllQuizzesRequest) (*GetAllQuizzesResponse, error)
	AddQuestionToQuiz(context.Context, *AddQuestionToQuizRequest) (*AddQuestionToQuizResponse, error)
	GetQuizQuestions(context.Context, *GetQuizQuestionsRequest) (*GetQuizQuestionsResponse, error)
	EditQuestionInQuiz(context.Context, *EditQuestionInQuizRequest) (*EditQuestionInQuizResponse, error)
	mustEmbedUnimplementedQuizServiceServer()
}

// UnimplementedQuizServiceServer must be embedded to have forward compatible implementations.
type UnimplementedQuizServiceServer struct {
}

func (UnimplementedQuizServiceServer) CreateQuiz(context.Context, *CreateQuizRequest) (*CreateQuizResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateQuiz not implemented")
}
func (UnimplementedQuizServiceServer) GetAllQuizzes(context.Context, *GetAllQuizzesRequest) (*GetAllQuizzesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllQuizzes not implemented")
}
func (UnimplementedQuizServiceServer) AddQuestionToQuiz(context.Context, *AddQuestionToQuizRequest) (*AddQuestionToQuizResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddQuestionToQuiz not implemented")
}
func (UnimplementedQuizServiceServer) GetQuizQuestions(context.Context, *GetQuizQuestionsRequest) (*GetQuizQuestionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetQuizQuestions not implemented")
}
func (UnimplementedQuizServiceServer) EditQuestionInQuiz(context.Context, *EditQuestionInQuizRequest) (*EditQuestionInQuizResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditQuestionInQuiz not implemented")
}
func (UnimplementedQuizServiceServer) mustEmbedUnimplementedQuizServiceServer() {}

// UnsafeQuizServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QuizServiceServer will
// result in compilation errors.
type UnsafeQuizServiceServer interface {
	mustEmbedUnimplementedQuizServiceServer()
}

func RegisterQuizServiceServer(s grpc.ServiceRegistrar, srv QuizServiceServer) {
	s.RegisterService(&QuizService_ServiceDesc, srv)
}

func _QuizService_CreateQuiz_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateQuizRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuizServiceServer).CreateQuiz(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QuizService_CreateQuiz_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuizServiceServer).CreateQuiz(ctx, req.(*CreateQuizRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QuizService_GetAllQuizzes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllQuizzesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuizServiceServer).GetAllQuizzes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QuizService_GetAllQuizzes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuizServiceServer).GetAllQuizzes(ctx, req.(*GetAllQuizzesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QuizService_AddQuestionToQuiz_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddQuestionToQuizRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuizServiceServer).AddQuestionToQuiz(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QuizService_AddQuestionToQuiz_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuizServiceServer).AddQuestionToQuiz(ctx, req.(*AddQuestionToQuizRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QuizService_GetQuizQuestions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetQuizQuestionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuizServiceServer).GetQuizQuestions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QuizService_GetQuizQuestions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuizServiceServer).GetQuizQuestions(ctx, req.(*GetQuizQuestionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QuizService_EditQuestionInQuiz_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EditQuestionInQuizRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuizServiceServer).EditQuestionInQuiz(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QuizService_EditQuestionInQuiz_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuizServiceServer).EditQuestionInQuiz(ctx, req.(*EditQuestionInQuizRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// QuizService_ServiceDesc is the grpc.ServiceDesc for QuizService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var QuizService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "editor.QuizService",
	HandlerType: (*QuizServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateQuiz",
			Handler:    _QuizService_CreateQuiz_Handler,
		},
		{
			MethodName: "GetAllQuizzes",
			Handler:    _QuizService_GetAllQuizzes_Handler,
		},
		{
			MethodName: "AddQuestionToQuiz",
			Handler:    _QuizService_AddQuestionToQuiz_Handler,
		},
		{
			MethodName: "GetQuizQuestions",
			Handler:    _QuizService_GetQuizQuestions_Handler,
		},
		{
			MethodName: "EditQuestionInQuiz",
			Handler:    _QuizService_EditQuestionInQuiz_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/editor/editor.proto",
}
