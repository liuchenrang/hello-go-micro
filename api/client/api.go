package client

import (
	"context"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/server"
	 //api "xiaoshijie.com/micro/hello/api/proto/api"
	 api "xiaoshijie.com/micro/hello/srv/proto/helloService"
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
	client := api.NewHelloService("go.micro.srv.helloService", service.Client())

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