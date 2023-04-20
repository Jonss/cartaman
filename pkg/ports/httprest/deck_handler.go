package httprest

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (a App) Create(c *fiber.Ctx) error {
	// TODO: add validation
	deck, err := a.DeckUseCase.Create(c.UserContext())
	if err != nil {
		c.Status(http.StatusInternalServerError).SendString("error")
	}
	return c.Status(http.StatusCreated).JSON(deck)
}

// TODO
func (a App) Open(c *fiber.Ctx) error {
	return c.SendString("TODO: open deck")
}

// TODO
func (a App) Draw(c *fiber.Ctx) error {
	return c.SendString("TODO: draw card")
}
