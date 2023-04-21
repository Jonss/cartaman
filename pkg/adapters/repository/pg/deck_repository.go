package pg

import (
	"context"
	"database/sql"
	"errors"

	"github.com/Jonss/cartaman/pkg/adapters/repository"
	_ "github.com/golang-migrate/migrate/v4/source/file" // migration
	_ "github.com/lib/pq"                                // postgres
)

type PGDeckRepository struct {
	DB *sql.DB
}

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
