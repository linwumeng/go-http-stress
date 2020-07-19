package main

import (
	"fmt"
	"github.com/linwumeng/go-test-stress/stress"
	"github.com/urfave/cli"
	"os"
	"time"
)

func main() {
	var url string
	var v int
	var n int
	var t int
	app := cli.NewApp()
	app.Name = "stress"
	app.Usage = "http test"
	app.Version = "0.0.1"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "url, u",
			Usage:       "specify target url",
			Value:       "https://www.baidu.com",
			Destination: &url,
		},
		cli.IntFlag{
			Name:        "concurrency, c",
			Usage:       "specify concurrency level (virutal user)",
			Value:       1,
			Destination: &v,
		},
		cli.IntFlag{
			Name:        "requests, n",
			Usage:       "specify request number per user",
			Value:       1,
			Destination: &n,
		},
		cli.IntFlag{
			Name:        "interval, t",
			Usage:       "specify interval(s) of printing a row",
			Value:       1,
			Destination: &t,
		},
	}

	app.Action = func(c *cli.Context) error {
		e := stress.NewEmitter(v, n)

		now := time.Now()
		e.Emit(url, t)
		fmt.Printf("总耗时:%6.3vs", time.Now().Sub(now).Seconds())
		return nil
	}

	app.Run(os.Args)
}
