package delivery

import "net/http"

func (h *Handler) SetEndpoints(mux *http.ServeMux) {
	mux.HandleFunc("/", h.MainPage)
	mux.HandleFunc("/artists/", h.ArtistPage)
	mux.HandleFunc("/filter/", h.FilterPage)

	mux.Handle("/ui/css/", http.StripPrefix("/ui/css/", http.FileServer(http.Dir("./ui/css/"))))
	mux.Handle("/ui/assets/", http.StripPrefix("/ui/assets/", http.FileServer(http.Dir("./ui/assets/"))))
	mux.Handle("/ui/js/", http.StripPrefix("/ui/js/", http.FileServer(http.Dir("./ui/js/"))))
}
