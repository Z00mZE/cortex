package v1

import (
	"context"
	"fmt"
	"github.com/Z00mZE/cortex/internal/auth/pkg/routing"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

type AuthUseCase interface {
	SignUp(ctx context.Context, email, password string) error
	//Confirm(ctx context.Context, token string) error
	//SignIn(ctx context.Context, email, password string) (accessToken, refreshToken string, signInError error)
	//ResetPassword(ctx context.Context, email string) error
	//Recovery(ctx context.Context, token, password string) error
	//User(ctx context.Context, token string) (entity.User, error)
	//RefreshToken(ctx context.Context, token string) (accessToken, refreshToken string, refreshError error)
}

type signUpRequest struct {
	Email             string `json:"email"`
	Password          string `json:"password"`
	PasswordDuplicate string `json:"password_duplicate"`
}

func NewRouter(handler *echo.Echo, useCases AuthUseCase) {
	routing.BindHttpRouteGroup(
		handler,
		routing.NewHttpRouteGroup(
			"/auth/v1",
			[]routing.IHttpHandler{
				routing.NewHttpRouteHandler(
					"/sign-up",
					http.MethodPost,
					func(ctx echo.Context) error {
						var req signUpRequest
						if err := ctx.Bind(&req); err != nil {
							return echo.NewHTTPError(signUpValidationError.Code, signUpValidationError)
						}
						err := useCases.SignUp(context.Background(), req.Email, req.Password)
						fmt.Println("signUp error", err)
						return ctx.JSON(http.StatusCreated, nil)
					},
				),
			},
			[]echo.MiddlewareFunc{
				middleware.Logger(),
				middleware.Recover(),
				middleware.RequestID(),
			},
		),
	)
}
