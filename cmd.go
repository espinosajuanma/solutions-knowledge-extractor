package notebook

import (
	"fmt"

	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/conf"
	"github.com/rwxrob/help"
	"github.com/rwxrob/term"
	"github.com/rwxrob/vars"
)

var solutions = NewSolutions()
var prefix = solutions.App.Name + "." + string(solutions.App.Env)

func init() {
	Z.Vars.SoftInit()
	token := Z.Vars.Get(getTokenKey())
	if token != "" {
		solutions.SetToken(token)
	}
	email := Z.Vars.Get(getEmailKey())
	if email != "" {
		solutions.User.Email = email
	}
}

var Cmd = &Z.Cmd{
	Name: `notebook`,
	Commands: []*Z.Cmd{
		// Standard Bonzai commands
		help.Cmd, vars.Cmd, conf.Cmd,
		// Custom commands
		loginCmd,
		ticketCmd,
		tasksCmd,
	},
	Shortcuts:   Z.ArgMap{},
	Version:     `v0.0.1`,
	Source:      `https://github.com/espinosajuanma/slingr-notebook`,
	Issues:      `https://github.com/espinosajuanma/slingr-notebook/issues`,
	Summary:     help.S(_main),
	Description: help.D(_main),
}

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

var loginCmd = &Z.Cmd{
	Name:        `login`,
	Aliases:     []string{},
	Commands:    []*Z.Cmd{help.Cmd},
	Summary:     help.S(_login),
	Description: help.D(_login),
	Usage:       ``,
	MinArgs:     0,

	Call: func(x *Z.Cmd, args ...string) error {
		email, _ := x.Caller.Get(getEmailKey())
		return login(email)
	},
}

func login(email string) error {
	if email != "" {
		solutions.User.Email = email
	}
	if solutions.User.Email == "" {
		email = term.Prompt("Email: ")
		if email == "" {
			return fmt.Errorf("email is required")
		}
		key := getEmailKey()
		Z.Vars.Set(key, email)
		solutions.User.Email = email
	}
	password := term.PromptHidden("Password: ")
	if password == "" {
		return fmt.Errorf("password is required")
	}
	solutions.User.Password = password

	err := solutions.Login()
	if err != nil {
		return err
	}

	return Z.Vars.Set(getTokenKey(), solutions.App.Token)
}

func getEmailKey() string {
	return prefix + ".email"
}

func getTokenKey() string {
	return prefix + ".token"
}

// --- Help Documentation ---

var _main = `
The solutions command line tool allows you to query tickets and tasks
directly from the terminal.
`

var _tickets = `
The tickets command fetches all support tickets associated with the
specified pool name.
`

var _tasks = `
The tasks command fetches specific tasks associated with the
specified project name.
`

var _login = `
`
