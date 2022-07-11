package models

type Relations struct {
	Concerts []Concert `json:"index"`
}

type Concert struct {
	DatesLocations map[string][]string `json:"datesLocations"`
}

type Artist struct {
	ID             int      `json:"id"`
	Image          string   `json:"image"`
	Name           string   `json:"name"`
	Members        []string `json:"members"`
	CreationDate   int      `json:"creationDate"`
	FirstAlbum     string   `json:"firstAlbum"`
	DatesLocations map[string][]string
}

type Error struct {
	ErrorCode string
	ErrorText string
}
