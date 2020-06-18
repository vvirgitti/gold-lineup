package players

import (
	"encoding/base64"
	"github.com/vvirgitti/gold-lineup/pkg/google"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"time"
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

func (s Server) Authorize(w http.ResponseWriter, r *http.Request) {
	oauthState := generateOAuthCookie(w)
	u := google.OauthConfig().AuthCodeURL(oauthState)
	http.Redirect(w, r, u, http.StatusTemporaryRedirect)
}

func generateOAuthCookie(w http.ResponseWriter) string {
	var expiration = time.Now().Add(365 * 24 * time.Hour)

	b := make([]byte, 16)
	rand.Read(b)
	state := base64.URLEncoding.EncodeToString(b)
	cookie := http.Cookie{Name: "oauthstate", Value: state, Expires: expiration}
	http.SetCookie(w, &cookie)

	return state
}

func (s Server) Token(w http.ResponseWriter, r *http.Request)  {
	oauthState, _ := r.Cookie("oauthstate")

	if r.FormValue("state") != oauthState.Value {
		log.Println("invalid oauth google state")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}


}
