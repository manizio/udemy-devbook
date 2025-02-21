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
}
