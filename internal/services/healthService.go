package services

import (
	"encoding/json"
	"github.com/arpit006/go_txns/internal/models"
	"net/http"
	"time"
)

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	health := models.Health{
		App: "GO TXN app",
		Code: 202,
		Time: time.Now(),
	}

	healthStr, err := json.Marshal(health)
	if err != nil {

	}
	// write to response
	_, err = w.Write(healthStr)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusAccepted)
}
