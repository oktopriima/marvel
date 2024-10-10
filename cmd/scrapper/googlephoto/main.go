package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

const apiKey = "AIzaSyDGdLAwAsQ20DL7pQKcYVEIDS3qzAdIj8c"
const placesAPIURL = "https://maps.googleapis.com/maps/api/place/findplacefromtext/json"

type PlaceResponse struct {
	Candidates []struct {
		Photos []struct {
			PhotoReference string `json:"photo_reference"`
		} `json:"photos"`
	} `json:"candidates"`
	Status string `json:"status"`
}

func main() {
	places := []string{
		"Komplek Aph Seturan Baru E-1",
	}

	for _, place := range places {
		fetchImage(place)
	}
}

func fetchImage(placeName string) {
	// Create the request URL
	url := fmt.Sprintf("%s?input=%s&inputtype=textquery&fields=photos&key=%s", placesAPIURL, strings.ReplaceAll(placeName, " ", "%20"), apiKey)

	// Make the HTTP request
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching place data:", err)
		return
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error fetching place: ", resp.StatusCode)
		return
	}

	// Parse the JSON response
	var placeResponse PlaceResponse
	if err := json.Unmarshal(body, &placeResponse); err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}

	// Check if we received any candidates
	if len(placeResponse.Candidates) == 0 {
		fmt.Println("No places found.")
		return
	}

	// Get the photo reference and create the photo URL
	for key, ref := range placeResponse.Candidates[0].Photos {
		imageURL := fmt.Sprintf("https://maps.googleapis.com/maps/api/place/photo?maxwidth=400&photoreference=%s&key=%s", ref.PhotoReference, apiKey)
		downloadImage(imageURL, fmt.Sprintf("%s-%d", placeName, key))
	}

	// Download the image
}

func downloadImage(url string, placeName string) {
	// Make the HTTP request for the image
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error downloading image:", err)
		return
	}
	defer resp.Body.Close()

	// Create a file to save the image
	outFile, err := os.Create("images/" + strings.ReplaceAll(strings.ToLower(placeName), " ", "-") + ".jpg")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer outFile.Close()

	// Write the image data to the file
	_, err = io.Copy(outFile, resp.Body)
	if err != nil {
		fmt.Println("Error saving image:", err)
		return
	}

	fmt.Printf("Image saved as %s.jpg\n", placeName)
}
