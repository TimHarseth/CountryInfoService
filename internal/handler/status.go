package handler

import (
	"assignment-1/internal/api"
	"assignment-1/internal/constants"
	"assignment-1/internal/models"
	"encoding/json"
	"net/http"
	"time"
)

var Starttime time.Time

func GetStatus(w http.ResponseWriter, r *http.Request) {
	countriesNowStatus := api.GetStatusApi(constants.CountriesnowHealthApiPath)
	restCountriesStatus := api.GetStatusApi(constants.RestcountriesHealthApiPath)

	Status := models.Status{
		CountriesNowStatus:  countriesNowStatus,
		RestCountriesStatus: restCountriesStatus,
		Version:             "v1",
		Uptime:              int(time.Since(Starttime).Seconds()),
	}

	// Write the combined information as a JSON response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	out, _ := json.Marshal(Status)
	w.Write(out)
}
