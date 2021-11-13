package entity

import "time"

type User struct {
	Id        string
	Username  string
	Email     string
	Roles     []string
	CreatedAt time.Time
	UpdatedAt time.Time
}
