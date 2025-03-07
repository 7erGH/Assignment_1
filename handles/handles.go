package handles

import (
	"encoding/json"
	"net/http"
	"log"
	"time"

	"countryinfo/services"
)

// Manages http requests of country information
type Handler struct {
	countryService *services.CountryService	// fetches country data
	startTime		time.Time				// starts time for status check
}

func NewHandler(countryService *services.CountryService) *Handler {
	return &Handler{
		countryService: countryService,
		startTime:		time.Now(),
	}
}

// RespondWithJSON sends response with status
func (h *Handler) respondWithJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Printf("Failed to encode response: %v", err)
	}
}

//checks if the request method matches the expected
func (h *Handler) validateMethod(w http.ResponseWriter, r *http.Request, method string) bool{
	if r.Method != method {
		h.respondWithJSON(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return false
	}
	return true
}

