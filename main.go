package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Nepooomuk/weatherservice-go/weather"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var weatherapiResponseTime = prometheus.NewSummary(
	prometheus.SummaryOpts{Name: "weather_api_response_time", Help: "Response time for weatherapi requests"},
)

func init() {
	prometheus.MustRegister(weatherapiResponseTime)

}

func main() {
	http.HandleFunc("/weather", handler)
	http.Handle("/metrics", promhttp.Handler())
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Port 8080 is already used")
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	tnow := time.Now()
	report, err := weather.CreateForecast()
	weatherapiResponseTime.Observe(time.Now().Sub(tnow).Seconds())
	fmt.Fprintln(w, report, err)
}
