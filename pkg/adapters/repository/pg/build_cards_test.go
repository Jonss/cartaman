package pg

import (
	"testing"

	"github.com/matryer/is"
)

func TestBuildCards(t *testing.T) {
	is := is.New(t)
	got := buildCards()
	want := " ('SPADES', 'ACE', 'AS'), ('SPADES', '2', '2S'), ('SPADES', '3', '3S'), ('SPADES', '4', '4S'), ('SPADES', '5', '5S'), ('SPADES', '6', '6S'), ('SPADES', '7', '7S'), ('SPADES', '8', '8S'), ('SPADES', '9', '9S'), ('SPADES', '10', '1S'), ('SPADES', 'JACK', 'JS'), ('SPADES', 'QUEEN', 'QS'), ('SPADES', 'KING', 'KS'), ('CLUBS', 'ACE', 'AC'), ('CLUBS', '2', '2C'), ('CLUBS', '3', '3C'), ('CLUBS', '4', '4C'), ('CLUBS', '5', '5C'), ('CLUBS', '6', '6C'), ('CLUBS', '7', '7C'), ('CLUBS', '8', '8C'), ('CLUBS', '9', '9C'), ('CLUBS', '10', '1C'), ('CLUBS', 'JACK', 'JC'), ('CLUBS', 'QUEEN', 'QC'), ('CLUBS', 'KING', 'KC'), ('DIAMONDS', 'ACE', 'AD'), ('DIAMONDS', '2', '2D'), ('DIAMONDS', '3', '3D'), ('DIAMONDS', '4', '4D'), ('DIAMONDS', '5', '5D'), ('DIAMONDS', '6', '6D'), ('DIAMONDS', '7', '7D'), ('DIAMONDS', '8', '8D'), ('DIAMONDS', '9', '9D'), ('DIAMONDS', '10', '1D'), ('DIAMONDS', 'JACK', 'JD'), ('DIAMONDS', 'QUEEN', 'QD'), ('DIAMONDS', 'KING', 'KD'), ('HEARTS', 'ACE', 'AH'), ('HEARTS', '2', '2H'), ('HEARTS', '3', '3H'), ('HEARTS', '4', '4H'), ('HEARTS', '5', '5H'), ('HEARTS', '6', '6H'), ('HEARTS', '7', '7H'), ('HEARTS', '8', '8H'), ('HEARTS', '9', '9H'), ('HEARTS', '10', '1H'), ('HEARTS', 'JACK', 'JH'), ('HEARTS', 'QUEEN', 'QH'), ('HEARTS', 'KING', 'KH')"

	is.Equal(want, got)
}
