package dto

type UpdateUserDTO struct {
	Name string `json:"name" binding:"required"`
}