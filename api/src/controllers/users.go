package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repo"
	"api/src/responses"
	"encoding/json"
	"io"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User
	if err = json.Unmarshal(requestBody, &user); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	if err = user.Prepare(); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	repo := repo.NewUsersRepository(db)
	user.ID, err = repo.Create(user)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusCreated, user)

}
func SearchAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando Todos Usuários!"))
}
func SearchUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando Usuário!"))
}
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualizando Usuário!"))
}
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deletando Usuário!"))
}
