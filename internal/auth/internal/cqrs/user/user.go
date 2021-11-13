package user

import (
	"github.com/Z00mZE/cortex/internal/auth/internal/cqrs/user/command"
	"github.com/Z00mZE/cortex/internal/auth/internal/cqrs/user/query"
)

type User struct {
	Query   Query
	Command Command
}

type Query struct {
	FindByEmailAndPasswordHandler *query.FindByEmailAndPasswordHandler
}
type Command struct {
	CreateUserHandler *command.CreateUserHandler
}
