package query

import "context"

type IsEmailExistChecker interface {
	IsEmailExist(ctx context.Context, email string) (bool, error)
}
type IsEmailRegisteredHandler struct {
	repository IsEmailExistChecker
}

func NewIsEmailRegistered(repository IsEmailExistChecker) *IsEmailRegisteredHandler {
	return &IsEmailRegisteredHandler{repository: repository}
}

func (h *IsEmailRegisteredHandler) Handler(ctx context.Context, email string) (bool, error) {
	return h.repository.IsEmailExist(ctx, email)
}
