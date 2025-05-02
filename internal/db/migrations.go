package db

import (
	"log"
	"os/user"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func RunMigrations(dbPath, migrationDir string) {
	log.Println("Running migrations...")

	usr, _ := user.Current()
	m, err := migrate.New("file://"+usr.HomeDir+migrationDir, "sqlite3://"+dbPath)
	if err != nil {
		log.Fatalf("Migration Setup Failed: %v", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Migration Failed: %v", err)
	}

	log.Println("All migrations run successfully")
}
