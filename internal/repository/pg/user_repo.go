package pg

import (
	"context"

	"github.com/tuannm-sns/auth-svc/domain"
	"github.com/tuannm-sns/auth-svc/internal/model"
	"gorm.io/gorm"
)

type pgUserRepository struct {
	connection *gorm.DB
}

// Delete implements domain.UserRepository.
func (pg *pgUserRepository) Delete(ctx context.Context, id int64) error {
	panic("unimplemented")
}

// Fetch implements domain.UserRepository.
func (pg *pgUserRepository) Fetch(ctx context.Context, cursor string, num int64) (res []domain.User, nextCursor string, err error) {
	// var usr model.User
	// pg.db.Get(&usr)
	panic("unimplemented")
}

// GetByTitle implements domain.UserRepository.
func (pg *pgUserRepository) GetByTitle(ctx context.Context, title string) (domain.User, error) {
	panic("unimplemented")
}

// Store implements domain.UserRepository.
func (pg *pgUserRepository) Store(ctx context.Context, a *domain.User) error {
	panic("unimplemented")
}

// Update implements domain.UserRepository.
func (pg *pgUserRepository) Update(ctx context.Context, ar *domain.User) error {
	// pg.db.Create(&model.User{Username: "usr1", Password: "pass1"})
	panic("unimplemented")
}

// GetByID implements domain.UserRepository.
func (pg *pgUserRepository) GetByID(ctx context.Context, id int64) (domain.User, error) {
	var usr model.User
	result := pg.connection.First(&usr, id)

	return domain.User{
			ID:       usr.ID,
			Username: usr.Username,
		},
		result.Error

}

func NewPgArticleRepository(conn *gorm.DB) domain.UserRepository {
	return &pgUserRepository{conn}
}
