package scraper

import (
	"fmt"
	"strconv"

	"log"
	"regexp"

	"github.com/PuerkitoBio/goquery"
	"github.com/headzoo/surf/browser"
)

var VICTORI_URL = "http://victori.semarangkota.go.id/info?tanggal="

func GetAvailableVaccineVenue(bow *browser.Browser, date string) (interface{}, error) {
	var url string

	if len(date) > 0 {
		url = VICTORI_URL + date
	} else {
		url = VICTORI_URL
	}

	err := bow.Open(url)

	columns := []string{}

	bow.Find("tr:contains('Aksi')").First().Children().Each(func(i int, s *goquery.Selection) {
		columns = append(columns, s.Text())
	})

	vaccinePlacesData := []interface{}{}

	bow.Find("tr:contains('Ambil Kupon')").Each(func(dataIndex int, rowData *goquery.Selection) {
		vaccinePlaceData := make(map[string]interface{})

		rowData.Children().Each(func(columnIndex int, s *goquery.Selection) {

			if columns[columnIndex] == "Kuota" || columns[columnIndex] == "Terisi" {

				reg, err := regexp.Compile("[^a-zA-Z0-9]+")

				if err != nil {
					log.Fatal(err)
				}

				processedString := reg.ReplaceAllString(s.Text(), "")

				value, err := strconv.Atoi(processedString)

				if err != nil {
					vaccinePlaceData[columns[columnIndex]] = s.Text()
				}

				vaccinePlaceData[columns[columnIndex]] = value
			} else {
				vaccinePlaceData[columns[columnIndex]] = s.Text()
			}

		})

		vaccinePlacesData = append(vaccinePlacesData, vaccinePlaceData)
	})

	bow.Find("tr:contains('Mendaftar')").Each(func(dataIndex int, rowData *goquery.Selection) {
		vaccinePlaceData := make(map[string]interface{})

		rowData.Children().Each(func(columnIndex int, columnData *goquery.Selection) {
			if columns[columnIndex] == "Kuota" || columns[columnIndex] == "Terisi" {

				reg, err := regexp.Compile("[^a-zA-Z0-9]+")

				if err != nil {
					log.Fatal(err)
				}

				processedString := reg.ReplaceAllString(columnData.Text(), "")

				value, err := strconv.Atoi(processedString)

				if err != nil {
					fmt.Println("hello world")
					vaccinePlaceData[columns[columnIndex]] = columnData.Text()
				}

				vaccinePlaceData[columns[columnIndex]] = value

			} else {
				vaccinePlaceData[columns[columnIndex]] = columnData.Text()
			}

		})

		vaccinePlacesData = append(vaccinePlacesData, vaccinePlaceData)
	})

	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	}

	return vaccinePlacesData, err
}

func GetAllVaccineVenue(bow *browser.Browser, date string) (interface{}, error) {
	var url string

	if len(date) > 0 {
		url = VICTORI_URL + date
	} else {
		url = VICTORI_URL
	}

	err := bow.Open(url)

	columns := []string{}

	vaccinePlacesData := []interface{}{}

	bow.Find("tr:contains('Aksi')").Each(func(i int, rowData *goquery.Selection) {
		vaccinePlaceData := make(map[string]interface{})

		rowData.Children().Each(func(columnIndex int, columnData *goquery.Selection) {
			if i == 0 {
				columns = append(columns, columnData.Text())
			} else {
				if columns[columnIndex] == "Kuota" || columns[columnIndex] == "Terisi" {

					reg, err := regexp.Compile("[^a-zA-Z0-9]+")

					if err != nil {
						log.Fatal(err)
					}

					processedString := reg.ReplaceAllString(columnData.Text(), "")

					value, err := strconv.Atoi(processedString)

					if err != nil {
						fmt.Println("hello world")
						vaccinePlaceData[columns[columnIndex]] = columnData.Text()
					}

					vaccinePlaceData[columns[columnIndex]] = value

				} else {
					vaccinePlaceData[columns[columnIndex]] = columnData.Text()
				}
			}
		})

		if i > 0 {
			vaccinePlacesData = append(vaccinePlacesData, vaccinePlaceData)
		}
	})

	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	}

	return vaccinePlacesData, err
}
