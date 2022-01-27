package postgres

import (
	"context"
	"github.com/Z00mZE/cortex/backend/auth/ent"
)

//NewPostgresORM -.
func NewPostgresORM(url string, opts ...ent.DBOption) (*ent.Client, error) {
	orm, ormError := ent.Open(ent.PgPgx, url, opts...)
	if ormError != nil {
		return nil, ormError
	}
	if migrateError := orm.Schema.Create(context.Background()); migrateError != nil {
		return nil, migrateError
	}
	return orm, nil
}
