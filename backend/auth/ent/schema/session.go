package schema

import (
	"database/sql/driver"
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgtype"
	"net"
	"time"
)

// Session holds the schema definition for the Session entity.
type Session struct {
	ent.Schema
	UserID                   int
	SessionID                uuid.UUID
	AccessToken              uuid.UUID
	AccessTokenExpirationAt  time.Time
	RefreshToken             uuid.UUID
	RefreshTokenExpirationAt time.Time
	IP                       *pgtype.Inet
	Agent                    string
}

// Inet represents a single IP address
type Inet struct {
	net.IP
}

// Scan implements the Scanner interface
func (i *Inet) Scan(value interface{}) (err error) {
	switch v := value.(type) {
	case nil:
	case []byte:
		if i.IP = net.ParseIP(string(v)); i.IP == nil {
			err = fmt.Errorf("invalid value for ip %q", v)
		}
	case string:
		if i.IP = net.ParseIP(v); i.IP == nil {
			err = fmt.Errorf("invalid value for ip %q", v)
		}
	default:
		err = fmt.Errorf("unexpected type %T", v)
	}
	return
}

// Value implements the driver Valuer interface
func (i Inet) Value() (driver.Value, error) {
	return i.IP.String(), nil
}

// Fields of the Session.
func (Session) Fields() []ent.Field {
	return []ent.Field{
		field.
			Int("user_id").Optional(),
		field.
			UUID("session_id", uuid.UUID{}).
			Default(uuid.New),
		field.
			String("access_token"),
		field.
			Time("access_token_expiration_at"),
		field.
			String("refresh_token"),
		field.
			Time("refresh_token_expiration_at"),
		field.
			String("ip").
			GoType(&Inet{}).
			SchemaType(map[string]string{
				dialect.Postgres: "inet",
			}).
			Validate(func(s string) error {
				if net.ParseIP(s) == nil {
					return fmt.Errorf("invalid value for ip %q", s)
				}
				return nil
			}),
		field.
			String("agent"),
		field.
			Time("created_at").
			Default(time.Now).
			Immutable(),
		field.
			Time("last_activity_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

// Edges of the Session.
func (Session) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("sessions").
			Field("user_id").
			Unique(),
	}
}
