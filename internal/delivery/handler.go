package delivery

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"

	"groupie-tracker/internal/models"
)

const (
	artistPath     = "/artists/"
	templateArtist = "artist.html"
	templateIndex  = "index.html"
	templateError  = "error.html"
)

type Handler struct {
	tmp     *template.Template
	storage *models.Store
}

func NewHandler(templatePath string) (*Handler, error) {
	tmpl, err := template.ParseGlob(templatePath)
	if err != nil {
		return nil, err
	}
	return &Handler{
		tmp:     tmpl,
		storage: &models.Store{},
	}, nil
}

func (h *Handler) MainPage(page http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		h.errorPage(page, http.StatusMethodNotAllowed, nil)
		return
	}

	if r.URL.Path != "/" {
		h.errorPage(page, http.StatusNotFound, nil)
		return
	}

	if err := h.jsonData(); err != nil {
		h.errorPage(page, http.StatusInternalServerError, err)
		return
	}

	if err := h.tmp.ExecuteTemplate(page, templateIndex, h.storage.Artist); err != nil {
		http.Error(page, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}

func (h *Handler) getIDFromURL(url, prefix string) (int, error) {
	return strconv.Atoi(strings.TrimPrefix(url, prefix))
}

func (h *Handler) ArtistPage(page http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		h.errorPage(page, http.StatusMethodNotAllowed, nil)
		return
	}

	if err := h.jsonData(); err != nil {
		h.errorPage(page, http.StatusInternalServerError, err)
		return
	}

	id, err := h.getIDFromURL(r.URL.Path, artistPath)
	if err != nil {

		log.Println(err)
		h.errorPage(page, http.StatusBadRequest, err)
		return
	}

	artist, err := h.storage.GetArtistByID(id)
	if err != nil {
		h.errorPage(page, http.StatusNotFound, err)
		return
	}

	if err := h.tmp.ExecuteTemplate(page, templateArtist, artist); err != nil {
		http.Error(page, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}

func (h *Handler) FilterPage(page http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		h.errorPage(page, http.StatusMethodNotAllowed, nil)
		return
	}

	r.ParseForm()

	if r.Form.Has("search_input") {
		newData := h.artistSearch(r.Form.Get("search_input"))

		if err := h.tmp.ExecuteTemplate(page, templateIndex, newData); err != nil {
			http.Error(page, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
	}
}

func (h *Handler) errorPage(page http.ResponseWriter, status int, err error) {
	errStruct := models.Error{
		ErrorCode: status,
		ErrorText: http.StatusText(status),
	}
	page.WriteHeader(status)
	if err := h.tmp.ExecuteTemplate(page, templateError, errStruct); err != nil {
		http.Error(page, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}
