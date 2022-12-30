package args

import (
	"os"
	"strings"
)

// ParseArgs
// Deprecated: use Parse() instead
func ParseArgs() (map[string]string, []string) {
	return parse()
}

func parse() (map[string]string, []string) {
	var args = os.Args[1:]
	var i = 0
	var k string
	var single []string
	var kv = map[string]string{}
	for i < len(args) {
		val := args[i]
		if strings.HasPrefix(val, "-") {
			if k != "" {
				kv[k] = ""
			}
			k = val[1:]
		} else {
			if k != "" {
				kv[k] = val
				k = ""
			} else {
				single = append(single, val)
			}
		}
		i++
	}
	if k != "" {
		kv[k] = ""
	}
	return kv, single
}

type KV map[string]string
type Cmd []string

func (kv KV) Is(options ...string) bool {
	for _, opt := range options {
		if _, ok := kv[opt]; ok {
			return ok
		}
	}
	return false
}

func (kv KV) Val(options ...string) string {
	for _, opt := range options {
		if val, ok := kv[opt]; ok {
			return val
		}
	}
	return ""
}

func (c Cmd) Cmd() string {
	if len(c) > 0 {
		return c[0]
	}
	return ""
}

func (c Cmd) Options() []string {
	if len(c) > 1 {
		return c[1:]
	}
	return []string{}
}

func Parse() (KV, Cmd) {
	return parse()
}
