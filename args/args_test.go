package args

import "testing"

func TestArgs(t *testing.T) {
	ctx := Parse()

	c := Chain{
		&Handler{
			Match: Name("build"),
			Handler: func(c *Context) (string, error) {
				return "build success", nil
			},
			Usage: "build:\n\tbuild this project\n\tusage: demo build -o <output dir>",
		},
		&Handler{
			Match: Or(Name("help"), Option("h")),
			Handler: func(c *Context) (string, error) {
				return "this is help message", nil
			},
			Usage: "help:\n\tshow help message",
		},
	}

	t.Log(c.Do(ctx))

}
