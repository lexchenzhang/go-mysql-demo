package routes

import (
	"go-mysql/api/controller"
	"go-mysql/infrastructure"
)

//UserRoute -> Route for user module
type UserRoute struct {
	Handler    infrastructure.GinRouter
	Controller controller.UserController
}

//NewUserRoute -> initializes new instance of UserRoute
func NewUserRoute(
	controller controller.UserController,
	handler infrastructure.GinRouter,
) UserRoute {
	return UserRoute{
		Handler:    handler,
		Controller: controller,
	}
}

//Setup -> setups user routes
func (u UserRoute) Setup() {
	user := u.Handler.Gin.Group("/auth")
	{
		user.POST("/register", u.Controller.CreateUser)
		user.POST("/login", u.Controller.LoginUser)
	}
}
