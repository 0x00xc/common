package sqlutil

import (
	"strings"
)

type AND struct {
	where  []string
	values []interface{}
}

func (a *AND) Where(where string, val ...interface{}) *AND {
	//where = placeholder.ReplaceAllString(where, "?")
	where = strings.TrimSpace(where)
	if !strings.Contains(where, " ") {
		where = where + " = ?"
	}
	a.where = append(a.where, where)
	a.values = append(a.values, val...)
	return a
}

func (a *AND) build() (string, []interface{}) {
	return strings.Join(a.where, " AND "), a.values
}

func (a *AND) empty() bool {
	return a == nil || len(a.where) == 0
}

func Where(where string, val ...interface{}) *AND {
	a := new(AND)
	return a.Where(where, val...)
}

type condition struct {
	where *AND
	or    []*AND
}

func (c *condition) Where(where string, val ...interface{}) {
	if c.where == nil {
		c.where = new(AND)
	}
	c.where.Where(where, val...)
}

func (c *condition) Or(where *AND) {
	c.or = append(c.or, where)
}

func (c *condition) build(builder *strings.Builder) []interface{} {
	var values []interface{}
	if !c.where.empty() {
		builder.WriteString(" WHERE ")
		st, val := c.where.build()
		builder.WriteString(st)
		values = append(values, val...)
		for _, or := range c.or {
			if or.empty() {
				continue
			}
			orSt, orVal := or.build()
			builder.WriteString(" OR " + orSt)
			values = append(values, orVal...)
		}
	}
	return values
}
