package controller

import (
	"fmt"
	"strconv"
	todoDTO "study_go/dto/todo"
	"study_go/service"

	"github.com/gin-gonic/gin"
)

type TodoController interface {
	FindOneTodo(*gin.Context) *gin.H
	FindAllTodo(*gin.Context) *gin.H
	AddTodo(*gin.Context) *gin.H
	ModifyTodo(*gin.Context) *gin.H
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

func (controller *todoController) FindOneTodo(ctx *gin.Context) *gin.H {
	todoId := ctx.Param("todoId")
	intTodoId, err := strconv.Atoi(todoId)
	errorHandler.ErrorHandling(err, "failed to convert string to int at controller")
	return &gin.H{
		"todo": controller.service.FindOneTodo(intTodoId, ctx),
	}
}

func (controller *todoController) FindAllTodo(ctx *gin.Context) *gin.H {
	return &gin.H{
		"todos": controller.service.FindAllTodo(ctx),
	}
}

func (controller *todoController) AddTodo(ctx *gin.Context) *gin.H {
	var createTodoDTO todoDTO.CreateTodoDTO
	err := ctx.ShouldBindJSON(&createTodoDTO)
	errorHandler.ErrorHandling(err, "failed to create todo at controller")

	userId := ctx.Keys["userId"]
	intUserId, err := strconv.Atoi(fmt.Sprint(userId))
	errorHandler.ErrorHandling(err, "failed to convert string to int at controller")

	return &gin.H{
		"todo": controller.service.AddTodo(createTodoDTO, intUserId, ctx),
	}
}

func (controller *todoController) ModifyTodo(ctx *gin.Context) *gin.H {
	var updateTodoDTO todoDTO.UpdateTodoDTO
	err := ctx.ShouldBindJSON(&updateTodoDTO)
	errorHandler.ErrorHandling(err, "failed to update todo at controller")

	todoId := ctx.Param("todoId")
	intTodoId, err := strconv.Atoi(todoId)
	errorHandler.ErrorHandling(err, "failed to convert string to int at controller")

	return &gin.H{
		"todo": controller.service.ModifyTodo(updateTodoDTO, intTodoId, ctx),
	}
}

func (controller *todoController) DeleteTodo(ctx *gin.Context) {
	todoId := ctx.Param("todoId")
	intTodoId, err := strconv.Atoi(todoId)
	errorHandler.ErrorHandling(err, "failed to convert string to int at controller")
	controller.service.DeleteTodo(intTodoId, ctx)
}