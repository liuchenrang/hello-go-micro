package client

import (
	"context"
	"github.com/micro/go-plugins/wrapper/trace/opentracing"
	"github.com/uber/jaeger-client-go/config"
	"time"
	
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/server"
	 //api "duoduo.com/micro/hello/api/proto/api"
	 api "duoduo.com/micro/hello/srv/proto/helloService"
	 //api "../../srv/proto/helloService"
	
)

type apiKey struct{}
type apiHi struct{}

// FromContext retrieves the client from the Context
func ApiFromContext(ctx context.Context) (api.HelloService, bool) {
	c, ok := ctx.Value(apiKey{}).(api.HelloService)
	return c, ok
}

// Client returns a wrapper for the ApiClient
func ApiWrapper(service micro.Service) server.HandlerWrapper {
	cfg := config.Configuration{
		ServiceName: "apiCall",//自定义服务名称
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
		return nil
	}
	defer closer.Close()
	
	clientWapper := opentracing.NewClientWrapper(tracer)
	
	client := api.NewHelloService("go.micro.srv.helloService",clientWapper( service.Client()))
	
	return func(fn server.HandlerFunc) server.HandlerFunc {
		return func(ctx context.Context, req server.Request, rsp interface{}) error {
			ctx = context.WithValue(ctx, apiKey{}, client)
			ctx = context.WithValue(ctx, apiHi{}, service)
			
			return fn(ctx, req, rsp)
		}
	}
}

func ServiceFromContext(ctx context.Context) (micro.Service, bool) {
	c, ok := ctx.Value(apiHi{}).(micro.Service)
	return c, ok
}