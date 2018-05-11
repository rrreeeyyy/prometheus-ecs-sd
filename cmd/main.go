package cmd

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/rrreeeyyy/prometheus-ecs-hako-sd/context"
)

var (
	successExitCode = 0
	errorExitCode   = 1
)

type command interface {
	Name() string           // "foobar"
	Args() string           // "<baz> [quux...]"
	ShortHelp() string      // "Foo the first bar"
	LongHelp() string       // "Foo the first bar meeting the following conditions..."
	Register(*flag.FlagSet) // command-specific flags
	Hidden() bool           // indicates whether the command should be hidden from help output
	Run(*context.Ctx, []string) error
}

func main() int {
	wd, err := os.Getwd()
	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to get working directory", err)
		os.Exit(1)
	}

	args := append([]string{os.Args[0]}, flag.Args()...)
	c := &Config{
		Args:       args,
		Stdout:     os.Stdout,
		Stderr:     os.Stderr,
		WorkingDir: wd,
		Env:        os.Environ(),
	}

	exit := c.Run()
	return (exit)
}

type Config struct {
	WorkingDir     string    // Where to execute
	Args           []string  // Command-line arguments, starting with the program name.
	Env            []string  // Environment variables
	Stdout, Stderr io.Writer // Log output
}

func (c *Config) Run() int {
	commands := [...]command{
		&versionCommand{},
	}
}
