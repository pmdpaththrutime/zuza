package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func getDiscordAuth(w http.ResponseWriter, r *http.Request) {
	// TODO redirect to Discord
	resp := struct {
		Status string `json:"status"`
	}{
		Status: "healthy",
	}

	respBytes, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Failed to marshal response: %v", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(respBytes)
}

func getDiscordAuthCallback(w http.ResponseWriter, r *http.Request) {
	// TODO sets session cookie, redirects to `continue`
	resp := struct {
		Status string `json:"status"`
	}{
		Status: "healthy",
	}

	respBytes, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Failed to marshal response: %v", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(respBytes)
}
