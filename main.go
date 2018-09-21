// +build go1.8

package main

import (
	"os"

	"github.com/linuxsuren/ren/commands"
)

func main() {
	err := commands.CmdRunner.Execute()
	if err.Ran {
		os.Exit(err.ExitCode)
	}
}
