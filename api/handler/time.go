package handler

import (
	"fmt"
	"net/http"

	"github.com/sithumonline/red-timer/transact/time"
)

//var mainTime *time.Time

// GetTime
// @Summary Get time diff
// @Tags Puso
// @Accept json
// @Produce json
// @Success 200 {object} time.Time
// @Failure 404 {string} string
// @Router	/time	[get]
func GetTime(w http.ResponseWriter, r *http.Request) {
	t := time.Time{}

	o := t.GetTime()
	fmt.Println(o)
	RespondWithJSON(w, http.StatusOK, &o)
}
