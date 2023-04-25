package httprest

import (
	"github.com/Jonss/cartaman/pkg/usecase/deck"
	"github.com/gofiber/fiber/v2"
)

type app struct {
	FiberApp    *fiber.App
	DeckService deck.DeckService
}

func NewApp(fiberApp *fiber.App, deckService deck.DeckService) app {
	return app{FiberApp: fiberApp, DeckService: deckService}
}

func (a app) Routes() {
	a.FiberApp.Post("/decks", func(c *fiber.Ctx) error {
		return a.Create(c)
	})

	a.FiberApp.Get("/decks/:id", func(c *fiber.Ctx) error {
		return a.Open(c)
	})

	a.FiberApp.Patch("/decks/:id/draw/:count", func(c *fiber.Ctx) error {
		return a.Draw(c)
	})
}

type ErrorResponse struct {
	Message string `json:"message"`
}

func newErrorMessage(message string) ErrorResponse {
	return ErrorResponse{Message: message}
}
