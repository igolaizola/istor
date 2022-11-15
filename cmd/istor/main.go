package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/igolaizola/istor"
	"github.com/peterbourgon/ff/v3"
	"github.com/peterbourgon/ff/v3/ffcli"
)

func main() {
	// Create signal based context
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	// Launch command
	cmd := newCommand()
	if err := cmd.ParseAndRun(ctx, os.Args[1:]); err != nil {
		log.Fatal(err)
	}
}

func newCommand() *ffcli.Command {
	fs := flag.NewFlagSet("istor", flag.ExitOnError)
	ip := fs.String("ip", "", "ip address")

	return &ffcli.Command{
		ShortUsage: "istor -ip <ip>",
		Options: []ff.Option{
			ff.WithConfigFileFlag("config"),
			ff.WithConfigFileParser(ff.PlainParser),
			ff.WithEnvVarPrefix("istor"),
		},
		FlagSet: fs,
		Exec: func(ctx context.Context, args []string) error {
			if *ip == "" {
				return errors.New("missing ip")
			}
			check, err := istor.IsExitNode(ctx, *ip)
			if err != nil {
				return err
			}
			fmt.Printf("%v\n", check)
			return nil
		},
	}
}
