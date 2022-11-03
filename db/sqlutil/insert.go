package sqlutil

import (
	"database/sql"
	"strings"
)

type InsertBuilder struct {
	table   string
	columns []string
	values  []interface{}

	onDuplicateColumns []string
	onDuplicateValues  []interface{}
}

func Insert(table string) *InsertBuilder {
	return &InsertBuilder{table: table}
}

func (b *InsertBuilder) Value(column string, value interface{}) *InsertBuilder {
	b.columns = append(b.columns, column)
	b.values = append(b.values, value)
	return b
}

func (b *InsertBuilder) OnDuplicate(kvs ...*KVPair) *InsertBuilder {
	for _, kv := range kvs {
		col, val := kv.KV()
		b.onDuplicateColumns = append(b.onDuplicateColumns, col)
		b.onDuplicateValues = append(b.onDuplicateValues, val)
	}
	return b
}

func (b *InsertBuilder) OnDuplicateColumn(column string, val interface{}) *InsertBuilder {
	b.OnDuplicate(KV(column, val))
	return b
}

func (b *InsertBuilder) Build() (string, []interface{}) {
	var values = b.values
	builder := &strings.Builder{}
	builder.WriteString("INSERT INTO " + b.table + " (")
	builder.WriteString(strings.Join(b.columns, ","))
	builder.WriteString(") VALUES (")
	v := strings.Repeat("?,", len(b.columns))
	builder.WriteString(v[:len(v)-1])
	builder.WriteString(")")
	if len(b.onDuplicateColumns) > 0 {
		builder.WriteString(" ON DUPLICATE KEY UPDATE ")
		builder.WriteString(strings.Join(b.onDuplicateColumns, ", "))
		values = append(values, b.onDuplicateValues...)
	}
	statement := builder.String()
	return statement, values
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
