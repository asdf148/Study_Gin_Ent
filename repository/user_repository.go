package repository

import (
	"context"
	"fmt"
	"study_go/DB"
	dto "study_go/dto/user"
	"study_go/ent"
	"study_go/ent/user"
)

var (
	database *ent.Client = DB.GetConnector()
)

type UserRepository interface {
	FindOne()
	FindAll()
	Create(dto.JoinDTO) (*ent.User, error)
	Update()
	Delete()
}

type userRepository struct {}

func NewUserRepository() UserRepository {
	return &userRepository{}
}

// TODO:
// 인터페이스 수정
// context 찾아서 공부, 수정하기
func (repository *userRepository) FindOne(userId int) (*ent.User, error) {
	user, err := database.User.Query().Where(user.ID(userId)).Only(context.Context)
	if err != nil {
		fmt.Println("failed creating user: %w", err)
		return nil, err
	}
	return user, nil
}

func (repository *userRepository) FindAll() {
	
}

func (repository *userRepository) Create(joinDTO dto.JoinDTO) (*ent.User, error) {
	user, err := database.User.Create().SetName(joinDTO.Name).SetEmail(joinDTO.Email).SetPassword(joinDTO.Password).Save(context.TODO())
	if err != nil {
		fmt.Println("failed creating user: %w", err)
		return nil, err
	}
	return user, nil
}

func (repository *userRepository) Update() {
	
}

func (repository *userRepository) Delete() {
	
}