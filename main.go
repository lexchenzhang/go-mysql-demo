package main

import (
	"go-mysql/api/controller"
	"go-mysql/api/repository"
	"go-mysql/api/routes"
	"go-mysql/api/service"
	"go-mysql/infrastructure"
	"go-mysql/models"
)

func main() {
	router := infrastructure.NewGinRouter()                     //router has been initialized and configured
	db := infrastructure.NewDatabase()                          // databse has been initialized and configured
	postRepository := repository.NewPostRepository(db)          // repository are being setup
	postService := service.NewPostService(postRepository)       // service are being setup
	postController := controller.NewPostController(postService) // controller are being set up
	postRoute := routes.NewPostRoute(postController, router)    // post routes are initialized
	postRoute.Setup()                                           // post routes are being setup

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService)
	userRoute := routes.NewUserRoute(userController, router)
	userRoute.Setup()

	db.DB.AutoMigrate(&models.Post{}, &models.User{})
	router.Gin.Run(":8000") //server started on 8000 port
}
