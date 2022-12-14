package db

import (
	"database/sql"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file" // migration
	_ "github.com/lib/pq"                                // postgres
)

func Migrate(db *sql.DB, dbName string) error {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("error when migrate.WithInstance(): error=(%w)", err)
	}
	m, err := migrate.NewWithDatabaseInstance("file://pkg/db/migrations", dbName, driver)
	if err != nil {
		return fmt.Errorf("error when migrate.NewWithDatabaseInstance(): error=(%w)", err)
	}

	err = m.Up()
	if err != nil {
		if err == migrate.ErrNoChange || err.Error() == "first .: file does not exist" {
			return nil
		}
	
		return fmt.Errorf("error on migrate.Up(): error=(%w)", err)
	}
	return nil
}