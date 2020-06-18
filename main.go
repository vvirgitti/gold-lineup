package main

import (
	"github.com/vvirgitti/gold-lineup/pkg/config"
	"github.com/vvirgitti/gold-lineup/pkg/players"
	"github.com/vvirgitti/gold-lineup/pkg/store"
	"log"
	"net/http"
)

func main() {
	conf := config.NewConfig()

	playerStore := store.NewStore(conf)
	server := players.NewServer(playerStore)

	http.HandleFunc("/", server.ReturnPlayersStats)

	http.HandleFunc("/auth/google/login", server.Authorize)
	http.HandleFunc("/auth/google/token", server.Token)

	log.Println("Client is running on port 1000")
	http.ListenAndServe(":1000", nil)
}