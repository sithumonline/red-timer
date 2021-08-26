package handler

import (
	"fmt"
	"net/http"

	"github.com/sithumonline/red-timer/transact/time"
)

//var mainTime *time.Time

// GetTime
// @Summary Get time diff
// @Tags Time
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

// ResetTime
// @Summary Create a new puso
// @Tags Time
// @Accept json
// @Produce json
// @Success 200 {string} string	"successfully reset"
// @Router	/time	[post]
func ResetTime(w http.ResponseWriter, r *http.Request) {
	t := time.Time{}

	t.ResetTime()

	RespondWithJSON(w, http.StatusOK, "successfully reset")
}

func Ui(w http.ResponseWriter, r *http.Request) {
	http.StripPrefix(r.RequestURI, http.FileServer(http.Dir("build"))).ServeHTTP(w, r)
}
