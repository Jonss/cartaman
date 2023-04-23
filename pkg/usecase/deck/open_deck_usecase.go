package deck

import (
	"context"

	"github.com/google/uuid"
)

func (r *deckService) Open(ctx context.Context, deckID uuid.UUID) (*OpenDeck, error) {
	openDeck, err := r.DeckRepository.FetchDeck(ctx, deckID)
	if err != nil {
		return nil, err
	}
	return &OpenDeck{
		DeckID:    openDeck.Deck.ExternalID,
		Remaining: openDeck.Deck.Remaining,
		Cards:     mapCards(openDeck.Cards),
	}, nil
}
