package cmd

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/rrreeeyyy/prometheus-ecs-sd/context"
)

var (
	successExitCode = 0
	errorExitCode   = 1
)

type Config struct {
	Env            []string
	Stdout, Stderr io.Writer
}

func Run(args []string) int {
	_, err := os.Getwd()
	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to get working directory", err)
		os.Exit(1)
	}

	c := &Config{
		Stdout: os.Stdout,
		Stderr: os.Stderr,
		Env:    os.Environ(),
	}

	options := CommandLineOptions{}
	flagSet := SetupFlagSet(args[0], &options)

	if err := flagSet.Parse(args[1:]); err != nil {
		return errorExitCode
	}

	outLogger := log.New(c.Stdout, "", 0)
	errLogger := log.New(c.Stderr, "", 0)

	ctx := &context.Ctx{
		Out:     outLogger,
		Err:     errLogger,
		Verbose: options.Verbose,
	}

	if options.ShowVersion {
		ShowVersion(ctx, args)
		return errorExitCode
	}

	return successExitCode
}
