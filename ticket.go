package notebook

import (
	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/help"
	"github.com/rwxrob/term"
)

var ticketCmd = &Z.Cmd{
	Name:        `tickets`,
	Aliases:     []string{"ticket", "support"},
	Commands:    []*Z.Cmd{help.Cmd},
	Summary:     help.S(_tickets),
	Description: help.D(_tickets),
	Usage:       `<pool name>`,
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

		poolName := "Non-LIMS"
		outputMode := "markdown"
		if len(args) > 0 {
			poolName = args[0]
		}
		if len(args) > 1 {
			outputMode = args[1]
		}
		out, err := solutions.GetTicketsByPoolName(poolName, outputMode)
		if err != nil {
			return err
		}
		term.Print(out)
		return nil
	},
}

var _tickets = `
The tickets command fetches all support tickets associated with the
specified pool name.
`
