// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/Z00mZE/cortex/backend/auth/ent/schema"
	"github.com/Z00mZE/cortex/backend/auth/ent/session"
	"github.com/Z00mZE/cortex/backend/auth/ent/user"
	"github.com/google/uuid"
)

// Session is the model entity for the Session schema.
type Session struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID int `json:"user_id,omitempty"`
	// SessionID holds the value of the "session_id" field.
	SessionID uuid.UUID `json:"session_id,omitempty"`
	// AccessToken holds the value of the "access_token" field.
	AccessToken string `json:"access_token,omitempty"`
	// AccessTokenExpirationAt holds the value of the "access_token_expiration_at" field.
	AccessTokenExpirationAt time.Time `json:"access_token_expiration_at,omitempty"`
	// RefreshToken holds the value of the "refresh_token" field.
	RefreshToken string `json:"refresh_token,omitempty"`
	// RefreshTokenExpirationAt holds the value of the "refresh_token_expiration_at" field.
	RefreshTokenExpirationAt time.Time `json:"refresh_token_expiration_at,omitempty"`
	// IP holds the value of the "ip" field.
	IP *schema.Inet `json:"ip,omitempty"`
	// Agent holds the value of the "agent" field.
	Agent string `json:"agent,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// LastActivityAt holds the value of the "last_activity_at" field.
	LastActivityAt time.Time `json:"last_activity_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the SessionQuery when eager-loading is set.
	Edges SessionEdges `json:"edges"`
}

// SessionEdges holds the relations/edges for other nodes in the graph.
type SessionEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e SessionEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.User == nil {
			// The edge user was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Session) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case session.FieldIP:
			values[i] = new(schema.Inet)
		case session.FieldID, session.FieldUserID:
			values[i] = new(sql.NullInt64)
		case session.FieldAccessToken, session.FieldRefreshToken, session.FieldAgent:
			values[i] = new(sql.NullString)
		case session.FieldAccessTokenExpirationAt, session.FieldRefreshTokenExpirationAt, session.FieldCreatedAt, session.FieldLastActivityAt:
			values[i] = new(sql.NullTime)
		case session.FieldSessionID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Session", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Session fields.
func (s *Session) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case session.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			s.ID = int(value.Int64)
		case session.FieldUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				s.UserID = int(value.Int64)
			}
		case session.FieldSessionID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field session_id", values[i])
			} else if value != nil {
				s.SessionID = *value
			}
		case session.FieldAccessToken:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field access_token", values[i])
			} else if value.Valid {
				s.AccessToken = value.String
			}
		case session.FieldAccessTokenExpirationAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field access_token_expiration_at", values[i])
			} else if value.Valid {
				s.AccessTokenExpirationAt = value.Time
			}
		case session.FieldRefreshToken:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field refresh_token", values[i])
			} else if value.Valid {
				s.RefreshToken = value.String
			}
		case session.FieldRefreshTokenExpirationAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field refresh_token_expiration_at", values[i])
			} else if value.Valid {
				s.RefreshTokenExpirationAt = value.Time
			}
		case session.FieldIP:
			if value, ok := values[i].(*schema.Inet); !ok {
				return fmt.Errorf("unexpected type %T for field ip", values[i])
			} else if value != nil {
				s.IP = value
			}
		case session.FieldAgent:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field agent", values[i])
			} else if value.Valid {
				s.Agent = value.String
			}
		case session.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				s.CreatedAt = value.Time
			}
		case session.FieldLastActivityAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field last_activity_at", values[i])
			} else if value.Valid {
				s.LastActivityAt = value.Time
			}
		}
	}
	return nil
}

// QueryUser queries the "user" edge of the Session entity.
func (s *Session) QueryUser() *UserQuery {
	return (&SessionClient{config: s.config}).QueryUser(s)
}

// Update returns a builder for updating this Session.
// Note that you need to call Session.Unwrap() before calling this method if this Session
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Session) Update() *SessionUpdateOne {
	return (&SessionClient{config: s.config}).UpdateOne(s)
}

// Unwrap unwraps the Session entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *Session) Unwrap() *Session {
	tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Session is not a transactional entity")
	}
	s.config.driver = tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Session) String() string {
	var builder strings.Builder
	builder.WriteString("Session(")
	builder.WriteString(fmt.Sprintf("id=%v", s.ID))
	builder.WriteString(", user_id=")
	builder.WriteString(fmt.Sprintf("%v", s.UserID))
	builder.WriteString(", session_id=")
	builder.WriteString(fmt.Sprintf("%v", s.SessionID))
	builder.WriteString(", access_token=")
	builder.WriteString(s.AccessToken)
	builder.WriteString(", access_token_expiration_at=")
	builder.WriteString(s.AccessTokenExpirationAt.Format(time.ANSIC))
	builder.WriteString(", refresh_token=")
	builder.WriteString(s.RefreshToken)
	builder.WriteString(", refresh_token_expiration_at=")
	builder.WriteString(s.RefreshTokenExpirationAt.Format(time.ANSIC))
	builder.WriteString(", ip=")
	builder.WriteString(fmt.Sprintf("%v", s.IP))
	builder.WriteString(", agent=")
	builder.WriteString(s.Agent)
	builder.WriteString(", created_at=")
	builder.WriteString(s.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", last_activity_at=")
	builder.WriteString(s.LastActivityAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Sessions is a parsable slice of Session.
type Sessions []*Session

func (s Sessions) config(cfg config) {
	for _i := range s {
		s[_i].config = cfg
	}
}