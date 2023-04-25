package httprest

import (
	"errors"
	"net/http"
	"strings"

	"github.com/Jonss/cartaman/pkg/adapters/repository"
	"github.com/Jonss/cartaman/pkg/usecase/deck"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

var deckNotFoundMessage = "deck not found"
var deckIdIsInvalidMessage = "deck id is invalid"
var unexpectedErrorMessage = "unexpected error"

func (a app) Create(c *fiber.Ctx) error {
	cardCodes := getCardCodes(c.Query("cards", ""))
	shuffled := c.QueryBool("shuffled", false)

	deck, err := a.DeckService.Create(c.UserContext(), deck.CreateParams{
		CardCodes: cardCodes,
		Shuffled:  shuffled,
	})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(newErrorMessage(unexpectedErrorMessage))
	}
	return c.Status(http.StatusCreated).JSON(deck)
}

func getCardCodes(codes string) []string {
	if len(codes) == 0 {
		return []string{}
	}
	return strings.Split(codes, ",")
}

func (a app) Open(c *fiber.Ctx) error {
	deckID, err := getDeckID(c)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(newErrorMessage(deckIdIsInvalidMessage))
	}

	openDeck, err := a.DeckService.Open(c.UserContext(), deckID)
	if err != nil {
		if err == repository.ErrorDeckNotFound {
			return c.Status(http.StatusNotFound).JSON(newErrorMessage(deckNotFoundMessage))
		}
		return c.Status(http.StatusInternalServerError).JSON(newErrorMessage(unexpectedErrorMessage))
	}
	return c.JSON(openDeck)
}

type DrawCardsResponse struct {
	Cards []deck.Card `json:"cards"`
}

func (a app) Draw(c *fiber.Ctx) error {
	deckID, err := getDeckID(c)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(newErrorMessage(deckIdIsInvalidMessage))
	}

	count, err := c.ParamsInt("count", 0)
	if err != nil || (count <= 0) {
		return c.Status(http.StatusBadRequest).JSON(newErrorMessage("count should be above 0"))
	}
	drawDeck, err := a.DeckService.Draw(c.UserContext(), deckID, count)
	if err != nil {
		if err == repository.ErrorDeckNotFound {
			return c.Status(http.StatusNotFound).JSON(newErrorMessage(deckNotFoundMessage))
		}
		return c.Status(http.StatusBadRequest).JSON(newErrorMessage(unexpectedErrorMessage))
	}
	return c.JSON(DrawCardsResponse{drawDeck})
}

func getDeckID(c *fiber.Ctx) (uuid.UUID, error) {
	paramID := c.Params("id")

	deckID, err := uuid.Parse(paramID)
	if err != nil {
		return uuid.Nil, errors.New("error id pattern unexpected")
	}
	return deckID, nil
}
