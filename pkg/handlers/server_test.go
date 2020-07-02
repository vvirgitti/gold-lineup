package handlers_test

import (
	"encoding/json"
	"github.com/vvirgitti/gold-lineup/pkg/handlers"
	"net/http"
	"net/http/httptest"
	"testing"
)

type stubStore struct {
	players []handlers.Player
}

func (s stubStore) GetStats() []handlers.Player  {
	return s.players
}

func TestPlayerServer(t *testing.T) {
	t.Run("returns 200", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		store := stubStore{}

		server := handlers.NewServer(store)
		server.StatsHandler(response, req)

		if response.Code != 200 {
			t.Errorf("expected 200, got %d", response.Code)
		}
	})

	t.Run("returns the stats for all players", func(t *testing.T) {
		store := stubStore{players: []handlers.Player{{"Sawamura", "Male", "0.10", "0.10"}, {"Furuya", "Male", "0.12", "0.12"}}}

		req, _ := http.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		server := handlers.NewServer(store)
		server.StatsHandler(response, req)

		var newPlayers []handlers.Player

		err := json.NewDecoder(response.Body).Decode(&newPlayers)
		if err != nil {
			t.Errorf("errored while decoding json")
		}

		if len(newPlayers) != 2 {
			t.Errorf("expected 2, got %d", len(newPlayers))
		}

		if newPlayers[0].Name != "Sawamura" {
			t.Errorf("expected Sawamura, got %s", newPlayers[0].Name)
		}
	})
}

