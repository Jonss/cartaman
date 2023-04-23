package deck

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

type OpenDeck struct {
	DeckID    uuid.UUID `json:"deck_id"`
	Shuffled  bool      `json:"shuffled"`
	Remaining int       `json:"remaining"`
	Cards     []Card    `json:"cards"`
}

type deckService struct {
	DeckRepository repository.DeckRepository
	CardRepository repository.CardRepository
}

type DeckService interface {
	Create(context.Context, CreateParams) (*Deck, error)
	Open(context.Context, uuid.UUID) (*OpenDeck, error)
	Draw(context.Context, uuid.UUID, int) ([]Card, error)
}

func NewDeckService(
	deckRepository repository.DeckRepository,
	cardRepository repository.CardRepository) deckService {
	return deckService{
		DeckRepository: deckRepository,
		CardRepository: cardRepository,
	}
}

var _ DeckService = (*deckService)(nil)
