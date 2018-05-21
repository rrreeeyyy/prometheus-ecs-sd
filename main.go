package main

import (
	"os"

	"github.com/rrreeeyyy/prometheus-ecs-sd/cmd"
)

func main() {
	exit := cmd.Run(os.Args)
	os.Exit(exit)
}
