package google

import (
	"context"
	"golang.org/x/oauth2/google"
	"io/ioutil"
	"log"
	"net/http"
)


func GoogleClient() (*http.Client){
	ctx := context.Background()
	b, err := ioutil.ReadFile("service_account.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	config, err := google.JWTConfigFromJSON(b, "https://www.googleapis.com/auth/spreadsheets.readonly")
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}

	return config.Client(ctx)
}
