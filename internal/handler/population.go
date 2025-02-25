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

	// new struct for limiting by year
	var limitedPopulation []models.PopulationValue

	yearLimit := r.URL.Query().Get("limit")
	if yearLimit != "" {
		// splitting up the startYear and endYear
		years := strings.Split(yearLimit, "-")
		if len(years) != 2 {
			http.Error(w, "Invalid limit format. Use xxxx-xxxx, e.g., 2008-2015.", http.StatusBadRequest)
			return
		}
		startYear, err1 := strconv.Atoi(years[0])
		endYear, err2 := strconv.Atoi(years[1])

		// Checks for parsing errors
		if err1 != nil || err2 != nil {
			http.Error(w, "Invalid year format. Years must be numbers.", http.StatusBadRequest)
			return
		}

		// Makes sure startYear is less or equal to endYear
		if startYear > endYear {
			http.Error(w, "Start year must be less than or equal to end year.", http.StatusBadRequest)
			return
		}

		// appends correct year and corresponding population value according to user inputted limit
		for i := 0; i < len(populationInfo.Data.PopulationValues); i++ {
			if populationInfo.Data.PopulationValues[i].Year >= startYear && populationInfo.Data.PopulationValues[i].Year <= endYear {
				limitedPopulation = append(limitedPopulation, populationInfo.Data.PopulationValues[i])
			}
		}

	} else {
		limitedPopulation = populationInfo.Data.PopulationValues // if user does not enter limit, use default length
	}

	// calculate mean
	var sum int
	for i := 0; i < len(limitedPopulation); i++ {
		sum += limitedPopulation[i].Value
	}
	sum = sum / len(limitedPopulation)

	// final struct for writing out
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
