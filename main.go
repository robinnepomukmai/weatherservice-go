package main

import (
	"fmt"
	"github.com/Nepooomuk/weatherservice-go/weather"
	"net/http"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/client_golang/prometheus"
	"time"
)

var 	weatherapiResponseTime = prometheus.NewSummary(
	prometheus.SummaryOpts{Name: "weather_api_response_time", Help: "Response time for weatherapi requests"},
)

func init() {
	prometheus.MustRegister(weatherapiResponseTime)

}

func main() {
	http.HandleFunc("/weather", func(w http.ResponseWriter, r *http.Request) {
		tnow := time.Now()
		report, err := weather.CreateForecast()
		weatherapiResponseTime.Observe(time.Now().Sub(tnow).Seconds())
		fmt.Fprintln(w, report, err)
	})
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":8080", nil)
}