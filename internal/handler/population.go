package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strings"
)

type population_info struct {
	PopulationValues []map[string]int `json:"populationCounts"`
	// fixthis
}

func getPopulation(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	country_code := vars["two_letter_country_code"]

	// Validate the country code length
	if len(country_code) != 2 {
		http.Error(w, "Invalid country code", http.StatusNotFound)
		return
	}

	res, err := http.Get("http://129.241.150.113:8080/v3.1/alpha/" + country_code + "?fields=name")
	if err != nil {
		// Log the error and return an internal server error response
		log.Printf("Error making HTTP request: %v", err)
		http.Error(w, "Internal Server Error: Unable to fetch data", http.StatusInternalServerError)
		return
	}
	defer res.Body.Close() // Ensure the response body is closed

	// Handle 404 Not Found response from the first API
	if res.StatusCode == 404 {
		http.Error(w, "Country not found", http.StatusNotFound)
		return
	}

	var getName string
	// Decode the country name  from the first API response
	if err := json.NewDecoder(res.Body).Decode(&getName); err != nil {
		log.Printf("Error decoding response: %v", err)
		http.Error(w, "Internal Server Error: Unable to decode country data", http.StatusInternalServerError)
		return
	}

	// Fetch city information from the second API
	url := "http://129.241.150.113:3500/api/v0.1/countries/population"
	method := "POST"

	// Create the payload for the second API request
	payload := strings.NewReader(fmt.Sprintf(`{"country": "%s"}`, getName))

	// Create a new HTTP request for the second API
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		log.Printf("Error creating request: %v", err)
		http.Error(w, "Internal Server Error: Unable to fetch data", http.StatusInternalServerError)
		return
	}
	req.Header.Set("Content-Type", "application/json") // Set the content type header

	// Send the request to the second API
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		http.Error(w, "Failed to fetch population", http.StatusInternalServerError)
		return
	}
	defer response.Body.Close() // Ensure the response body is closed

	// Handle non-200 status codes from the second API
	if response.StatusCode != http.StatusOK {
		log.Printf("Unexpected status code from second API: %d", response.StatusCode)
		http.Error(w, "Internal Server Error: Unable to fetch population", http.StatusInternalServerError)
		return
	}

	// Decode the city information from the second API response
	var populationInfo population_info
	if err := json.NewDecoder(response.Body).Decode(&cityInfo); err != nil {
		log.Printf("Error decoding response: %v", err)
		http.Error(w, "Internal Server Error: Unable to decode population data", http.StatusInternalServerError)
		return
	}

}
