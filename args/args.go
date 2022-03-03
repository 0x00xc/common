package args

import (
	"os"
	"strings"
)

func ParseArgs() (map[string]string, []string) {
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
