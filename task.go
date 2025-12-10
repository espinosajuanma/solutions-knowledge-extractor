package notebook

import (
	"fmt"

	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/help"
	"github.com/rwxrob/term"
)

var tasksCmd = &Z.Cmd{
	Name:        `tasks`,
	Aliases:     []string{"task", "dev"},
	Commands:    []*Z.Cmd{help.Cmd},
	Summary:     help.S(_tasks),
	Description: help.D(_tasks),
	Usage:       `<project name>`,
	MinArgs:     0,

	Call: func(x *Z.Cmd, args ...string) error {
		_, err := solutions.GetCurrentUser()
		if err != nil {
			email, _ := x.Caller.Get(getEmailKey())
			err := login(email)
			if err != nil {
				return err
			}
		}
		if len(args) == 0 {
			return fmt.Errorf("project name is required")
		}
		term.Print("")

		return nil
	},
}

var _tasks = `
The tasks command fetches specific tasks associated with the
specified project name.
`
