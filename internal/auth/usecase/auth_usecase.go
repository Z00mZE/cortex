package usecase

import (
	"context"
	"github.com/Z00mZE/cortex/internal/auth/internal/cqrs/user"
)

type AuthUseCase struct {
	cqrs user.User
}

func NewAuthUseCase(cqrs user.User) *AuthUseCase {
	return &AuthUseCase{cqrs: cqrs}
}

func (a AuthUseCase) SignUp(ctx context.Context, email, password string) error {
	panic("implement me")
}
