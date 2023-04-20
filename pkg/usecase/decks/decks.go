package decks

import (
	"context"

	"github.com/Jonss/cartaman/pkg/adapters/repository"
	"github.com/google/uuid"
)

type Deck struct {
	DeckID    uuid.UUID `json:"deck_id"`
	Shuffled  bool      `json:"shuffled"`
	Remaining int       `json:"remaining"`
}

type deckUseCase struct {
	DeckRepository repository.DeckRepository
}

type DeckUseCase interface {
	Create(context.Context) (*Deck, error)
}

var _ DeckUseCase = (*deckUseCase)(nil)
