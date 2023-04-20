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

func (r PGDeckRepository) CreateDeck(ctx context.Context, cardIDs []int) (*repository.Deck, error) {
	if len(cardIDs) == 0 {
		return nil, errors.New("error expect cardIDs length > 0")
	}

	tx, err := r.DB.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	externalID := uuid.New()
	query := `INSERT INTO decks (external_id, is_shuffle) values ($1, $2) RETURNING id`

	var deckID int
	err = tx.QueryRowContext(ctx, query, externalID, false).Scan(&deckID)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	for _, cardID := range cardIDs {
		query = `INSERT INTO decks_cards (card_id, deck_id) VALUES ($1, $2)`
		_, err := tx.ExecContext(ctx, query, cardID, deckID)
		if err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	deck := &repository.Deck{
		ID:         deckID,
		ExternalID: externalID,
		Shuffled:   false, //hardcoded, change it
		Remaining:  len(cardIDs),
	}
	return deck, nil
}
