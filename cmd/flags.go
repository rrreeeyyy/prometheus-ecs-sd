package cmd

import (
	"flag"
)

type CommandLineOptions struct {
	Path            string
	Verbose         bool
	ShowVersion     bool
	OnlyECSSDEnable bool
	RefreshInterval int
}

func SetupFlagSet(name string, options *CommandLineOptions) *flag.FlagSet {
	flagSet := flag.NewFlagSet(name, flag.ContinueOnError)

	flagSet.BoolVar(&options.Verbose, "v", false, "enable verbose logging")
	flagSet.BoolVar(&options.ShowVersion, "version", false, "show the command version information")
	flagSet.BoolVar(&options.OnlyECSSDEnable, "only-ecs-sd-enable", false, "discovery only if container has `ECS_SD_ENABLE=1` environment variable")

	flagSet.StringVar(&options.Path, "path", "", "path of file to write service discovery file")

	flagSet.IntVar(&options.RefreshInterval, "refresh-interval", 60, "interval at which to scrape the AWS API for ECS service discovery")

	return flagSet
}
