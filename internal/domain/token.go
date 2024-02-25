package domain

import "context"

type Token struct {
	ID        int64  `json:"id"`
	Type      string `json:"type"`
	Name      string `json:"name"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type TokenRepository interface {
	GetByID(ctx context.Context, id int64) (Token, error)
}
