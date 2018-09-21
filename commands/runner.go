package commands

import "os"

type ExecError struct {
	Err      error
	Ran      bool
	ExitCode int
}

type Runner struct {
	commands map[string]*Command
	execute  func() error
}

func NewRunner() *Runner {
	return &Runner{}
}

func (r *Runner) Execute() ExecError {
	args := NewArgs(os.Args[1:])
	args.ProgramPath = os.Args[0]

	if args.Command == "" && len(args.GlobalFlags) == 0 {
		args.Command = "help"
	}
	return ExecError{}
}
