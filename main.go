package main

import (
    "encoding/json"
    "log"
    "net/http"
    "time"
)

type TimeResponse struct {
    Time string `json:"time"`
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
    currentTime := time.Now().Format(time.RFC3339)
    response := TimeResponse{Time: currentTime}

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}

func main() {
    http.HandleFunc("/time", timeHandler)
    log.Println("Starting server on :8795...")
    if err := http.ListenAndServe(":8795", nil); err != nil {
        log.Fatal(err)
    }
}
