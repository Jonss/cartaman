package repository

import (
	"context"

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
	Suite string
	Code  string
}

type CreateDeckParams struct {
	CardIDs    []int
	Shuffled   bool
	ExternalID uuid.UUID
}

type DeckRepository interface {
	CreateDeck(context.Context, CreateDeckParams) (*Deck, error)
}

type CardRepository interface {
	SeedCards(context.Context) error
	GetCardIDs(context.Context, []string) ([]int, error)
}
