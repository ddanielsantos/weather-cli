package app

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	dotenv "github.com/joho/godotenv"
)

var API_KEY = getEnv("API_KEY")

func getEnv(key string) string {
	path, _ := os.Getwd()
	err := dotenv.Load(strings.Split(path, "weather-cli")[0] + "weather-cli/.env")
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

type Options struct {
	daysFromNowOn int
}

func CityForecast(cityName string, options *Options) ([]Forecast, error) {
	return []Forecast{}, nil
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

func UnixToDate(unixTime int) string {
	return time.Unix(int64(unixTime), 0).Format("2006-01-02")
}

func RepeatChar(char string, howManyTimes int) string {
	return strings.Repeat(char, howManyTimes)
}

func PrintWeather(city string) {
	w, err := CityCurrentWeather(city)
	HandleError(err)

	date := UnixToDate(w.Date)

	// 42 minimum width
	cityNameLength := len(w.CityName)
	countryAbbrLength := 2
	firstLineBaseLenght := 42
	boxWidth := firstLineBaseLenght + cityNameLength + countryAbbrLength

	// the width of this line will be = boxWidth
	fmt.Printf("\n┏━━━━━━━┫ Current weather in %s - %s ┣━━━━━━━┓\n", city, w.Info.Country)

	// the width of this line will be = boxWith - 10 (harcoded string) + 10 (date lenght)
	fmt.Printf("┃%s┃\n", RepeatChar(" ", boxWidth-2))
	fmt.Printf("┃ Date: %s%s ┃\n", date, RepeatChar(" ", boxWidth-10-10))

	latLenght := len(fmt.Sprintf("%.2f", w.Coordinates.Latitude))
	lonLenght := len(fmt.Sprintf("%.2f", w.Coordinates.Longitude))
	fmt.Printf("┃ Lat: %.2fº%s ┃\n┃ Lon: %.2fº%s ┃\n", w.Coordinates.Latitude, RepeatChar(" ", boxWidth-10-latLenght), w.Coordinates.Longitude, RepeatChar(" ", boxWidth-10-lonLenght))

	feelsLikeLenght := len(fmt.Sprintf("%.2f", w.Forecast.FeelsLike))
	minTemperatureLength := len(fmt.Sprintf("%.2f", w.Forecast.Min))
	maxTemperatureLength := len(fmt.Sprintf("%.2f", w.Forecast.Max))
	fmt.Printf("┃ Feels like: %.2f Kº%s ┃\n┃ Min: %.2f Kº%s ┃\n┃ Max: %.2f Kº%s ┃\n", w.Forecast.FeelsLike, RepeatChar(" ", boxWidth-19-feelsLikeLenght), w.Forecast.Min, RepeatChar(" ", boxWidth-12-minTemperatureLength), w.Forecast.Max, RepeatChar(" ", boxWidth-12-maxTemperatureLength))

	pressureLenght := len(fmt.Sprintf("%d", w.Forecast.Pressure))
	humidityLenght := len(fmt.Sprintf("%d", w.Forecast.Humidity))
	fmt.Printf("┃ Pressure: %d hPa%s ┃\n┃ Humidity: %d%%%s ┃\n", w.Forecast.Pressure, RepeatChar(" ", boxWidth-18-pressureLenght), w.Forecast.Humidity, RepeatChar(" ", boxWidth-15-humidityLenght))

	windLenght := len(fmt.Sprintf("%.2f", w.Wind.Speed))
	degreesLenght := len(fmt.Sprintf("%.2f", w.Wind.Deg))
	gustLenght := len(fmt.Sprintf("%.2f", w.Wind.Gust))
	fmt.Printf("┃ Wind speed: %.2f m/s%s ┃\n┃ Deg: %.2f%s ┃\n┃ Gust: %.2f m/s%s ┃\n", w.Wind.Speed, RepeatChar(" ", boxWidth-20-windLenght), w.Wind.Deg, RepeatChar(" ", boxWidth-9-degreesLenght), w.Wind.Gust, RepeatChar(" ", boxWidth-14-gustLenght))
	fmt.Printf("┗%s┛\n\n", RepeatChar("━", boxWidth-2))
}
