# Go Forecast

Enter an address or location name and see the current weather for that place.

Uses Google Maps API to convert the location into geographical coordinates that Dark Sky can recognize.

This project was created as an exercise while learning Go, and is based on [Consuming JSON APIs with Go](https://medium.com/@IndianGuru/consuming-json-apis-with-go-d711efc1dcf9) by Satish Manohar Talim.

## Creating an ENV File

This project uses [Go Dot Env](github.com/joho/godotenv) by [John Barton](https://github.com/joho). To get started, create a .env file and add your Dark Sky and Google Maps API keys. A sample file has been provided. Place the .env file in the same directory that you will be running the app from.

## Planned Updates

A planned future update of this project will include a table view that shows the forecast for multiple days.
