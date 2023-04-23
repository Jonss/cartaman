package httprest

import (
	"github.com/Jonss/cartaman/pkg/usecase/decks"
	"github.com/gofiber/fiber/v2"
)

type app struct {
	FiberApp    *fiber.App
	DeckService decks.DeckService
}

func NewApp(fiberApp *fiber.App, deckService decks.DeckService) app {
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
