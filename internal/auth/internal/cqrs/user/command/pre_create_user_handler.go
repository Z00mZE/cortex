package command

import "context"

type CreateUserMethod interface {
	CreateUser(ctx context.Context, email, password string) error
}
type CreateUserHandler struct {
	repository CreateUserMethod
}

func NewCreateUser(repo CreateUserMethod) *CreateUserHandler {
	return &CreateUserHandler{repository: repo}
}
func (h *CreateUserHandler) Handler(ctx context.Context, email, password string) error {
	return h.repository.CreateUser(ctx, email, password)
}
