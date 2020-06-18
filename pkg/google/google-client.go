package google

import (
	"context"
	"github.com/vvirgitti/gold-lineup/pkg/config"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"log"
	"net/http"
)

func OauthConfig() *oauth2.Config{
	conf := config.NewConfig()
	scope := []string{"https://www.googleapis.com/auth/spreadsheets"}

	return &oauth2.Config{
		ClientID:     conf.GoogleClientId,
		ClientSecret: conf.GoogleClientSecret,
		RedirectURL:  "http://localhost:1000/authentication/google/token",
		Scopes:       scope,
		Endpoint: google.Endpoint,
	}
}


func GoogleClient(code string) *http.Client{
	conf := OauthConfig()
	ctx := context.Background()
	tok, err := conf.Exchange(ctx, "authorization-code")
	if err != nil {
		log.Fatal("Unable to get token from OAuth", err)
	}

	client := conf.Client(ctx, tok)
	return client
}
