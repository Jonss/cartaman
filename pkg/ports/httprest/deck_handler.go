package httprest

import (
	"net/http"
	"strings"

	"github.com/Jonss/cartaman/pkg/usecase/decks"
	"github.com/gofiber/fiber/v2"
)

func (a App) Create(c *fiber.Ctx) error {
	cardCodes := strings.Split(c.Query("cards", ""), ",")
	shuffled := c.QueryBool("shuffled", false)

	deck, err := a.DeckUseCase.Create(c.UserContext(), decks.CreateParams{
		CardCodes: cardCodes,
		Shuffled:  shuffled,
	})
	if err != nil {
		c.Status(http.StatusInternalServerError).SendString("error")
		return err
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
