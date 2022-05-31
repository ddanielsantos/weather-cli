package app

import (
	"testing"
)

type mockWeatherService struct {
	cityName string
	option   Options
}

func teste() {
}

func TestCityCurrentWeather(t *testing.T) {
	validWeatherTests := []mockWeatherService{
		{cityName: "London", option: struct{ daysFromNowOn int }{5}},
		{cityName: "Belém", option: struct{ daysFromNowOn int }{2}},
		{cityName: "São Paulo", option: struct{ daysFromNowOn int }{1}},
	}

	// t.Run("should return error if invalid quantity of days", func(t *testing.T) {

	// })

	// t.Run("should return error if empty city name", func(t *testing.T) {

	// })

	// t.Run("should fetch the forecast for x days from now", func(t *testing.T) {
	// 	for _, test := range validWeatherTests {
	// 		_, err := CityForecastWeather(test.cityName, &test.option)
	// 		if err != nil {
	// 			t.Errorf("error fetching weather for %s", test.cityName)
	// 		}

	// 		if result.CityName != expected {
	// 			t.Errorf("City name should be %s, but was %s", expected, result.CityName)
	// 		}
	// 	}
	// })

	// unixDateToDate := func(unixDate int) string {
	// 	t.Helper()

	// 	date := time.Unix(int64(unixDate), 0).Format("2006-01-02")

	// 	return date
	// }

	t.Run("should fetch the actual weather", func(t *testing.T) {
		for _, test := range validWeatherTests {
			expected := test.cityName
			result, err := CityCurrentWeather(expected)

			if err != nil {
				t.Errorf("error fetching weather for %s", test.cityName)
			}

			if result.CityName != expected {
				t.Errorf("City name should be %s, but was %s", expected, result.CityName)
			}
		}
	})
}
