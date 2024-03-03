package usecase

import (
	"github.com/motchai-sns/sn-mono/internal/domain"
)

type AuthUsecase struct{}

func NewAuthUsecase() domain.IAuthUsecase {
	return &AuthUsecase{}
}

// GithubLogin implements domain.IAuthUsecase.
func (authUC *AuthUsecase) GithubLogin() error {
	panic("unimplemented")
}

// GoogleLogin implements domain.IAuthUsecase.
func (authUC *AuthUsecase) GoogleLogin() error {
	panic("unimplemented")
}

// Register implements domain.IAuthUsecase.
func (authUC *AuthUsecase) Register(id uint) error {
	panic("unimplemented")
}

// GithubCallback implements domain.IAuthUsecase.
func (authUC *AuthUsecase) GithubCallback(user domain.GithubUserEntity) error {
	panic("unimplemented")
}

// GoogleCallback implements domain.IAuthUsecase.
func (authUC *AuthUsecase) GoogleCallback(user domain.GoogleUserEntity) error {
	panic("unimplemented")
}
