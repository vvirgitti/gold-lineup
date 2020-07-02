package handlers

import (
	"html/template"
	"log"
	"net/http"
)

type Player struct {
	Name   string
	Gender string
	Obp    string
	Slg    string
}

type Store interface {
	GetStats() []Player
}

type Server struct {
	store Store
}

func NewServer(store Store) *Server {
	return &Server{
		store: store,
	}
}

func (s Server) StatsHandler(w http.ResponseWriter, _ *http.Request) {
	players := s.store.GetStats()

	files := []string{"./frontend/stats.tmpl", "./frontend/layout.tmpl", "./frontend/header.tmpl"}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	err = ts.Execute(w, players)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}

}

func (s Server) HomeHandler(w http.ResponseWriter, _ *http.Request)  {
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
