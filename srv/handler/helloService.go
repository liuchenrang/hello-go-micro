package handler

import (
	"context"

	"github.com/micro/go-micro/util/log"

	helloService "xiaoshijie.com/micro/hello/srv/proto/helloService"
)

type HelloService struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *HelloService) Call(ctx context.Context, req *helloService.Request, rsp *helloService.Response) error {
	log.Log("Received HelloService.Call request " + req.Name)
	println("11" + req.Name)
	rsp.Msg = "Hello " + req.Name
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
