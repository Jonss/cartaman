package decks

import (
	"context"

	"github.com/google/uuid"
)

func (r *deckUseCase) Open(ctx context.Context, deckID uuid.UUID) (*OpenCard, error) {
	// r.DeckRepository.FetchDeck(ctx, deckID)
	return nil, nil
}
