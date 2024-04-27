package schema

import (
	"context"
	"encoding/hex"
	"fmt"

	"entgo.io/contrib/entgql"
	"entgo.io/contrib/entoas"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"gocloud.dev/secrets"

	emixin "github.com/datumforge/entx/mixin"

	"github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/ent/generated/hook"
	"github.com/datumforge/datum/internal/ent/generated/intercept"
	"github.com/datumforge/datum/internal/ent/mixin"
)

// Hush maps configured integrations (github, slack, etc.) to organizations
type Hush struct {
	ent.Schema
}

// Fields of the Hush
func (Hush) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Comment("the logical name of the corresponding hush secret or it's general grouping").
			NotEmpty().
			Annotations(
				entgql.OrderField("name"),
			),
		field.String("description").
			Comment("a description of the hush value or purpose, such as github PAT").
			Optional().
			Annotations(
				entgql.Skip(entgql.SkipWhereInput),
			),
		field.String("kind").
			Comment("the kind of secret, such as sshkey, certificate, api token, etc.").
			Optional().
			Annotations(
				entgql.OrderField("kind"),
			),
		field.String("secret_name").
			Comment("the generic name of a secret associated with the organization").
			Immutable().
			Optional(),
		field.String("secret_value").
			Comment("the secret value").
			Sensitive().
			Annotations(
				entgql.Skip(entgql.SkipWhereInput),
			).
			Optional().
			Immutable(),
	}
}

// Edges of the Hush
func (Hush) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("integrations", Integration.Type).
			Comment("the integration associated with the secret").
			Ref("secrets").
			Annotations(entoas.Skip(true)),
	}
}

// Annotations of the Hush
func (Hush) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
		entgql.RelayConnection(),
		entgql.Mutations(entgql.MutationCreate(), (entgql.MutationUpdate())),
	}
}

// Mixin of the Hush
func (Hush) Mixin() []ent.Mixin {
	return []ent.Mixin{
		emixin.AuditMixin{},
		emixin.IDMixin{},
		mixin.SoftDeleteMixin{},
	}
}

// Hooks of the Hush
func (Hush) Hooks() []ent.Hook {
	return []ent.Hook{
		hook.If(
			func(next ent.Mutator) ent.Mutator {
				return hook.HushFunc(func(ctx context.Context, m *generated.HushMutation) (generated.Value, error) {
					v, ok := m.SecretValue()
					if !ok || v == "" {
						return nil, fmt.Errorf("unexpected 'secret_name' value") // nolint: goerr113
					}

					c, err := m.Secrets.Encrypt(ctx, []byte(v))
					if err != nil {
						return nil, err
					}

					m.SetName(hex.EncodeToString(c))
					u, err := next.Mutate(ctx, m)
					if err != nil {
						return nil, err
					}

					if u, ok := u.(*generated.Hush); ok {
						err = decrypt(ctx, m.Secrets, u)
					}
					return u, err
				})
			},
			hook.HasFields("secret_value"),
		),
	}
}

// Interceptors of the User.
func (Hush) Interceptors() []ent.Interceptor {
	return []ent.Interceptor{
		ent.InterceptFunc(func(next ent.Querier) ent.Querier {
			return intercept.HushFunc(func(ctx context.Context, query *generated.HushQuery) (generated.Value, error) {
				v, err := next.Query(ctx, query)
				if err != nil {
					return nil, err
				}
				hush, ok := v.([]*generated.Hush)
				// Skip all query types besides node queries (e.g., Count, Scan, GroupBy).
				if !ok {
					return v, nil
				}
				for _, u := range hush {
					if err := decrypt(ctx, query.Secrets, u); err != nil {
						return nil, err
					}
				}
				return hush, nil
			})
		}),
	}
}

func decrypt(ctx context.Context, k *secrets.Keeper, u *generated.Hush) error {
	b, err := hex.DecodeString(u.SecretValue)
	if err != nil {
		return err
	}

	plain, err := k.Decrypt(ctx, b)
	if err != nil {
		return err
	}

	u.Name = string(plain)

	return nil
}
