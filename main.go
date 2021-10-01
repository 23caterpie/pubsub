package main

import (
	"log"
	"os"

	"github.com/23caterpie/pubsub/cmd/publish"

	"github.com/urfave/cli"
)

func main() {
	app := cli.App{
		Name: "pubsub",
		Commands: []cli.Command{
			publish.Command(),
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Println("Exited with error:", err.Error())
		os.Exit(1)
	}
}
