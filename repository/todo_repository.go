package repository

import (
	todoDTO "study_go/dto/todo"
	"study_go/ent"
	"study_go/ent/todo"

	"github.com/gin-gonic/gin"
)

type TodoRepository interface {
	FindOne(int, *gin.Context) *ent.Todo
	FindAll(*gin.Context) []*ent.Todo
	Create(todoDTO.CreateTodoDTO, int, *ent.User, *gin.Context) *ent.Todo
	Update(int, todoDTO.UpdateTodoDTO, *gin.Context) *ent.Todo
	Delete(int, *gin.Context)
}

type todoRepository struct {}

func NewTodoRepository() TodoRepository {
	return &todoRepository{}
}

func (repository *todoRepository) FindOne(todoId int, ctx *gin.Context) *ent.Todo {
	todo, err := database.Todo.Query().Where(todo.ID(todoId)).Only(ctx)
	errorHandler.ErrorHandling(err, "failed finding todo at repository")
	return todo
}

func (repository *todoRepository) FindAll(ctx *gin.Context) []*ent.Todo {
	todos, err := database.Todo.Query().All(ctx)
	errorHandler.ErrorHandling(err, "failed finding todos at repository")
	return todos
}

func (repository *todoRepository) Create(createTodoDTO todoDTO.CreateTodoDTO, userId int, user *ent.User, ctx *gin.Context) *ent.Todo {
	todo, err := database.Todo.Create().SetTitle(createTodoDTO.Title).SetContent(createTodoDTO.Content).SetUserID(userId).SetUser(user).Save(ctx)
	errorHandler.ErrorHandling(err, "failed creating todo at repository")
	return todo
}

func (repository *todoRepository) Update(todoId int, updateUserDTO todoDTO.UpdateTodoDTO, ctx *gin.Context) *ent.Todo {
	todo, err := database.Todo.UpdateOneID(todoId).SetTitle(updateUserDTO.Title).SetContent(updateUserDTO.Content).Save(ctx)
	errorHandler.ErrorHandling(err, "failed updating todo at repository")
	return todo
}

func (repository *todoRepository) Delete(todoId int, ctx *gin.Context) {
	err := database.Todo.DeleteOneID(todoId).Exec(ctx)
	errorHandler.ErrorHandling(err, "failed deleting todo at repository")
}