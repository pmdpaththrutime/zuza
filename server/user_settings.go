package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func getCurrentUser(w http.ResponseWriter, r *http.Request) {
	// TODO Get the current user + settings
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

func updateCurrentUserSettings(w http.ResponseWriter, r *http.Request) {
	// TODO update the current user's settings
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
