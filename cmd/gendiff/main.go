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
            &cli.BoolFlag{
                Name:  "format",
                Aliases: []string{"f"},
                Usage: "output format (default: \"stylish\")",
            },
        },
        Action: func(ctx context.Context, cmd *cli.Command) error {
            filepath1 := cmd.Args().Get(0)
            filepath2 := cmd.Args().Get(1)
            format := cmd.Args().Get(2)
    
            result, err := code.GenDiff(filepath1, filepath2, format)
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