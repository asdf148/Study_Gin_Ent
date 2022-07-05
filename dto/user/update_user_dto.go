package user

type UpdateUserDTO struct {
	Name string `json:"name" binding:"required"`
}