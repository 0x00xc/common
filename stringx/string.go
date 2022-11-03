package stringx

import (
	"bytes"
	"strings"
	"unicode/utf8"
)

func ToSnake(s string) string {
	if !isASCII(s) {
		return s
	}
	var b []byte
	var last = 0
	for _, v := range []byte(s) {
		var n int
		if v >= 'A' && v <= 'Z' {
			n = 1
		} else {
			n = 0
		}
		if n == 0 {
			b = append(b, v)
		} else {
			if last == 1 || len(b) == 0 {
				b = append(b, v+32)
			} else {
				b = append(b, '_', v+32)
			}
		}
		last = n
	}
	return string(b)
}

func ToCamel(s string, split ...string) string {
	if !isASCII(s) {
		return s
	}

	sp := "_"
	if len(split) > 0 {
		sp = split[0]
	}
	b := bytes.Split([]byte(s), []byte(sp))
	for _, v := range b {
		if v[0] >= 'a' && v[0] <= 'z' {
			v[0] = v[0] - 32
		}
	}
	return string(bytes.Join(b, []byte("")))
}

func isASCII(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] >= utf8.RuneSelf {
			return false
		}
	}
	return true
}

func Unique(array []string) []string {
	if len(array) < 2 {
		return array
	}
	tmp := make(map[string]uint8)
	for _, v := range array {
		tmp[v] = 1
	}
	var a = make([]string, len(tmp))
	var i = 0
	for k := range tmp {
		a[i] = k
		i++
	}
	return a
}

func ReplaceFunc(s, old string, new func(i int) string) string {
	// Compute number of replacements.
	m := strings.Count(s, old)
	if m == 0 {
		return s // avoid allocation
	}

	// Apply replacements to buffer.
	var b strings.Builder
	//b.Grow(len(s) + n*(len(new)-len(old)))

	start := 0
	for i := 0; i < m; i++ {
		j := start
		if len(old) == 0 {
			if i > 0 {
				_, wid := utf8.DecodeRuneInString(s[start:])
				j += wid
			}
		} else {
			j += strings.Index(s[start:], old)
		}
		b.WriteString(s[start:j])
		b.WriteString(new(i))
		start = j + len(old)
	}
	b.WriteString(s[start:])
	return b.String()
}
