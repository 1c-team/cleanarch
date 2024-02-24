package usecase

import (
	"context"
	"time"

	"github.com/tuannm-sns/auth-svc/domain"
)

type userUsecase struct {
	userRepo       domain.UserRepository
	contextTimeout time.Duration
}

// Delete implements domain.UserRepository.
func (uuc *userUsecase) Delete(ctx context.Context, id int64) error {
	return uuc.userRepo.Delete(ctx, id)
}

// Fetch implements domain.UserRepository.
func (uuc *userUsecase) Fetch(ctx context.Context, cursor string, num int64) (res []domain.User, nextCursor string, err error) {
	return uuc.userRepo.Fetch(ctx, cursor, num)
}

// GetByID implements domain.UserRepository.
func (uuc *userUsecase) GetByID(ctx context.Context, id int64) (domain.User, error) {
	return uuc.userRepo.GetByID(ctx, id)
}

// GetByTitle implements domain.UserRepository.
func (uuc *userUsecase) GetByTitle(ctx context.Context, title string) (domain.User, error) {
	return uuc.userRepo.GetByTitle(ctx, title)
}

// Store implements domain.UserRepository.
func (uuc *userUsecase) Store(ctx context.Context, user *domain.User) error {
	return uuc.userRepo.Store(ctx, user)
}

// Update implements domain.UserRepository.
func (uuc *userUsecase) Update(ctx context.Context, user *domain.User) error {
	return uuc.userRepo.Update(ctx, user)
}

func NewUserUsecase(ur domain.UserRepository, timeout time.Duration) domain.UserRepository {
	return &userUsecase{
		userRepo:       ur,
		contextTimeout: timeout,
	}
}
