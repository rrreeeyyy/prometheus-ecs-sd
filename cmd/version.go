package main

import (
	"flag"
	"runtime"
)

var (
	version    = "devel"
	buildDate  string
	commitHash string
)

const versionHelp = `Show the prometheus-ecs-hako-sd version information`

func (cmd *versionCommand) Name() string { return "version" }
func (cmd *versionCommand) Args() string {
	return ""
}
func (cmd *versionCommand) ShortHelp() string { return versionHelp }
func (cmd *versionCommand) LongHelp() string  { return versionHelp }
func (cmd *versionCommand) Hidden() bool      { return false }

func (cmd *versionCommand) Register(fs *flag.FlagSet) {}

type versionCommand struct{}

func (cmd *versionCommand) Run(ctx *main.Ctx, args []string) error {
	ctx.Out.Printf(`prometheus-ecs-hako-sd:
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
