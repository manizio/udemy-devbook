package routes

import (
	"net/http"
	"webapp/src/controllers"
)

var postRoutes = []Route {
	{
		URI: "/posts",
		Method: http.MethodPost,
		F: controllers.CreatePost,
		Authentication: true,
	},
	{
		URI: "/posts/{postID}/like",
		Method: http.MethodPost,
		F: controllers.LikePost,
		Authentication: true,
	},
	{
		URI: "/posts/{postID}/unlike",
		Method: http.MethodPost,
		F: controllers.UnlikePost,
		Authentication: true,
	},
	{
		URI: "/posts/{postID}/edit",
		Method: http.MethodGet,
		F: controllers.LoadEditPostPage,
		Authentication: true,
	},
	{
		URI: "/posts/{postID}",
		Method: http.MethodPut,
		F: controllers.UpdatePost,
		Authentication: true,
	},
	{
		URI: "/posts/{postID}",
		Method: http.MethodDelete,
		F: controllers.DeletePost,
		Authentication: true,
	},
	
}
