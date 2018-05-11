package main

import (
	"os"

	"github.com/rrreeeyyy/prometheus-ecs-hako-sd/cmd"
)

func main() {
	exit := cmd.main()
	os.Exit(exit)
}
