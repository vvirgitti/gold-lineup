package main

import (
	"github.com/vvirgitti/gold-lineup/pkg/config"
	"github.com/vvirgitti/gold-lineup/pkg/google"
	"github.com/vvirgitti/gold-lineup/pkg/players"
	"github.com/vvirgitti/gold-lineup/pkg/store"
	"html/template"
	"log"
	"net/http"
)

func main() {
	conf := config.NewConfig()
	playerStore := store.NewStore(conf, google.GoogleClient())

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", homeHandler)

	server := players.NewServer(playerStore)
	http.HandleFunc("/stats", server.ReturnPlayersStats)

	log.Println("Client is running on port 1000")
	err := http.ListenAndServe(":1000", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func homeHandler(w http.ResponseWriter, _ *http.Request)  {
	files := []string{"./frontend/home.tmpl", "./frontend/layout.tmpl", "./frontend/header.tmpl"}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}