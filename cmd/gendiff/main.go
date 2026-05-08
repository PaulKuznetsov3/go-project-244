// Package main предоставляет инструмент командной строки для сравнения JSON/YAML файлов
package main

import (
	"code"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v3"
)

// Точка входа в приложение
func main() {
	cmd := &cli.Command{
		Name:  "gendiff",
		Usage: "Compares two configuration files and shows a difference",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "format",
				Aliases: []string{"f"},
				Usage:   "output format (default: \"stylish\")",
			},
		},
		Action: func(_ context.Context, cmd *cli.Command) error {
			fpath1 := cmd.Args().Get(0)
			fpath2 := cmd.Args().Get(1)
			format := cmd.String("format")

			result, err := code.GenDiff(fpath1, fpath2, format)
			if err != nil {
				return err
			}
			fmt.Print(result)
			return nil

		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
