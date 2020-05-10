package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

var (
	author  string
	version string
)

func main() {

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
		Usage:     "Tame all hte bits",
		Action: func(c *cli.Context) error {
			fmt.Println("send me the bits")
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
