package httprest_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Jonss/cartaman/pkg/ports/httprest"
	"github.com/gofiber/fiber/v2"
)

func TestCreateDeck(t *testing.T) {

	testCases := []struct {
		name string
	}{
		{
			name: "should get dummy message",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			app := httprest.App{
				fiber.New(),
			}
			app.Routes()

			r := httptest.NewRequest(http.MethodPost, "/decks", nil)

			res, err := app.FiberApp.Test(r, 1000)
			if err != nil {
				t.Fatalf("unexpected error %q", err)
			}

			// TODO: add lib to assertions, 'matryer/is' maybe
			if res.StatusCode != 200 {
				t.Fatalf("want 200, got %d", res.StatusCode)
			}
		})
	}
}
