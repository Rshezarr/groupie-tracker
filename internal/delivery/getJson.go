package delivery

import (
	"encoding/json"
	"net/http"
	"time"
)

const (
	artistURL   = "https://groupietrackers.herokuapp.com/api/artists"
	relationURL = "https://groupietrackers.herokuapp.com/api/relation"
	timeout     = 30 * time.Second
)

func getJson(url string, result interface{}) error {
	client := &http.Client{
		Timeout: timeout,
	}

	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	response, err := client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	return json.NewDecoder(response.Body).Decode(result)
}

func (h *Handler) jsonData() error {
	if len(h.storage.Artist) != 0 {
		return nil
	}

	if err := getJson(artistURL, &h.storage.Artist); err != nil {
		return err
	}

	if err := getJson(relationURL, &h.storage.Relations); err != nil {
		return err
	}

	for i, relation := range h.storage.Relations.Concerts {
		h.storage.Artist[i].DatesLocations = relation.DatesLocations
	}
	return nil
}
