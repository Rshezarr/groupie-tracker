package internal

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"gr/internal/models"
	"gr/internal/storage"
)

// var client *http.Client

type Handler struct {
	counter int
	client  *http.Client
}

func NewHandler() *Handler {
	return &Handler{
		counter: 0,
		client:  &http.Client{},
	}
}

func (h *Handler) GetJson(url string, val interface{}) error {
	resp, err := h.client.Get(url)
	if err != nil {
		log.Fatalln(errors.New("WTF3"))
	}
	defer resp.Body.Close()
	return json.NewDecoder(resp.Body).Decode(val)
}

func (h *Handler) GetArtistInfo() []models.Artist {
	var artist []models.Artist

	err := h.GetJson("https://groupietrackers.herokuapp.com/api/artists", &artist)
	go storage.SaveInStorage(artist)
	if err != nil {
		log.Fatal(err)
	}
	return artist
}

func (h *Handler) HomePage(w http.ResponseWriter, r *http.Request) {
	// tmpl, err := template.ParseFiles("./templates/index.html")
	// if err != nil {
	// 	log.Println(err)
	// }
	if r.URL.Path != "/" {
		w.WriteHeader(400)
		return
	}
	var artists []models.Artist
	if h.counter == 0 {
		artists = h.GetArtistInfo()
		h.counter++
	} else {
		artists = storage.Artists
	}
	fmt.Println(h.counter)
	// if err := tmpl.Execute(w, artists); err != nil {
	// 	log.Println(errors.New("execute"))
	// }
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(artists)
}

func (h *Handler) GetArtistById(w http.ResponseWriter, r *http.Request) {
	id_str := strings.TrimPrefix(r.URL.Path, "/artist/")
	id, err := strconv.Atoi(id_str)
	if err != nil {
		w.WriteHeader(400)
		return
	}
	artist, ok := storage.Storage[id]
	if !ok {
		w.WriteHeader(400)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(artist)
}

func (h *Handler) GetRelationsByArtist(w http.ResponseWriter, r *http.Request) {
	id_str := strings.TrimPrefix(r.URL.Path, "/relation/")
	id, err := strconv.Atoi(id_str)
	if err != nil {
		w.WriteHeader(400)
		return
	}
	artist, ok := storage.Storage[id]
	if !ok {
		w.WriteHeader(400)
		return
	}
	var relation models.Relation
	if err := h.GetJson(artist.Relations, &relation); err != nil {
		w.WriteHeader(500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(relation)
}
