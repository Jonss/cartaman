package httprest

import (
	"net/http"
	"strings"

	"github.com/Jonss/cartaman/pkg/adapters/repository"
	"github.com/Jonss/cartaman/pkg/usecase/decks"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
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

func (a App) Open(c *fiber.Ctx) error {
	paramID := c.Params("id")

	deckID, err := uuid.Parse(paramID)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString("error id pattern unexpected")
	}

	openDeck, err := a.DeckUseCase.Open(c.UserContext(), deckID)
	if err != nil {
		if err == repository.ErrorDeckNotFound {
			return c.Status(http.StatusNotFound).SendString("deck not found")
		}
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}
	return c.JSON(openDeck)
}

// TODO
func (a App) Draw(c *fiber.Ctx) error {
	return c.SendString("TODO: draw card")
}
