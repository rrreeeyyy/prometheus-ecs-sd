package main

import (
	"log"
)

type Ctx struct {
	GOPATH   string
	Out, Err *log.Logger
	Verbose  bool
}
