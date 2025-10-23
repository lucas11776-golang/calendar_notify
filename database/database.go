package database

import (
	models "github.com/lucas11776-golang/calendar_notify/database/migrations"
	"github.com/lucas11776-golang/http/utils/env"
	"github.com/lucas11776-golang/orm"
	"github.com/lucas11776-golang/orm/drivers/sqlite"
	"github.com/lucas11776-golang/orm/migrations"
)

// Comment
func Setup() {
	db := sqlite.Connect(env.Env("DATABASE"))

	orm.DB.Add(orm.DefaultDatabaseName, db) // SQLite database config...

	Migrate()
}

// Comment
func Migrate() {
	migrations.Migrations(
		&models.TokenMigration{},
		&models.EventMigration{},
	).Up()
}
