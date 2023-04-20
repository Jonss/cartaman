package pg

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
)

type Card struct {
	ID    int
	Value string
	Suit  string
	Code  string
}

type PGCardRepository struct {
	DB *sql.DB
}

func (r PGCardRepository) SeedCards(ctx context.Context) error {
	var quantity int
	err := r.DB.QueryRowContext(ctx, "select count(1) from cards").Scan(&quantity)
	if err != nil {
		return err
	}
	if quantity > 0 {
		return nil
	}

	query := fmt.Sprintf(`INSERT INTO cards (suit, value, code) VALUES %s;`, buildCards())
	_, err = r.DB.ExecContext(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func buildCards() string {
	cardSuits := []string{"CLUBS", "DIAMONDS", "HEARTS", "SPADES"}
	cardValues := []string{"ACE", "2", "3", "4", "5", "6", "7", "8", "9", "10", "JACK", "QUEEN", "KING"}

	var strBuilder strings.Builder
	for _, s := range cardSuits {
		for _, v := range cardValues {
			code := fmt.Sprintf("%c%c", s[0], v[0])
			strBuilder.WriteString(fmt.Sprintf(" ('%s', '%s', '%s'),", s, v, code))
		}
	}
	cards := strBuilder.String()
	return cards[:len(cards)-1]
}
