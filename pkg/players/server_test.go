package players_test

import (
	"encoding/json"
	"github.com/vvirgitti/gold-lineup/pkg/players"
	"net/http"
	"net/http/httptest"
	"testing"
)

type stubStore struct {
	players []players.Player
}

func (s stubStore) GetStats() []players.Player  {
	return s.players
}

func TestPlayerServer(t *testing.T) {
	t.Run("returns 200", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		store := stubStore{}

		server := players.NewServer(store)
		server.ReturnPlayersStats(response, req)

		if response.Code != 200 {
			t.Errorf("expected 200, got %d", response.Code)
		}
	})

	t.Run("returns the stats for all players", func(t *testing.T) {
		store := stubStore{players: []players.Player{{"Sawamura", "Male", 0.10, 0.10}, {"Furuya", "Male", 0.12, 0.12}}}

		req, _ := http.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		server := players.NewServer(store)
		server.ReturnPlayersStats(response, req)

		var newPlayers []players.Player

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

