package httprest_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Jonss/cartaman/pkg/ports/httprest"
	"github.com/Jonss/cartaman/pkg/usecase/decks"
	mock_decks "github.com/Jonss/cartaman/pkg/usecase/decks/mocks"
	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/matryer/is"
)

func TestCreateDeck(t *testing.T) {
	is := is.New(t)
	deckID := uuid.Must(uuid.Parse("c723d533-8612-4cde-bd5e-438c03f6204a"))
	testCases := []struct {
		name           string
		endpoint       string
		buildStubs     func(deckRepo *mock_decks.MockDeckUseCase)
		want           string
		wantStatusCode int
	}{
		{
			name: "should get a valid response",
			buildStubs: func(deckRepo *mock_decks.MockDeckUseCase) {
				deckRepo.EXPECT().Create(context.Background(), gomock.Any()).Times(1).Return(&decks.Deck{
					DeckID:    deckID,
					Shuffled:  false,
					Remaining: 56,
				}, nil)
			},
			want:           `{"deck_id":"c723d533-8612-4cde-bd5e-438c03f6204a","shuffled":false,"remaining":56}`,
			wantStatusCode: http.StatusCreated,
		},
		{
			name:     "should get a shuffled when it's true",
			endpoint: "?shuffled=true",
			buildStubs: func(deckRepo *mock_decks.MockDeckUseCase) {
				deckRepo.EXPECT().Create(context.Background(), gomock.Any()).Times(1).Return(&decks.Deck{
					DeckID:    deckID,
					Shuffled:  true,
					Remaining: 56,
				}, nil)
			},
			want:           `{"deck_id":"c723d533-8612-4cde-bd5e-438c03f6204a","shuffled":true,"remaining":56}`,
			wantStatusCode: http.StatusCreated,
		},
		{
			name:     "should get 3 remaining cards as response",
			endpoint: "?cards=AS,KD,AC",
			buildStubs: func(deckRepo *mock_decks.MockDeckUseCase) {
				params := decks.CreateParams{CardCodes: []string{"AS", "KD", "AC"}, Shuffled: false}
				deckRepo.EXPECT().Create(context.Background(), params).Times(1).Return(&decks.Deck{
					DeckID:    deckID,
					Shuffled:  false,
					Remaining: 3,
				}, nil)
			},
			want:           `{"deck_id":"c723d533-8612-4cde-bd5e-438c03f6204a","shuffled":false,"remaining":3}`,
			wantStatusCode: http.StatusCreated,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// given
			ctrl := gomock.NewController(t)
			deckUseCase := mock_decks.NewMockDeckUseCase(ctrl)
			tc.buildStubs(deckUseCase)

			app := httprest.App{
				fiber.New(),
				deckUseCase,
			}
			app.Routes()

			// when
			r := httptest.NewRequest(http.MethodPost, "/decks"+tc.endpoint, nil)
			res, err := app.FiberApp.Test(r, 1000)
			is.NoErr(err)

			got := make([]byte, res.ContentLength)
			_, _ = res.Body.Read(got)

			// then
			is.Equal(tc.wantStatusCode, res.StatusCode)
			is.Equal(tc.want, string(got))
		})
	}
}
