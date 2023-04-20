package httprest_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Jonss/cartaman/pkg/ports/httprest"
	"github.com/Jonss/cartaman/pkg/usecases/decks"
	mock_decks "github.com/Jonss/cartaman/pkg/usecases/decks/mocks"
	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/matryer/is"
)

func TestCreateDeck(t *testing.T) {
	is := is.New(t)
	testCases := []struct {
		name           string
		buildStubs     func(deckRepo *mock_decks.MockDeckUseCase)
		want           string
		wantStatusCode int
	}{
		{
			name: "should get dummy message",
			buildStubs: func(deckRepo *mock_decks.MockDeckUseCase) {
				deckID := uuid.Must(uuid.Parse("c723d533-8612-4cde-bd5e-438c03f6204a"))
				deckRepo.EXPECT().Create().Times(1).Return(&decks.Deck{
					DeckID:    deckID,
					Shuffled:  false,
					Remaining: 56,
				}, nil)
			},
			want:           `{"deck_id":"c723d533-8612-4cde-bd5e-438c03f6204a","shuffled":false,"remaining":56}`,
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
			r := httptest.NewRequest(http.MethodPost, "/decks", nil)
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
