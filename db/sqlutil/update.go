package sqlutil

import (
	"database/sql"
	"strings"
)

type UpdateBuilder struct {
	table   string
	columns []string
	values  []interface{}
	c       condition
}

func Update(table string) *UpdateBuilder {
	return &UpdateBuilder{table: table}
}

func (b *UpdateBuilder) Set(column string, val interface{}) *UpdateBuilder {
	//column = placeholder.ReplaceAllString(column, "?")
	if !strings.Contains(column, "=") {
		column = column + " = ?"
	}
	b.columns = append(b.columns, column)
	b.values = append(b.values, val)
	return b
}

func (b *UpdateBuilder) Where(where string, val interface{}) *UpdateBuilder {
	b.c.Where(where, val)
	return b
}

func (b *UpdateBuilder) Or(or *AND) *UpdateBuilder {
	b.c.Or(or)
	return b
}

func (b *UpdateBuilder) Build() (string, []interface{}) {
	var values []interface{}
	builder := &strings.Builder{}
	builder.WriteString("UPDATE " + b.table + " SET ")
	builder.WriteString(strings.Join(b.columns, ", "))
	values = append(values, b.values...)
	values = append(values, b.c.build(builder)...)
	statement := builder.String()
	//statement = replace(statement)
	return statement, values
}

func (b *UpdateBuilder) Exec(db *sql.DB) (sql.Result, error) {
	st, val := b.Build()
	return db.Exec(st, val...)
}

func (b *UpdateBuilder) Update(db *sql.DB) (int64, error) {
	re, err := b.Exec(db)
	if err != nil {
		return 0, err
	}
	return re.RowsAffected()
}
