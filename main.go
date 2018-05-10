package main

import (
	"fmt"
)

func main() {
	fmt.Println("vim-go")
}

func (c *Config) Run() int {
	commands := [...]command{
		&versionCommand(),
	}
	return 0
}
