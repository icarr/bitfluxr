package main

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

var (
	author  string
	version string
)

func main() {
	initLogging()
	app := initCLI()

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func initCLI() *cli.App {
	app := &cli.App{
		Name:    "BitFluxR",
		Version: "0.0.1",
		Authors: []*cli.Author{
			{
				Name:  "Ian Carr",
				Email: "ian@bitfluxr.com",
			},
		},
		Copyright: "(c) 2020 BitFluxR",
		Usage:     "Tame all the bits",
		//Flags
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "verbose",
				Aliases: []string{"vv"},
				Usage:   "verbose output",
			},
			&cli.StringFlag{
				Name:    "config, c",
				Aliases: []string{"c"},
				Usage:   "load config from `FILE`",
			},
			&cli.StringFlag{
				Name: "log",
			},
		},
		//Commands
		Commands: []*cli.Command{
			{
				Name: "add",
				//Aliases: []string{"a"},
				Usage: "add data source",
				Action: func(c *cli.Context) error {
					fmt.Println("added source: ", c.Args().First())
					return nil
				},
			},
			{
				Name: "complete",
				//Aliases: []string{"c"},
				Usage: "complete a task on the list",
				Action: func(c *cli.Context) error {
					fmt.Println("completed task: ", c.Args().First())
					return nil
				},
			},
		},
	}
	return app
}

func initLogging() {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
}
