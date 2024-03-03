package domain

// Entity
type (
	GoogleUserEntity struct {
		ID            uint   `json:"id"`
		Email         string `json:"email"`
		VerifiedEmail string `json:"verified_email"`
		Name          string `json:"name"`
		GivenName     string `json:"given_name"`
		FamilyName    string `json:"family_name"`
		Picture       string `json:"picture"`
		Locale        string `json:"locate"`
	}

	GithubUserEntity struct {
		ID            uint   `json:"id"`
		Email         string `json:"email"`
		VerifiedEmail string `json:"verified_email"`
		Name          string `json:"name"`
		GivenName     string `json:"given_name"`
		FamilyName    string `json:"family_name"`
		Picture       string `json:"picture"`
		Locale        string `json:"locate"`
	}
)

func NewGoogleUser(username, password, email, status string) GoogleUserEntity {
	return GoogleUserEntity{}
}

func NewGithubUser(username, password, email, status string) GithubUserEntity {
	return GithubUserEntity{}
}

// Usecase Interface
type (
	IAuthUsecase interface {
		GoogleLogin() error
		GithubLogin() error

		GoogleCallback(user GoogleUserEntity) error
		GithubCallback(user GithubUserEntity) error

		Register(id uint) error
	}
)
