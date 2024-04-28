package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	if err := app.Listen(":4001"); err != nil {
		if err = fmt.Errorf("error: %v", err); err != nil {
			panic(err)
		}
	}

}