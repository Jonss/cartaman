package pg_test

import (
	"context"
	"testing"

	"github.com/Jonss/cartaman/pkg/adapters/repository/pg"
	"github.com/matryer/is"
)

func TestSeedCards(t *testing.T) {
	// given
	is := is.New(t)
	conn, tearDown := pg.NewDbTestSetup(t)
	defer tearDown()

	repo := pg.PGCardRepository{DB: conn}

	// when
	err := repo.SeedCards(context.Background())

	// then
	is.NoErr(err)
}

func TestSeedCards_WhenSeedIsAlreadyCalled(t *testing.T) {
	// given
	is := is.New(t)
	conn, tearDown := pg.NewDbTestSetup(t)
	defer tearDown()

	repo := pg.PGCardRepository{DB: conn}

	// call once to seed
	err := repo.SeedCards(context.Background())
	is.NoErr(err)

	// when
	// call twice to check if got errors
	err = repo.SeedCards(context.Background())

	// then
	is.NoErr(err)
}

func TestGetCardIDs(t *testing.T) {
	// given
	is := is.New(t)
	conn, tearDown := pg.NewDbTestSetup(t)
	defer tearDown()

	repo := pg.PGCardRepository{DB: conn}

	err := repo.SeedCards(context.Background())
	is.NoErr(err)

	testCases := []struct {
		name          string
		cardIDs       []string
		expectedCards int
	}{
		{
			name:          "should get 52 cards",
			cardIDs:       []string{},
			expectedCards: 52,
		},
		{
			name:          "should get 3 cards when receive valid codes",
			cardIDs:       []string{"AS", "KD", "AC"},
			expectedCards: 3,
		},
		{
			name:          "should get no cards when receive invalid codes",
			cardIDs:       []string{"XD", "KO", "XP"},
			expectedCards: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// when
			cardIDs, err := repo.GetCardIDs(context.Background(), tc.cardIDs)
			is.NoErr(err)

			// then
			is.Equal(tc.expectedCards, len(cardIDs))
		})
	}
}
