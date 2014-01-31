package main

import (
	"flag"
	"fmt"
	"os"
	"text/template"
)

// shamelessly snagged from the go tool
// each command gets its own set of args,
// defines its own entry point, and provides its own help
type Command struct {
	Run  func(cmd *Command, args ...string)
	Flag flag.FlagSet

	Name  string
	Usage string

	Summary string
	Help    string
}

func (c *Command) Exec(args []string) {
	c.Flag.Usage = func() {
		// c.Usage()
	}

	help := c.Flag.Bool("help", false, "print help for this command")

	err := c.Flag.Parse(args)
	if err != nil {
		fmt.Println("error parsing flags: ", err)
		return
	}

	if *help {
		c.PrintDetailedHelp()
		return
	}

	c.Run(c, c.Flag.Args()...)
}

func (c *Command) GetFlagValue(key string) interface{} {
	f := c.Flag.Lookup(key)
	if f == nil {
		return nil
	}

	g, ok := f.Value.(flag.Getter)
	if !ok {
		return nil
	}
	return g.Get()
}

func (c *Command) PrintDetailedHelp() {
	detailedTemplate.Execute(os.Stdout, c)
	c.Flag.PrintDefaults()
}

var detailedTemplate = template.Must(template.New("detailed-help").Parse(
	` {{.Name}} usage:

    $ goose {{.Name}} {{.Usage}}

    {{.Summary}}

    Additional flags:

`))
