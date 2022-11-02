package sqlutil

import "testing"

func TestCreate(t *testing.T) {
	b := Create("example").
		Column(
			C("id", "varchar(32)", "not null"),
			C("username", "varchar(32)").NotNull().Default("''").Comment("username"),
			C("created_at", "timestamp(0)").NotNull().Default("now()"),
		).
		C("uid varchar(32) not null comment 'uid'").
		Unique("idx_username", "username").
		PK("id").
		FK("uid", "table_uid", "id")
	t.Log(b.Build())
}
