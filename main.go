package main

import (
	"fmt"
	"math"
	"os"
	"strings"

	"log"

	"github.com/dionaditya/victory-scrape/scraper"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/headzoo/surf"

	"net/http"
	"net/url"
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
		data, _, err := scraper.GetAvailableVaccineVenue(bow, c.Query("tanggal"))

		if err != nil {
			return c.SendString("Error")
		}

		return c.JSON(data)
	})

	app.Get("/api/v1/vaccine-scrape", func(c *fiber.Ctx) error {
		data, err := scraper.GetVaccinationDate(bow)

		if err != nil {
			return c.SendString("Error")
		}

		for _, date := range data {
			data, columns, err := scraper.GetAllVaccineVenue(bow, date)

			if err != nil {
				return c.SendString("Error")
			}

			var client = &http.Client{}

			botToken := os.Getenv("BOT_TOKEN")

			channelAccount := os.Getenv("CHANNEL_ACCOUNT")

			botUrl := "https://api.telegram.org/bot" + botToken + "/sendMessage?chat_id=@" + channelAccount + "&parse_mode=Markdown&text="

			for batch := 0; float64(batch) < math.Ceil(float64(len(data))/float64(2)); batch++ {
				rawMessage := "Data vaksinasi tanggal " + "*" + date + "*" + "\n" + "\n" + ""

				for i, s := range data {

					if i < 2*(batch+1) && i >= (2*batch) {
						columnIndex := 0

						for _, column := range columns {

							checkedEmoji := ""

							if column == "Aksi" && (strings.Contains(s[column].(string), "Mendaftar") || strings.Contains(s[column].(string), "Ambil Kupon")) {
								checkedEmoji = " \U00002714"
							}

							if column == "Aksi" && (strings.Contains(s[column].(string), "Kuota Telah Terpenuhi")) {
								checkedEmoji = " 	\U0000274c"
							}

							updatedMessage := rawMessage + "*" + column + "*" + ": " + fmt.Sprintf("%v", s[column]) + checkedEmoji + "\n"

							if len(s)-1 == columnIndex {
								rawMessage = updatedMessage + "\n"
							} else {
								rawMessage = updatedMessage
							}

							columnIndex++

						}
					}

				}

				message := url.QueryEscape(rawMessage)

				request, err := http.NewRequest("POST", botUrl+message, nil)

				if err != nil {
					log.Fatal(err)
				}

				response, err := client.Do(request)

				if err != nil {
					log.Fatal(err)
				}

				defer response.Body.Close()
			}
		}

		return c.JSON(data)
	})

	app.Get("/api/v1/vaccine-venue", func(c *fiber.Ctx) error {
		data, _, err := scraper.GetAllVaccineVenue(bow, c.Query("tanggal"))

		if err != nil {
			return c.SendString("Error")
		}

		return c.JSON(data)
	})

	app.Get("/api/v1/vaccine-venue/bot", func(c *fiber.Ctx) error {

		tanggalVaksinasi := c.Query("tanggal")

		data, columns, err := scraper.GetAllVaccineVenue(bow, tanggalVaksinasi)

		if err != nil {
			return c.SendString("Error")
		}

		var client = &http.Client{}

		botToken := os.Getenv("BOT_TOKEN")

		channelAccount := os.Getenv("CHANNEL_ACCOUNT")

		botUrl := "https://api.telegram.org/bot" + botToken + "/sendMessage?chat_id=@" + channelAccount + "&parse_mode=Markdown&text="

		for batch := 0; float64(batch) < math.Ceil(float64(len(data))/float64(2)); batch++ {
			rawMessage := "Data vaksinasi tanggal " + "*" + tanggalVaksinasi + "*" + "\n" + "\n" + ""

			for i, s := range data {

				if i < 2*(batch+1) && i >= (2*batch) {
					columnIndex := 0

					for _, column := range columns {

						checkedEmoji := ""

						if column == "Aksi" && (s[column] == "Mendaftar" || s[column] == "Ambil Kupon") {
							checkedEmoji = " \U00002714"
						}

						if column == "Aksi" && s[column] == "Kuota Telah Terpenuhi" {
							checkedEmoji = " 	\U0000274c"
						}

						updatedMessage := rawMessage + "*" + column + "*" + ": " + fmt.Sprintf("%v", s[column]) + checkedEmoji + "\n"

						if len(s)-1 == columnIndex {
							rawMessage = updatedMessage + "\n"
						} else {
							rawMessage = updatedMessage
						}

						columnIndex++

					}
				}

			}

			message := url.QueryEscape(rawMessage)

			request, err := http.NewRequest("POST", botUrl+message, nil)

			if err != nil {
				log.Fatal(err)
			}

			response, err := client.Do(request)

			if err != nil {
				log.Fatal(err)
			}

			defer response.Body.Close()
		}

		return c.JSON(data)

	})

	app.Listen(":" + port)
}
