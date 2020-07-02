package players

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

type PlayerStore interface {
	GetStats() []Player
}

type Server struct {
	store PlayerStore
}

func NewServer(store PlayerStore) *Server {
	return &Server{
		store: store,
	}
}

func (s Server) ReturnPlayersStats(w http.ResponseWriter, _ *http.Request) {
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
