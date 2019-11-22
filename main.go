package main

import (
	"errors"
	"fmt"
	"github.com/urfave/cli"
	"os"
	"strings"
	"tcping/tcping"
	"time"
)

func main() {
	app := cli.NewApp()

	app.Name = "tcping"
	app.EnableBashCompletion = true
	app.Version = "1.0.0"
	app.UsageText = "tcping [global options] host"
	app.Description = "Use the tcp protocol to ping the specified port of the specified website."
	app.Flags = []cli.Flag{
		cli.IntFlag{
			Name:   "port,p",
			Usage:  "Set the port to request",
			Hidden: false,
			Value:  80,
		},
		cli.IntFlag{
			Name:        "timeout,t",
			Usage:       "Connection timeout",
			Hidden:      false,
			Value:       5,
			Destination: nil,
		},
	}

	app.Action = Run

	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
	}
}

func Run(c *cli.Context) (err error) {
	if len(c.Args()) < 1 {
		return errors.New("parameter cannot be empty")
	}
	host := c.Args()[len(c.Args()) - 1]

	if host == "" {
		return errors.New("parameter cannot be empty")
	}

	if strings.Index(host, ".") == -1 {
		return errors.New("wrong parameter")
	}

	client := tcping.New(tcping.Config{
		Host:        host,
		Port:        c.Int("port"),
		ConnTimeout: time.Duration(c.Int("timeout")) * time.Second,
	})

	var (
		i = 0
	)
	fmt.Printf("start ping %s:%d\n", host, c.Int("port"))
	for {
		data := client.Ping()
		if data.Err != nil {
			fmt.Printf("%s\n", data.Err)
		} else {
			fmt.Printf("Connected to %s sqe=%d time=%s \n", data.RemoteAddr, i, data.HumanTime())
		}
		i++
		time.Sleep(time.Second * 1)
	}
	return
}