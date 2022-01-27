package v1

import (
	"context"
	"fmt"
	"github.com/Z00mZE/cortex/backend/auth/pkg/routing"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

type UseCaseError interface {
	StatusCode() int
	Message() string
}

type AuthUseCase interface {
	SignUp(ctx context.Context, email, password, duplicate string) UseCaseError
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

func NewRouter(app *echo.Echo, useCases AuthUseCase) {
	routing.BindHttpRouteGroup(
		app,
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
						err := useCases.SignUp(
							context.Background(),
							req.Email,
							req.Password,
							req.PasswordDuplicate,
						)
						fmt.Println("signUp error", err)
						return ctx.JSON(err.StatusCode(), echo.Map{"message": err.Message()})
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
