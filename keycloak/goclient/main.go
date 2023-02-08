package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	oidc "github.com/coreos/go-oidc"
	"golang.org/x/oauth2"
)

func main() {
	// context é uma biblioteca que permite a gente pare solicitações quando a gente quiser.
	ctx := context.Background()
	provider, err := oidc.NewProvider(ctx, "http://localhost:8080/auth/realms/MyRealmTest")
	if err != nil {
		log.Fatal(err)
	}
	config := oauth2.Config{
		ClientID:     "myclient",
		ClientSecret: "LHR32AH5t7fw4EZ6pwA3m1xkNgLSg8aa",
		RedirectURL:  "http://localhost:8081/auth/callback",
		Endpoint:     provider.Endpoint(),
		Scopes:       []string{oidc.ScopeOpenID, "profile", "email"},
	}
	state := "abc123"

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, config.AuthCodeURL(state), http.StatusFound)
	})

	http.HandleFunc("/auth/callback", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("state") != state {
			http.Error(w, "State did not match", http.StatusBadRequest)
			return
		}
		token, err := config.Exchange(ctx, r.URL.Query().Get("code"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		rawIDToken, ok := token.Extra("id_token").(string)
		if !ok {
			http.Error(w, "No id_token field in oauth2 token.", http.StatusInternalServerError)
		}

		userInfo, err := provider.UserInfo(ctx, oauth2.StaticTokenSource(token))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		resp := struct {
			AccessToken *oauth2.Token
			RawIdToken  string
			UserInfo    *oidc.UserInfo
		}{
			AccessToken: token,
			RawIdToken:  rawIDToken,
			UserInfo:    userInfo,
		}
		data, err := json.MarshalIndent(resp, "", "  ")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Verifica se o token é válido.
		_, err = provider.Verifier(&oidc.Config{ClientID: config.ClientID}).Verify(ctx, rawIDToken)
		if err != nil {
			log.Fatal(err)
		}
		w.Write(data)
	})

	log.Fatal(http.ListenAndServe(":8081", nil))
}
