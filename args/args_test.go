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
			Usage: "build this project",
		},
		&Handler{
			Match: Or(Name("help"), Option("h")),
			Handler: func(c *Context) (string, error) {
				return "help message", nil
			},
			Usage: "show help message",
		},
	}

	t.Log(c.Do(ctx))
	t.Log(c.Usage())

}
