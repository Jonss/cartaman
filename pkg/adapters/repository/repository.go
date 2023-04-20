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

type DeckRepository interface {
	CreateDeck(context.Context) (*Deck, error)
}

type CardRepository interface {
	SeedCards(context.Context) error
}
