package cmd

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/prometheus/prometheus/documentation/examples/custom-sd/adapter"
	"github.com/rrreeeyyy/prometheus-ecs-sd/sd"
)

var (
	successExitCode = 0
	errorExitCode   = 1
)

type RuntimeEnvironment struct {
	Env            []string
	Stdout, Stderr io.Writer
}

func Run(args []string) int {
	_, err := os.Getwd()
	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to get working directory", err)
		os.Exit(1)
	}

	re := &RuntimeEnvironment{
		Stdout: os.Stdout,
		Stderr: os.Stderr,
		Env:    os.Environ(),
	}

	options := CommandLineOptions{}
	flagSet := SetupFlagSet(args[0], &options)

	if err := flagSet.Parse(args[1:]); err != nil {
		return errorExitCode
	}

	outLogger := log.New(re.Stdout, "", 0)
	errLogger := log.New(re.Stderr, "", 0)

	ctx := &Ctx{
		Out:     outLogger,
		Err:     errLogger,
		Verbose: options.Verbose,
	}

	if options.ShowVersion {
		ShowVersion(ctx, args)
		return errorExitCode
	}

	ctx := context.Background()

	sdc := &SDConfig{
		RefreshInterval: options.RefreshInterval,
		OnlyECSEnable:   options.OnlyECSSDEnable,
	}

	sdAdapter := adapter.NewAdapter(ctx, *outputFile, "httpSD", disc, logger)
	sdAdapter.Run()

	return successExitCode
}
