package migrations

import (
	"github.com/lucas11776-golang/orm"
	"github.com/lucas11776-golang/orm/migrations"
)

// TODO move to e.g drive.google.com to share all devices...
type TokenMigration struct{}

// Comment
func (ctx *TokenMigration) Up() {
	migrations.Create(orm.DefaultDatabaseName, "tokens", func(table *migrations.Table) {
		table.String("name")
		table.BigInteger("expires")
		table.Text("access_token")
	})
}

// Comment
func (ctx *TokenMigration) Down() {
	migrations.Drop("sqlite", "tokens")
}
