package delivery

import (
	"encoding/json"
	"net/http"

	"grTracker/internal/models"
)

const (
	artistsURL   = "https://groupietrackers.herokuapp.com/api/artists"
	relationsURL = "https://groupietrackers.herokuapp.com/api/relation"
)

var Artist []models.Artist

func getJson(url string, result interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(result)
}

func GetData() error {
	if len(Artist) != 0 {
		return nil
	}
	if err := getJson(artistsURL, &Artist); err != nil {
		return err
	}
	var relations models.Relations
	if err := getJson(relationsURL, &relations); err != nil {
		return err
	}
	for i, v := range relations.Concerts {
		Artist[i].DatesLocations = v.DatesLocations
	}
	return nil
}
