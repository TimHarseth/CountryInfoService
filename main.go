package main

import (
	"assignment-1/internal/handler"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
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

	r.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/countryinfo/v1/", http.StatusMovedPermanently)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Println("Listening on port", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
