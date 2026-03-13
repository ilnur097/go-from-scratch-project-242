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
		Usage: "print size of a file or directory; supports -r (recursive), -H (human-readable), -a (include hidden)",
		Flags: []cli.Flag{
                       
                        &cli.BoolFlag{
				Name:    "recursive",
				Aliases: []string{"r"},
				Usage:   "recursive size of directories",
				Value:   false,
                                     },
			&cli.BoolFlag{
				Name:    "human",
				Aliases: []string{"H"},
				Usage:   "human-readable sizes (auto-select unit)",
                                Value:   false,
			},
			&cli.BoolFlag{
				Name:    "all",
				Aliases: []string{"a"},
				Usage:   "include hidden files and directories",
                                Value:   false,
			},

		},
		Action: func(ctx context.Context, c *cli.Command) error {
			if c.Args().Len() < 1 {
				return nil
			}

			path := c.Args().First()
                        recursive := c.Bool("recursive")
			human := c.Bool("human")
			all := c.Bool("all")

			result, err := code.GetPathSize(path, recursive, human, all)
			if err != nil {
				fmt.Fprintf(os.Stderr, "error: %v\n", err)
				os.Exit(1)
			}

			fmt.Println(result)
			return nil
		},
	}

	
	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
