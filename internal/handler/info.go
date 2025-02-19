package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// Struct to represent country information from the first API
type country_info struct {
	Name struct {
		Common string `json:"common"` // Common name of the country
	} `json:"name"`
	Continent  []string          `json:"continents"` // List of continents the country belongs to
	Population int               `json:"population"` // Population of the country
	Languages  map[string]string `json:"languages"`  // Languages spoken in the country
	Borders    []string          `json:"borders"`    // List of bordering countries
	Capital    []string          `json:"capital"`    // List of capital cities
	Flags      struct {
		PNG string `json:"png"` // URL to the country's flag in PNG format
	} `json:"flags"`
}

// Struct to represent city information from the second API
type city_info struct {
	Cities []string `json:"data"` // List of cities in the country
}

// Struct to combine country and city information for the final response
type combined_info struct {
	Name       string            `json:"name"`       // Common name of the country
	Continent  []string          `json:"continents"` // List of continents the country belongs to
	Population int               `json:"population"` // Population of the country
	Languages  map[string]string `json:"languages"`  // Languages spoken in the country
	Borders    []string          `json:"borders"`    // List of bordering countries
	Flag       string            `json:"flag"`       // URL to the country's flag in PNG format
	Capital    string            `json:"capital"`    // Capital city of the country
	Cities     []string          `json:"cities"`     // List of cities in the country
}

// GetInfo Handler function to fetch and combine country and city information
func GetInfo(w http.ResponseWriter, r *http.Request) {
	// Extract the two-letter country code from the URL path
	vars := mux.Vars(r)
	country_code := vars["two_letter_country_code"]

	// Validate the country code length
	if len(country_code) != 2 {
		http.Error(w, "Invalid country code", http.StatusNotFound)
		return
	}

	// Extract the "limit" query parameter to limit the number of cities returned
	userLimit := r.URL.Query().Get("limit")
	limit := 10 // Default limit
	if userLimit != "" {
		var err error
		limit, err = strconv.Atoi(userLimit) // Convert the limit to an integer
		if err != nil || limit < 1 {         // Validate the limit
			http.Error(w, "Invalid limit", http.StatusBadRequest)
			return
		}
	}

	// Fetch country information from the first API
	res, err := http.Get("http://129.241.150.113:8080/v3.1/alpha/" + country_code + "?fields=name,capital,continents,population,languages,borders,flags")
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

	// Decode the country information from the first API response
	var countryInfo country_info
	if err := json.NewDecoder(res.Body).Decode(&countryInfo); err != nil {
		log.Printf("Error decoding response: %v", err)
		http.Error(w, "Internal Server Error: Unable to decode country data", http.StatusInternalServerError)
		return
	}

	// Fetch city information from the second API
	url := "http://129.241.150.113:3500/api/v0.1/countries/cities"
	method := "POST"

	// Create the payload for the second API request
	payload := strings.NewReader(fmt.Sprintf(`{"country": "%s"}`, countryInfo.Name.Common))

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
		http.Error(w, "Failed to fetch cities", http.StatusInternalServerError)
		return
	}
	defer response.Body.Close() // Ensure the response body is closed

	// Handle non-200 status codes from the second API
	if response.StatusCode != http.StatusOK {
		log.Printf("Unexpected status code from second API: %d", response.StatusCode)
		http.Error(w, "Internal Server Error: Unable to fetch cities", http.StatusInternalServerError)
		return
	}

	// Decode the city information from the second API response
	var cityInfo city_info
	if err := json.NewDecoder(response.Body).Decode(&cityInfo); err != nil {
		log.Printf("Error decoding response: %v", err)
		http.Error(w, "Internal Server Error: Unable to decode cities data", http.StatusInternalServerError)
		return
	}

	// Limit the number of cities based on the user-provided limit
	if len(cityInfo.Cities) > limit {
		cityInfo.Cities = cityInfo.Cities[:limit]
	}

	// Use the first capital city if available
	var capital string
	if len(countryInfo.Capital) > 0 {
		capital = countryInfo.Capital[0]
	}

	// Combine country and city information into a single struct
	CombinedInfo := combined_info{
		Name:       countryInfo.Name.Common,
		Continent:  countryInfo.Continent,
		Population: countryInfo.Population,
		Languages:  countryInfo.Languages,
		Borders:    countryInfo.Borders,
		Flag:       countryInfo.Flags.PNG,
		Capital:    capital,
		Cities:     cityInfo.Cities,
	}

	// Write the combined information as a JSON response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	out, _ := json.Marshal(CombinedInfo)
	w.Write(out)
}
