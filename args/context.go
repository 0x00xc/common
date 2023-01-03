package args

import "strings"

type Context struct {
	args Cmd
	cmd  Cmd
	kvs  KV
}

func (c *Context) KVs() KV {
	return c.kvs
}

func (c *Context) Origin() string {
	return strings.Join(c.args, " ")
}

func (c *Context) Exe() string {
	return c.args.Cmd()
}

func (c *Context) Cmd() Cmd {
	return c.cmd
}
