package main

import (
	"os"

	"log"

	"github.com/dionaditya/victory-scrape/scraper"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/headzoo/surf"
)

func main() {

	bow := surf.NewBrowser()

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	app := fiber.New()

	app.Use(cors.New())

	app.Get("/api/v1/vaccine-venue/available", func(c *fiber.Ctx) error {
		data, err := scraper.GetAvailableVaccineVenue(bow, c.Query("tanggal"))

		if err != nil {
			return c.SendString("Error")
		}

		return c.JSON(data)
	})

	app.Get("/api/v1/vaccine-venue", func(c *fiber.Ctx) error {
		data, err := scraper.GetAllVaccineVenue(bow, c.Query("tanggal"))

		if err != nil {
			return c.SendString("Error")
		}

		return c.JSON(data)
	})

	app.Listen(":" + port)
}
