package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func joinRecordingRoom(w http.ResponseWriter, r *http.Request) {
	// TODO validate + return WS join token, TURN creds
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

func getRecordingRoomInfo(w http.ResponseWriter, r *http.Request) {
	// TODO get room metadata (name, participants, host)
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

func createNewRecordingRoom(w http.ResponseWriter, r *http.Request) {
	// TODO create a new recording room
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
