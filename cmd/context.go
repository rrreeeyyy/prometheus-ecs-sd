package cmd

import (
	"github.com/go-kit/kit/log"
)

// Supporing context of prometheus-ecs-sd
type Ctx struct {
	GOPATH   string
	Out, Err log.Logger
	Verbose  bool
}
