package groupie

import (
	"encoding/json"
	"net/http"
)

// FetchHandler fetches data from a given URL and decodes it into the provided data structure
func FetchHandler(url string, data interface{}, id string, w http.ResponseWriter, r *http.Request) {
	response, err := http.Get(url + id)
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError, "Internal Server Error")
		return
	}
	defer response.Body.Close()

	// Decode JSON response into the provided data structure
	err = json.NewDecoder(response.Body).Decode(data)
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError, "Failed to decode JSON")
		return
	}
}