package routes

import (
	"context"
	"encoding/json"
	"net/http"
	"session-logger/app/db"
	"session-logger/app/event"
	"session-logger/app/models"
	"time"

	"github.com/gorilla/mux"
)

func LogEvent(w http.ResponseWriter, r *http.Request) {
	var logEntry models.SessionLog
	err := json.NewDecoder(r.Body).Decode(&logEntry)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	logEntry.Timestamp = time.Now()

	_, err = db.LogsCollection.InsertOne(context.Background(), logEntry)
	if err != nil {
		http.Error(w, "Failed to save log", http.StatusInternalServerError)
		return
	}

	event.SendEventToKafka(logEntry.Username + " performed " + logEntry.Action)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"msg": "Log saved successfully"})
}

func RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/log", LogEvent).Methods("POST")
}
