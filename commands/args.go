package commands

import "github.com/github/hub/cmd"

type Args struct {
	Executable  string
	GlobalFlags []string
	Command     string
	ProgramPath string
	Params      []string
	beforeChain []*cmd.Cmd
	afterChain  []*cmd.Cmd
	Noop        bool
	Terminator  bool
	noForward   bool
	Callbacks   []func() error
}

func NewArgs(args []string) *Args {
	return &Args{
		Executable: "git",
		beforeChain: make([]*cmd.Cmd, 0)
		afterChain: make([]*cmd.Cmd, 0),
	}
}
