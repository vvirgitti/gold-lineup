package main

import (
	"github.com/vvirgitti/gold-lineup/pkg/config"
	"github.com/vvirgitti/gold-lineup/pkg/google"
	"github.com/vvirgitti/gold-lineup/pkg/handlers"
	"github.com/vvirgitti/gold-lineup/pkg/store"
	"log"
	"net/http"
)

func main() {
	conf := config.NewConfig()
	playerStore := store.NewStore(conf, google.GoogleClient())
	server := handlers.NewServer(playerStore)

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", server.HomeHandler)
	http.HandleFunc("/stats", server.StatsHandler)

	log.Println("Client is running on port 1000")
	err := http.ListenAndServe(":1000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
