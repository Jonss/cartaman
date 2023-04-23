package httprest_test

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Jonss/cartaman/pkg/adapters/repository"
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
		buildStubs     func(deckRepo *mock_decks.MockDeckService)
		want           string
		wantStatusCode int
	}{
		{
			name: "should get a valid response",
			buildStubs: func(deckRepo *mock_decks.MockDeckService) {
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
			buildStubs: func(deckRepo *mock_decks.MockDeckService) {
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
			buildStubs: func(deckRepo *mock_decks.MockDeckService) {
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
			deckUseCase := mock_decks.NewMockDeckService(ctrl)
			tc.buildStubs(deckUseCase)

			app := httprest.NewApp(fiber.New(), deckUseCase)
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

func TestOpenDeck(t *testing.T) {
	is := is.New(t)
	testCases := []struct {
		name           string
		deckID         string
		buildStubs     func(deckRepo *mock_decks.MockDeckService)
		want           string
		wantStatusCode int
	}{
		{
			name:   "should get deck when deckID is valid",
			deckID: "c723d533-8612-4cde-bd5e-438c03f6204a",
			buildStubs: func(deckRepo *mock_decks.MockDeckService) {
				deckID := "c723d533-8612-4cde-bd5e-438c03f6204a"
				deckRepo.EXPECT().Open(context.Background(), gomock.Any()).
					Times(1).
					Return(&decks.OpenDeck{
						DeckID:    uuid.MustParse(deckID),
						Shuffled:  false,
						Remaining: 1,
						Cards: []decks.Card{
							{
								Value: "KING",
								Suit:  "HEARTS",
								Code:  "KH",
							},
						},
					}, nil)
			},
			want:           `{"deck_id":"c723d533-8612-4cde-bd5e-438c03f6204a","shuffled":false,"remaining":1,"cards":[{"value":"KING","suit":"HEARTS","code":"KH"}]}`,
			wantStatusCode: http.StatusOK,
		},
		{
			name:           "should get status code 400 when deckID is invalid",
			deckID:         "invalid-deck-id",
			buildStubs:     func(deckRepo *mock_decks.MockDeckService) {},
			want:           `error id pattern unexpected`,
			wantStatusCode: http.StatusBadRequest,
		},
		{
			name:   "should get status code 404 when deck is not found",
			deckID: "c723d533-8612-4cde-bd5e-438c03f6204a",
			buildStubs: func(deckRepo *mock_decks.MockDeckService) {
				deckRepo.EXPECT().Open(context.Background(), gomock.Any()).
					Times(1).
					Return(&decks.OpenDeck{}, repository.ErrorDeckNotFound)
			},
			want:           `deck not found`,
			wantStatusCode: http.StatusNotFound,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// given
			ctrl := gomock.NewController(t)
			deckUseCase := mock_decks.NewMockDeckService(ctrl)
			tc.buildStubs(deckUseCase)

			app := httprest.NewApp(fiber.New(), deckUseCase)
			app.Routes()

			// when
			r := httptest.NewRequest(http.MethodGet, "/decks/"+tc.deckID, nil)
			res, err := app.FiberApp.Test(r, 1000)
			is.NoErr(err)

			got := make([]byte, res.ContentLength)
			_, _ = res.Body.Read(got)

			// then
			// is.Equal(tc.wantStatusCode, res.StatusCode)
			is.Equal(tc.want, string(got))
		})
	}
}

func TestDrawDeck(t *testing.T) {
	is := is.New(t)
	testCases := []struct {
		name           string
		deckID         string
		count          string
		buildStubs     func(deckRepo *mock_decks.MockDeckService)
		want           string
		wantStatusCode int
	}{
		{
			name:   "should draw 1 card",
			deckID: "c723d533-8612-4cde-bd5e-438c03f6204a",
			count:  "1",
			buildStubs: func(deckRepo *mock_decks.MockDeckService) {
				deckRepo.EXPECT().Draw(context.Background(), gomock.Any(), gomock.Any()).
					Times(1).
					Return([]decks.Card{
						{
							Suit:  "SPADES",
							Value: "ACE",
							Code:  "AS",
						},
					}, nil)
			},
			want:           `{"cards":[{"value":"ACE","suit":"SPADES","code":"AS"}]}`,
			wantStatusCode: http.StatusOK,
		},
		{
			name:   "should get error when draw number is 0",
			deckID: "c723d533-8612-4cde-bd5e-438c03f6204a",
			count:  "0",
			buildStubs: func(deckRepo *mock_decks.MockDeckService) {
			},
			want:           `count should be above 0`,
			wantStatusCode: http.StatusBadRequest,
		},

		{
			name:   "should get error when draw number is invalid",
			deckID: "c723d533-8612-4cde-bd5e-438c03f6204a",
			count:  "-90",
			buildStubs: func(deckRepo *mock_decks.MockDeckService) {
			},
			want:           `count should be above 0`,
			wantStatusCode: http.StatusBadRequest,
		},
		{
			name:   "should get error when draw number is not a number",
			deckID: "c723d533-8612-4cde-bd5e-438c03f6204a",
			count:  "invalid-number",
			buildStubs: func(deckRepo *mock_decks.MockDeckService) {
			},
			want:           `count should be above 0`,
			wantStatusCode: http.StatusBadRequest,
		},
		{
			name:   "should get error when deck was not found",
			deckID: "c723d533-8612-4cde-bd5e-438c03f6204a",
			count:  "1",
			buildStubs: func(deckRepo *mock_decks.MockDeckService) {
				deckRepo.EXPECT().Draw(context.Background(), gomock.Any(), gomock.Any()).
					Times(1).
					Return([]decks.Card{}, repository.ErrorDeckNotFound)
			},
			want:           `deck not found`,
			wantStatusCode: http.StatusNotFound,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// given
			ctrl := gomock.NewController(t)
			deckUseCase := mock_decks.NewMockDeckService(ctrl)
			tc.buildStubs(deckUseCase)

			app := httprest.NewApp(fiber.New(), deckUseCase)
			app.Routes()

			// when
			r := httptest.NewRequest(http.MethodPatch, fmt.Sprintf("/decks/%s/draw/%s", tc.deckID, tc.count), nil)
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
