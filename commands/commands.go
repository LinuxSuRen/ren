package commands

import (
	"flag"

	flag "github.com/ogier/pflag"
)

var (
	CmdRunner = NewRunner()
)

type Command struct {
	Run  func(cmd *Command, args *Args)
	Flag flag.FlagSet

	Key          string
	Usage        string
	Long         string
	GitExtension bool

	subCommands map[string]*Command
}
