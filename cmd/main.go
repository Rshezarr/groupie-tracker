package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"grTracker/internal/delivery"
)

const port = "8090"

func main() {
	mux := http.NewServeMux()

	server := &http.Server{
		Addr:           ":" + port,
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	handler := delivery.NewHandler()

	mux.HandleFunc("/", handler.HandleMainPage)
	mux.HandleFunc("/artists/", handler.ArtistPage)
	mux.Handle("/templates/css/", http.StripPrefix("/templates/css/", http.FileServer(http.Dir("./templates/css/"))))
	mux.Handle("/templates/assets/", http.StripPrefix("/templates/assets/", http.FileServer(http.Dir("./templates/assets/"))))
	mux.Handle("/templates/js/", http.StripPrefix("/templates/js/", http.FileServer(http.Dir("./templates/js/"))))

	fmt.Printf("Starting server at port %v\nhttp://localhost%v\n", port, ":"+port)
	log.Fatal(server.ListenAndServe())
}
