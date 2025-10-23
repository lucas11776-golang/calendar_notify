package migrations

import (
	"github.com/lucas11776-golang/orm"
	"github.com/lucas11776-golang/orm/migrations"
)

type EventMigration struct{}

// Comment
func (ctx *EventMigration) Up() {
	migrations.Create(orm.DefaultDatabaseName, "events", func(table *migrations.Table) {
		table.String("id").PrimaryKey().Unique()
		table.BigInteger("start_timestamp")
		table.BigInteger("end_timestamp")
		table.String("link")
		table.String("title")
		table.Text("description").Nullable()
	})
}

// Comment
func (ctx *EventMigration) Down() {
	migrations.Drop("sqlite", "events")
}
