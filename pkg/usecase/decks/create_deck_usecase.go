package decks

import "context"

func (r *deckUseCase) Create(ctx context.Context) (*Deck, error) {
	deck, err := r.DeckRepository.CreateDeck(ctx)
	if err != nil {
		return nil, err
	}
	return &Deck{
		DeckID:    deck.ExternalID,
		Shuffled:  deck.Shuffled,
		Remaining: deck.Remaining,
	}, nil
}
