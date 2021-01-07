package exampleGrpcStarter

import (
	"fmt"
	"github.com/bb-orz/goinfras"
	"goapp/protobuf/generate/examplePb"
	"goapp/services"
	"google.golang.org/grpc"
	"net"
)

func NewGrpcStarter() *GrpcStarter {
	s := new(GrpcStarter)
	s.cfg = &Config{}
	return s
}

type GrpcStarter struct {
	goinfras.BaseStarter
	cfg        *Config
	listener   net.Listener
	grpcServer *grpc.Server
}

func (s *GrpcStarter) Name() string {
	return "ExampleGRpc"
}

func (s *GrpcStarter) Init(sctx *goinfras.StarterContext) {
	var err error
	var define Config
	viperCfg := sctx.Configs()
	if viperCfg != nil {
		err = viperCfg.UnmarshalKey("Example", &define)
		sctx.PassWarning(s.Name(), goinfras.StepInit, err)
	}

	s.cfg = &define
	sctx.Logger().Debug(s.Name(), goinfras.StepInit, fmt.Sprintf("Config: %+v ", define))

}

func (s *GrpcStarter) Setup(sctx *goinfras.StarterContext) {
	// 注册rpc服务
	var err error
	// 先启动一个监听端口
	if s.listener, err = net.Listen("tcp", fmt.Sprintf("%s:%d", s.cfg.GrpcServerHost, s.cfg.GrpcServerPort)); err != nil {
		sctx.Logger().Error(s.Name(), goinfras.StepSetup, err)
	}

	// 创建grpc服务并注册
	s.grpcServer = grpc.NewServer()
	// 注册服务
	examplePb.RegisterExampleRpcServiceServer(s.grpcServer, &services.ExampleRpcService{})

	sctx.Logger().OK(s.Name(), goinfras.StepSetup, "Register GRpc Service Server Successful!")

}

func (s *GrpcStarter) Start(sctx *goinfras.StarterContext) {
	// 服务运行
	sctx.Logger().OK(s.Name(), goinfras.StepSetup, fmt.Sprintf("GRpc Service Server Startup ... Listening Port: %v", s.cfg.GrpcServerPort))
	if err := s.grpcServer.Serve(s.listener); err != nil {
		sctx.Logger().Error(s.Name(), goinfras.StepStart, err)
	}
}

func (s *GrpcStarter) Check(sctx *goinfras.StarterContext) bool {
	// 检查服务
	// var err error
	// var conn *grpc.ClientConn
	// var result *examplePb.CommonRespMsg
	// // 获得一个grpc服务端的连接
	// if conn, err = grpc.Dial(fmt.Sprintf(":%d", s.cfg.GrpcServerPort), grpc.WithInsecure()); err != nil {
	// 	sctx.Logger().Error(s.Name(), goinfras.StepStart, err)
	// 	return false
	// }
	//
	// // 创建客户端
	// client := examplePb.NewExampleRpcServiceClient(conn)
	// if result, err = client.Ping(context.TODO(), &examplePb.EmptyReqMsg{}); err != nil {
	// 	sctx.Logger().Error(s.Name(), goinfras.StepStart, err)
	// }
	//
	// return result.Result

	return true
}

func (s *GrpcStarter) StartBlocking() bool {
	return true
}

func (s *GrpcStarter) Stop() {
	s.grpcServer.Stop()

}

func (s *GrpcStarter) PriorityGroup() goinfras.PriorityGroup {
	return goinfras.AppGroup
}

func (s *GrpcStarter) Priority() int {
	return goinfras.DEFAULT_PRIORITY
}
