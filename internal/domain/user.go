package domain

import (
	"context"
	"time"
)

type User struct {
	ID        uint      `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Email     string    `json:"email"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewUser(username, password, email, status string) User {
	return User{
		Username: username,
		Password: password,
		Email:    email,
		Status:   status,
	}
}

func (u *User) SetStatus(s string) {
	u.Status = s
}

type IUserUsecase interface {
	GetByID(ctx context.Context, id uint) (User, error)
}

type IUserRepository interface {
	Insert(user *User) error
	Fetch(ctx context.Context, cursor string, num int64) (res []User, nextCursor string, err error)
	GetByID(ctx context.Context, id uint) (User, error)
	GetByTitle(ctx context.Context, title string) (User, error)
	Update(ctx context.Context, user *User) error
	Store(ctx context.Context, user *User) error
	Delete(ctx context.Context, id int64) error
}
