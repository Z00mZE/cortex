package user

import (
	"github.com/Z00mZE/cortex/backend/auth/internal/cqrs/user/command"
	"github.com/Z00mZE/cortex/backend/auth/internal/cqrs/user/query"
)

type User struct {
	Query   Query
	Command Command
}

type Query struct {
	IsEmailRegistered             *query.IsEmailRegisteredHandler
	FindByEmailAndPasswordHandler *query.FindByEmailAndPasswordHandler
}
type Command struct {
	CreateUserHandler *command.CreateUserHandler
}
