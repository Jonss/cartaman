package deck

import (
	"context"

	"github.com/google/uuid"
)

func (r *deckService) Draw(ctx context.Context, deckID uuid.UUID, count int) ([]Card, error) {
	err := r.DeckRepository.DrawCardFromDeck(ctx, deckID, count)
	if err != nil {
		return []Card{}, err
	}

	cards, err := r.DeckRepository.FetchDrewCards(ctx, deckID)
	if err != nil {
		return []Card{}, err
	}

	return mapCards(cards), nil
}
