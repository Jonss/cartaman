package pg

import (
	"context"
	"database/sql"
	"testing"

	pg_container "github.com/jonss/testcontainers-go-wrapper/pg"
)

/*
	newDbTestSetup does:

- configures a container with postgres,
- creates a connection
- run migration
*/
func NewDbTestSetup(t *testing.T) (*sql.DB, func()) {
	cfg := pg_container.PostgresCfg{
		ImageName: "postgres:15-alpine",
		Password:  "a_secret_password",
		UserName:  "test",
		DbName:    "posterr_test",
	}

	pgInfo, err := pg_container.Container(context.Background(), cfg)
	if err != nil {
		t.Fatalf("error creating pgContainer. error=(%v)", err)
	}

	dbConn, err := NewConnection(pgInfo.DbURL)
	if err != nil {
		t.Fatalf("error connecting db. error=(%v)", err)
	}

	err = Migrate(dbConn, cfg.DbName)

	if err != nil {
		t.Fatalf("error connecting db. error=(%v)", err)
	}

	return dbConn, pgInfo.TearDown
}
