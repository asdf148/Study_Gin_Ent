// user_controller
package controller

import (
	dto "study_go/dto/user"
	myutils "study_go/myUtils"
	"study_go/service"

	"github.com/gin-gonic/gin"
)

var (
	errorHandler myutils.ErrorHandler = myutils.NewErrorHandler()
)

type UserController interface {
	Join(*gin.Context) gin.H
	Login(*gin.Context) gin.H
}

type userController struct {
	service service.UserService
}

func NewUserController(service service.UserService) UserController {
	return &userController {
		service: service,
	}
}

func (controller *userController) Join(ctx *gin.Context) gin.H {
	var joinDTO dto.JoinDTO
	err := ctx.ShouldBindJSON(&joinDTO)
	errorHandler.ErrorHandling(err, "")
	return gin.H{
		"user": controller.service.Join(joinDTO, ctx),
	}
}

func (controller *userController) Login(ctx *gin.Context) gin.H {
	var loginDTO dto.LoginDTO
	err := ctx.ShouldBindJSON(&loginDTO)
	errorHandler.ErrorHandling(err, "")
	token := controller.service.Login(loginDTO, ctx)
	if token == "" {
		return nil
	}
	return gin.H{
		"token": controller.service.Login(loginDTO, ctx),
	}
}