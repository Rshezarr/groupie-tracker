package storage

import "gr/internal/models"

var (
	Storage = make(map[int]models.Artist)
	Artists = []models.Artist{}
)

func SaveInStorage(artists []models.Artist) {
	for id, artist := range artists {
		Storage[id+1] = artist
		Artists = append(Artists, artist)
	}
}
