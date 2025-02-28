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

func CreatePost(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	post, err := json.Marshal(map[string]string{
		"title":   r.FormValue("title"),
		"content": r.FormValue("content"),
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

	url := fmt.Sprintf("%s/posts", config.ApiURL)

	response, err := requests.MakeAuthRequest(
		r,
		http.MethodPost,
		url,
		bytes.NewBuffer(post),
	)

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

func LikePost(w http.ResponseWriter, r *http.Request) {
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

	url := fmt.Sprintf("%s/posts/%d/like", config.ApiURL, postID)

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

func UnlikePost(w http.ResponseWriter, r *http.Request) {
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

	url := fmt.Sprintf("%s/posts/%d/unlike", config.ApiURL, postID)

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

func UpdatePost(w http.ResponseWriter, r *http.Request) {
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

	r.ParseForm()

	post, err := json.Marshal(map[string]string{
		"title":   r.FormValue("title"),
		"content": r.FormValue("content"),
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

	url := fmt.Sprintf("%s/posts/%d", config.ApiURL, postID)
	response, err := requests.MakeAuthRequest(r, http.MethodPut, url, bytes.NewBuffer(post))

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

func DeletePost(w http.ResponseWriter, r *http.Request) {
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

	if response.StatusCode >= 400 {
		responses.HandleErrorStatusCode(w, response)
		return
	}
	defer response.Body.Close()

	responses.JSON(w, response.StatusCode, nil)
}
