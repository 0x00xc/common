package sqlutil

import (
	"database/sql"
	"strings"
)

type InsertBuilder struct {
	table   string
	columns []string
	values  []interface{}
}

func Insert(table string) *InsertBuilder {
	return &InsertBuilder{table: table}
}

func (b *InsertBuilder) Value(column string, value interface{}) *InsertBuilder {
	b.columns = append(b.columns, column)
	b.values = append(b.values, value)
	return b
}

func (b *InsertBuilder) Build() (string, []interface{}) {
	builder := &strings.Builder{}
	builder.WriteString("INSERT INTO " + b.table + " (")
	builder.WriteString(strings.Join(b.columns, ","))
	builder.WriteString(") VALUES (")
	v := strings.Repeat("?,", len(b.columns))
	builder.WriteString(v[:len(v)-1])
	builder.WriteString(")")
	statement := builder.String()
	return statement, b.values
}

func (b *InsertBuilder) Exec(db *sql.DB) (sql.Result, error) {
	st, val := b.Build()
	return db.Exec(st, val...)
}

func (b *InsertBuilder) Insert(db *sql.DB) (int64, error) {
	re, err := b.Exec(db)
	if err != nil {
		return 0, err
	}
	return re.LastInsertId()
}
