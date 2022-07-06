package service

import (
	todoDTO "study_go/dto/todo"
	"study_go/ent"
	"study_go/repository"

	"github.com/gin-gonic/gin"
)

type TodoService interface {
	FindOneTodo(int, *gin.Context) *ent.Todo
	FindAllTodo(*gin.Context) []*ent.Todo
	AddTodo(todoDTO.CreateTodoDTO, int, *gin.Context) *ent.Todo
	ModifyTodo(todoDTO.UpdateTodoDTO, int, *gin.Context) *ent.Todo
	DeleteTodo(int, *gin.Context)
}

type todoService struct {
	repository repository.TodoRepository
	userRepository repository.UserRepository
}

func NewTodoService(repository repository.TodoRepository, userRepository repository.UserRepository) TodoService {
	return &todoService{
		repository: repository,
		userRepository: userRepository,
	}
}

func (service *todoService) FindOneTodo(todoId int, ctx *gin.Context) *ent.Todo {
	return service.repository.FindOne(todoId, ctx)
}

func (service *todoService) FindAllTodo(ctx *gin.Context) []*ent.Todo {
	return service.repository.FindAll(ctx)
}

func (service *todoService) AddTodo(createTodoDTO todoDTO.CreateTodoDTO, userId int, ctx *gin.Context) *ent.Todo {
	user := service.userRepository.FindOne(userId, ctx)
	return service.repository.Create(createTodoDTO, userId, user, ctx)
}

func (service *todoService) ModifyTodo(updateTodoDTO todoDTO.UpdateTodoDTO, todoId int, ctx *gin.Context) *ent.Todo {
	return service.repository.Update(updateTodoDTO, todoId, ctx)
}

func (service *todoService) DeleteTodo(todoId int, ctx *gin.Context) {
	service.repository.Delete(todoId, ctx)
}