package delivery

import (
	"strconv"
	"strings"

	"groupie-tracker/internal/models"
)

func (h *Handler) artistSearch(inputFromSearch string) []models.Artist {
	var searchResult []models.Artist
	bandInfo := strings.Split(inputFromSearch, " - ")
	var isFound bool

	for i := range h.storage.Artist {
		// Search by name
		if containsCaseInsensitive(h.storage.Artist[i].Name, bandInfo[0]) {
			searchResult = append(searchResult, h.storage.Artist[i])
			isFound = true
			continue
		}

		// Search by members
		for _, member := range h.storage.Artist[i].Members {
			if containsCaseInsensitive(member, bandInfo[0]) {
				searchResult = append(searchResult, h.storage.Artist[i])
				isFound = true
				break
			}
		}

		// Search by first album
		if containsCaseInsensitive(h.storage.Artist[i].FirstAlbum, bandInfo[0]) {
			searchResult = append(searchResult, h.storage.Artist[i])
			isFound = true
			continue
		}

		// Search by creation date
		creationDate := strconv.Itoa(h.storage.Artist[i].CreationDate)
		if creationDate == bandInfo[0] {
			searchResult = append(searchResult, h.storage.Artist[i])
			isFound = true
			continue
		}

		// Search by locations and dates
		for city, dates := range h.storage.Artist[i].DatesLocations {
			if containsCaseInsensitive(city, bandInfo[0]) {
				searchResult = append(searchResult, h.storage.Artist[i])
				isFound = true
				continue
			}

			for _, date := range dates {
				if containsCaseInsensitive(date, bandInfo[0]) {
					searchResult = append(searchResult, h.storage.Artist[i])
					isFound = true
					break
				}
			}
		}
	}

	if isFound {
		return searchResult
	}

	return nil
}

func containsCaseInsensitive(a string, b string) bool {
	return strings.Contains(
		strings.ToLower(a),
		strings.ToLower(b),
	)
}
