package cmd

import (
	"runtime"

	"github.com/rrreeeyyy/prometheus-ecs-sd/context"
)

var (
	version    = "devel"
	buildDate  string
	commitHash string
)

func ShowVersion(ctx *context.Ctx, args []string) error {
	ctx.Out.Printf(`prometheus-ecs-sd:
	version     : %s
	build date  : %s
	git hash    : %s
	go version  : %s
	go compiler : %s
	platform    : %s/%s
`, version, buildDate, commitHash,
		runtime.Version(), runtime.Compiler, runtime.GOOS, runtime.GOARCH)
	return nil
}
