package handler

import (
	"net/http"
	"os"
)

// GetHealth godoc
// @Summary Returns health of the service
// @Router /healthz [get]
func GetHealth(w http.ResponseWriter, r *http.Request) {
	RespondWithJSON(w, http.StatusOK, "Go Puso up and running version: "+os.Getenv("VERSION"))
}
