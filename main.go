package main

import (
	"fucketcd/cli"
	"log"
	"os"
)

func main() {
	app := cli.Parse()

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
