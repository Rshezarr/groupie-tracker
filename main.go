package main

import (
	"log"
	"net/http"
	"time"

	"gr/internal"
)

func main() {
	// tmpl, _ := template.ParseFiles("templates/index.html")
	// http.HandleFunc("/", internal.HomePage)
	// fmt.Printf("Server goes brrrr\n")
	// log.Fatal(http.ListenAndServe(":8090", nil))

	router := http.NewServeMux()
	server := &http.Server{
		Addr:           ":8090",
		Handler:        router,
		WriteTimeout:   10 * time.Second,
		ReadTimeout:    10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	handler := internal.NewHandler()
	router.HandleFunc("/", handler.HomePage)
	router.HandleFunc("/artist/", handler.GetArtistById)
	router.HandleFunc("/relation/", handler.GetRelationsByArtist)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Server listen and serve error: %v\n", err)
	}
}
