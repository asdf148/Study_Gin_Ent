package repository

import (
	"study_go/DB"
	dto "study_go/dto/user"
	"study_go/ent"
	"study_go/ent/user"
	myutils "study_go/myUtils"

	"github.com/gin-gonic/gin"
)

var (
	database *ent.Client = DB.GetConnector()
	errorHandler myutils.ErrorHandler = myutils.NewErrorHandler()
)

type UserRepository interface {
	FindOne(int, *gin.Context) *ent.User
	FindOneByEmail(string, *gin.Context) *ent.User
	FindAll(*gin.Context) []*ent.User
	Create(dto.JoinDTO, *gin.Context) *ent.User
	Update(int, dto.UpdateUserDTO, *gin.Context) *ent.User
	Delete(int, *gin.Context)
}

type userRepository struct {}

func NewUserRepository() UserRepository {
	return &userRepository{}
}

func (repository *userRepository) FindOne(userId int, ctx *gin.Context) *ent.User {
	user, err := database.User.Query().Where(user.ID(userId)).Only(ctx)
	errorHandler.ErrorHandling(err, "failed finding user at repository")
	return user
}

func (repository *userRepository) FindOneByEmail(userEmail string, ctx *gin.Context) *ent.User {
	user, err := database.User.Query().Where(user.Email(userEmail)).Only(ctx)
	errorHandler.ErrorHandling(err, "failed finding user at repository")
	return user
}

func (repository *userRepository) FindAll(ctx *gin.Context) []*ent.User {
	users, err := database.User.Query().All(ctx)
	errorHandler.ErrorHandling(err, "failed finding users at repository")
	return users
}

func (repository *userRepository) Create(joinDTO dto.JoinDTO, ctx *gin.Context) *ent.User {
	user, err := database.User.Create().SetName(joinDTO.Name).SetEmail(joinDTO.Email).SetPassword(joinDTO.Password).Save(ctx)
	errorHandler.ErrorHandling(err, "failed creating user at repository")
	return user
}

func (repository *userRepository) Update(userId int, updateUserDTO dto.UpdateUserDTO, ctx *gin.Context) *ent.User {
	user, err := database.User.UpdateOneID(userId).SetName(updateUserDTO.Name).Save(ctx)
	errorHandler.ErrorHandling(err, "failed updating user at repository")
	return user
}

func (repository *userRepository) Delete(userId int, ctx *gin.Context) {
	err := database.User.DeleteOneID(userId).Exec(ctx)
	errorHandler.ErrorHandling(err, "failed deleting user at repository")
	return
}