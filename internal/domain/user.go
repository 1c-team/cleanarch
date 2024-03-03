package domain

import (
	"context"
	"time"
)

type UserEntity struct {
	ID        uint      `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Email     string    `json:"email"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewUser(username, password, email, status string) UserEntity {
	return UserEntity{
		Username: username,
		Password: password,
		Email:    email,
		Status:   status,
	}
}

func (u *UserEntity) SetStatus(s string) {
	u.Status = s
}

type IUserUsecase interface {
	GetUserByID(ctx context.Context, id uint) (UserEntity, error)
}

type IUserRepository interface {
	Insert(user *UserEntity) error
	Fetch(ctx context.Context, cursor string, num int64) (res []UserEntity, nextCursor string, err error)
	GetByID(ctx context.Context, id uint) (UserEntity, error)
	GetByTitle(ctx context.Context, title string) (UserEntity, error)
	Update(ctx context.Context, user *UserEntity) error
	Store(ctx context.Context, user *UserEntity) error
	Delete(ctx context.Context, id int64) error
}
