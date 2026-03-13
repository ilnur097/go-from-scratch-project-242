package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"code/code"
        "github.com/urfave/cli/v3"
)

func main() {
	
	cmd := &cli.Command{
		Name:  "hexlet-path-size",
		Usage: "calculate path size",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "human",
				Aliases: []string{"H"},
				Usage:   "human-readable sizes (auto-select unit)",
			},
			&cli.BoolFlag{
				Name:    "all",
				Aliases: []string{"a"},
				Usage:   "include hidden files and directories",
			},
		},
		Action: func(ctx context.Context, c *cli.Command) error {
			if c.Args().Len() < 1 {
				return nil
			}

			path := c.Args().First()
			human := c.Bool("human")
			all := c.Bool("all"

			result, err := code.GetPathSize(path, false, human, all)
			if err != nil {
				return err
			}

			fmt.Println(result)
			return nil
		},
	}

	
	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
