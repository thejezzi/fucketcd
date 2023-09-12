package cli

import (
	"github.com/thejezzi/fucketcd/pkg/manager"
	"github.com/urfave/cli/v2"
)

func Parse() *cli.App {
	return &cli.App{
		Name:        "fucketcd",
		Description: "If you hate etcd and REDA components as much as me, this is for you",
		Commands: []*cli.Command{
			{
				Name:    "import",
				Aliases: []string{"i"},
				Usage:   "push a structure into the fucker",
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name:    "yes",
						Value:   false,
						Usage:   "prevents confirmation message",
						Aliases: []string{"y"},
					},
				},
				Action: manager.Import,
			},
			{
				Name:    "export",
				Aliases: []string{"e"},
				Usage:   "choke him to death",
				Action:  manager.Export,
			},
		},
	}
}
