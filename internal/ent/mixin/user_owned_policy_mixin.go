package mixin

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/mixin"

	"github.com/datumforge/datum/internal/ent/generated/privacy"
	"github.com/datumforge/datum/internal/ent/privacy/rule"
	"github.com/datumforge/datum/internal/ent/privacy/utils"
)

type (
	UserOwnedMutationPolicyMixin struct {
		mixin.Schema
		AllowAdminMutation bool
	}

	UserOwnedQueryPolicyMixin struct {
		mixin.Schema
	}
)

func (mixin UserOwnedMutationPolicyMixin) Policy() ent.Policy {
	return privacy.Policy{
		Mutation: privacy.MutationPolicy{
			privacy.OnMutationOperation(
				utils.NewMutationPolicyWithoutNil(privacy.MutationPolicy{
					rule.DenyIfNoViewer(),
					rule.AllowMutationAfterApplyingOwnerFilter(),
					privacy.AlwaysDenyRule(),
				}),
				ent.OpCreate,
			),
			privacy.OnMutationOperation(
				utils.NewMutationPolicyWithoutNil(privacy.MutationPolicy{
					rule.DenyIfNoViewer(),
					rule.AllowMutationAfterApplyingOwnerFilter(),
					privacy.AlwaysDenyRule(),
				}),
				ent.OpUpdateOne|ent.OpUpdate|ent.OpDeleteOne|ent.OpDelete,
			),
		},
	}
}

func (mixin UserOwnedQueryPolicyMixin) Policy() ent.Policy {
	return privacy.Policy{
		Query: privacy.QueryPolicy{
			rule.AllowIfAdmin(),
			rule.AllowIfOwnedByViewer(),
			privacy.AlwaysDenyRule(),
		},
	}
}
