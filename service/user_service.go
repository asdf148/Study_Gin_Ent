package service

import (
	dto "study_go/dto/user"
)

type UserService interface {
	Join(dto.JoinDTO) string
	Login(dto.LoginDTO) string
}

type userService struct {}

func NewUserService() UserService {
	return &userService{}
}

func (service *userService) Join(joinDTO dto.JoinDTO) string {
	// database.s
	return "sdf"
}

func (service *userService) Login(loginDTO dto.LoginDTO) string {
	return "dfs"
}