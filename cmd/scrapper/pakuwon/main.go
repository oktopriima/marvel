package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"log"
	"strings"
)

func main() {
	// Create a new collector
	c := colly.NewCollector(
		colly.AllowedDomains("pakuwonmalljogja.com"), // Replace with your target website's domain
	)

	var result []Restaurant

	mainAddr := "Pakuwon Mall Jogja"
	detailAddr := "Jl. Ring Road Utara, Kaliwaru, Condongcatur, Kec. Depok, Kabupaten Sleman, Daerah Istimewa Yogyakarta 55281"

	baseURL := "https://pakuwonmalljogja.com/tenant/dine/page"
	for i := 1; i <= 8; i++ {
		url := fmt.Sprintf("%s/%d", baseURL, i)
		c.OnHTML("div.mbr-section__col", func(e *colly.HTMLElement) {
			name := e.ChildText("h3.mbr-header__text")
			level := strings.ReplaceAll(e.ChildText("p"), "\t", "")
			level = strings.ReplaceAll(level, "\n", "")
			address := fmt.Sprintf("%s %s. %s", mainAddr, level, detailAddr)

			restaurant := Restaurant{
				Name:    name,
				Address: address,
			}
			result = append(result, restaurant)
		})

		// Error handling
		c.OnError(func(_ *colly.Response, err error) {
			log.Println("Error:", err)
		})

		err := c.Visit(url)
		if err != nil {
			log.Println("Error:", err)
		}
		c.Wait()
	}

	fmt.Println(result)
}

type Restaurant struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}
