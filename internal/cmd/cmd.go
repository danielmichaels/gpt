package cmd

import (
	"fmt"
	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/conf"
	"github.com/rwxrob/help"
	"github.com/rwxrob/vars"
	"os"
	"text/template"
)

func init() {
	_ = Z.Conf.SoftInit()
	_ = Z.Vars.SoftInit()
	err := checkRequirements()
	if err != nil {
		Z.ExitError(fmt.Errorf("requirement initialiser failed: %v", err))
	}
}

var Cmd = &Z.Cmd{
	Name:    `gpt`,
	Summary: `My GPT client CLI wrapper`,
	Usage:   `gpt <text>`,
	Version: `v0.1.0`,
	License: `Apache Software License 2.0`,
	Source:  `github.com/danielmichaels/gpt`,
	Description: `
		gpt is a simple command line wrapper for *mods* and *tgpt*
		`,
	Call: gptEntry,
	Commands: []*Z.Cmd{
		// external imports below
		help.Cmd, conf.Cmd, vars.Cmd,

		// internal commands below
		TgptEntry, GptEntry,
	},
}

var GptEntry = &Z.Cmd{
	Name:    `mods`,
	Summary: `Run mods`,
	Description: `
		Run *mods*.

		*mods* uses the OpenAI API which costs money. It does not offer
		an interactive mod out of the box. 

		This CLI gives that ability through your $EDITOR; "{{ editor }}". It
		will open a temp file where you can type, copy paste etc into it. Save
		and exit to run *mods* with that text as its input.
		`,
	Other: []Z.Section{
		{Title: `Examples`, Body: `
			gpt "what is gpt"

			gpt (starts interactive mode)`},
	},
	Aliases: []string{"m"},
	Dynamic: template.FuncMap{
		"editor": func() string { return os.Getenv("EDITOR") },
	},
	Commands: []*Z.Cmd{help.Cmd},
	Call:     gptEntry,
}

var TgptEntry = &Z.Cmd{
	Name:    `tgpt`,
	Summary: `Run tgpt`,
	Description: `
		Run the *tgpt* binary. *tgpt* is a free LLM client which returns
		adequate results.

		*tgpt* has a interactive mode which this CLI will drop in to if **no**
		arguments are supplied.`,
	Other: []Z.Section{
		{Title: `Examples`, Body: `
			gpt t "what is gpt"

			gpt t (starts interactive mode)`},
	},
	Aliases:  []string{"t"},
	Commands: []*Z.Cmd{help.Cmd},
	Call:     tgptEntry,
}
