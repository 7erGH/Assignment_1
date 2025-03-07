package services

import (
	"countryinfo/models"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	RestCountriesBaseURL = "http://129.241.150.113:8080/v3.1"
)

// Handles country related API requests
type CountryService struct {
	client *http.Client
}

func NewCountryService() *CountryService {
	return &CountryService{client: &http.Client{},}
}

// Retrieves country info from country code
func (s *CountryService) GetCountryInfo(countryCode string) (*models.RestCountries, error) {
	url := fmt.Sprintf("%s/alpha/%s", RestCountriesBaseURL, countryCode)
	var countryRecords []models.RestCountries

	if err := s.makeRequest(url, &countryRecords); err != nil {
		return nil, err
	}

	if len(countryRecords) == 0 {
		return nil, fmt.Errorf("no country data found")
	}

	return &countryRecords[0], nil
}

// Performs get requests and decodes JSON responses
func (s *CountryService) makeRequest(url string, target interface{}) error {
	resp, err := s.client.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK{
		return fmt.Errorf("request failed with status: %d", resp.StatusCode)
	}

	return json.NewDecoder(resp.Body).Decode(target)
}