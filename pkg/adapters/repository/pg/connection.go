package pg

import (
	"database/sql"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file" // migration
	_ "github.com/lib/pq"                                // postgres
)

func NewConnection(datasource string) (*sql.DB, error) {
	conn, err := sql.Open("postgres", datasource)
	if err != nil {
		return nil, err
	}
	return conn, err
}

func Migrate(db *sql.DB, dbName, path string) error {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("error when migrate.WithInstance(): error=(%w)", err)
	}
	m, err := migrate.NewWithDatabaseInstance("file://"+path, dbName, driver)
	if err != nil {
		return fmt.Errorf("error when migrate.NewWithDatabaseInstance(): error=(%w)", err)
	}

	err = m.Up()
	if err != nil {
		if err == migrate.ErrNoChange {
			return nil
		}
		return fmt.Errorf("error on migrate.Up(): error=(%w)", err)
	}
	return nil
}
