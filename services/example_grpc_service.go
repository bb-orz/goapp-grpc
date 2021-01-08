package services

import (
	"context"
	"fmt"
	"goapp/protobuf/generate/examplePb"
)

// TODO 请实现 examplePb.ExampleGrpcServiceServer 接口
type ExampleGrpcService struct{}

func (*ExampleGrpcService) Ping(context.Context, *examplePb.EmptyReqMsg) (*examplePb.CommonRespMsg, error) {
	return &examplePb.CommonRespMsg{
		Result: true,
	}, nil
}

func (*ExampleGrpcService) Foo(ctx context.Context, userId *examplePb.FooReqMsg) (*examplePb.FooRespMsg, error) {
	fmt.Println("Receive Request:", userId.Id)
	return &examplePb.FooRespMsg{
		Result: "Foo",
	}, nil
}

func (*ExampleGrpcService) Bar(ctx context.Context, user *examplePb.BarReqMsg) (*examplePb.BarRespMsg, error) {
	return &examplePb.BarRespMsg{Result: true}, nil
}
