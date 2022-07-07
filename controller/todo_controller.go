package controller

import (
	"fmt"
	"net/http"
	"strconv"
	todoDTO "study_go/dto/todo"
	"study_go/service"

	"github.com/gin-gonic/gin"
)

type TodoController interface {
	FindOneTodo(*gin.Context)
	FindAllTodo(*gin.Context)
	AddTodo(*gin.Context)
	ModifyTodo(*gin.Context)
	DeleteTodo(*gin.Context)
}

type todoController struct {
	service service.TodoService
}

func NewTodoController(service service.TodoService) TodoController {
	return &todoController{
		service: service,
	}
}

func (controller *todoController) FindOneTodo(ctx *gin.Context) {
	defer routerErrorHandler(ctx)

	todoId := ctx.Param("todoId")
	intTodoId, err := strconv.Atoi(todoId)
	errorHandler.ErrorHandling(err, "failed to convert string to int at controller")

	ctx.IndentedJSON(http.StatusOK, gin.H{
		"todo": controller.service.FindOneTodo(intTodoId, ctx),
	})
}

func (controller *todoController) FindAllTodo(ctx *gin.Context) {
	defer routerErrorHandler(ctx)

	ctx.IndentedJSON(http.StatusOK, gin.H{
		"todos": controller.service.FindAllTodo(ctx),
	})
}

func (controller *todoController) AddTodo(ctx *gin.Context) {
	defer routerErrorHandler(ctx)

	var createTodoDTO todoDTO.CreateTodoDTO
	err := ctx.ShouldBindJSON(&createTodoDTO)
	errorHandler.ErrorHandling(err, "failed to create todo at controller")

	userId := ctx.Keys["userId"]
	intUserId, err := strconv.Atoi(fmt.Sprint(userId))
	errorHandler.ErrorHandling(err, "failed to convert string to int at controller")

	ctx.IndentedJSON(http.StatusCreated, gin.H{
		"todo": controller.service.AddTodo(createTodoDTO, intUserId, ctx),
	})
}

func (controller *todoController) ModifyTodo(ctx *gin.Context) {
	defer routerErrorHandler(ctx)

	var updateTodoDTO todoDTO.UpdateTodoDTO
	err := ctx.ShouldBindJSON(&updateTodoDTO)
	errorHandler.ErrorHandling(err, "failed to update todo at controller")

	todoId := ctx.Param("todoId")
	intTodoId, err := strconv.Atoi(todoId)
	errorHandler.ErrorHandling(err, "failed to convert string to int at controller")

	ctx.IndentedJSON(http.StatusOK, gin.H{
		"todo": controller.service.ModifyTodo(updateTodoDTO, intTodoId, ctx),
	})
}

func (controller *todoController) DeleteTodo(ctx *gin.Context) {
	defer routerErrorHandler(ctx)

	todoId := ctx.Param("todoId")
	intTodoId, err := strconv.Atoi(todoId)
	errorHandler.ErrorHandling(err, "failed to convert string to int at controller")
	controller.service.DeleteTodo(intTodoId, ctx)

	ctx.IndentedJSON(http.StatusOK, nil)
}
