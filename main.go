package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var API_KEY = GetEnv("API_KEY")

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

func cityCurrentWeather(cityName string) (CurrentWeather, error) {
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

func main() {
	// TODO: change to command line args
	w, err := cityCurrentWeather("Frankfurt")
	HandleError(err)

	// j, err := json.MarshalIndent(w, "", "  ")
	// HandleError(err)
	// fmt.Println(string(j))

	fmt.Println(w)
}
