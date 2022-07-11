package delivery

import (
	"errors"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
)

var tmpl, tmplErr = template.ParseGlob("templates/*.html")

var (
	ErrNotFound    = errors.New("page does not exist")
	ErrBadRequest  = errors.New("bad request")
	ErrWrongMethod = errors.New("method not allowed")
	ErrServer      = errors.New("internal server error")
)

type Handler struct {
	client *http.Client
}

func NewHandler() *Handler {
	return &Handler{
		client: &http.Client{},
	}
}

func (h *Handler) HandleMainPage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		ErrorHandler(w, http.StatusMethodNotAllowed, ErrWrongMethod)
		return
	}

	if r.URL.Path != "/" {
		ErrorHandler(w, http.StatusNotFound, ErrNotFound)
		return
	}

	if err := GetData(); err != nil {
		ErrorHandler(w, http.StatusInternalServerError, ErrServer)
		return
	}

	if tmplErr != nil {
		w.WriteHeader(500)
		http.Error(w, strconv.Itoa(http.StatusInternalServerError)+" "+http.StatusText(http.StatusInternalServerError), 500)
		return
	}

	if err := tmpl.ExecuteTemplate(w, "index.html", &Artist); err != nil {
		ErrorHandler(w, http.StatusInternalServerError, ErrServer)
		return
	}
}

func (h *Handler) ArtistPage(w http.ResponseWriter, r *http.Request) {
	if tmplErr != nil {
		ErrorHandler(w, http.StatusInternalServerError, ErrServer)
		return
	}

	idStr := strings.TrimPrefix(r.URL.Path, "/artists/")

	id, err := strconv.Atoi(idStr)
	if err != nil || id < 1 || id > len(Artist) {
		ErrorHandler(w, http.StatusNotFound, ErrNotFound)
		return
	}

	if err := tmpl.ExecuteTemplate(w, "artist.html", Artist[id-1]); err != nil {
		log.Println(err)
	}
}
