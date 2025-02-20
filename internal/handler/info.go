package handler

import (
	"assignment-1/internal/api"
	"assignment-1/internal/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

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
	// Fetch, decode and return struct of Country API
	countryInfo, err := api.GetCountryApi(country_code)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// Fetch, decode and return struct of Cities API
	cityInfo, err := api.GetCitiesApi(countryInfo.Name.Common)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
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
	CombinedInfo := models.Combined_info{
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
