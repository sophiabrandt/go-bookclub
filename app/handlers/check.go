package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sophiabrandt/go-bookclub/business/logger"
)

type checkGroup struct {
	logger *logger.Logger
}

func (cg checkGroup) readiness(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	status := struct {
		Status string
	}{
		Status: "OK",
	}
	json.NewEncoder(w).Encode(status)
	cg.logger.Info(fmt.Sprint(status))

	return nil
}
