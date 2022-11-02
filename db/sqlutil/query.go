package sqlutil

import (
	"database/sql"
	"strings"
)

type QueryBuilder struct {
	table   string
	selects []string
	order   []string
	group   string
	offset  int
	limit   int

	c condition
}

func Query(table string) *QueryBuilder {
	return &QueryBuilder{
		table: table,
	}
}

func (b *QueryBuilder) Select(sel ...string) *QueryBuilder {
	b.selects = append(b.selects, sel...)
	return b
}

func (b *QueryBuilder) Where(w string, val ...interface{}) *QueryBuilder {
	b.c.Where(w, val...)
	return b
}

func (b *QueryBuilder) Or(where *AND) *QueryBuilder {
	b.c.Or(where)
	return b
}

func (b *QueryBuilder) Order(order ...string) *QueryBuilder {
	b.order = append(b.order, order...)
	return b
}

func (b *QueryBuilder) Group(groupBy string) *QueryBuilder {
	b.group = groupBy
	return b
}
func (b *QueryBuilder) Limit(limit int) *QueryBuilder {
	b.limit = limit
	return b
}

func (b *QueryBuilder) Offset(offset int) *QueryBuilder {
	b.offset = offset
	return b
}

func (b *QueryBuilder) Build() (string, []interface{}) {
	builder := &strings.Builder{}
	builder.WriteString("SELECT ")
	if len(b.selects) > 0 {
		builder.WriteString(strings.Join(b.selects, ","))
	} else {
		builder.WriteString("*")
	}
	builder.WriteString(" FROM " + b.table)
	var values = b.c.build(builder)
	if len(b.order) > 0 {
		builder.WriteString(" ORDER BY " + strings.Join(b.order, ","))
	}
	if b.group != "" {
		builder.WriteString(" GROUP BY " + b.group)
	}
	if b.offset > 0 {
		builder.WriteString(" OFFSET ? ")
		values = append(values, b.offset)
	}
	if b.limit > 0 {
		builder.WriteString(" LIMIT ?")
		values = append(values, b.limit)
	}
	statement := builder.String()
	//statement = replace(statement)
	return statement, values
}

func (b *QueryBuilder) Query(db *sql.DB) (*sql.Rows, error) {
	var statement, values = b.Build()
	return db.Query(statement, values)
}

func (b *QueryBuilder) QueryRow(db *sql.DB) *sql.Row {
	var statement, values = b.Build()
	return db.QueryRow(statement, values)
}
