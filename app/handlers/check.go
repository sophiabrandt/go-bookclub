package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sophiabrandt/go-bookclub/business/logger"
)

type checkGroup struct {
	logger *logger.Logger
}

func (cg checkGroup) readiness(w http.ResponseWriter, r *http.Request) {

	status := struct {
		Status string
	}{
		Status: "OK",
	}
	json.NewEncoder(w).Encode(status)
	cg.logger.Info(fmt.Sprint(status))
}
