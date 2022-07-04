package models

type Artist struct {
	ConcertDates string   `json:"concertDates"`
	CreationDate int64    `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	ID           int64    `json:"id"`
	Image        string   `json:"image"`
	Locations    string   `json:"locations"`
	Members      []string `json:"members"`
	Name         string   `json:"name"`
	Relations    string   `json:"relations"`
}

type Relation struct {
	DatesLocations map[string][]string `json:"datesLocations"`
	ID             int64               `json:"id"`
}
