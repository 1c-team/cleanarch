package usecase

import (
	"context"

	"github.com/motchai-sns/sn-mono/internal/domain"
)

type UserUsecase struct {
	userRepo domain.IUserRepository
}

func NewUserUsecase(ur domain.IUserRepository) domain.IUserUsecase {
	return &UserUsecase{
		userRepo: ur,
	}
}

// GetUserByID implements domain.IUserUsecase.
func (uuc *UserUsecase) GetUserByID(ctx context.Context, id uint) (domain.UserEntity, error) {
	return uuc.userRepo.GetByID(ctx, id)
}
