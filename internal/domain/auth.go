package domain

// Entity
type (
	GoogleUser struct {
		ID            uint   `json:"id"`
		Email         string `json:"email"`
		VerifiedEmail string `json:"verified_email"`
		Name          string `json:"name"`
		GivenName     string `json:"given_name"`
		FamilyName    string `json:"family_name"`
		Picture       string `json:"picture"`
		Locale        string `json:"locate"`
	}

	GithubUser struct {
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

func NewGoogleUser(username, password, email, status string) GoogleUser {
	return GoogleUser{}
}

func NewGithubUser(username, password, email, status string) GithubUser {
	return GithubUser{}
}

// Usecase Interface
type (
	IAuthUsecase interface {
		GoogleLogin() error
		GithubLogin() error

		GoogleCallback(user GoogleUser) error
		GithubCallback(user GithubUser) error

		Register(id uint) error
	}
)
