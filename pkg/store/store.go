package store

import (
	"context"
	"github.com/vvirgitti/gold-lineup/pkg/config"
	"github.com/vvirgitti/gold-lineup/pkg/handlers"
	_ "google.golang.org/api/compute/v1"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
	"log"
	"net/http"
)

type PlayerStore struct {
	config       config.Config
	googleClient *http.Client
}

func NewStore(config config.Config, googleClient *http.Client) *PlayerStore {
	return &PlayerStore{
		config:       config,
		googleClient: googleClient,
	}
}

func (ps PlayerStore) GetStats() []handlers.Player {
	ctx := context.Background()
	client := ps.googleClient

	srv, err := sheets.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Fatal("Unable to create a Google sheet client", err)
	}

	spreadsheetId := "1lyWJfbw2-2AKXTpkjlsjaMTRu4NqNJQA4DrPIEiM7Mw"
	readRange := "stats!B3:U16"
	resp, err := srv.Spreadsheets.Values.Get(spreadsheetId, readRange).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve data from sheet: %v", err)
	}

	var playersList []handlers.Player

	if len(resp.Values) == 0 {
		log.Fatalf("data not found")
	}

	for _, value := range resp.Values {
		player := handlers.Player{
			Name:   value[0].(string),
			Gender: value[1].(string),
			Obp:    value[18].(string),
			Slg:    value[19].(string),
		}
		playersList = append(playersList, player)
	}

	return playersList
}
