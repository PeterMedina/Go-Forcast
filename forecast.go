package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// Load the .env file to get API keys
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file. Make sure you have a .env file with Dark Sky and Google Maps API keys.")
	}

	// Declare the API keys for Dark Sky and Google Maps.
	darkskyKey := os.Getenv("DARKSKY")
	googleKey := os.Getenv("GOOGLE_MAPS")

	// Declare a variable to hold the user's location
	var userLocation string

	// Ask for the user's location and store it in userLocation
	fmt.Println("Enter Location:")
	reader := bufio.NewReader(os.Stdin)
	userLocation, err = reader.ReadString('\n')

	// Remove the line break at the end of userLocation
	userLocation = userLocation[:len(userLocation)-1]

	// Declare a variable to hold an escaped version of
	// the user's location
	escapedUserLocation := url.QueryEscape(userLocation)

	url := fmt.Sprintf("https://maps.googleapis.com/maps/api/geocode/json?address=%s&key=%s", escapedUserLocation, googleKey)

	// Create a GET request for the GoogleMaps URL
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return
	}

	// Create an HTTP client
	client := &http.Client{}

	// Send the client request, and return the response
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return
	}

	// End the request
	defer resp.Body.Close()

	// Store the GoogleMaps JSON response in mapData
	var mapData GoogleMapsData

	// Decode the JSON response in mapData
	if err := json.NewDecoder(resp.Body).Decode(&mapData); err != nil {
		log.Println(err)
	}

	// Store the latitude and longitude data in lat and lng
	lat := mapData.Results[0].Geometry.Location.Lat
	lng := mapData.Results[0].Geometry.Location.Lng

	// Store the Dark Sky API url with the key and coordinates
	url = fmt.Sprintf("https://api.darksky.net/forecast/%s/%f,%f", darkskyKey, lat, lng)

	// Create a GET request for Dark Sky
	req, err = http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return
	}

	client = &http.Client{}

	// Send the request
	resp, err = client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return
	}

	// Close the request
	defer resp.Body.Close()

	// Store the Dark Sky JSON response in weatherForecast
	var weatherForecast DarkSkyData

	// Decode weatherForecast JSON
	if err := json.NewDecoder(resp.Body).Decode(&weatherForecast); err != nil {
		log.Println(err)
	}

	// Dark Sky returns humidity as a decimal; convert to whole number
	relativeHumidity := weatherForecast.Currently.Humidity * 100

	// Output the forecast; round numbers to nearest whole
	fmt.Printf("Summary: %s | Temp: %.0f° | Humidity: %.0f%%\n", weatherForecast.Currently.Summary, weatherForecast.Currently.Temperature, relativeHumidity)
}
