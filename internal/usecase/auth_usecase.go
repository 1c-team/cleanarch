package usecase

import (
	"github.com/motchai-sns/sn-mono/internal/domain"
)

type AuthUsecase struct {
	userRepo domain.IUserRepository
}

// Login implements domain.IAuthUsecase.
func (*AuthUsecase) Login(strategy domain.LoginStrategy) error {
	panic("unimplemented")
}

// register implements domain.IAuthUsecase.
func (*AuthUsecase) Register(id uint) error {
	panic("unimplemented")
}

func NewAuthUsecase(ur domain.IUserRepository) domain.IAuthUsecase {
	return &AuthUsecase{
		userRepo: ur,
	}
}
