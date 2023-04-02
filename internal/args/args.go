package args

import (
	"fmt"

	"github.com/akamensky/argparse"
)

// StartArgs arguments of start command
type StartArgs struct {
	Message string
	Project string
	Tag     string
}

// Args represent the arguments and call command
type Args interface {
	Exec(cmd Command) error
}

// Command represent the commands of timetrack
type Command interface {
	Auth() error
	Start(args StartArgs) error
	Stop() error
}

type parsed struct {
	parser *argparse.Parser
	auth   struct {
		Cmd *argparse.Command
	}
	start struct {
		Args StartArgs
		Cmd  *argparse.Command
	}
	stop struct {
		Cmd *argparse.Command
	}
}

// Parse parse args and create commands
func Parse(args []string) (Args, error) {
	p := new(parsed)
	p.parser = argparse.NewParser("commands", "Track your time")

	p.auth.Cmd = p.parser.NewCommand("auth", "Auth on toggl")

	p.start.Cmd = p.parser.NewCommand("start", "Start timetrack")
	message := p.start.Cmd.StringPositional(nil)
	project := p.start.Cmd.String("p", "project", nil)
	tag := p.start.Cmd.String("t", "tag", nil)

	p.stop.Cmd = p.parser.NewCommand("stop", "Stop current timetrack")

	err := p.parser.Parse(args)
	if err != nil {
		fmt.Print(p.parser.Usage(err))
		return nil, err
	}

	p.start.Args = StartArgs{Message: *message, Project: *project, Tag: *tag}

	return p, nil
}

// Exec run selected command
func (p *parsed) Exec(cmd Command) error {
	if p.auth.Cmd.Happened() {
		return cmd.Auth()
	} else if p.start.Cmd.Happened() {
		return cmd.Start(p.start.Args)
	} else if p.stop.Cmd.Happened() {
		return cmd.Stop()
	} else {
		return fmt.Errorf("bad arguments, please check usage")
	}
}
