package routes

import (
	"api/src/controllers"
	"net/http"
)

var userRoutes = []Route{
	{
		URI: "/usuarios",
		Method: http.MethodPost,
		F: controllers.CreateUser,
		Authenticate: false,
	},

	{
		URI: "/usuarios",
		Method: http.MethodGet,
		F: controllers.SearchAllUsers,
		Authenticate: false,
	},

	{
		URI: "/usuarios/{userId}",
		Method: http.MethodGet,
		F: controllers.SearchUser,
		Authenticate: false,
	},

	{
		URI: "/usuarios/{userID}",
		Method: http.MethodPut,
		F: controllers.UpdateUser,
		Authenticate: false,
	},

	{
		URI: "/usuarios/{userId}",
		Method: http.MethodDelete,
		F: controllers.DeleteUser,
		Authenticate: false,
	},

}
