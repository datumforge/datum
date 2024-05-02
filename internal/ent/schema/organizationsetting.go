package schema

import (
	"context"
	"fmt"
	"net/mail"
	"net/url"
	"regexp"

	"entgo.io/contrib/entgql"
	"entgo.io/contrib/entoas"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"

	emixin "github.com/datumforge/entx/mixin"
	"github.com/datumforge/fgax/entfga"

	"github.com/datumforge/datum/internal/ent/enums"
	"github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/ent/generated/privacy"
	"github.com/datumforge/datum/internal/ent/interceptors"
	"github.com/datumforge/datum/internal/ent/mixin"
	"github.com/datumforge/datum/internal/ent/privacy/rule"
)

// OrganizationSetting holds the schema definition for the OrganizationSetting entity
type OrganizationSetting struct {
	ent.Schema
}

// Fields of the OrganizationSetting
func (OrganizationSetting) Fields() []ent.Field {
	return []ent.Field{
		field.Strings("domains").
			Comment("domains associated with the organization").
			Validate(func(domains []string) error {
				for _, domain := range domains {
					u, err := url.Parse("http://" + domain)
					if err != nil || u.Scheme == "" || u.Host == "" {
						return fmt.Errorf("invalid domain: %s", domain) // nolint: goerr113
					}
				}
				return nil
			}).
			Optional(),
		field.String("billing_contact").
			Comment("Name of the person to contact for billing").
			Optional(),
		field.String("billing_email").
			Comment("Email address of the person to contact for billing").
			Validate(func(email string) error {
				_, err := mail.ParseAddress(email)
				return err
			}).
			Optional(),
		field.String("billing_phone").
			Comment("Phone number to contact for billing").
			Validate(func(phone string) error {
				regex := `^\+[1-9]{1}[0-9]{3,14}$`
				_, err := regexp.MatchString(regex, phone)
				return err
			}).
			Optional(),
		field.String("billing_address").
			Comment("Address to send billing information to").
			Optional(),
		field.String("tax_identifier").
			Comment("Usually government-issued tax ID or business ID such as ABN in Australia").
			Optional(),
		field.Strings("tags").
			Comment("tags associated with the object").
			Default([]string{}).
			Optional(),
		field.Enum("geo_location").
			GoType(enums.Region("")).
			Comment("geographical location of the organization").
			Default(enums.Amer.String()).
			Optional(),
		field.String("organization_id").
			Comment("the ID of the organization the settings belong to").
			Optional(),
	}
}

// Edges of the OrganizationSetting
func (OrganizationSetting) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("organization", Organization.Type).Ref("setting").Field("organization_id").Unique(),
	}
}

// Annotations of the OrganizationSetting
func (OrganizationSetting) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
		entgql.RelayConnection(),
		entgql.Mutations(entgql.MutationCreate(), (entgql.MutationUpdate())),
		entoas.Skip(true),
		entoas.CreateOperation(entoas.OperationPolicy(entoas.PolicyExclude)),
		entoas.ReadOperation(entoas.OperationPolicy(entoas.PolicyExclude)),
		entoas.UpdateOperation(entoas.OperationPolicy(entoas.PolicyExclude)),
		entoas.DeleteOperation(entoas.OperationPolicy(entoas.PolicyExclude)),
		entoas.ListOperation(entoas.OperationPolicy(entoas.PolicyExclude)),
		entfga.Annotations{
			ObjectType:      "organization",
			IncludeHooks:    false,
			IDField:         "OrganizationID",
			NillableIDField: true,
		},
	}
}

// Interceptors of the Subscriber
func (OrganizationSetting) Interceptors() []ent.Interceptor {
	return []ent.Interceptor{
		interceptors.InterceptorOrganizationSetting(),
	}
}

// Mixin of the OrganizationSetting
func (OrganizationSetting) Mixin() []ent.Mixin {
	return []ent.Mixin{
		emixin.AuditMixin{},
		emixin.IDMixin{},
		mixin.SoftDeleteMixin{},
	}
}

// Policy defines the privacy policy of the OrganizationSetting
func (OrganizationSetting) Policy() ent.Policy {
	return privacy.Policy{
		Mutation: privacy.MutationPolicy{
			rule.DenyIfNoSubject(),
			privacy.OrganizationSettingMutationRuleFunc(func(ctx context.Context, m *generated.OrganizationSettingMutation) error {
				return m.CheckAccessForEdit(ctx)
			}),
			privacy.AlwaysDenyRule(),
		},
		Query: privacy.QueryPolicy{
			privacy.OrganizationSettingQueryRuleFunc(func(ctx context.Context, q *generated.OrganizationSettingQuery) error {
				return q.CheckAccess(ctx)
			}),
			privacy.AlwaysDenyRule(),
		},
	}
}
