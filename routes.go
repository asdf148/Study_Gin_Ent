package main

import (
	"fmt"
	"net/http"
	"study_go/controller"
	"study_go/dto"
	"study_go/repository"
	"study_go/service"

	"github.com/gin-gonic/gin"
)

var (
	userRepository repository.UserRepository = repository.NewUserRepository()
	userService service.UserService = service.NewUserService(userRepository)
	userController controller.UserController = controller.NewUserController(userService)
)

func initializeRoutes() *gin.Engine {
	router := gin.Default()

	userRoutes := router.Group("user")
	{
		userRoutes.POST("/join", func(ctx *gin.Context) {
			ctx.IndentedJSON(http.StatusCreated, userController.Join(ctx))
		})

		userRoutes.POST("/login", func(ctx *gin.Context) {
			defer func() {
				if r := recover(); r != nil {
					error_message := fmt.Sprint(r)
					errorResponse := dto.NewErrorResponse(error_message)
					ctx.IndentedJSON(http.StatusBadRequest, errorResponse)
				}
			}()
			result := userController.Login(ctx)
			ctx.IndentedJSON(http.StatusCreated, result)
		})
	}

	return router
}