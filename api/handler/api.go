package handler

import (
	"context"
	"encoding/json"
	"github.com/micro/go-micro"
	client2 "github.com/micro/go-micro/client"
	"github.com/micro/go-micro/util/log"
	
	api "github.com/micro/go-micro/api/proto"
	"github.com/micro/go-micro/errors"
	"duoduo.com/micro/hello/api/client"
	api2 "duoduo.com/micro/hello/srv/proto/helloService"
)

type Api struct{}

func extractValue(pair *api.Pair) string {
	if pair == nil {
		return ""
	}
	if len(pair.Values) == 0 {
		return ""
	}
	return pair.Values[0]
}

// Api.Call is called by the API as /api/call with post body {"name": "foo"}
func (e *Api) Call(ctx context.Context, req *api.Request, rsp *api.Response) error {
	log.Log("Received Api.Call request")

	ser,_ := client.ServiceFromContext(ctx)
	p := micro.NewPublisher("go.micro.srv.helloService",ser.Client())
	p.Publish(context.TODO(), &api2.Request{Name:"who"})
	// extract the client from the context
	apiClient, ok := client.ApiFromContext(ctx)
	

	
	if !ok {
		return errors.InternalServerError("go.micro.api.api.api.call", "api client not found")
	}
	
	// make request
	response, err := apiClient.Call(ctx, &api2.Request{
		Name: extractValue(req.Post["name"]),
	}, func(options *client2.CallOptions) {
		options.RequestTimeout = 3000000
	})
	if err != nil {
		return errors.InternalServerError("go.micro.api.api.api.call", err.Error())
	}

	b, _ := json.Marshal(response)
	rsp.StatusCode = 200
	rsp.Body = string(b)

	return nil
}
