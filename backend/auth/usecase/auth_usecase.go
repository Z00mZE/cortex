package usecase

import (
	"context"
	"fmt"
	"github.com/Z00mZE/cortex/backend/auth/ent"
	"github.com/Z00mZE/cortex/backend/auth/ent/user"
	v1 "github.com/Z00mZE/cortex/backend/auth/internal/controller/http/v1"
	"net/http"
)

type AuthUseCase struct {
	sessionClient *ent.SessionClient
	userClient    *ent.UserClient
	hasher        hasher
	mailService   MailService
}

type hasher interface {
	GenerateFromPassword([]byte) ([]byte, error)
	CompareHashAndPassword(hashedPassword, password []byte) bool
}

// MailService @todo сервис отправки писем подтверждения регистрации
type MailService interface {
	SendRegistrationMail(context.Context, string) error
}

func NewAuthUseCase(hasher hasher, mailSender MailService, userClient *ent.UserClient, sessionClient *ent.SessionClient) (*AuthUseCase, error) {
	switch true {
	case hasher == nil:
		return nil, fmt.Errorf("password hash-function undefined")
	case mailSender == nil:
		return nil, fmt.Errorf("mailsend service undefined")
	case userClient == nil:
		return nil, fmt.Errorf("orm.user service undefined")
	case sessionClient == nil:
		return nil, fmt.Errorf("orm.session service undefined")
	}
	return &AuthUseCase{
		hasher:        hasher,
		mailService:   mailSender,
		userClient:    userClient,
		sessionClient: sessionClient,
	}, nil
}

func (a *AuthUseCase) SignUp(ctx context.Context, email, password, duplicate string) v1.UseCaseError {
	if password != duplicate {
		return NewAuthUseCaseError(http.StatusConflict, "пароли не совпадают")
	}
	{
		isExist, err := a.userClient.Query().Where(user.Email(email)).Exist(ctx)
		if err != nil {
			return NewAuthUseCaseError(http.StatusInternalServerError, err.Error())
		}
		if isExist {
			return NewAuthUseCaseError(http.StatusConflict, "регистрация для указанногот email невозможна")
		}
	}

	hashPass, hashPassError := a.hasher.GenerateFromPassword([]byte(password))
	if hashPassError != nil {
		return NewAuthUseCaseError(http.StatusInternalServerError, hashPassError.Error())
	}
	userEntity, userEntitySaveError := a.
		userClient.
		Create().
		SetEmail(email).
		SetActive(false).
		SetPassword(string(hashPass)).
		Save(ctx)

	if userEntitySaveError != nil {
		return NewAuthUseCaseError(http.StatusInternalServerError, userEntitySaveError.Error())
	}
	//@todo обработка ошибки отправки письма
	if sendMailError := a.mailService.SendRegistrationMail(ctx, userEntity.Email); sendMailError != nil {
		fmt.Println("send registration mail error:", sendMailError)
	}
	return NewAuthUseCaseError(http.StatusCreated, "учётная запись создана")
}
