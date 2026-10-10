package main

import (
	"code"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:                   "hexlet-path-size",
		Usage:                  "print size of a file or directory",
		ArgsUsage:              "<path>",
		UseShortOptionHandling: true,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:        "human",
				Usage:       "human-readable sizes (auto-select unit)",
				Aliases:     []string{"H"},
				DefaultText: "false",
			},
			&cli.BoolFlag{
				Name:        "all",
				Usage:       "include hidden files and directories",
				Aliases:     []string{"a"},
				DefaultText: "false",
			},
			&cli.BoolFlag{
				Name:        "recursive",
				Usage:       "recursive size of directories",
				Aliases:     []string{"r"},
				DefaultText: "false",
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			path, recursive, human, all := "", false, false, false
			if cmd.NArg() > 0 {
				path = cmd.Args().Get(0)
			}
			if cmd.Bool("human") || cmd.Bool("H") {
				human = true
			}
			if cmd.Bool("all") || cmd.Bool("a") {
				all = true
			}
			if cmd.Bool("recursive") || cmd.Bool("r") {
				recursive = true
			}
			ret, err := code.GetPathSize(path, recursive, human, all)
			if err != nil {
				fmt.Printf("%v\n", err)
			} else {
				fmt.Printf("%s\t%s\n", ret, path)
			}
			return nil
		},
	}
	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
