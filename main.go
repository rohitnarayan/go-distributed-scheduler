package main

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	version = "0.0.1"
	commit  = "n/a"
)

func main() {
	cli := &cobra.Command{
		Use:   "go-distributed-scheduler",
		Short: "Distributed Scheduler in Go",
	}

	cli.AddCommand(commands()...)
	cli.Version = fmt.Sprintf("%s (Commit: %s)", version, commit)
}

func commands() []*cobra.Command {
	return []*cobra.Command{
		{
			Use:     "server",
			Short:   "start server",
			Aliases: []string{"server", "start"},
			RunE: func(_ *cobra.Command, _ []string) error {
				return nil
			},
		},
	}
}
