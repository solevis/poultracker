package migrations

import (
	"embed"
	"log"

	"git.sula.io/solevis/poultracker/internal/database"
	migrate "github.com/rubenv/sql-migrate"
)

//go:embed *
var migrationsFiles embed.FS

func Run() error {
	db := database.GetDB()

	migrations := &migrate.EmbedFileSystemMigrationSource{
		FileSystem: migrationsFiles,
		Root:       ".",
	}

	if _, err := migrate.Exec(db, "sqlite3", migrations, migrate.Up); err != nil {
		return err
	}

	log.Println("Migrated database")

	return nil
}
