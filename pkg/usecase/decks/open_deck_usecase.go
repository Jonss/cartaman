package decks

import (
	"context"

	"github.com/Jonss/cartaman/pkg/adapters/repository"
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
		Cards:     buildCards(openDeck.Cards),
	}, nil
}

func buildCards(repoCards []repository.Card) []Card {
	cards := make([]Card, len(repoCards))
	for i, repoCard := range repoCards {
		cards[i] = Card{
			Value: repoCard.Value,
			Suit:  repoCard.Suit,
			Code:  repoCard.Code,
		}
	}
	return cards
}
