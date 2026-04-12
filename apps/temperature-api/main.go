package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"time"
)

type Response struct {
	Location    string  `json:"location"`
	Temperature float64 `json:"temperature"`
	Unit        string  `json:"unit"`
	Timestamp   int64   `json:"timestamp"`
}

func temperatureHandler(w http.ResponseWriter, r *http.Request) {
	location := r.URL.Query().Get("location")
	if location == "" {
		location = "unknown"
	}

	temp := 10 + rand.Float64()*25

	resp := Response{
		Location:    location,
		Temperature: temp,
		Unit:        "C",
		Timestamp:   time.Now().Unix(),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func main() {
	rand.Seed(time.Now().UnixNano())

	http.HandleFunc("/temperature", temperatureHandler)

	server := &http.Server{
		Addr:         ":8081",
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	println("temperature-api running on :8081")
	server.ListenAndServe()
}
