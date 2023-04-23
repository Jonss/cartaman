package deck

import "github.com/Jonss/cartaman/pkg/adapters/repository"

func mapCards(repoCards []repository.Card) []Card {
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
