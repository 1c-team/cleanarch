package model

import (
	"github.com/motchai-sns/sn-mono/internal/domain"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string
	Password string
	Email    string
	Status   string
}

func NewUserModel(user *domain.UserEntity) *User {
	return &User{
		Username: user.Username,
		Password: user.Password,
		Email:    user.Email,
		Status:   user.Status,
	}
}

func (userModel *User) ToEntity() *domain.UserEntity {
	return &domain.UserEntity{
		ID:        userModel.ID,
		Username:  userModel.Username,
		Password:  userModel.Password,
		Email:     userModel.Email,
		Status:    userModel.Status,
		CreatedAt: userModel.CreatedAt,
		UpdatedAt: userModel.UpdatedAt,
	}
}
