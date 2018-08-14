package cmd

import (
	"fmt"
	"os"

	"github.com/rrreeeyyy/prometheus-ecs-sd/discovery"
	"github.com/rrreeeyyy/prometheus-ecs-sd/log"
)

var (
	successExitCode = 0
	errorExitCode   = 1
)

func Run(args []string) int {
	_, err := os.Getwd()
	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to get working directory", err)
		os.Exit(1)
	}

	options := CommandLineOptions{}
	flagSet := SetupFlagSet(args[0], &options)

	if err := flagSet.Parse(args[1:]); err != nil {
		return errorExitCode
	}

	logger := logger.NewLogger(os.Stdout, os.Stderr)

	if options.ShowVersion {
		ShowVersion(args)
		return errorExitCode
	}

	cfg := &SDConfig{
		RefreshInterval: options.RefreshInterval,
		OnlyECSEnable:   options.OnlyECSSDEnable,
		Region:          options.Region,
		AccessKey:       options.AccessKey,
		SecretKey:       options.SecretKey,
		Profile:         options.Profile,
		RoleARN:         options.RoleARN,
	}

	disc, err := NewDiscovery(cfg)

	sdAdapter := adapter.NewAdapter(ctx, *outputFile, "ECSSD", disc, logger)
	sdAdapter.Run()

	return successExitCode
}
