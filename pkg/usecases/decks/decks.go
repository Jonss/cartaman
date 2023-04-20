package decks

import (
	"github.com/Jonss/cartaman/pkg/repository"
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
	Create() (*Deck, error)
}

var _ DeckUseCase = (*deckUseCase)(nil)
