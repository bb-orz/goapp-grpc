package exampleGrpcStarter

type Config struct {
	Name           string
	GrpcServerHost string
	GrpcServerPort int64
	GrpcClientPort int64
}

func DefaultConfig() *Config {
	return &Config{
		Name:           "ExampleGRpc",
		GrpcServerHost: "127.0.0.1",
		GrpcServerPort: 8899,
		GrpcClientPort: 8889,
	}
}
