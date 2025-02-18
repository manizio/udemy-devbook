package controllers

import (
	"api/src/auth"
	"api/src/database"
	"api/src/models"
	"api/src/repo"
	"api/src/responses"
	"encoding/json"
	"io"
	"net/http"
)

func CreatePost(w http.ResponseWriter, r *http.Request) {
	userID, err := auth.ExtractUserID(r)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var post models.Post
	if err = json.Unmarshal(reqBody, &post); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return 
	}

	post.AuthorID = userID

	if err = post.Prepare(); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := repo.NewPostsRepository(db)

	post.ID, err = repo.Create(post)

	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusCreated, post)
}
func SearchPosts(w http.ResponseWriter, r *http.Request) {

}
func SearchPost(w http.ResponseWriter, r *http.Request) {

}
func UpdatePost(w http.ResponseWriter, r *http.Request) {

}
func DeletePost(w http.ResponseWriter, r *http.Request) {

}
