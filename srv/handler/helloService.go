package handler

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/util/log"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	"google.golang.org/grpc/metadata"
	
	"time"
	helloService "xiaoshijie.com/micro/hello/srv/proto/helloService"
	"xiaoshijie.com/micro/hello/srv/services"
)

type HelloService struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *HelloService) Call(ctx context.Context, req *helloService.Request, rsp *helloService.Response) error {
	log.Log("Received HelloService.Call request 11 " + req.Name)
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
	md, _ := metadata.FromIncomingContext(ctx)
	badHeader := md.Get(jaeger.TraceContextHeaderName)
	
	fmt.Printf("%+v",md)
	tracer, closer, err := cfg.NewTracer()
	if err != nil {
		println(err.Error())
		return nil
	}
	defer closer.Close()
	badHeader := md.Get(jaeger.TraceContextHeaderName)
	fmt.Printf("%+v",badHeader)
	
	tmc := make(opentracing.TextMapCarrier)
	tmc[jaeger.TraceContextHeaderName] = badHeader[0]
	spanContext, _ := tracer.Extract(opentracing.TextMap, tmc)
	
	span := tracer.StartSpan("FindCode", opentracing.ChildOf(spanContext))
	err = tracer.Inject(
		span.Context(),
		opentracing.TextMap,
		tmc,
	)
	
	
	opentracing.SetGlobalTracer(tracer)
	
	opentracing.StartSpanFromContext(ctx, "FindCode",)
	
	//span.LogKV("SqbApp","FindCode",req.Name)
	span = span.SetBaggageItem("SqbApp", "FindCode")
	span = span.SetBaggageItem("req.Name", req.Name)
	defer span.Finish()
	
	sqbApp, err := services.SqbApp.FindCode(req.Name)
	if err != nil {
		log.Fatal(err.Error())
		rsp.Msg = "Hello " + req.Name + " host "
	}else{
		rsp.Msg = "code name is " + sqbApp.Qq
	}
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *HelloService) Stream(ctx context.Context, req *helloService.StreamingRequest, stream helloService.HelloService_StreamStream) error {
	log.Logf("Received HelloService.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Logf("Responding: %d", i)
		if err := stream.Send(&helloService.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *HelloService) PingPong(ctx context.Context, stream helloService.HelloService_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Logf("Got ping %v", req.Stroke)
		if err := stream.Send(&helloService.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
