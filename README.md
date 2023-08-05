# gpt

> My terminal GPT client
This CLI leverages other LLM based applications and packages them
into a workable application for me.
 
![demo.gif](_docs%2Fdemo.gif)

## Install

Standalone

```
go install github.com/danielmichaels/gpt/cmd/gpt@latest
```

Composed

```go
package z

import (
	Z "github.com/rwxrob/bonzai/z"
	"github.com/danielmichaels/gpt"
)

var Cmd = &Z.Cmd{
	Name:     `gpt`,
	Commands: []*Z.Cmd{help.Cmd, gpt.Cmd},
}
```

## Requirements

The following binaries must exist in the `$PATH`:

- [tgpt](https://github.com/aandrew-me/tgpt)
- [mods](https://github.com/charmbracelet/mods)

To use the default `gpt` command you'll need a OpenAI paid account and key.
Set the following environment variable for `mods` to work. Unfortunately,
this is how `mods` wants the key.

- `$OPENAI_API_KEY`

## Tab Completion

To activate bash completion just use the `complete -C` option from your
`.bashrc` or command line. There is no messy sourcing required. All the
completion is done by the program itself.

```
complete -C gpt gpt
```

If you don't have bash or tab completion check use the shortcut
commands instead.

## Embedded Documentation

All documentation (like manual pages) has been embedded into the source
code of the application. See the source or run the program with help to
access it.

## Other Examples

* <https://github.com/rwxrob/z> - the one that started it all by Bonzai's creator.
* <https://github.com/danielmichaels/zet-cmd> - my own zettelkasten bonzai app
