package pg_test

import (
	"context"
	"testing"

	"github.com/Jonss/cartaman/pkg/adapters/repository/pg"
	"github.com/matryer/is"
)

func TestSeedCards(t *testing.T) {
	is := is.New(t)
	conn, tearDown := pg.NewDbTestSetup(t)
	defer tearDown()

	repo := pg.PGCardRepository{DB: conn}
	err := repo.SeedCards(context.Background())
	is.NoErr(err)
}

func TestSeedCards_WhenSeedIsAlreadyCalled(t *testing.T) {
	is := is.New(t)
	conn, tearDown := pg.NewDbTestSetup(t)
	defer tearDown()

	repo := pg.PGCardRepository{DB: conn}
	// call once to seed
	err := repo.SeedCards(context.Background())
	is.NoErr(err)

	// call twice to check if got errors
	err = repo.SeedCards(context.Background())
	is.NoErr(err)
}
