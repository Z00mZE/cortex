package postgre

import (
	"context"
	"github.com/Z00mZE/cortex/internal/auth/internal/entity"
	"github.com/Z00mZE/cortex/internal/auth/pkg/postgres"
)

type UserRepository struct {
	connection *postgres.Postgres
}

func NewUserRepository(pg *postgres.Postgres) *UserRepository {
	return new(UserRepository)
}

func (u UserRepository) FindByEmailAndPassword(ctx context.Context, email, password string) (entity.User, error) {
	panic("implement me")
}
