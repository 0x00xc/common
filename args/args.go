package args

import (
	"os"
	"strings"
)

func parse() (map[string][]string, []string) {
	var args = os.Args[1:]
	var i = 0
	var k string
	var single []string
	var kv = map[string][]string{}
	for i < len(args) {
		val := args[i]
		if strings.HasPrefix(val, "-") {
			if k != "" {
				if _, ok := kv[k]; !ok {
					kv[k] = []string{}
				} else {

				}
			}
			k = val[1:]
		} else {
			if k != "" {
				kv[k] = append(kv[k], val)
				k = ""
			} else {
				single = append(single, val)
			}
		}
		i++
	}
	if k != "" {
		if _, ok := kv[k]; !ok {
			kv[k] = []string{}
		}
	}
	return kv, single
}

type KV map[string][]string
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
			if len(val) > 0 {
				return val[0]
			} else {
				return ""
			}
		}
	}
	return ""
}

func (kv KV) Vals(options ...string) []string {
	for _, opt := range options {
		if val, ok := kv[opt]; ok {
			return val
		}
	}
	return nil
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
