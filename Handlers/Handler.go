package groupie

import (
	"encoding/json"
	"log"
	"net/http"
	"io"
	groupie "groupie/data"
)

func Handler(url string, data []groupie.Band) []groupie.Band {
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err.Error())
	}
	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(responseData, &data)
	return data
}
