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

type TwoSlices struct {
	main_slice  []int
	other_slice []int
}

type SortByOther TwoSlices

func (sbo SortByOther) Len() int {
	return len(sbo.main_slice)
}

func (sbo SortByOther) Swap(i, j int) {
	sbo.main_slice[i], sbo.main_slice[j] = sbo.main_slice[j], sbo.main_slice[i]
	sbo.other_slice[i], sbo.other_slice[j] = sbo.other_slice[j], sbo.other_slice[i]
}

func (sbo SortByOther) Less(i, j int) bool {
	return sbo.other_slice[i] < sbo.other_slice[j]
}

func GetAvailableVaccineVenue(bow *browser.Browser, date string) ([]map[string]interface{}, []string, error) {
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

	vaccinePlacesData := []map[string]interface{}{}

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
			} else if columns[columnIndex] == "Aksi" {
				vaccinePlaceData[columns[columnIndex]] = s.Children().First().Text()
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

			} else if columns[columnIndex] == "Aksi" {
				vaccinePlaceData[columns[columnIndex]] = columnData.Children().First().Text()
			} else {
				vaccinePlaceData[columns[columnIndex]] = columnData.Text()
			}

		})

		vaccinePlacesData = append(vaccinePlacesData, vaccinePlaceData)
	})

	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	}

	return vaccinePlacesData, columns, err
}

func GetAllVaccineVenue(bow *browser.Browser, date string) ([]map[string]interface{}, []string, error) {
	var url string

	if len(date) > 0 {
		url = VICTORI_URL + date
	} else {
		url = VICTORI_URL
	}

	err := bow.Open(url)

	columns := []string{}

	vaccinePlacesData := []map[string]interface{}{}

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

				} else if columns[columnIndex] == "Aksi" {
					vaccinePlaceData[columns[columnIndex]] = columnData.Children().First().Text()
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

	return vaccinePlacesData, columns, err
}

func GetVaccinationDate(bow *browser.Browser) ([]string, error) {
	err := bow.Open(VICTORI_URL)

	if err != nil {
		return nil, err
	}

	vaccinationDate := []string{}

	bow.Find("select[name='tanggal']").Children().Each(func(i int, s *goquery.Selection) {

		if len(s.Text()) == 0 {
			return
		}

		vaccinationDate = append(vaccinationDate, s.Text())

	})

	return vaccinationDate, err
}
