package main

import (
	"fmt"
	"os"

	"github.com/pharmpress/azurectl/command"
	"github.com/pharmpress/azurectl/version"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "azurectl"
	app.Version = version.Version
	app.Usage = "A simple command line client for azure storage."
	app.Flags = []cli.Flag{
		cli.StringFlag{Name: "account-name", Value: "", Usage: "Account name"},
		cli.StringFlag{Name: "account-key", Value: "", Usage: "Account key"},
	}
	app.Commands = []cli.Command{
		command.NewLsCommand(),
		command.NewUploadCommand(),
		command.NewDownloadCommand(),
	}

	app.Run(os.Args)
	fmt.Println("Finish.")
}
