package main

import (
	"encoding/json"
	"fcompressor/app"
	_ "fcompressor/app/api"
	"fcompressor/db"
	_ "fcompressor/env"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	db.Open()

	fiber := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	fiber.Use(logger.New())
	fiber.Use(cors.New())

	api := app.New(fiber)
	api.Start()
}
