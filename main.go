package main

import (
	"fmt"
	"github.com/urfave/cli"
	"os"
)

func main() {
	app := cli.NewApp()

	app.Name = "tcping"
	app.EnableBashCompletion = true
	app.Description = "Use the tcp protocol to ping the specified port of the specified website."
	app.Flags = []cli.Flag{
		cli.IntFlag{
			Name:   "port,p",
			Usage:  "Set the port to request",
			EnvVar: "port",
			Hidden: false,
			Value:  80,
		},
		cli.IntFlag{
			Name:        "timeout,t",
			Usage:       "Connection timeout",
			EnvVar:      "timeout",
			Hidden:      false,
			Value:       5,
			Destination: nil,
		},
	}

	app.Action = func(c *cli.Context) (err error) {
		return
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
	}
}
