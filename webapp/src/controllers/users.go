package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"webapp/src/config"
	"webapp/src/cookies"
	"webapp/src/requests"
	"webapp/src/responses"
	"webapp/src/utils"

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

func EditUser(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	user, err := json.Marshal(map[string]string{
		"name":  r.FormValue("name"),
		"email": r.FormValue("email"),
		"nick":  r.FormValue("nick"),
	})

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

	cookie, _ := cookies.Read(r)
	userID, _ := strconv.ParseUint(cookie["id"], 10, 64)

	url := fmt.Sprintf("%s/usuarios/%d", config.ApiURL, userID)

	response, err := requests.MakeAuthRequest(r, http.MethodPut, url, bytes.NewBuffer(user))

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

func LoadUpdatePasswordPage(w http.ResponseWriter, r *http.Request) {
	utils.ExecTemplate(w, "update-password.html", nil)
}

func UpdatePassword(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	passwords, err := json.Marshal(map[string]string{
		"current": r.FormValue("current"),
		"new": r.FormValue("new"),
	})

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

	cookie, _ := cookies.Read(r)
	userID , _ := strconv.ParseUint(cookie["id"], 10, 64)

	url := fmt.Sprintf("%s/usuarios/%d/atualizar-senha", config.ApiURL, userID)

	response, err := requests.MakeAuthRequest(r, http.MethodPost, url, bytes.NewBuffer(passwords))

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

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	cookie, _ := cookies.Read(r)
	userID, _ := strconv.ParseUint(cookie["id"], 10, 64)

	url := fmt.Sprintf("%s/usuarios/%d", config.ApiURL, userID)

	response, err := requests.MakeAuthRequest(r, http.MethodDelete, url, nil)

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
