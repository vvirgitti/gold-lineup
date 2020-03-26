package main

import (
	"github.com/vvirgitti/gold-lineup/pkg/players"
	store "github.com/vvirgitti/gold-lineup/pkg/store"
	"net/http"
)

func main() {
	store := store.NewStore()
	server := players.NewServer(store)

	http.HandleFunc("/", server.ReturnPlayersStats)
	http.ListenAndServe(":1000", nil)
}