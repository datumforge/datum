package schema

import (
	"context"
	"net/mail"
	"regexp"

	"entgo.io/contrib/entgql"
	"entgo.io/contrib/entoas"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"

	"github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/ent/generated/privacy"
	"github.com/datumforge/datum/internal/ent/hooks"
	"github.com/datumforge/datum/internal/ent/interceptors"
	"github.com/datumforge/datum/internal/ent/mixin"
	"github.com/datumforge/datum/internal/ent/privacy/rule"
	"github.com/datumforge/datum/internal/ent/privacy/token"
	"github.com/datumforge/fgax/entfga"
)

// Subscriber holds the schema definition for the Subscriber entity
type Subscriber struct {
	ent.Schema
}

// Fields of the Subscriber
func (Subscriber) Fields() []ent.Field {
	return []ent.Field{
		field.String("email").
			Comment("email address of the subscriber").
			Validate(func(email string) error {
				_, err := mail.ParseAddress(email)
				return err
			}),
		field.String("phone_number").
			Comment("phone number of the subscriber").
			Optional().
			Validate(func(phone string) error {
				regex := `^\+[1-9]{1}[0-9]{3,14}$`
				_, err := regexp.MatchString(regex, phone)
				return err
			}),
		field.Bool("verified_email").
			Comment("indicates if the email address has been verified").
			Default(false),
		field.Bool("verified_phone").
			Comment("indicates if the phone number has been verified").
			Default(false),
		field.Bool("active").
			Comment("indicates if the subscriber is active or not, active users will have at least one verified contact method").
			Default(false),
		field.String("token").
			Comment("the verification token sent to the user via email which should only be provided to the /subscribe endpoint + handler").
			Unique().
			Annotations(entgql.Skip()).
			NotEmpty(),
		field.Time("ttl").
			Comment("the ttl of the verification token which defaults to 7 days").
			Annotations(entgql.Skip()).
			Nillable(),
		field.Bytes("secret").
			Comment("the comparison secret to verify the token's signature").
			NotEmpty().
			Annotations(entgql.Skip()).
			Nillable(),
	}
}

// Mixin of the Subscriber
func (Subscriber) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.AuditMixin{},
		mixin.IDMixin{},
		mixin.SoftDeleteMixin{},
		OrgOwnerMixin{ // empty org means Datum system Subscriber
			Ref:        "subscribers",
			Optional:   true,
			AllowWhere: true,
		},
	}
}

// Edges of the Subscriber
func (Subscriber) Edges() []ent.Edge {
	return []ent.Edge{
		// Edges go here
	}
}

func (Subscriber) Hooks() []ent.Hook {
	return []ent.Hook{
		hooks.HookSubscriber(),
	}
}

// Indexes of the Subscriber
func (Subscriber) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("email", "owner_id").
			Unique().
			Annotations(
				entsql.IndexWhere("deleted_at is NULL"),
			),
	}
}

// Annotations of the Subscriber
func (Subscriber) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
		entgql.RelayConnection(),
		entgql.Mutations(entgql.MutationCreate(), (entgql.MutationUpdate())),
		entoas.CreateOperation(entoas.OperationPolicy(entoas.PolicyExclude)),
		entoas.UpdateOperation(entoas.OperationPolicy(entoas.PolicyExclude)),
		entoas.DeleteOperation(entoas.OperationPolicy(entoas.PolicyExclude)),
		entfga.Annotations{
			ObjectType:      "organization",
			IncludeHooks:    false,
			IDField:         "OwnerID",
			NillableIDField: true,
		},
	}
}

// Interceptors of the Subscriber
func (Subscriber) Interceptors() []ent.Interceptor {
	return []ent.Interceptor{
		interceptors.InterceptorSubscriber(),
	}
}

// Policy of the Subscriber
func (Subscriber) Policy() ent.Policy {
	return privacy.Policy{
		Mutation: privacy.MutationPolicy{
			rule.AllowIfContextHasPrivacyTokenOfType(&token.SignUpToken{}),
			rule.AllowIfContextHasPrivacyTokenOfType(&token.VerifyToken{}),
			privacy.SubscriberMutationRuleFunc(func(ctx context.Context, m *generated.SubscriberMutation) error {
				return m.CheckAccessForEdit(ctx)
			}),
			privacy.AlwaysDenyRule(),
		},
		Query: privacy.QueryPolicy{
			rule.AllowIfContextHasPrivacyTokenOfType(&token.SignUpToken{}),
			rule.AllowIfContextHasPrivacyTokenOfType(&token.VerifyToken{}),
			privacy.SubscriberQueryRuleFunc(func(ctx context.Context, q *generated.SubscriberQuery) error {
				return q.CheckAccess(ctx)
			}),
			privacy.AlwaysDenyRule(),
		},
	}
}
