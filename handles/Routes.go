package handles

import (
	"countryinfo/models"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// Retrives detailed country info
func (h *Handler) HandleCountryInfo(w http.ResponseWriter, r *http.Request) {
	if !h.validateMethod(w, r, http.MethodGet) {
		return
	}

	// Takes country code from url
	countryCode := strings.TrimPrefix(r.URL.Path, "/countryinfo/v1/info/")
	if countryCode == ""{
		h.respondWithJSON(w, http.StatusBadRequest, map[string]string{"error": "country code is required"})
		return
	}

	// Parse Limit query parameter
	limitStr := r.URL.Query().Get("Limit")
	limit := 0
	if limitStr != "" {
		if parsedLimit, err := strconv.Atoi(limitStr); err != nil || parsedLimit < 0 {
			h.respondWithJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid limit parameter"})
			return
		} else {
			limit = parsedLimit
		}
	}

	// Fetch country info
	countryInfo, err := h.countryService.GetCountryInfo(countryCode)
	if err != nil {
		h.respondWithJSON(w, http.StatusNotFound, map[string]string{"error": err.Error()})
		return
	}

	// Fetch cities for selected country with optional limit
	cities, err := h.countryService.GetCities(countryInfo.Name.Common, limit)
	if err != nil {
		h.respondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "failed to get cities"})
		return
	}
	// Construct response
	response := models.CountryInfo{
		Name:		countryInfo.Name.Common,
		Continents:	countryInfo.Continents,
		Population:	countryInfo.Population,
		Languages:	countryInfo.Languages,
		Borders:	countryInfo.Borders,
		Flag:		countryInfo.Flags.PNG,
		Capital:	countryInfo.Capital[0],
		Cities:		cities, 
	}

	h.respondWithJSON(w, http.StatusOK, response)
}

// Retrives population and calculates mean population
func (h *Handler) HandlePopulation(w http.ResponseWriter, r *http.Request) {
	if !h.validateMethod(w, r, http.MethodGet) {
		return
	}

	// Fetches country name from url
	countryName := strings.TrimPrefix(r.URL.Path, "/countryinfo/v1/population/")
	if countryName == "" {
		h.respondWithJSON(w, http.StatusBadRequest, map[string]string{"error": "country name is required"})
		return
	}

	// Fetch population data
	populationData, err := h.countryService.GetPopulation(countryName)
	if err != nil {
		h.respondWithJSON(w, http.StatusNotFound, map[string]string{"error": err.Error()})
		return
	}

	// If no data is found mean returns 0
	if len(populationData) == 0 {
		h.respondWithJSON(w, http.StatusOK, models.PopulationData{Mean: 0, Values: populationData})
		return
	}

	// calculates mean population
	var sum int
	for _, p := range populationData {
		sum += p.Value
	}
	mean := sum / len(populationData)

	response := models.PopulationData{
		Mean:	mean,
		Values:	populationData,
	}

	h.respondWithJSON(w, http.StatusOK, response)
}

// Checks health status for external services
func (h *Handler) HandleStatus(w http.ResponseWriter, r *http.Request) {
	if !h.validateMethod(w, r, http.MethodGet) {
		return
	}

	restCountriesStatus := http.StatusOK
	if _, err := h.countryService.GetCountryInfo("NOR"); err != nil{
		restCountriesStatus = http.StatusServiceUnavailable
	}

	countriesNowStatus := http.StatusOK
	if _, err := h.countryService.GetCities("Norway", 0); err != nil{
		countriesNowStatus = http.StatusServiceUnavailable
	}

	response := models.StatusUpdate{
		CountriesNowAPI:	countriesNowStatus,
		RestCountriesAPI:	restCountriesStatus,
		Version:			"v1",
		Uptime:				time.Since(h.startTime).Milliseconds(),
	}

	h.respondWithJSON(w, http.StatusOK, response)
}