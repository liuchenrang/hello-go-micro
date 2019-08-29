package subscriber

import (
	"context"
	"github.com/micro/go-micro/util/log"

	helloService "xiaoshijie.com/micro/hello/srv/proto/helloService"
)

type HelloService struct{}

func (e *HelloService) Handle(ctx context.Context, msg *helloService.Message) error {
	log.Log("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *helloService.Message) error {
	log.Log("Function Received message: ", msg.Say)
	return nil
}
