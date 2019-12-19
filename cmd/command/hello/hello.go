package hello

import (
	"context"
	"fmt"
	"github.com/micro/cli"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/config/cmd"
	tbServiceSrv "xiaoshijie.com/micro/common/proto/tbServiceSrv"
)



func  run(ctx *cli.Context)  {
	webClient := tbServiceSrv.NewTbServiceSrvService("go.micro.srv.tbServiceSrv", client.DefaultClient)
	rsp, err := webClient.QueryAppId(context.TODO(), &tbServiceSrv.QueryAppIdRequest{
		Mobile:   "123",
	})
	if err != nil {
		println(err.Error())
		return
	}
	fmt.Printf("AppId %+d",rsp.AppId)
}

func Commands(options ...cmd.Option) []cli.Command {
	command := cli.Command{
		Name:  "hello",
		Usage: "Run Hello",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:   "name",
				Usage:  "say hello name",
				EnvVar: "MICRO_HELLO_NAME",
			},
		},
		Action: func(ctx *cli.Context) {
			run(ctx)
		},
	}
	return []cli.Command{command}
}
