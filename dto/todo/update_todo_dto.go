package todo

type UpdateTodoDTO struct {
	Title string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}