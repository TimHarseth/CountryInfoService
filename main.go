package main

import (
	"assignment-1/internal/handler"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/countryinfo/v1/info/{two_letter_country_code}", handler.GetInfo)
	//http.HandleFunc("/countryinfo/v1/population/{:two_letter_country_code}{?limit={:startYear-endYear}}", handler.population)
	//http.HandleFunc("/countryinfo/v1/status/", handler.status)

	fmt.Println("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
