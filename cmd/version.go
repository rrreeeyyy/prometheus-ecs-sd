package cmd

import (
	"fmt"
	"runtime"
)

var (
	version    = "devel"
	buildDate  string
	commitHash string
)

func ShowVersion(args []string) error {
	fmt.Printf(`prometheus-ecs-sd:
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
