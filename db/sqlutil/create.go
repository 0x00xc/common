package sqlutil

import (
	"database/sql"
	"fmt"
	"strings"
)

type CreateBuilder struct {
	table   string
	columns []string
	comment string
}

func Create(table string) *CreateBuilder {
	return &CreateBuilder{table: table}
}

func (b *CreateBuilder) Column(columns ...*ColumnBuilder) *CreateBuilder {
	for _, col := range columns {
		b.C(col.build())
	}
	return b
}

func (b *CreateBuilder) C(columns ...string) *CreateBuilder {
	b.columns = append(b.columns, columns...)
	return b
}

func (b *CreateBuilder) PK(column ...string) *CreateBuilder {
	pk := "PRIMARY KEY (" + strings.Join(column, ",") + ")"
	b.columns = append(b.columns, pk)
	return b
}

func (b *CreateBuilder) FK(column string, refTable, refColumn string) *CreateBuilder {
	b.columns = append(b.columns, fmt.Sprintf("FOREIGN KEY (%s) REFERENCES %s(%s)", column, refTable, refColumn))
	return b
}

func (b *CreateBuilder) Index(indexName string, column string, more ...string) *CreateBuilder {
	if len(more) > 0 {
		column = column + "," + strings.Join(more, ",")
	}
	idx := "INDEX " + indexName + " (" + column + ")"
	b.columns = append(b.columns, idx)
	return b
}

func (b *CreateBuilder) Unique(indexName string, column string, more ...string) *CreateBuilder {
	if len(more) > 0 {
		column = column + "," + strings.Join(more, ",")
	}
	idx := "UNIQUE INDEX " + indexName + " (" + column + ")"
	b.columns = append(b.columns, idx)
	return b
}

func (b *CreateBuilder) Comment(comment string) *CreateBuilder {
	b.comment = comment
	return b
}

func (b *CreateBuilder) Build() (string, []interface{}) {
	builder := new(strings.Builder)
	builder.WriteString("CREATE TABLE " + b.table + " (")
	builder.WriteString(strings.Join(b.columns, ", "))
	builder.WriteString(")")
	if b.comment != "" {
		builder.WriteString(" COMMENT '" + b.comment + "'")
	}
	return builder.String(), nil
}

func (b *CreateBuilder) Exec(db *sql.DB) (sql.Result, error) {
	st, val := b.Build()
	return db.Exec(st, val...)
}

func (b *CreateBuilder) Create(db *sql.DB) error {
	_, err := b.Exec(db)
	return err
}

type ColumnBuilder struct {
	column string
	typ    string
	ext    []string
}

func Column(column string, typ string, ext ...string) *ColumnBuilder {
	return &ColumnBuilder{
		column: column,
		typ:    typ,
		ext:    ext,
	}
}

func C(column string, typ string, ext ...string) *ColumnBuilder {
	return Column(column, typ, ext...)
}

func (b *ColumnBuilder) NotNull() *ColumnBuilder {
	b.ext = append(b.ext, "NOT NULL")
	return b
}

func (b *ColumnBuilder) AutoIncrement() *ColumnBuilder {
	b.ext = append(b.ext, "AUTO_INCREMENT")
	return b
}

func (b *ColumnBuilder) Default(v string) *ColumnBuilder {
	var val string
	if v == "" {
		v = "''"
	}
	val = fmt.Sprintf("DEFAULT %s", v)
	b.ext = append(b.ext, val)
	return b
}

func (b *ColumnBuilder) Comment(comment string) *ColumnBuilder {
	b.ext = append(b.ext, "COMMENT '"+comment+"'")
	return b
}

func (b *ColumnBuilder) build() string {
	return fmt.Sprintf("%s %s %s", b.column, b.typ, strings.Join(b.ext, " "))
}
