package routes

import (
	"api/src/controllers"
	"net/http"
)

var postsRoutes = []Route{
	{
		URI:          "/posts",
		Method:       http.MethodPost,
		F:            controllers.CreatePost,
		Authenticate: true,
	},
	{
		URI:          "/posts",
		Method:       http.MethodGet,
		F:            controllers.SearchPosts,
		Authenticate: true,
	},
	{
		URI:          "/posts/{postID}",
		Method:       http.MethodGet,
		F:            controllers.SearchPost,
		Authenticate: true,
	},
	{
		URI:          "/posts/{postID}",
		Method:       http.MethodPut,
		F:            controllers.UpdatePost,
		Authenticate: true,
	},
	{
		URI:          "/posts/{postID}",
		Method:       http.MethodDelete,
		F:            controllers.DeletePost,
		Authenticate: true,
	},
}
