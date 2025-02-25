package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"webapp/src/config"
	"webapp/src/models"
	"webapp/src/requests"
	"webapp/src/responses"
	"webapp/src/utils"
)

func LoadLoginScreen(w http.ResponseWriter, r *http.Request) {
	utils.ExecTemplate(w, "login.html", nil)
}

func LoadUserSignInPage(w http.ResponseWriter, r *http.Request) {
	utils.ExecTemplate(w, "cadastro.html", nil)
}

func LoadHomePage(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("%s/posts", config.ApiURL)

	response, err := requests.MakeAuthRequest(r, http.MethodGet, url, nil)
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, err)
		return
	}

	defer response.Body.Close()
	if response.StatusCode >= 400 {
		responses.HandleErrorStatusCode(w, response)
		return
	}

	var posts []models.Post
	if err = json.NewDecoder(response.Body).Decode(&posts); err != nil {
		responses.JSON(w,
			http.StatusUnprocessableEntity,
			responses.APIError{
				Error: err.Error(),
			},
		)
		return
	}

	utils.ExecTemplate(w, "home.html", posts)
}
