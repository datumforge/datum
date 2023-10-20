package schema

import (
	"encoding/base64"
	"fmt"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
	"lukechampine.com/frand"
)

// Sessions are authentication sessions. They can either be first-party web auth sessions or OAuth sessions.
// Sessions should persist in the database for some time duration after expiration, but with the "disabled" boolean set to true.
type Session struct {
	ent.Schema
}

func (Session) Fields() []ent.Field {
	return []ent.Field{
		// NOTE: the created_at and updated_at fields are automatically created by the AuditMixin, you do not need to re-declare / add them in these fields
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Unique(),
		field.Enum("type").
			Comment("Sessions can derrive from the local (password auth), oauth, or app_password").
			Values(
				"local",
				"oauth",
				"app_password",
			).
			Immutable(),

		field.Bool("disabled").
			Comment("The session may be disabled by the user or by automatic security policy"),

		// Session expiry can be determined by the application at runtime based on the created_at field.

		field.String("token").
			Comment("random 32 bytes encoded as base64").
			Unique().
			Immutable().
			DefaultFunc(func() string {
				b := make([]byte, 20)
				_, _ = frand.Read(b)
				out := make([]byte, 27)
				base64.RawStdEncoding.Encode(out, b)
				return string(out)
			}).
			Validate(func(s string) error {
				v, err := base64.RawStdEncoding.DecodeString(s)
				if err != nil {
					return err
				}
				if len(v) != 32 {
					return fmt.Errorf("invalid token size")
				}
				return nil
			}),

		// TODO: OAuth fields
		field.String("user_agent").
			Comment("The last known user-agent").
			Optional(),

		field.String("ips").
			Comment("All IPs that have been associated with this session. Reverse-chornological order. The current IP is the first item in the slice"),
	}
}

// Indexes of the Session
func (Session) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id").
			Unique(), // enforce globally unique ids
	}
}

// Edges of the Session
func (Session) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("users", User.Type).
			Unique().
			Comment("Sessions belong to users"),
	}
}

// Mixins of the Session
func (Session) Mixin() []ent.Mixin {
	return []ent.Mixin{
		AuditMixin{},
	}
}
