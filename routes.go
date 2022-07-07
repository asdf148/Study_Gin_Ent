package main

import (
	"fmt"
	"net/http"
	"study_go/controller"
	"study_go/dto"
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
		userRoutes.POST("/join", func(ctx *gin.Context) {
			defer routerErrorHandler(ctx)
			ctx.IndentedJSON(http.StatusCreated, userController.Join(ctx))
		})

		userRoutes.POST("/login", func(ctx *gin.Context) {
			defer routerErrorHandler(ctx)
			result := userController.Login(ctx)
			ctx.IndentedJSON(http.StatusCreated, result)
		})
	}

	todoRoutes := router.Group("todo")
	{
		todoRoutes.GET("/", func(ctx *gin.Context) {
			defer routerErrorHandler(ctx)
			ctx.IndentedJSON(http.StatusOK, todoController.FindAllTodo(ctx))
		})

		todoRoutes.GET("/detail/:todoId", func(ctx *gin.Context) {
			defer routerErrorHandler(ctx)
			ctx.IndentedJSON(http.StatusOK, todoController.FindOneTodo(ctx))
		})

		todoRoutes.Use(tokenMiddleware.TokenVerify)
		todoRoutes.POST("/", func(ctx *gin.Context) {
			defer routerErrorHandler(ctx)
			ctx.IndentedJSON(http.StatusCreated, todoController.AddTodo(ctx))
		})

		todoRoutes.PUT("/todoId", func(ctx *gin.Context) {
			defer routerErrorHandler(ctx)
			ctx.IndentedJSON(http.StatusOK, todoController.ModifyTodo(ctx))
		})

		todoRoutes.DELETE("/:todoId", func(ctx *gin.Context) {
			defer routerErrorHandler(ctx)
			todoController.DeleteTodo(ctx)
			ctx.IndentedJSON(http.StatusOK, nil)
		})
	}

	return router
}

func routerErrorHandler(ctx *gin.Context) {
	if r := recover(); r != nil {
		error_message := fmt.Sprint(r)
		errorResponse := dto.NewErrorResponse(error_message)
		ctx.IndentedJSON(http.StatusBadRequest, errorResponse)
	}
}
