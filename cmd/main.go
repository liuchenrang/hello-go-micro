package main

import (
	"github.com/micro/go-micro/config/cmd"
	"duoduo.com/micro/tbservice/tb-service-cmd/command"
)

// Init initialised the command line
func Init(options ...cmd.Option) {
	command.Setup(cmd.App(), options...)
	err := cmd.Init(
		cmd.Name("fx-cmd"),
		cmd.Description("fx-cmd"),
		cmd.Version("latest"),
	)
	if err != nil {
		panic(err.Error())
	}

}

func main() {
	Init()
	
}
