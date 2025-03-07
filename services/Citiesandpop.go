package services

import (
	"bytes"
	"countryinfo/models"
	"fmt"
	"encoding/json"
	"net/http"
	"sort"
)

const (
	CountriesNowBaseURL = "http://129.241.150.113:3500/api/v0.1/countries"
)

// Retrives a sorted list of cities for a given country
func (s *CountryService) GetCities(countryName string, limit int) ([]string, error) {
	var citiesResp models.CountriesNCities
	err := s.makePostRequest(fmt.Sprintf("%s/cities", CountriesNowBaseURL), models.CountriesNow{Country: countryName}, &citiesResp)
	if err != nil {
		return nil, err
	}

	if citiesResp.Error {
		return nil, fmt.Errorf("cities API error: %s", citiesResp.Msg)
	}

	cities := citiesResp.Data
	sort.Strings(cities)

	if limit > 0 && limit < len(cities) {
		cities = cities[:limit]
	}

	return cities, nil
}

// fetches population data for a given country
func (s *CountryService) GetPopulation(countryName string) ([]models.PopulationValue, error) {
	var popResp models.CountriesNPopulation
	err := s.makePostRequest(fmt.Sprintf("%s/population", CountriesNowBaseURL), models.CountriesNow{Country: countryName}, &popResp)
	if err != nil{
		return nil, err
	}

	if popResp.Error {
		return nil, fmt. Errorf("population API error: %s", popResp.Msg)
	}

	return popResp.Data.Population, nil
}

// Sends a Post request with JSON
func (s *CountryService) makePostRequest(url string, body interface{}, target interface{}) error {
	jsonData, err := json.Marshal(body)
	if err != nil {
		return err
	}

	resp, err := s.client.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil{
		return err
	}
	 defer resp.Body.Close()

	 if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("request failed with status: %d", resp.StatusCode)
	 }

	 return json.NewDecoder(resp.Body).Decode(target)
}