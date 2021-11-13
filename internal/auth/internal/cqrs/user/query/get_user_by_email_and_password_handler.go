package query

import (
	"context"
	"github.com/Z00mZE/cortex/internal/auth/internal/entity"
)

type FindByEmailAndPasswordReadModel interface {
	FindByEmailAndPassword(ctx context.Context, email, password string) (entity.User, error)
}
type FindByEmailAndPasswordHandler struct {
	repository FindByEmailAndPasswordReadModel
}

func NewFindByEmailAndPasswordHandler(repo FindByEmailAndPasswordReadModel) *FindByEmailAndPasswordHandler {
	return &FindByEmailAndPasswordHandler{repository: repo}
}
func (h *FindByEmailAndPasswordHandler) Handler(ctx context.Context, email, password string) (entity.User, error) {
	return h.repository.FindByEmailAndPassword(ctx, email, password)
}
