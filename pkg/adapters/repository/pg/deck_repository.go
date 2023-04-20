package pg

import (
	"context"
	"database/sql"

	"github.com/Jonss/cartaman/pkg/adapters/repository"
	_ "github.com/golang-migrate/migrate/v4/source/file" // migration
	_ "github.com/lib/pq"                                // postgres
)

type PGDeckRepository struct {
	db *sql.DB
}

func (r PGDeckRepository) CreateDeck(ctx context.Context) (*repository.Deck, error) {
	return nil, nil
}
