package app

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	dotenv "github.com/joho/godotenv"
)

var API_KEY = getEnv("API_KEY")

func getEnv(key string) string {
	err := dotenv.Load("../.env")
	HandleError(err)

	return os.Getenv(key)
}

func getRequest(url string) ([]byte, error) {
	resp, e := http.Get(url)

	if e != nil {
		return []byte{}, e
	}

	defer resp.Body.Close()

	bytes, e := ioutil.ReadAll(resp.Body)

	if e != nil {
		return []byte{}, e
	}

	return bytes, e
}

func CityCurrentWeather(cityName string) (CurrentWeather, error) {
	ForecastSearchURL := "https://api.openweathermap.org/data/2.5/weather?q=" + cityName + "&appId=" + API_KEY

	resp, e := getRequest(ForecastSearchURL)

	if e != nil {
		return CurrentWeather{}, e
	}

	var currentWeather CurrentWeather
	e = json.Unmarshal(resp, &currentWeather)

	if e != nil {
		return CurrentWeather{}, e
	}

	return currentWeather, e
}

func PrintWeather(city string) {
	w, err := CityCurrentWeather(city)
	HandleError(err)

	fmt.Println(w)
}
