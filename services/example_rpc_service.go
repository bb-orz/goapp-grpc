package services

import (
	"context"
	"fmt"
	"goapp/protobuf/generate/examplePb"
)

// TODO 请实现 examplePb.ExampleRpcServiceServer 接口
type ExampleRpcService struct{}

func (*ExampleRpcService) Ping(context.Context, *examplePb.EmptyReqMsg) (*examplePb.CommonRespMsg, error) {
	return &examplePb.CommonRespMsg{
		Result: true,
	}, nil
}

func (*ExampleRpcService) Foo(ctx context.Context, userId *examplePb.FooReqMsg) (*examplePb.FooRespMsg, error) {
	fmt.Println("Receive Request:", userId.Id)
	return &examplePb.FooRespMsg{
		Result: "Foo",
	}, nil
}

func (*ExampleRpcService) Bar(ctx context.Context, user *examplePb.BarReqMsg) (*examplePb.BarRespMsg, error) {
	return &examplePb.BarRespMsg{Result: true}, nil
}
