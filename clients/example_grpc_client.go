package clients

import (
	"fmt"
	"goapp/protobuf/generate/examplePb"
	"goapp/starter/exampleGrpc"
	"google.golang.org/grpc"
)

func GetExampleGrpcClient() (examplePb.ExampleGrpcServiceClient, error) {
	// 检查服务
	var err error
	var conn *grpc.ClientConn
	// 获得一个grpc服务端的连接
	if conn, err = grpc.Dial(fmt.Sprintf("%s:%d", exampleGrpc.Cfg.ServerHost, exampleGrpc.Cfg.ServerPort), grpc.WithInsecure()); err != nil {
		return nil, err
	}

	// 创建客户端
	client := examplePb.NewExampleGrpcServiceClient(conn)
	return client, nil
}
