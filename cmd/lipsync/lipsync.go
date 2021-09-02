package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "lipsync",
		Usage: "create podcast RSS files on the fly!",
		Commands: []*cli.Command{
			{
				Name:  "generate",
				Usage: "generates a podcast RSS file",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "config",
						Usage: "the path to the configuration",
						Value: "lipsync.yaml",
					},
					&cli.StringFlag{
						Name:  "out-file",
						Usage: "the name of the output file",
						Value: "out.rss",
					},
				},
				Action: generateAction,
			},
			{
				Name:  "serve",
				Usage: "starts a HTTP server with the podcast information",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "config",
						Usage: "the path to the configuration",
						Value: "lipsync.yaml",
					},
					&cli.IntFlag{
						Name:  "port",
						Usage: "the port for the web server",
						Value: 8080,
					},
				},
				Action: serveAction,
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
