package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/goccy/go-json"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello Go lang")
	}) 

	// JSON stucture 
	type JsonData struct {
		word string
	}

	app.Get("/json", func(c *fiber.Ctx) error{
		// create data struct
		data := JsonData {
			word: "Fiber",
		}

		return c.JSON(data)
		// => Content-Type: application/json
  		// => "{"word": "Fiber"}"


	})

	// GET http://127.0.0.1:3000/test
	app.Get("/:value", func(c *fiber.Ctx) error{
		return c.SendString("value: " + c.Params("value"))
	})


	app.Listen(":3000")
}