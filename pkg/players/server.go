package players

import (
	"html/template"
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

	t := template.Must(template.ParseFiles("frontend/index.gohtml"))

	t.Execute(w, players)

}
