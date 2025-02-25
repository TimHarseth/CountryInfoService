package main

import (
	"assignment-1/internal/handler"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func main() {
	handler.Starttime = time.Now() // start timer when application starts for Uptime in status
	r := mux.NewRouter()

	r.HandleFunc("/countryinfo/v1/info/{two_letter_country_code}", handler.GetInfo)
	r.HandleFunc("/countryinfo/v1/population/{two_letter_country_code}", handler.GetPopulation)
	r.HandleFunc("/countryinfo/v1/status/", handler.GetStatus)
	r.HandleFunc("/countryinfo/v1/", handler.Root)
	r.HandleFunc("/countryinfo/v1/info/", handler.DefaultInfo)
	r.HandleFunc("/countryinfo/v1/population/", handler.DefaultPopulation)

	fmt.Println("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
