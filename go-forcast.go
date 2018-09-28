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

// GoogleMapsData uses the Google Maps API to return information
// about a location entered by the user. The DarkSky API can only
// accept geographical coordinates, which are provided in the
// GoogleMapsData struct as Lat and Lng.
//
// GoogleMapsData.Results[0].Geometry.Location.Lat
// GoogleMapsData.Results[0].Geometry.Location.Lng
//
// Add a Google Maps API key to your .env file.
type GoogleMapsData struct {
	Results []struct {
		AddressComponents []struct {
			LongName  string   `json:"long_name"`
			ShortName string   `json:"short_name"`
			Types     []string `json:"types"`
		} `json:"address_components"`
		FormattedAddress string `json:"formatted_address"`
		Geometry         struct {
			Location struct {
				Lat float64 `json:"lat"`
				Lng float64 `json:"lng"`
			} `json:"location"`
			LocationType string `json:"location_type"`
			Viewport     struct {
				Northeast struct {
					Lat float64 `json:"lat"`
					Lng float64 `json:"lng"`
				} `json:"northeast"`
				Southwest struct {
					Lat float64 `json:"lat"`
					Lng float64 `json:"lng"`
				} `json:"southwest"`
			} `json:"viewport"`
		} `json:"geometry"`
		PlaceID  string `json:"place_id"`
		PlusCode struct {
			CompoundCode string `json:"compound_code"`
			GlobalCode   string `json:"global_code"`
		} `json:"plus_code"`
		Types []string `json:"types"`
	} `json:"results"`
	Status string `json:"status"`
}

// DarkSkyData uses the Dark Sky API to get the forcast for
// any geographical coordinate on earth. The DarkSkyData struct
// includes the all the available data. Note that you can change
// which type of weather data you want Dark Sky to return by
// changing the quary parameters. If you do that, you'll want to
// change which data gets included in the DarkSkyData struct.
//
// Add a Dark Sky API key to your .env file.
type DarkSkyData struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Timezone  string  `json:"timezone"`
	Currently struct {
		Time                 int     `json:"time"`
		Summary              string  `json:"summary"`
		Icon                 string  `json:"icon"`
		NearestStormDistance int     `json:"nearestStormDistance"`
		PrecipIntensity      float64 `json:"precipIntensity"`
		PrecipIntensityError float64 `json:"precipIntensityError"`
		PrecipProbability    float64 `json:"precipProbability"`
		PrecipType           string  `json:"precipType"`
		Temperature          float64 `json:"temperature"`
		ApparentTemperature  float64 `json:"apparentTemperature"`
		DewPoint             float64 `json:"dewPoint"`
		Humidity             float64 `json:"humidity"`
		Pressure             float64 `json:"pressure"`
		WindSpeed            float64 `json:"windSpeed"`
		WindGust             float64 `json:"windGust"`
		WindBearing          int     `json:"windBearing"`
		CloudCover           float64 `json:"cloudCover"`
		UvIndex              int     `json:"uvIndex"`
		Visibility           float64 `json:"visibility"`
		Ozone                float64 `json:"ozone"`
	} `json:"currently"`
	Minutely struct {
		Summary string `json:"summary"`
		Icon    string `json:"icon"`
		Data    []struct {
			Time                 int     `json:"time"`
			PrecipIntensity      int     `json:"precipIntensity"`
			PrecipProbability    int     `json:"precipProbability"`
			PrecipIntensityError float64 `json:"precipIntensityError,omitempty"`
			PrecipType           string  `json:"precipType,omitempty"`
		} `json:"data"`
	} `json:"minutely"`
	Hourly struct {
		Summary string `json:"summary"`
		Icon    string `json:"icon"`
		Data    []struct {
			Time                int     `json:"time"`
			Summary             string  `json:"summary"`
			Icon                string  `json:"icon"`
			PrecipIntensity     float64 `json:"precipIntensity"`
			PrecipProbability   float64 `json:"precipProbability"`
			PrecipType          string  `json:"precipType,omitempty"`
			Temperature         float64 `json:"temperature"`
			ApparentTemperature float64 `json:"apparentTemperature"`
			DewPoint            float64 `json:"dewPoint"`
			Humidity            float64 `json:"humidity"`
			Pressure            float64 `json:"pressure"`
			WindSpeed           float64 `json:"windSpeed"`
			WindGust            float64 `json:"windGust"`
			WindBearing         int     `json:"windBearing"`
			CloudCover          float64 `json:"cloudCover"`
			UvIndex             int     `json:"uvIndex"`
			Visibility          float64 `json:"visibility"`
			Ozone               float64 `json:"ozone"`
		} `json:"data"`
	} `json:"hourly"`
	Daily struct {
		Summary string `json:"summary"`
		Icon    string `json:"icon"`
		Data    []struct {
			Time                        int     `json:"time"`
			Summary                     string  `json:"summary"`
			Icon                        string  `json:"icon"`
			SunriseTime                 int     `json:"sunriseTime"`
			SunsetTime                  int     `json:"sunsetTime"`
			MoonPhase                   float64 `json:"moonPhase"`
			PrecipIntensity             float64 `json:"precipIntensity"`
			PrecipIntensityMax          float64 `json:"precipIntensityMax"`
			PrecipIntensityMaxTime      int     `json:"precipIntensityMaxTime"`
			PrecipProbability           float64 `json:"precipProbability"`
			PrecipType                  string  `json:"precipType"`
			TemperatureHigh             float64 `json:"temperatureHigh"`
			TemperatureHighTime         int     `json:"temperatureHighTime"`
			TemperatureLow              float64 `json:"temperatureLow"`
			TemperatureLowTime          int     `json:"temperatureLowTime"`
			ApparentTemperatureHigh     float64 `json:"apparentTemperatureHigh"`
			ApparentTemperatureHighTime int     `json:"apparentTemperatureHighTime"`
			ApparentTemperatureLow      float64 `json:"apparentTemperatureLow"`
			ApparentTemperatureLowTime  int     `json:"apparentTemperatureLowTime"`
			DewPoint                    float64 `json:"dewPoint"`
			Humidity                    float64 `json:"humidity"`
			Pressure                    float64 `json:"pressure"`
			WindSpeed                   float64 `json:"windSpeed"`
			WindGust                    float64 `json:"windGust"`
			WindGustTime                int     `json:"windGustTime"`
			WindBearing                 int     `json:"windBearing"`
			CloudCover                  float64 `json:"cloudCover"`
			UvIndex                     int     `json:"uvIndex"`
			UvIndexTime                 int     `json:"uvIndexTime"`
			Visibility                  float64 `json:"visibility"`
			Ozone                       float64 `json:"ozone"`
			TemperatureMin              float64 `json:"temperatureMin"`
			TemperatureMinTime          int     `json:"temperatureMinTime"`
			TemperatureMax              float64 `json:"temperatureMax"`
			TemperatureMaxTime          int     `json:"temperatureMaxTime"`
			ApparentTemperatureMin      float64 `json:"apparentTemperatureMin"`
			ApparentTemperatureMinTime  int     `json:"apparentTemperatureMinTime"`
			ApparentTemperatureMax      float64 `json:"apparentTemperatureMax"`
			ApparentTemperatureMaxTime  int     `json:"apparentTemperatureMaxTime"`
		} `json:"data"`
	} `json:"daily"`
	Flags struct {
		Sources        []string `json:"sources"`
		NearestStation float64  `json:"nearest-station"`
		Units          string   `json:"units"`
	} `json:"flags"`
	Offset int `json:"offset"`
}

func main() {
	// Load the .env file to get API keys
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
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

	// Store the Dark Sky JSON response in weatherForcast
	var weatherForcast DarkSkyData

	// Decode weatherForcast JSON
	if err := json.NewDecoder(resp.Body).Decode(&weatherForcast); err != nil {
		log.Println(err)
	}

	// Dark Sky returns humidity as a decimal; convert to whole number
	relativeHumidity := weatherForcast.Currently.Humidity * 100

	// Output the forcast; round numbers to nearest whole
	fmt.Printf("Summary: %s | Temp: %.0f° | Humidity: %.0f%%\n", weatherForcast.Currently.Summary, weatherForcast.Currently.Temperature, relativeHumidity)
}
