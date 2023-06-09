package repository

import (
	"context"
	"errors"

	"github.com/google/uuid"
)

type Deck struct {
	ID         int
	ExternalID uuid.UUID
	Shuffled   bool
	Remaining  int
}

type Card struct {
	ID    int
	Value string
	Suit  string
	Code  string
}

type OpenDeck struct {
	Deck  Deck
	Cards []Card
}

type CreateDeckParams struct {
	CardIDs    []int
	Shuffled   bool
	ExternalID uuid.UUID
}

type DeckRepository interface {
	CreateDeck(context.Context, CreateDeckParams) (*Deck, error)
	FetchDeck(context.Context, uuid.UUID) (*OpenDeck, error)
	DrawCardFromDeck(context.Context, uuid.UUID, int) error
	FetchDrewCards(ctx context.Context, deckID uuid.UUID) ([]Card, error)
}

type CardRepository interface {
	SeedCards(context.Context) error
	GetCardIDs(context.Context, []string) ([]int, error)
}

var ErrorDeckNotFound = errors.New("error deck not found")
var ErrorCardIDsInvalid = errors.New("error expect cardIDs length > 0")
