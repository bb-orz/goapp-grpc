package exampleGrpc

import (
	"fmt"
	"github.com/bb-orz/goinfras"
	"goapp/protobuf/generate/examplePb"
	"goapp/services"
	"google.golang.org/grpc"
	"net"
)

var Cfg *Config

func NewStarter() *starter {
	s := new(starter)
	s.cfg = &Config{}
	return s
}

type starter struct {
	goinfras.BaseStarter
	cfg        *Config
	listener   net.Listener
	grpcServer *grpc.Server
}

func (s *starter) Name() string {
	return "ExampleGRpc"
}

func (s *starter) Init(sctx *goinfras.StarterContext) {
	var err error
	var define Config
	viper := sctx.Configs()
	if viper != nil {
		err = viper.UnmarshalKey("ExampleGrpc", &define)
		sctx.PassWarning(s.Name(), goinfras.StepInit, err)
	}
	s.cfg = &define
	Cfg = &define
	sctx.Logger().Debug(s.Name(), goinfras.StepInit, fmt.Sprintf("Config: %+v ", define))

}

func (s *starter) Setup(sctx *goinfras.StarterContext) {
	// 注册rpc服务
	var err error
	// 先启动一个监听端口
	if s.listener, err = net.Listen("tcp", fmt.Sprintf("%s:%d", s.cfg.ServerHost, s.cfg.ServerPort)); err != nil {
		sctx.Logger().Error(s.Name(), goinfras.StepSetup, err)
	}

	// 创建grpc服务并注册
	s.grpcServer = grpc.NewServer()
	// 注册服务
	examplePb.RegisterExampleRpcServiceServer(s.grpcServer, &services.ExampleRpcService{})

	sctx.Logger().OK(s.Name(), goinfras.StepSetup, "Register GRpc Service Server Successful!")

}

func (s *starter) Check(sctx *goinfras.StarterContext) bool {
	if s.grpcServer != nil && s.listener != nil {
		return true
	}
	return false
}

func (s *starter) Start(sctx *goinfras.StarterContext) {
	// 服务运行
	sctx.Logger().OK(s.Name(), goinfras.StepStart, fmt.Sprintf("GRpc Service Server Startup ... Listening: %s:%d", s.cfg.ServerHost, s.cfg.ServerPort))
	if err := s.grpcServer.Serve(s.listener); err != nil {
		sctx.Logger().Error(s.Name(), goinfras.StepStart, err)
	}
}

func (s *starter) StartBlocking() bool {
	return true
}

func (s *starter) Stop() {
	s.grpcServer.Stop()

}

func (s *starter) PriorityGroup() goinfras.PriorityGroup {
	return goinfras.AppGroup
}

func (s *starter) Priority() int {
	return goinfras.DEFAULT_PRIORITY
}
