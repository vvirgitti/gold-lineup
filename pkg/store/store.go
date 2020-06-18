package store

import (
	"context"
	"github.com/vvirgitti/gold-lineup/pkg/config"
	"github.com/vvirgitti/gold-lineup/pkg/players"
	"golang.org/x/oauth2/google"
	_ "google.golang.org/api/compute/v1"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
	"io/ioutil"
	"log"
)

type PlayerStore struct {
	config      config.Config
}

func NewStore(config config.Config) *PlayerStore {
	return &PlayerStore{
		config:      config,
	}
}

func (ps PlayerStore) GetStats() []players.Player {
	ctx := context.Background()
	b, err := ioutil.ReadFile("service_account.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	config, err := google.JWTConfigFromJSON(b, "https://www.googleapis.com/auth/spreadsheets.readonly")
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}

	client := config.Client(ctx)

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

		var playersList []players.Player

		if len(resp.Values) == 0 {
			log.Fatalf("data not found")
		}

		for _, value := range resp.Values {
			player := players.Player{
				Name:   value[0].(string),
				Gender: value[1].(string),
				Obp:    value[18].(string),
				Slg:    value[19].(string),
			}
			playersList = append(playersList, player)
		}

		return playersList
}