package confirm

import (
	"os"
)

// confirm holds configuration values
type Confirm struct{
	Port					string
	RestCountriesBaseURL	string
	CountriesNowBaseURL		string
}

// Reconfirm starts up and returns a confirm instance
func Reconfirm() *Confirm {
	cfg := &Confirm{}
	cfg.Port=							getEnv("port", "8080")
	cfg.RestCountriesBaseURL= 			getEnv("REST_COUNTRIES_URL","http://129.241.150.113:8080/v3.1")
	cfg.CountriesNowBaseURL=			getEnv("COUNTRIES_NOW_URL","http://129.241.150.113:3500/api/v0.1")
	
	return cfg
}
// Function to get environment variable values
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
