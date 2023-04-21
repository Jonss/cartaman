package pg

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
)

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

func (r PGCardRepository) GetCardIDs(ctx context.Context, codes []string) ([]int, error) {
	query := "SELECT id FROM cards"
	if len(codes) > 0 {
		query += fmt.Sprintf(" WHERE code IN ('%s');", strings.Join(codes, "','"))
	}

	var IDs []int
	rows, err := r.DB.QueryContext(ctx, query)
	if err != nil {
		return IDs, err
	}

	for rows.Next() {

		var id int
		err := rows.Scan(&id)
		if err != nil {
			return IDs, err
		}
		IDs = append(IDs, id)
	}

	return IDs, nil
}

func buildCards() string {
	cardSuits := []string{"CLUBS", "DIAMONDS", "HEARTS", "SPADES"}
	cardValues := []string{"ACE", "2", "3", "4", "5", "6", "7", "8", "9", "10", "JACK", "QUEEN", "KING"}

	var strBuilder strings.Builder
	for _, s := range cardSuits {
		for _, v := range cardValues {
			code := fmt.Sprintf("%c%c", v[0], s[0])
			strBuilder.WriteString(fmt.Sprintf(" ('%s', '%s', '%s'),", s, v, code))
		}
	}
	cards := strBuilder.String()
	return cards[:len(cards)-1]
}
