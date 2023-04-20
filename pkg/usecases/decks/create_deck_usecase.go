package decks

func (r *deckUseCase) Create() (*Deck, error) {
	deck, err := r.DeckRepository.CreateDeck()
	if err != nil {
		return nil, err
	}
	return &Deck{
		DeckID:    deck.ExternalID,
		Shuffled:  deck.Shuffled,
		Remaining: deck.Remaining,
	}, nil
}
