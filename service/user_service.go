package service

import (
	"fmt"
	dto "study_go/dto/user"
	"study_go/ent"
	myutils "study_go/myUtils"
	"study_go/repository"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

var customUtil myutils.JwtUtil = myutils.NewJwtUtil()

type UserService interface {
	Join(dto.JoinDTO, *gin.Context) (*ent.User, error)
	Login(dto.LoginDTO, *gin.Context) (string, error)
}

type userService struct {
	repository repository.UserRepository
}

func NewUserService(repository repository.UserRepository) UserService {
	return &userService{
		repository: repository,
	}
}

func (service *userService) Join(joinDTO dto.JoinDTO, ctx *gin.Context) (*ent.User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(joinDTO.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("failed password conversion at service: %w", err)
		return nil, err
	}
	joinDTO.Password = string(hash)

	user, err := service.repository.Create(joinDTO, ctx)
	if err != nil {
		fmt.Println("failed creating user at service: %w", err)
		return nil, err
	}
	return user, nil
}

func (service *userService) Login(loginDTO dto.LoginDTO, ctx *gin.Context) (string, error) {
	user, err := service.repository.FindOneByEmail(loginDTO.Email, ctx)
	if err != nil {
		fmt.Println("failed finding user at service: %w", err)
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginDTO.Password))
	if err != nil {
		fmt.Println("password is incorrect: %w", err)
		return "", err
	}

	token, err := customUtil.CreateAccessToken(user.ID)
	if err != nil {
		fmt.Println("password is incorrect: %w", err)
		return "", err
	}

	return token, nil
}