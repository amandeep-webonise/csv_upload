package bundle

import (
	"github.com/gobuffalo/packr"
	migrate "github.com/rubenv/sql-migrate"
	"github.com/webonise/csv_upload/pkg/migrator"
)

// BundleInitialiser implements the bundle methods
type BundleInitialiser interface {
	Pack() packr.Box
	Migration() *migrator.DBMigratorImpl
}

// PackerConfig contains Packr configurations
type PackerConfig struct {
	Path string
}

// Pack bundling the assets and returns the box
func (pc *PackerConfig) Pack() packr.Box {
	return packr.NewBox(pc.Path)
}

// Migration bundling the migrations
func (pc *PackerConfig) Migration() *migrator.DBMigratorImpl {
	return &migrator.DBMigratorImpl{
		PackrMigrationSource: migrate.PackrMigrationSource{
			Box: pc.Pack(),
		},
	}
}
