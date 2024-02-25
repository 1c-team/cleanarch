package usecase

import (
	"context"

	"github.com/motchai-sns/auth-svc/internal/domain"
)

type UserUsecase struct {
	userRepo domain.IUserRepository
}

func NewUserUsecase(ur domain.IUserRepository) domain.IUserUsecase {
	return &UserUsecase{
		userRepo: ur,
	}
}

// Delete implements domain.UserRepository.
func (uuc *UserUsecase) Delete(ctx context.Context, id int64) error {
	return uuc.userRepo.Delete(ctx, id)
}

// Fetch implements domain.UserRepository.
func (uuc *UserUsecase) Fetch(ctx context.Context, cursor string, num int64) (res []domain.User, nextCursor string, err error) {
	return uuc.userRepo.Fetch(ctx, cursor, num)
}

// GetByID implements domain.UserRepository.
func (uuc *UserUsecase) GetByID(ctx context.Context, id uint) (domain.User, error) {
	return uuc.userRepo.GetByID(ctx, id)
}

// GetByTitle implements domain.UserRepository.
func (uuc *UserUsecase) GetByTitle(ctx context.Context, title string) (domain.User, error) {
	return uuc.userRepo.GetByTitle(ctx, title)
}

// Store implements domain.UserRepository.
func (uuc *UserUsecase) Store(ctx context.Context, user *domain.User) error {
	return uuc.userRepo.Store(ctx, user)
}

// Update implements domain.UserRepository.
func (uuc *UserUsecase) Update(ctx context.Context, user *domain.User) error {
	return uuc.userRepo.Update(ctx, user)
}
