/*
Copyright © 2022 Johnson Shi <Johnson.Shi@microsoft.com>
*/
package main

import (
	"flag"
	"io"
	"os"

	"github.com/spf13/cobra"
)

func newRootCmd(stdin io.Reader, stdout io.Writer, stderr io.Writer, args []string) *cobra.Command {
	cobraCmd := &cobra.Command{
		Use:     "command",
		Short:   "TODO command description",
		Example: `TODO command example`,
	}

	cobraCmd.PersistentFlags().AddGoFlagSet(flag.CommandLine)
	flags := cobraCmd.PersistentFlags()

	cobraCmd.AddCommand(
		newSubcommandCmd(stdin, stdout, stderr, args),
	)

	_ = flags.Parse(args)

	return cobraCmd
}

func execute() {
	rootCmd := newRootCmd(os.Stdin, os.Stdout, os.Stderr, os.Args[1:])
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
