package sqlutil

import (
	"database/sql"
	"strings"
)

type DeleteBuilder struct {
	table string
	c     condition
}

func Delete(table string) *DeleteBuilder {
	return &DeleteBuilder{table: table}
}

func (b *DeleteBuilder) Where(where string, val ...interface{}) *DeleteBuilder {
	b.c.Where(where, val...)
	return b
}

func (b *DeleteBuilder) Or(or *AND) *DeleteBuilder {
	b.c.Or(or)
	return b
}

func (b *DeleteBuilder) Build() (string, []interface{}) {
	builder := new(strings.Builder)
	builder.WriteString("DELETE FROM " + b.table)
	values := b.c.build(builder)
	statement := builder.String()
	//return replace(builder.String()), values
	return statement, values
}

func (b *DeleteBuilder) Exec(db *sql.DB) (sql.Result, error) {
	st, val := b.Build()
	return db.Exec(st, val...)
}

func (b *DeleteBuilder) Delete(db *sql.DB) (int64, error) {
	re, err := b.Exec(db)
	if err != nil {
		return 0, err
	}
	return re.RowsAffected()
}
