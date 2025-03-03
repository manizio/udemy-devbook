package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"webapp/src/config"
	"webapp/src/cookies"
	"webapp/src/models"
	"webapp/src/requests"
	"webapp/src/responses"
	"webapp/src/utils"

	"github.com/gorilla/mux"
)

func LoadLoginScreen(w http.ResponseWriter, r *http.Request) {
	cookies, _ := cookies.Read(r)

	if cookies["token"] != "" {
		http.Redirect(w, r, "/home", 302)
		return
	}

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

	cookie, _ := cookies.Read(r)
	userID, _ := strconv.ParseUint(cookie["id"], 10, 64)

	utils.ExecTemplate(w, "home.html", struct {
		Posts  []models.Post
		UserID uint64
	}{
		Posts:  posts,
		UserID: userID,
	})
}

func LoadEditPostPage(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	postID, err := strconv.ParseUint(params["postID"], 10, 64)
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

	url := fmt.Sprintf("%s/posts/%d", config.ApiURL, postID)
	response, err := requests.MakeAuthRequest(r, http.MethodGet, url, nil)

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

	var post models.Post
	if err = json.NewDecoder(response.Body).Decode(&post); err != nil {
		responses.JSON(w, http.StatusUnprocessableEntity,
			responses.APIError{
				Error: err.Error(),
			})
		return
	}
	utils.ExecTemplate(w, "edit-post.html", post)
}
func LoadSearchedUsersPage(w http.ResponseWriter, r *http.Request) {
	nameOrNick := strings.ToLower(r.URL.Query().Get("user"))
	url := fmt.Sprintf("%s/usuarios?user=%s", config.ApiURL, nameOrNick)

	response, err := requests.MakeAuthRequest(r, http.MethodGet, url, nil)

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

	var users []models.User
	if err = json.NewDecoder(response.Body).Decode(&users); err != nil {
		responses.JSON(
			w,
			http.StatusInternalServerError,
			responses.APIError{
				Error: err.Error(),
			},
		)
		return
	}

	utils.ExecTemplate(w, "usuarios.html", users)

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
	cookie, _ := cookies.Read(r)
	loggedUserID, _ := strconv.ParseUint(cookie["id"], 10, 64)

	utils.ExecTemplate(w, "usuario.html", struct {
		User         models.User
		LoggedUserID uint64
	}{
		User:         user,
		LoggedUserID: loggedUserID,
	})
}
