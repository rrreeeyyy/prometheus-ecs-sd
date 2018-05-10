package cmd

import (
	"flag"
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

	flags := flag.NewFlagSet(args[0], flag.ContinueOnError)
	verbose := flags.Bool("v", false, "enable verbose logging")
	version := flags.Bool("version", false, "show the command version information")
	path := flag.String("path", "", "path of file to write service discovery file")

	fmt.Printf("%s\n", *path)
	fmt.Printf("%s\n", *discovery_container_port)
	fmt.Printf("%s\n", *only_ecs_sd_enable)

	if err := flags.Parse(args[1:]); err != nil {
		return errorExitCode
	}

	outLogger := log.New(c.Stdout, "", 0)
	errLogger := log.New(c.Stderr, "", 0)

	ctx := &context.Ctx{
		Out:     outLogger,
		Err:     errLogger,
		Verbose: *verbose,
		Path:    string,
	}

	if *version {
		ShowVersion(ctx, args)
		return errorExitCode
	}

	return successExitCode
}
