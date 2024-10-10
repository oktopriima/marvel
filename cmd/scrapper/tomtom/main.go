package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

const apiKey = "RcKXGlSGvJcuhH8pbshiMlxcABrAoHy0"

func main() {
	// Define the URL for the GET request
	city := "papua"
	baseUrl := "https://api.tomtom.com/search/2/poiSearch/"
	url := fmt.Sprintf("%s%s.json?key=%s&limit=100", baseUrl, city, apiKey)

	// Create the GET request
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Failed to make GET request: %v", err)
	}
	defer resp.Body.Close() // Ensure the response body is closed after reading

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}

	var res TomTomResult
	err = json.Unmarshal(body, &res)
	if err != nil {
		log.Fatalf("Failed to unmarshal response: %v", err)
	}

	// Convert the struct to JSON
	jsonData, err := json.MarshalIndent(res, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling to JSON:", err)
		return
	}

	// Write JSON data to a file
	fileName := fmt.Sprintf("%s.json", city)
	err = os.WriteFile(fileName, jsonData, 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	// Print the response status and body
	fmt.Println("Response Status:", resp.Status)
	fmt.Printf("Data written to %s", fileName)
}

type TomTomResult struct {
	Summary struct {
		Query        string `json:"query"`
		QueryType    string `json:"queryType"`
		QueryTime    int    `json:"queryTime"`
		NumResults   int    `json:"numResults"`
		Offset       int    `json:"offset"`
		TotalResults int    `json:"totalResults"`
		FuzzyLevel   int    `json:"fuzzyLevel"`
	} `json:"summary"`
	Results []struct {
		Type  string  `json:"type"`
		Id    string  `json:"id"`
		Score float64 `json:"score"`
		Info  string  `json:"info"`
		Poi   struct {
			Name        string `json:"name"`
			Phone       string `json:"phone,omitempty"`
			CategorySet []struct {
				Id int `json:"id"`
			} `json:"categorySet"`
			Categories      []string `json:"categories"`
			Classifications []struct {
				Code  string `json:"code"`
				Names []struct {
					NameLocale string `json:"nameLocale"`
					Name       string `json:"name"`
				} `json:"names"`
			} `json:"classifications"`
			Url string `json:"url,omitempty"`
		} `json:"poi"`
		Address struct {
			StreetName                       string `json:"streetName,omitempty"`
			MunicipalitySubdivision          string `json:"municipalitySubdivision,omitempty"`
			Municipality                     string `json:"municipality"`
			MunicipalitySecondarySubdivision string `json:"municipalitySecondarySubdivision"`
			CountrySubdivision               string `json:"countrySubdivision"`
			CountrySubdivisionName           string `json:"countrySubdivisionName"`
			CountrySubdivisionCode           string `json:"countrySubdivisionCode"`
			PostalCode                       string `json:"postalCode"`
			CountryCode                      string `json:"countryCode"`
			Country                          string `json:"country"`
			CountryCodeISO3                  string `json:"countryCodeISO3"`
			FreeformAddress                  string `json:"freeformAddress"`
			LocalName                        string `json:"localName"`
		} `json:"address"`
		Position struct {
			Lat float64 `json:"lat"`
			Lon float64 `json:"lon"`
		} `json:"position"`
		Viewport struct {
			TopLeftPoint struct {
				Lat float64 `json:"lat"`
				Lon float64 `json:"lon"`
			} `json:"topLeftPoint"`
			BtmRightPoint struct {
				Lat float64 `json:"lat"`
				Lon float64 `json:"lon"`
			} `json:"btmRightPoint"`
		} `json:"viewport"`
		EntryPoints []struct {
			Type     string `json:"type"`
			Position struct {
				Lat float64 `json:"lat"`
				Lon float64 `json:"lon"`
			} `json:"position"`
		} `json:"entryPoints"`
		DataSources struct {
			Geometry struct {
				Id string `json:"id"`
			} `json:"geometry"`
		} `json:"dataSources,omitempty"`
	} `json:"results"`
}
