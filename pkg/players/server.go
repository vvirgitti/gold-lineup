package players

import (
	"bytes"
	json2 "encoding/json"
	"io"
	"log"
	"net/http"
)

type Player struct {
	Name string
	Gender string
	Obp string
	Slg string
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

	json, err := json2.Marshal(players)
	if err != nil {
		log.Fatalf("couldn't marshall into json %v", err)
	}

	playerData := bytes.NewReader(json)

	io.Copy(w, playerData)
}


