package main

import (
	"code"
	"context"
	"fmt"
	"log"
	"os"
	"github.com/urfave/cli/v3"
)

/** Точка входа в приложение. */
func main() {
    cmd := &cli.Command{
        Name:  "gendiff",
        Usage: "Compares two configuration files and shows a difference",
        Action: func(ctx context.Context, cmd *cli.Command) error {
			fmt.Println("boom! I say!")
            result, err := code.GenDiff("asfd", "hfgh", "dzfg")
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