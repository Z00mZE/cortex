package usecase

type AuthUseCaseError struct {
	code    int
	message string
}

func (a *AuthUseCaseError) StatusCode() int {
	return a.code
}

func (a *AuthUseCaseError) Message() string {
	return a.message
}

func (a *AuthUseCaseError) Error() string {
	return a.message
}

func (a *AuthUseCaseError) RuntimeError() {
	panic("implement me")
}

func NewAuthUseCaseError(code int, message string) *AuthUseCaseError {
	return &AuthUseCaseError{code: code, message: message}
}
