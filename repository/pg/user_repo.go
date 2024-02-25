package pg

import (
	"context"

	"github.com/motchai-sns/auth-svc/internal/domain"
	"github.com/motchai-sns/auth-svc/repository/models"
	"gorm.io/gorm"
)

type PgUserRepo struct {
	conn *gorm.DB
}

func NewPgUserRepository(conn *gorm.DB) domain.IUserRepository {
	return &PgUserRepo{conn}
}

// Delete implements domain.IUserRepository.
func (pg *PgUserRepo) Delete(ctx context.Context, id int64) error {
	var usr model.User
	result := pg.conn.Delete(&usr, id)

	return result.Error
}

// Fetch implements domain.IUserRepository.
func (pg *PgUserRepo) Fetch(ctx context.Context, cursor string, num int64) (res []domain.User, nextCursor string, err error) {
	panic("unimplemented")
}

// GetByTitle implements domain.IUserRepository.
func (pg *PgUserRepo) GetByTitle(ctx context.Context, title string) (domain.User, error) {
	panic("unimplemented")
}

// Insert implements domain.IUserRepository.
func (pg *PgUserRepo) Insert(user *domain.User) error {
	newUser := model.NewUserModelFromEntity(user)
	result := pg.conn.Create(newUser)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

// Store implements domain.IUserRepository.
func (pg *PgUserRepo) Store(ctx context.Context, user *domain.User) error {
	panic("unimplemented")
}

// Update implements domain.IUserRepository.
func (*PgUserRepo) Update(ctx context.Context, user *domain.User) error {
	panic("unimplemented")
}

// GetByID implements domain.UserRepository.
func (pg *PgUserRepo) GetByID(ctx context.Context, id uint) (domain.User, error) {
	var usr model.User
	result := pg.conn.First(&usr, id)

	return *usr.ToEntity(), result.Error
}
