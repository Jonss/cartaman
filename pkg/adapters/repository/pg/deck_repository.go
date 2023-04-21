package pg

import (
	"context"
	"database/sql"
	"errors"

	"github.com/Jonss/cartaman/pkg/adapters/repository"
	_ "github.com/golang-migrate/migrate/v4/source/file" // migration
	"github.com/google/uuid"
	_ "github.com/lib/pq" // postgres
)

type PGDeckRepository struct {
	DB *sql.DB
}

var _ repository.DeckRepository = (*PGDeckRepository)(nil)

func (r PGDeckRepository) CreateDeck(ctx context.Context, params repository.CreateDeckParams) (*repository.Deck, error) {
	if len(params.CardIDs) == 0 {
		return nil, errors.New("error expect cardIDs length > 0")
	}

	tx, err := r.DB.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	query := `INSERT INTO decks (external_id, is_shuffled) values ($1, $2) RETURNING id`

	var deckID int
	err = tx.QueryRowContext(ctx, query, params.ExternalID, params.Shuffled).Scan(&deckID)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	for _, cardID := range params.CardIDs {
		query = `INSERT INTO decks_cards (card_id, deck_id) VALUES ($1, $2)`
		_, err := tx.ExecContext(ctx, query, cardID, deckID)
		if err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	deck := &repository.Deck{
		ID:         deckID,
		ExternalID: params.ExternalID,
		Shuffled:   params.Shuffled,
		Remaining:  len(params.CardIDs),
	}
	return deck, nil
}

func (r PGDeckRepository) FetchDeck(ctx context.Context, deckID uuid.UUID) (*repository.OpenDeck, error) {
	query := `
	SELECT d.id, d.external_id, c.suit, c.value, c.code
	FROM decks d
	INNER JOIN decks_cards dc
	ON d.id = dc.deck_id
	INNER JOIN cards c
	ON c.id = dc.card_id
	WHERE d.external_id = $1
	AND dc.is_drew IS false
	`
	rows, err := r.DB.QueryContext(ctx, query, deckID)
	if err != nil {
		return nil, err
	}

	var openDeck repository.OpenDeck
	var cards []repository.Card
	for rows.Next() {
		var card repository.Card
		err := rows.Scan(&openDeck.Deck.ID, &openDeck.Deck.ExternalID, &card.Suit, &card.Value, &card.Code)
		if err != nil {
			return nil, err
		}
		cards = append(cards, card)
	}
	if openDeck.Deck.ID == 0 {
		return nil, repository.ErrorDeckNotFound
	}

	openDeck.Cards = cards
	return &openDeck, nil
}
