package args

import (
	"errors"
	"strings"
)

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

func (c *Context) Child() (*Context, error) {
	if len(c.cmd) == 0 {
		return nil, errors.New("child command not found")
	}
	child := &Context{
		args: c.args,
		cmd:  c.cmd[1:],
		kvs:  c.kvs,
	}
	return child, nil
}
