package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/client/grpc"
	_ "github.com/micro/go-micro/client/grpc"
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-plugins/wrapper/trace/opentracing"
	"github.com/uber/jaeger-client-go/config"
	"os"
	"time"
	"xiaoshijie.com/micro/hello/api/client"
	"xiaoshijie.com/micro/hello/api/handler"
	
	api "xiaoshijie.com/micro/hello/api/proto/api"
)
func GrpcClient() micro.Option {
	return func(o *micro.Options) {
		if(os.Getenv("MICRO_CLIENT") == "grpc"){
			o.Client = grpc.NewClient()
		}
	}
}
func main() {
	
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.api.apiHelloService"),
		micro.Version("latest"),
		GrpcClient(),
	)
	cfg := config.Configuration{
		ServiceName: "Http2RpcService",//自定义服务名称
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans:            true,
			BufferFlushInterval: 1 * time.Second,
			LocalAgentHostPort:  "127.0.0.1:5775",//jaeger agent
		},
	}
	tracer, closer, err := cfg.NewTracer()
	if err != nil {
		println(err.Error())
		return
	}
	defer closer.Close()
	// Initialise service
	service.Init(
		// create wrap for the Api srv client
		micro.WrapHandler(client.ApiWrapper(service)),
		
		
		micro.WrapHandler(
			opentracing.NewHandlerWrapper(tracer),
		),
	)

	// Register Handler
	api.RegisterApiHandler(service.Server(), new(handler.Api))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
