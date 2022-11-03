package sqlutil

import "strconv"

type Builder interface {
	Build() (string, []interface{})
}

type Wrapper struct {
	Placeholder func(i int) string
	Builder     Builder
}

func (w *Wrapper) Build() (string, []interface{}) {
	statement, values := w.Builder.Build()
	var s []byte
	var i = 0
	for _, v := range statement {
		if v == '?' {
			s = append(s, []byte(w.Placeholder(i))...)
			i++
		} else {
			s = append(s, byte(v))
		}
	}
	return string(s), values
}

func DMPlaceholder(i int) string {
	return ":" + strconv.Itoa(i+1)
}
