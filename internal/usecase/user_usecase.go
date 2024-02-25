package usecase

import (
	"context"

	"github.com/tuannm-sns/auth-svc/domain"
)

type UserUsecase struct {
	userRepo domain.UserRepository
}

func NewUserUsecase(ur domain.UserRepository) domain.UserUsecase {
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
func (uuc *UserUsecase) GetByID(ctx context.Context, id int64) (domain.User, error) {
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
