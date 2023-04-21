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

type Card struct {
	Value string `json:"value"`
	Suit  string `json:"suit"`
	Code  string `json:"code"`
}

type OpenCard struct {
	Deck
	Card
}

type deckUseCase struct {
	DeckRepository repository.DeckRepository
	CardRepository repository.CardRepository
}

type DeckUseCase interface {
	Create(context.Context, CreateParams) (*Deck, error)
	Open(context.Context, uuid.UUID) (*OpenCard, error)
}

var _ DeckUseCase = (*deckUseCase)(nil)
