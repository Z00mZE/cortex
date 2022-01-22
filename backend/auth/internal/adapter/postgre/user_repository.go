package postgre

import (
	"context"
	"fmt"
	"github.com/Z00mZE/cortex/backend/auth/internal/entity"
	"github.com/Z00mZE/cortex/backend/auth/pkg/postgres"
	"github.com/jackc/pgx/v4"
	"time"
)

type UserRepository struct {
	connection *postgres.Postgres
}

func NewUserRepository(pg *postgres.Postgres) *UserRepository {
	return &UserRepository{connection: pg}
}

func (u *UserRepository) FindByEmailAndPassword(ctx context.Context, email, password string) (entity.User, bool, error) {
	var userEntity entity.User
	query := `SELECT * FROM users WHERE username = $1 AND password = $2`
	result := u.connection.Pool.QueryRow(ctx, query, email, password)
	if rowScanError := result.Scan(&userEntity); rowScanError != nil {
		if rowScanError == pgx.ErrNoRows {
			return userEntity, false, nil
		}
		return userEntity, false, rowScanError
	}
	return userEntity, true, nil
}

func (u *UserRepository) CreateUser(ctx context.Context, email, password string) error {
	now := time.Now().Format(time.RFC3339)
	query := "insert into users(email,password,username,created_at,updated_at) values($1,$2,$3,$4::date,$5::date)"
	_, err := u.connection.Pool.Exec(ctx, query, email, password, "john doe", now, now)
	fmt.Println("")
	fmt.Println("insert", err)
	fmt.Println("")
	return err
}

func (u *UserRepository) IsEmailExist(ctx context.Context, email string) (bool, error) {
	query := "SELECT FROM users WHERE email = $1 LIMIT 1"
	result, err := u.connection.Pool.Exec(ctx, query, email)
	return result.RowsAffected() != 0, err
}
