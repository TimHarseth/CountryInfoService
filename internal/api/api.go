package api

import (
	"assignment-1/internal/constants"
	"assignment-1/internal/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

func GetCountryApi(code string) (models.Country_info, error) {
	// Fetch country information from the first API
	res, err := http.Get(constants.CountryApiPath + code + "?fields=name,capital,continents,population,languages,borders,flags")
	if err != nil {
		// Log the error
		log.Printf("Error making API request: %v", err)
		return models.Country_info{}, fmt.Errorf("failed to make API request: %v", err)
	}
	defer res.Body.Close() // Ensure the response body is closed

	// Handle 404 Not Found response from the first API
	if res.StatusCode == 404 {
		return models.Country_info{}, fmt.Errorf("country not found: %v", res.StatusCode)
	}

	// Decode the country information from the first API response
	var countryInfo models.Country_info
	if err := json.NewDecoder(res.Body).Decode(&countryInfo); err != nil {
		log.Printf("Error decoding response: %v", err)
		return models.Country_info{}, fmt.Errorf("Internal Server Error: Unable to decode country data: %v", err)
	}
	return countryInfo, err
}

func GetCitiesApi(code string) (models.City_info, error) {
	// Fetch city information from the CountriesNow API
	url := constants.CityApiPath
	method := "POST"

	// Create the payload for the CountriesNow API
	payload := strings.NewReader(fmt.Sprintf(`{"country": "%s"}`, code))

	// Create a new HTTP request for the CountriesNow API
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		log.Printf("Error creating request: %v", err)
		return models.City_info{}, fmt.Errorf("failed to make API request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json") // Set the content type header

	// Send the request to the CountriesNow API
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		log.Printf("Failed to fetch cities: %v", err)
		return models.City_info{}, fmt.Errorf("failed to fetch cities: %v", err)
	}
	defer response.Body.Close() // Ensure the response body is closed

	// Handle 404 Not Found response
	if response.StatusCode == 404 {
		return models.City_info{}, fmt.Errorf("country not found: %v", response.StatusCode)
	}

	// Decode the city information from the CountriesNow API
	var cityInfo models.City_info
	if err := json.NewDecoder(response.Body).Decode(&cityInfo); err != nil {
		log.Printf("Error decoding response: %v", err)
		return models.City_info{}, fmt.Errorf("Internal Server Error: Unable to decode country data: %v", err)
	}
	return cityInfo, err

}

func GetPopulationAPI(code string) (models.Population_info, error) {
	// Fetch population information from CountriesNow API
	url := constants.PopulationApiPath
	method := "POST"

	// Create the payload for the CountriesNow API
	payload := strings.NewReader(fmt.Sprintf(`{"country": "%s"}`, code))

	// Create a new HTTP request for the CountriesNow API
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		log.Printf("Error creating request: %v", err)
		return models.Population_info{}, fmt.Errorf("failed to make API request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json") // Set the content type header

	// Send the request to the CountriesNow API
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		log.Printf("Failed to fetch cities: %v", err)
		return models.Population_info{}, fmt.Errorf("failed to fetch cities: %v", err)
	}
	defer response.Body.Close() // Ensure the response body is closed

	// Handle 404 Not Found response
	if response.StatusCode == 404 {
		return models.Population_info{}, fmt.Errorf("country not found: %v", response.StatusCode)
	}

	// Decode the population information from the CountriesNow API
	var populationInfo models.Population_info
	if err := json.NewDecoder(response.Body).Decode(&populationInfo); err != nil {
		log.Printf("Error decoding response: %v", err)
		return models.Population_info{}, fmt.Errorf("Internal Server Error: Unable to decode country data: %v", err)
	}
	return populationInfo, err

}

/*
GetStatusApi returns status code for url
if unable to Get url, returns status code 503
*/
func GetStatusApi(url string) int {
	res, err := http.Get(url)
	if err != nil {
		return http.StatusServiceUnavailable
	}
	defer res.Body.Close()
	return res.StatusCode
}
