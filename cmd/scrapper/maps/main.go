package main

import (
	"fmt"
	"regexp"
)

func main() {
	// URL containing latitude and longitude
	url := "https://www.google.com/maps/place/Jogja+City+Mall/@-7.7856924,110.3164483,12.88z/data=!3m1!5s0x2e7a58f6fef5b2a1:0x2e6ec8c226d1d39c!4m6!3m5!1s0x2e7a58f6765d4a33:0x39d891d2e8da95a9!8m2!3d-7.7532424!4d110.3609091!16s%2Fg%2F1yp37k7pr?entry=ttu&g_ep=EgoyMDI0MTAwNS4yIKXMDSoASAFQAw%3D%3D"

	// Regular expression to extract latitude and longitude
	re := regexp.MustCompile(`@(-?\d+\.\d+),(-?\d+\.\d+)`)
	matches := re.FindStringSubmatch(url)

	// Check if matches were found
	if len(matches) == 3 {
		latitude := matches[1]
		longitude := matches[2]
		fmt.Printf("Latitude: %s\nLongitude: %s\n", latitude, longitude)
	} else {
		fmt.Println("Latitude and longitude not found in the URL.")
	}
}
