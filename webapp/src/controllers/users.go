package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"webapp/src/config"
	"webapp/src/requests"
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

func Unfollow(w http.ResponseWriter, r *http.Request) { 
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

	url := fmt.Sprintf("%s/usuarios/%d/deixar-de-seguir", config.ApiURL, userID)

	response, err := requests.MakeAuthRequest(r, http.MethodPost, url, nil)

	if err != nil {
		responses.JSON(
			w,
			http.StatusInternalServerError,
			responses.APIError{
				Error: err.Error(),
			},
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

func Follow(w http.ResponseWriter, r *http.Request) {
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

	url := fmt.Sprintf("%s/usuarios/%d/seguir", config.ApiURL, userID)

	response, err := requests.MakeAuthRequest(r, http.MethodPost, url, nil)
	if err != nil {
		responses.JSON(
			w,
			http.StatusInternalServerError,
			responses.APIError{
				Error: err.Error(),
			},
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
