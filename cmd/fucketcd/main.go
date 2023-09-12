package main

import (
	"log"
	"os"

	"github.com/thejezzi/fucketcd/pkg/cli"
)

func main() {
	app := cli.Parse()

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
