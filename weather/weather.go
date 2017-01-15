package weather

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"log"
	"encoding/json"
)

const apiKey = "aa34c9b93c8e923537921e29afefbd23"

type WeatherReport struct {
	Weather []struct {
		Main        string `json:"main"`
		Description string `json:"description"`
	}
	Main struct {
		Temperature    float64 `json:"temp"`
		TemperatureMin float64 `json:"temp_min"`
		TemperatureMax float64 `json:"temp_max"`
	}
	Sys struct {
		Country string `json:"country"`
		Sunrise int    `json:"sunrise"`
		Sunset  int    `json:"sunset"`
	}
	Name  string `json:"name"`
	Error string `json:"message"`
}

type urlType struct {
	city   string
	apiKey string
}

func buildUrl(u urlType) string {
	var city string = "cologne"
	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q={%s}&appid=%s", city, apiKey)
	return url
}

func getForecast() []byte {
	forecastUrl := buildUrl(urlType{city: "cologne", apiKey: apiKey})
	res, err := http.Get(forecastUrl)
	defer res.Body.Close()
	if err != nil {
		log.Print(err)
	}

	weather, err := ioutil.ReadAll(res.Body)
	return weather
}

func CreateForecast() (WeatherReport, error) {
	var report WeatherReport
	data := getForecast()

	if err := json.Unmarshal(data, &report); err != nil {
		return report, err
	}
	return report, nil
}

