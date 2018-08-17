package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/prometheus/prometheus/documentation/examples/custom-sd/adapter"
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

	logger := log.NewLogger(os.Stdout, os.Stderr)

	if options.ShowVersion {
		ShowVersion(args)
		return errorExitCode
	}

	cfg := discovery.SDConfig{
		RefreshInterval: options.RefreshInterval,
		OnlyECSEnable:   options.OnlyECSSDEnable,
		Region:          options.Region,
		AccessKey:       options.AccessKey,
		SecretKey:       options.SecretKey,
		Profile:         options.Profile,
		RoleARN:         options.RoleARN,
	}

	ctx := context.Background()

	disc, err := discovery.NewDiscovery(cfg, logger)

	sdAdapter := adapter.NewAdapter(ctx, options.Path, "ECSSD", disc, logger.Err)
	sdAdapter.Run()

	return successExitCode
}
