package cmd

import (
	"log"
)

// Supporing context of prometheus-ecs-sd
type Ctx struct {
	GOPATH   string
	Out, Err *log.Logger
	Verbose  bool
}
