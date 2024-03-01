package domain

type LoginStrategy int8

const (
	GoogleStrategy int = iota
	GithubStrategy
)

type IAuthUsecase interface {
	Login(strategy LoginStrategy) error
	Register(id uint) error
}
