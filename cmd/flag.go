package cmd

import (
	"flag"
)

type CommandLineOptions struct {
	ShowVersion bool
}

func SetupFlagSet(name string, options *CommandLineOptions) *flag.FlagSet {
	flagSet := flag.NewFlagSet(name, flag.ContinueOnError)
	flagSet.BoolVar(&options.ShowVersion, "version", false, "Show the command version information")

	return flagSet
}
