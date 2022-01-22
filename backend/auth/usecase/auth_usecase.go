package usecase

import (
	"context"
	"fmt"
	v1 "github.com/Z00mZE/cortex/backend/auth/internal/controller/http/v1"
	"github.com/Z00mZE/cortex/backend/auth/internal/cqrs/user"
	"net/http"
)

type AuthUseCase struct {
	cqrs user.User
}

func NewAuthUseCase(cqrs user.User) *AuthUseCase {
	return &AuthUseCase{cqrs: cqrs}
}

func (a *AuthUseCase) SignUp(ctx context.Context, email, password, duplicate string) v1.UseCaseError {
	if password != duplicate {
		return NewAuthUseCaseError(http.StatusConflict, "пароли не совпадают")
	}
	{
		isExist, err := a.cqrs.Query.IsEmailRegistered.Handler(ctx, email)

		fmt.Println("isExist", isExist, email, password)
		if err != nil {
			return NewAuthUseCaseError(http.StatusInternalServerError, err.Error())
		}
		if isExist {
			return NewAuthUseCaseError(http.StatusConflict, "регистрация для указанногот email невозможна")
		}
	}
	err := a.cqrs.Command.CreateUserHandler.Handler(ctx, email, password)
	if err != nil {
		fmt.Println("insert err", err)
		return NewAuthUseCaseError(http.StatusInternalServerError, err.Error())
	}

	return NewAuthUseCaseError(http.StatusCreated, "учётная запись создана")
}
