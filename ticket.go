package notebook

import (
	"github.com/espinosajuanma/solutions-knowledge-extractor/parser"
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

		// Get Tickets
		pool, err := solutions.GetPoolByName(poolName)
		if err != nil {
			return err
		}
		tickets, err := solutions.GetTicketsByPool(pool)
		if err != nil {
			return err
		}

		// Print Output
		var out string
		if outputMode == "html" {
			out, err = parser.ToHTML("tickets", tickets)
			if err != nil {
				return err
			}
		} else {
			out, err = parser.ToMarkdown("tickets", tickets)
			if err != nil {
				return err
			}
		}

		term.Print(out)
		return nil
	},
}

var _tickets = `
The tickets command fetches all support tickets associated with the
specified pool name.
`
