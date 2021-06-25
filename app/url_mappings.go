package app

import "microservices/controllers"

func mapUrls() {
	router.POST("users", controllers.UsersController.Create)
	router.GET("users/:id", controllers.UsersController.Get)
}
