/*
Copyright © 2022 Johnson Shi <Johnson.Shi@microsoft.com>
*/
package main

import (
	"io"

	"github.com/spf13/cobra"
)

type subcommandOpts struct {
	stdin  io.Reader
	stdout io.Writer
	stderr io.Writer
	arg1   string
}

func newSubcommandCmd(stdin io.Reader, stdout io.Writer, stderr io.Writer, args []string) *cobra.Command {
	opts := &subcommandOpts{
		stdin:  stdin,
		stdout: stdout,
		stderr: stderr,
	}

	cobraCmd := &cobra.Command{
		Use:     "subcommand",
		Short:   "TODO command description",
		Example: `TODO command example`,
		RunE: func(_ *cobra.Command, args []string) error {
			return opts.run()
		},
	}

	f := cobraCmd.Flags()

	f.StringVar(&opts.arg1, "arg1", "", "TODO arg1 description")
	cobraCmd.MarkFlagRequired("arg1")

	return cobraCmd
}

func (opts *subcommandOpts) run() error {
	return nil
}
