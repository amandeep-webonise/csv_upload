package migrator

import (
	"database/sql"
	"fmt"

	migrate "github.com/rubenv/sql-migrate"
)

// DbMigrator defines behaviour for database migrations
type DbMigrator interface {
	Migrate(driver string, db *sql.DB) error
}

// DBMigratorImpl implements DbMigrator for database migration
type DBMigratorImpl struct {
	PackrMigrationSource migrate.PackrMigrationSource
}

// Migrate performs migration by taking source of packer box
func (d *DBMigratorImpl) Migrate(driver string, db *sql.DB) error {

	if _, err := migrate.Exec(db, driver, d.PackrMigrationSource, migrate.Up); err != nil {
		return fmt.Errorf("failed to migrate %v", err)
	}

	return nil
}
