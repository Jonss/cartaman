package httprest

import "github.com/gofiber/fiber/v2"

func (a App) Create(c *fiber.Ctx) error {
	return c.SendString("TODO: create deck")
}

func (a App) Open(c *fiber.Ctx) error {
	return c.SendString("TODO: open deck")
}

func (a App) Draw(c *fiber.Ctx) error {
	return c.SendString("TODO: draw card")
}
