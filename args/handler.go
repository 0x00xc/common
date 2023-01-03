package args

import (
	"errors"
	"fmt"
)

var (
	ErrContinue = errors.New("continue")
)

func And(f ...func(c *Context) bool) func(c *Context) bool {
	return func(c *Context) bool {
		for _, v := range f {
			if !v(c) {
				return false
			}
		}
		return true
	}
}

func Or(f ...func(c *Context) bool) func(c *Context) bool {
	return func(c *Context) bool {
		for _, v := range f {
			if v(c) {
				return true
			}
		}
		return false
	}
}

func Name(name string) func(c *Context) bool {
	return func(c *Context) bool {
		return c.Cmd().Cmd() == name
	}
}

func Option(options ...string) func(c *Context) bool {
	return func(c *Context) bool {
		return c.KVs().Is(options...)
	}
}

type Handler struct {
	Match   func(c *Context) bool
	Handler func(c *Context) (string, error)
	Usage   string
}

func (h *Handler) do(c *Context) (string, error) {
	if h.Match(c) {
		return h.Handler(c)
	}
	return "", ErrContinue
}

func (h *Handler) Do(c *Context) (string, error) {
	return h.do(c)
}

func (h *Handler) usage() string {
	return h.Usage
}

type handler interface {
	do(c *Context) (string, error)
	usage() string
}

type Chain []handler

func (c Chain) Add(h ...handler) Chain {
	return append(c, h...)
}

func (c Chain) Do(ctx *Context) (string, error) {
	for _, h := range c {
		msg, err := h.do(ctx)
		if err == ErrContinue {
			continue
		}
		return msg, err
	}
	if Or(Name(""), Name("help"), Option("h", "-help"))(ctx) {
		usage := fmt.Sprintf("Usage for %s:\n%s", ctx.Exe(), c.usage())
		return usage, nil
	}
	exe := ctx.Exe()
	cmd := ctx.Cmd().Cmd()
	return "", fmt.Errorf("%s %s: unknown command", exe, cmd)
}

func (c Chain) usage() string {
	var usage string
	for _, h := range c {
		usage += h.usage() + "\n"
	}
	return usage
}
