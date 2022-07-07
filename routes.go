package main

import (
	"study_go/controller"
	"study_go/middleware"
	"study_go/repository"
	"study_go/service"

	"github.com/gin-gonic/gin"
)

var (
	userRepository repository.UserRepository = repository.NewUserRepository()
	userService    service.UserService       = service.NewUserService(userRepository)
	userController controller.UserController = controller.NewUserController(userService)

	todoRepository repository.TodoRepository = repository.NewTodoRepository()
	todoService    service.TodoService       = service.NewTodoService(todoRepository, userRepository)
	todoController controller.TodoController = controller.NewTodoController(todoService)

	tokenMiddleware middleware.TokenVerification = middleware.NewTokenVerification()
)

func initializeRoutes() *gin.Engine {
	router := gin.Default()

	userRoutes := router.Group("user")
	{
		userRoutes.POST("/join", userController.Join)

		userRoutes.POST("/login", userController.Login)
	}

	todoRoutes := router.Group("todo")
	{
		todoRoutes.GET("/", todoController.FindAllTodo)

		todoRoutes.GET("/detail/:todoId", todoController.FindOneTodo)

		todoRoutes.Use(tokenMiddleware.TokenVerify)
		todoRoutes.POST("/", todoController.AddTodo)

		todoRoutes.PUT("/:todoId", todoController.ModifyTodo)

		todoRoutes.DELETE("/:todoId", todoController.DeleteTodo)
	}

	return router
}
