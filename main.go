package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/sheets/v4"
)

type Player struct {
	name string
	gender string
	obp string
	slg string
}


func main() {
	b, err := ioutil.ReadFile("credentials.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	config, err := google.ConfigFromJSON(b, "https://www.googleapis.com/auth/spreadsheets.readonly")
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	client := GetClient(config)

	srv, err := sheets.New(client)
	if err != nil {
		log.Fatalf("Unable to retrieve Sheets client: %v", err)
	}

	spreadsheetId := "1cWTd6_AdLr8rMe7-BLKP61SK0LsQPAbLqVq5KU_BdHA"
	readRange := "2019 League Summary Sheet!B3:T16"
	resp, err := srv.Spreadsheets.Values.Get(spreadsheetId, readRange).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve data from sheet: %v", err)
	}

	var playerList []interface{}

	if len(resp.Values) == 0 {
		fmt.Println("No data found.")
	} else {
		fmt.Println("Player, OBP, SLG:")
		for _, row := range resp.Values {
			playerList = append(playerList, Player{name: row[0].(string), obp: row[17].(string), slg: row[18].(string)})
		}
		fmt.Println(">>>>>>>>> ", playerList)
	}
}