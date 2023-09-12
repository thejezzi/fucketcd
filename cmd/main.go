package main

import (
	"github/thejezzi/fucketcd/cli"
	"log"
	"os"
)

func main() {
	app := cli.Parse()

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
