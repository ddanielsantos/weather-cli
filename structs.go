package main

type Coordinates struct {
	Latitude  float64 `json:"lon"`
	Longitude float64 `json:"lat"`
}

type Wind struct {
	Speed float64 `json:"speed"`
	Deg   float64 `json:"deg"`
	Gust  float64 `json:"gust"`
}

type Forecast struct {
	Temp      float64 `json:"temp"`
	FeelsLike float64 `json:"feels_like"`
	Min       float64 `json:"temp_min"`
	Max       float64 `json:"temp_max"`
	Pressure  int     `json:"pressure"`
	Humidity  int     `json:"humidity"`
}

type Weather struct {
	Id          int    `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type CurrentWeather struct {
	Coordinates Coordinates `json:"coord"`
	Weather     []Weather   `json:"weather"`
	Forecast    Forecast    `json:"main"`
	Wind        Wind        `json:"wind"`
	CityName    string      `json:"name"`
	Info        struct {
		country string
	} `json:"sys"`
}
