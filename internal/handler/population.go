package handler

import (
	"assignment-1/internal/api"
	"assignment-1/internal/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"strings"
)

func GetPopulation(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	country_code := vars["two_letter_country_code"]

	// Validate the country code length
	if len(country_code) != 2 {
		http.Error(w, "Invalid country code", http.StatusNotFound)
		return
	}

	// Fetch, decode and return struct of Country API
	countryInfo, err := api.GetCountryApi(country_code)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// Fetch, decode and return struct of Population API
	populationInfo, err := api.GetPopulationAPI(countryInfo.Name.Common)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	var limitedPopulation []models.PopulationValue
	yearLimit := r.URL.Query().Get("limit")
	if yearLimit != "" {
		years := strings.Split(yearLimit, "-")
		if len(years) == 2 {
			startYear, _ := strconv.Atoi(years[0])
			endYear, _ := strconv.Atoi(years[1])

			for i := 0; i < len(populationInfo.Data.PopulationValues); i++ {
				if populationInfo.Data.PopulationValues[i].Year >= startYear && populationInfo.Data.PopulationValues[i].Year <= endYear {
					limitedPopulation = append(limitedPopulation, populationInfo.Data.PopulationValues[i])
				}
			}

		}

		if len(years) != 2 {
			http.Error(w, "Invalid limit", http.StatusNotFound)
		}

	} else {
		limitedPopulation = populationInfo.Data.PopulationValues
	}

	// calculate mean
	var sum int
	for i := 0; i < len(limitedPopulation); i++ {
		sum += limitedPopulation[i].Value
	}
	sum = sum / len(limitedPopulation)

	FinalPopulation := models.FinalPopulation{
		Mean:   sum,
		Values: limitedPopulation,
	}

	// Write the combined information as a JSON response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	out, _ := json.Marshal(FinalPopulation)
	w.Write(out)
}
