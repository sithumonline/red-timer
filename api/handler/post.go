package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"

	"github.com/sithumonline/red-timer/transact/post"
)

// CreatePost
// @Summary Create a new post
// @Tags Post
// @Accept json
// @Produce json
// @Param data body	post.Post	true	"data"
// @Success 200 {string} string	"successfully post created"
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router	/post	[post]
func CreatePost(w http.ResponseWriter, r *http.Request) {
	t := post.Post{}

	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := t.Save(); err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	RespondWithJSON(w, http.StatusOK, "successfully post created")
}

// PostList
// @Summary Get post list
// @Tags Post
// @Accept json
// @Produce json
// @Success 200 {object} []post.Post
// @Failure 404 {string} string
// @Router	/post	[get]
func PostList(w http.ResponseWriter, r *http.Request) {
	t := post.Post{}

	list, err := t.GetList()

	if err != nil {
		RespondWithError(w, http.StatusNotFound, err.Error())
		return
	}

	RespondWithJSON(w, http.StatusOK, list)
}

// GetPost
// @Summary Get post
// @Tags Post
// @Accept json
// @Produce json
// @Param   id	path	string	true	"ID"
// @Success 200 {object} post.Post
// @Failure 404 {string} string
// @Router /post/{id} [get]
func GetPost(w http.ResponseWriter, r *http.Request) {
	t := post.Post{}
	ID := chi.URLParam(r, "id")
	o, err := t.Get(ID)

	if err != nil {
		RespondWithError(w, http.StatusNotFound, err.Error())
		return
	}

	RespondWithJSON(w, http.StatusOK, o)
}

// DeletePost
// @Summary Delete post
// @Tags Post
// @Accept json
// @Produce json
// @Param   id	path	string	true	"ID"
// @Success 200 {nil}	nil
// @Failure 404 {string}	string
// @Router /post/{id} [delete]
func DeletePost(w http.ResponseWriter, r *http.Request) {
	t := post.Post{}
	ID := chi.URLParam(r, "id")

	if err := t.Delete(ID); err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	RespondWithJSON(w, http.StatusNoContent, nil)
}

// UpdatePost
// @Summary Update post
// @Tags Post
// @Description Update post
// @Accept  json
// @Produce  json
// @Param   id	path	string	true	"ID"
// @Success 200 {string} string	"successfully updated"
// @Failure 400 {string} string
// @Failure 404 {string} string
// @Router /post/{id} [put]
func UpdatePost(w http.ResponseWriter, r *http.Request) {
	t := post.Post{}
	ID := chi.URLParam(r, "id")

	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := t.Update(ID); err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	RespondWithJSON(w, http.StatusOK, "successfully updated")
}
