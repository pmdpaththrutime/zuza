package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	io.WriteString(w, "This is my website!\n")
}

func getHello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /hello request\n")
	io.WriteString(w, "Hello, HTTP!\n")
}

func getRecordingRoom(w http.ResponseWriter, r *http.Request) {
	room_id := r.PathValue("room")
	if room_id != "" {
		fmt.Fprintf(w, "You are in recording room %s\n", room_id)
	} else {
		// Unreachable
		http.Error(w, "A room ID is required", http.StatusBadRequest)
	}
}

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
	mux.HandleFunc("/", getRoot)
	mux.HandleFunc("/hello", getHello)
	mux.HandleFunc("/health", healthCheck)
	mux.HandleFunc("/record/{room}", getRecordingRoom)

	err := http.ListenAndServe(":3333", mux)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
