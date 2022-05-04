package main

import (
	"fmt"
	"github.com/iqoologic/goserve/action"
	"github.com/iqoologic/goserve/server"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	app := cli.App{
		Action: action.Run,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "host",
				Aliases:     []string{"u"},
				Usage:       "host",
				Value:       "0.0.0.0",
				DefaultText: "0.0.0.0",
				Required:    false,
			},
			&cli.IntFlag{
				Name:        "port",
				Aliases:     []string{"p"},
				Usage:       "port",
				Value:       server.RandomizePort(),
				DefaultText: "random port number",
				Required:    false,
			},
			&cli.StringFlag{
				Name:        "dir",
				Aliases:     []string{"d"},
				Usage:       "directory to serve",
				Value:       server.CurrentDirectory(),
				DefaultText: "current directory",
				Required:    false,
			},
			&cli.BoolFlag{
				Name:        "tls",
				Aliases:     []string{"s"},
				Usage:       "http over tls",
				Value:       false,
				DefaultText: "false",
				Required:    false,
			},
			&cli.StringFlag{
				Name:        "crt",
				Aliases:     []string{"c"},
				Usage:       "certificate file",
				Value:       "",
				DefaultText: "empty",
				Required:    false,
			},
			&cli.StringFlag{
				Name:        "key",
				Aliases:     []string{"k"},
				Usage:       "private key",
				Value:       "",
				DefaultText: "empty",
				Required:    false,
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatalf("Error running app: %v", err)
	}

	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig
	fmt.Printf("Shutting down ...")
}
