package models

import "errors"

func (s *Store) GetArtistByID(id int) (Artist, error) {
	if id < 1 || id > len(s.Artist) {
		return Artist{}, errors.New("not found")
	}
	return s.Artist[id-1], nil
}
