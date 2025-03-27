package mpm_server

import (
	"fmt"
	"log"
	"net/http"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	oauth2Config = oauth2.Config{
		ClientID:     "50701837760-u5l4uqdesfd7dnagnhrfjph9d1l3dlji.apps.googleusercontent.com", // Set your GitHub OAuth client ID
		ClientSecret: "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",                                     // Set your GitHub OAuth client secret
		RedirectURL:  "http://localhost:8080/callback",                                          // This must match the one you provided in GitHub OAuth
		Scopes:       []string{"user:email"},                                                    // Define the scope (what info to request)
		Endpoint:     google.Endpoint,                                                           // GitHub's OAuth2 endpoint
	}
	oauth2StateString = "random-state"
)

type mpm_server struct {
	certFile   string
	privateKey string
}

func (s mpm_server) test_handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"message": "grias de!"}`)
}

func (s mpm_server) handleLogin(w http.ResponseWriter, r *http.Request) {
	url := oauth2Config.AuthCodeURL(oauth2StateString, oauth2.AccessTypeOffline)
	http.Redirect(w, r, url, http.StatusFound)
}

func (s mpm_server) API_setup() {
	http.HandleFunc("/api", s.test_handler)
	http.HandleFunc("/login", s.handleLogin)
	log.Print("Server started at https://localhost:8080")
	if err := http.ListenAndServeTLS(":8080", s.certFile, s.privateKey, nil); err != nil {
		log.Fatalf("ListenAndServeTLS failec: %v", err)
	}
}

func NewMPMserver(certFile string, privateKey string) *mpm_server {
	server_inst := mpm_server{certFile, privateKey}
	return &server_inst

}
