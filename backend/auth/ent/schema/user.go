package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
	Username  string
	EMail     string
	Password  string
	Active    bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.
			String("username").
			Unique().
			Comment("username"),
		field.
			String("email").
			Unique().
			Comment("email"),
		field.
			String("password").
			Comment("password").
			Sensitive(),
		field.Bool("active").
			Default(false).
			Comment("account activated status"),
		field.
			Time("created_at").
			Default(time.Now).
			Immutable(),
		field.
			Time("updated_at").
			Default(time.Now).
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.
			To("sessions", Session.Type),
	}
}
