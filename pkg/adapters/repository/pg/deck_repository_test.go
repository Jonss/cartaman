package pg_test

import (
	"context"
	"testing"

	"github.com/Jonss/cartaman/pkg/adapters/repository"
	"github.com/Jonss/cartaman/pkg/adapters/repository/pg"
	"github.com/google/uuid"
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

	params := repository.CreateDeckParams{
		CardIDs:    []int{1, 2, 3},
		ExternalID: uuid.New(),
	}

	deck, err := deckRepo.CreateDeck(ctx, params)
	is.NoErr(err)

	is.Equal(false, deck.Shuffled)
	is.True(0 < deck.ID)
	is.Equal(3, deck.Remaining)
}

func TestCreateDeck_NoCardIDs(t *testing.T) {
	// given
	is := is.New(t)
	conn, tearDown := pg.NewDbTestSetup(t)
	defer tearDown()

	ctx := context.Background()
	deckRepo := pg.PGDeckRepository{DB: conn}
	cardsRepo := pg.PGCardRepository{DB: conn}

	err := cardsRepo.SeedCards(ctx)
	is.NoErr(err)

	params := repository.CreateDeckParams{
		ExternalID: uuid.New(),
	}
	// when
	deck, err := deckRepo.CreateDeck(ctx, params)

	// then
	is.Equal(err.Error(), "error expect cardIDs length > 0")
	is.Equal(nil, deck)
}

func TestFetchCards(t *testing.T) {
	// given
	is := is.New(t)
	conn, tearDown := pg.NewDbTestSetup(t)
	defer tearDown()

	ctx := context.Background()
	deckRepo := pg.PGDeckRepository{DB: conn}
	cardsRepo := pg.PGCardRepository{DB: conn}

	err := cardsRepo.SeedCards(ctx)
	is.NoErr(err)

	externalID := uuid.New()
	params := repository.CreateDeckParams{
		CardIDs:    []int{1},
		ExternalID: externalID,
	}
	_, err = deckRepo.CreateDeck(ctx, params)
	is.NoErr(err)

	openDeck, err := deckRepo.FetchDeck(ctx, externalID)
	is.NoErr(err)
	is.Equal(1, len(openDeck.Cards))
	is.Equal("ACE", openDeck.Cards[0].Value)
	is.Equal("AS", openDeck.Cards[0].Code)
	is.Equal("SPADES", openDeck.Cards[0].Suit)
	is.Equal(externalID, openDeck.Deck.ExternalID)
}

func TestFetchCards_ExternalIDNotFound(t *testing.T) {
	// given
	is := is.New(t)
	conn, tearDown := pg.NewDbTestSetup(t)
	defer tearDown()

	ctx := context.Background()
	deckRepo := pg.PGDeckRepository{DB: conn}
	cardsRepo := pg.PGCardRepository{DB: conn}

	err := cardsRepo.SeedCards(ctx)
	is.NoErr(err)

	externalID := uuid.New()

	// when
	openDeck, err := deckRepo.FetchDeck(ctx, externalID)

	// then
	is.Equal(err, repository.ErrorDeckNotFound)
	is.Equal(nil, openDeck)
}

func TestDrawCardFromDeck(t *testing.T) {
	// given
	is := is.New(t)
	conn, tearDown := pg.NewDbTestSetup(t)
	defer tearDown()

	ctx := context.Background()
	deckRepo := pg.PGDeckRepository{DB: conn}
	cardsRepo := pg.PGCardRepository{DB: conn}

	err := cardsRepo.SeedCards(ctx)
	is.NoErr(err)

	externalID := uuid.New()
	params := repository.CreateDeckParams{
		CardIDs:    []int{1, 2, 3, 4},
		ExternalID: externalID,
	}
	_, err = deckRepo.CreateDeck(ctx, params)
	is.NoErr(err)

	openDeck, err := deckRepo.FetchDeck(ctx, externalID)
	is.NoErr(err)
	is.Equal(4, len(openDeck.Cards))

	// when
	err = deckRepo.DrawCardFromDeck(ctx, externalID, 3)
	is.NoErr(err)

	// then
	openDeck, err = deckRepo.FetchDeck(ctx, externalID)
	is.NoErr(err)
	is.Equal(1, len(openDeck.Cards))
}

func TestDrawCardFromDeck_WhenDeckDoesNotExists(t *testing.T) {
	// given
	is := is.New(t)
	conn, tearDown := pg.NewDbTestSetup(t)
	defer tearDown()

	ctx := context.Background()
	deckRepo := pg.PGDeckRepository{DB: conn}
	cardsRepo := pg.PGCardRepository{DB: conn}

	err := cardsRepo.SeedCards(ctx)
	is.NoErr(err)

	externalID := uuid.New()

	// when
	err = deckRepo.DrawCardFromDeck(ctx, externalID, 3)
	is.Equal(err, repository.ErrorDeckNotFound)
}
