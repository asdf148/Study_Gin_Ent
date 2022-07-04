package repository

import (
	"fmt"
	"study_go/DB"
	dto "study_go/dto/user"
	"study_go/ent"
	"study_go/ent/user"

	"github.com/gin-gonic/gin"
)

var (
	database *ent.Client = DB.GetConnector()
)

type UserRepository interface {
	FindOne(int, *gin.Context) (*ent.User, error)
	FindOneByEmail(string, *gin.Context) (*ent.User, error)
	FindAll(*gin.Context) ([]*ent.User, error)
	Create(dto.JoinDTO, *gin.Context) (*ent.User, error)
	Update(int, dto.UpdateUserDTO, *gin.Context) (*ent.User, error)
	Delete(int, *gin.Context) error
}

type userRepository struct {}

func NewUserRepository() UserRepository {
	return &userRepository{}
}

func (repository *userRepository) FindOne(userId int, ctx *gin.Context) (*ent.User, error) {
	user, err := database.User.Query().Where(user.ID(userId)).Only(ctx)
	if err != nil {
		fmt.Println("failed finding user at repository: %w", err)
		return nil, err
	}
	return user, nil
}

func (repository *userRepository) FindOneByEmail(userEmail string, ctx *gin.Context) (*ent.User, error) {
	user, err := database.User.Query().Where(user.Email(userEmail)).Only(ctx)
	if err != nil {
		fmt.Println("failed finding user at repository: %w", err)
		return nil, err
	}
	return user, nil
}

func (repository *userRepository) FindAll(ctx *gin.Context) ([]*ent.User, error) {
	users, err := database.User.Query().All(ctx)
	if err != nil {
		fmt.Println("failed finding users at repository: %w", err)
		return nil, err
	}
	return users, nil
}

func (repository *userRepository) Create(joinDTO dto.JoinDTO, ctx *gin.Context) (*ent.User, error) {
	user, err := database.User.Create().SetName(joinDTO.Name).SetEmail(joinDTO.Email).SetPassword(joinDTO.Password).Save(ctx)
	if err != nil {
		fmt.Println("failed creating user at repository: %w", err)
		return nil, err
	}
	return user, nil
}

func (repository *userRepository) Update(userId int, updateUserDTO dto.UpdateUserDTO, ctx *gin.Context) (*ent.User, error) {
	user, err := database.User.UpdateOneID(userId).SetName(updateUserDTO.Name).Save(ctx)
	if err != nil {
		fmt.Println("failed updating user at repository: %w", err)
		return nil, err
	}
	return user, nil
}

func (repository *userRepository) Delete(userId int, ctx *gin.Context) error {
	err := database.User.DeleteOneID(userId).Exec(ctx)
	if err != nil {
		fmt.Println("failed deleting user at repository: %w", err)
		return err
	}
	return nil
}