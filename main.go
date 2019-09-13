package main

import (
	"log"
	"os"
	"github.com/wanhello/iris-admin-cli/cmd"
	"github.com/urfave/cli"
)

func main()  {
	
	app := cli.NewApp()
	app.Name = "gin-admin-cli"
	app.Description = "GinAdmin辅助工具"
	app.Version = "0.2.0"
	app.Commands = []cli.Command{
		cmd.NewCommand(),
		cmd.GenerateCommand(),
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatalf(err.Error())
	}
}