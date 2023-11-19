package main

import (

	"github.com/gofiber/fiber/v2"
)

type Resp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		resp := Resp{
			Code:    200,
			Message: "Hello, World 👋!",
		}

		return c.JSON(resp)
	})

	app.Listen(":3000")
}
