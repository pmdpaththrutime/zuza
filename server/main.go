package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
)

// Health check - used by Docker Compose
func healthCheck(w http.ResponseWriter, r *http.Request) {
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

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", healthCheck)

	mux.HandleFunc("GET /auth/discord", getDiscordAuth)
	mux.HandleFunc("GET /auth/discord/callback", getDiscordAuthCallback)

	mux.HandleFunc("GET /api/me", getCurrentUser)
	mux.HandleFunc("PATCH /api/me/settings", updateCurrentUserSettings)

	mux.HandleFunc("POST /api/rooms/{room_code}/join", joinRecordingRoom)
	mux.HandleFunc("GET /api/rooms/{room_code}", getRecordingRoomInfo)
	mux.HandleFunc("POST /api/rooms", createNewRecordingRoom)

	err := http.ListenAndServe(":3333", mux)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
