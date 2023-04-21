package decks

import (
	"context"

	"github.com/Jonss/cartaman/pkg/adapters/repository"
	"github.com/google/uuid"
)

type CreateParams struct {
	CardCodes []string
	Shuffled  bool
}

func (r *deckUseCase) Create(ctx context.Context, params CreateParams) (*Deck, error) {
	cardIDs, err := r.CardRepository.GetCardIDs(ctx, params.CardCodes)
	if err != nil {
		return nil, err
	}

	externalID := uuid.New()
	deck, err := r.DeckRepository.CreateDeck(ctx, repository.CreateDeckParams{
		ExternalID: externalID,
		CardIDs:    cardIDs,
		Shuffled:   params.Shuffled,
	})
	if err != nil {
		return nil, err
	}
	return &Deck{
		DeckID:    deck.ExternalID,
		Shuffled:  deck.Shuffled,
		Remaining: deck.Remaining,
	}, nil
}
