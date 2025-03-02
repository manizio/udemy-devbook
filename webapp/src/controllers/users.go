package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"webapp/src/config"
	"webapp/src/models"
	"webapp/src/responses"

	"github.com/gorilla/mux"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	user, err := json.Marshal(map[string]string{
		"name":     r.FormValue("name"),
		"email":    r.FormValue("email"),
		"nick":     r.FormValue("nick"),
		"password": r.FormValue("password"),
	})

	if err != nil {
		responses.JSON(
			w,
			http.StatusBadRequest,
			responses.APIError{Error: err.Error()},
		)

		return
	}

	url := fmt.Sprintf("%s/usuarios", config.ApiURL)
	response, err := http.Post(
		url,
		"application/json",
		bytes.NewBuffer(user),
	)

	if err != nil {
		responses.JSON(
			w,
			http.StatusInternalServerError,
			responses.APIError{Error: err.Error()},
		)

		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.HandleErrorStatusCode(w, response)
		return
	}

	responses.JSON(w, response.StatusCode, nil)

}

func LoadUserProfile(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID, err := strconv.ParseUint(params["userID"], 10, 64)
	if err != nil {
		responses.JSON(
			w,
			http.StatusBadRequest,
			responses.APIError{
				Error: err.Error(),
			},
		)
		return
	}

	user, err := models.SearchFullUser(userID, r)

	fmt.Println(user)
}
