package repository

import "github.com/google/uuid"

type Deck struct {
	ID         int
	ExternalID uuid.UUID
	Shuffled   bool
	Remaining  int
}

type DeckRepository interface {
	CreateDeck() (*Deck, error)
}
