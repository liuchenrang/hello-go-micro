package command

import (
	ccli "github.com/micro/cli"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/client/grpc"
	"github.com/micro/go-micro/config/cmd"
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-plugins/registry/etcdv3"
	"github.com/micro/micro/plugin"
	"os"
	"strings"
	"duoduo.com/micro/tbservice/tb-service-cmd/command/hello"
)

// Setup sets up a cli.App
func Setup(app *ccli.App, options ...cmd.Option) {
	// Add the various commands
	app.Commands = append(app.Commands, hello.Commands(options...)...)
	app.Action = func(context *ccli.Context) { ccli.ShowAppHelp(context) }
	

	setup(app)
}
func setup(app *ccli.App) {
	
	app.Flags = append(app.Flags,
		
		ccli.StringFlag{
			Name:   "api_address",
			Usage:  "Set the api address e.g 0.0.0.0:8080",
			EnvVar: "MICRO_API_ADDRESS",
		},
		ccli.StringFlag{
			Name:   "proxy_address",
			Usage:  "Proxy requests via the HTTP address specified",
			EnvVar: "MICRO_PROXY_ADDRESS",
		},
	
	)
	
	plugins := plugin.Plugins()
	
	for _, p := range plugins {
		if flags := p.Flags(); len(flags) > 0 {
			app.Flags = append(app.Flags, flags...)
		}
		
		if cmds := p.Commands(); len(cmds) > 0 {
			app.Commands = append(app.Commands, cmds...)
		}
	}
	
	before := app.Before
	
	app.Before = func(ctx *ccli.Context) error {
		for _, p := range plugins {
			if err := p.Init(ctx); err != nil {
				return err
			}
		}
		// now do previous before
		err := before(ctx)
		if err != nil {
			panic(err.Error())
		}
		defaultOptions := cmd.DefaultOptions()
		Register(&defaultOptions)
		return nil
	}

}
func Register(options *cmd.Options) {
	username := strings.TrimSpace(os.Getenv("REGISTER_AUTH_USERNAME"))
	password := strings.TrimSpace(os.Getenv("REGISTER_AUTH_PASSWORD"))
	register := *options.Registry
	err :=  register.Init(etcdv3.Auth(username,password))
	client.DefaultClient = grpc.NewClient(func(ops *client.Options) {
		ops.Registry = *options.Registry
	})
	if err != nil {
		log.Fatal(err)
	}
	println(username, password)
}
