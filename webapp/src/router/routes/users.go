package routes

import (
	"net/http"
	"webapp/src/controllers"
)

var usersRoutes = []Route{
	{
		URI:            "/cadastrar",
		Method:         http.MethodGet,
		F:              controllers.LoadUserSignInPage,
		Authentication: false,
	},
	{
		URI:            "/users",
		Method:         http.MethodPost,
		F:              controllers.CreateUser,
		Authentication: false,
	},
	{
		URI: "/search",
		Method: http.MethodGet,
		F: controllers.LoadSearchedUsersPage,
		Authentication: true,
	},
	{
		URI: "/users/{userID}",
		Method: http.MethodGet,
		F: controllers.LoadUserProfile,
		Authentication: true,
	},
}
