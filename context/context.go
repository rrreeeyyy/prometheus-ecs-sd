package context

import (
	"flag"
	"log"
)

type Ctx struct {
	GOPATH   string
	Out, Err *log.Logger
	Verbose  bool
}

type command interface {
	Name() string
	Args() string
	ShortHelp() string
	LongHelp() string
	Register(*flag.FlagSet)
	Hidden() bool
	Run(*Ctx, []string) error
}
