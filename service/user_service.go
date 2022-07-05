package service

import (
	dto "study_go/dto/user"
	"study_go/ent"
	myutils "study_go/myUtils"
	"study_go/repository"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

var (
	customUtil myutils.JwtUtil = myutils.NewJwtUtil()
	errorHandler myutils.ErrorHandler = myutils.NewErrorHandler()
)

type UserService interface {
	Join(dto.JoinDTO, *gin.Context) *ent.User
	Login(dto.LoginDTO, *gin.Context) string
}

type userService struct {
	repository repository.UserRepository
}

func NewUserService(repository repository.UserRepository) UserService {
	return &userService{
		repository: repository,
	}
}

func (service *userService) Join(joinDTO dto.JoinDTO, ctx *gin.Context) *ent.User {
	hash, err := bcrypt.GenerateFromPassword([]byte(joinDTO.Password), bcrypt.DefaultCost)
	errorHandler.ErrorHandling(err, "failed password conversion at service")
	joinDTO.Password = string(hash)

	user := service.repository.Create(joinDTO, ctx)
	return user
}

func (service *userService) Login(loginDTO dto.LoginDTO, ctx *gin.Context) string {
	user := service.repository.FindOneByEmail(loginDTO.Email, ctx)

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginDTO.Password))
	if err != nil {
		errorHandler.ErrorHandling(err, "password is incorrect at service")
		return ""
	}

	token, err := customUtil.CreateAccessToken(user.ID)
	errorHandler.ErrorHandling(err, "failed token create at service")

	return token
}