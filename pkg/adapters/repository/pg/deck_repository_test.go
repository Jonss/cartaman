package pg_test

import (
	"context"
	"testing"

	"github.com/Jonss/cartaman/pkg/adapters/repository/pg"
	"github.com/matryer/is"
)

func TestCreateDeck(t *testing.T) {
	is := is.New(t)
	conn, tearDown := pg.NewDbTestSetup(t)
	defer tearDown()

	ctx := context.Background()
	deckRepo := pg.PGDeckRepository{DB: conn}
	cardsRepo := pg.PGCardRepository{DB: conn}

	err := cardsRepo.SeedCards(ctx)
	is.NoErr(err)

	deck, err := deckRepo.CreateDeck(ctx, []int{1, 2, 3})
	is.NoErr(err)

	is.Equal(false, deck.Shuffled)
	is.True(0 < deck.ID)
	is.Equal(3, deck.Remaining)
}

func TestCreateDeck_NoCardIDs(t *testing.T) {
	is := is.New(t)
	conn, tearDown := pg.NewDbTestSetup(t)
	defer tearDown()

	ctx := context.Background()
	deckRepo := pg.PGDeckRepository{DB: conn}
	cardsRepo := pg.PGCardRepository{DB: conn}

	err := cardsRepo.SeedCards(ctx)
	is.NoErr(err)

	deck, err := deckRepo.CreateDeck(ctx, []int{})
	is.Equal(err.Error(), "error expect cardIDs length > 0")
	is.Equal(nil, deck)
}
