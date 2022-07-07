// user_controller
package controller

import (
	"fmt"
	"net/http"
	"study_go/dto"
	userDTO "study_go/dto/user"
	myutils "study_go/myUtils"
	"study_go/service"

	"github.com/gin-gonic/gin"
)

var (
	errorHandler myutils.ErrorHandler = myutils.NewErrorHandler()
)

type UserController interface {
	Join(*gin.Context)
	Login(*gin.Context)
}

type userController struct {
	service service.UserService
}

func NewUserController(service service.UserService) UserController {
	return &userController{
		service: service,
	}
}

func (controller *userController) Join(ctx *gin.Context) {
	defer routerErrorHandler(ctx)

	var joinDTO userDTO.JoinDTO
	err := ctx.ShouldBindJSON(&joinDTO)
	errorHandler.ErrorHandling(err, "failed get joinDTO at controller")

	ctx.IndentedJSON(http.StatusCreated, gin.H{
		"user": controller.service.Join(joinDTO, ctx),
	})
}

func (controller *userController) Login(ctx *gin.Context) {
	defer routerErrorHandler(ctx)

	var loginDTO userDTO.LoginDTO
	err := ctx.ShouldBindJSON(&loginDTO)
	errorHandler.ErrorHandling(err, "failed get loginDTO at controller")

	ctx.IndentedJSON(http.StatusCreated, gin.H{
		"token": controller.service.Login(loginDTO, ctx),
	})
}

func routerErrorHandler(ctx *gin.Context) {
	if r := recover(); r != nil {
		error_message := fmt.Sprint(r)
		errorResponse := dto.NewErrorResponse(error_message)
		ctx.IndentedJSON(http.StatusBadRequest, errorResponse)
	}
}
