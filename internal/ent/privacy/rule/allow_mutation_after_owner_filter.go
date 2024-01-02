package rule

import (
	"context"

	"github.com/datumforge/datum/internal/ent/generated/predicate"
	"github.com/datumforge/datum/internal/ent/generated/privacy"
	"github.com/datumforge/datum/internal/ent/generated/user"
	"github.com/datumforge/datum/internal/ent/privacy/viewer"
)

// AllowMutationAfterApplyingOwnerFilter defines a privacy rule for mutations in the context of an owner filter
func AllowMutationAfterApplyingOwnerFilter() privacy.MutationRule {
	type OwnerFilter interface {
		WhereHasOwnerWith(predicates ...predicate.User)
	}

	return privacy.FilterFunc(
		func(ctx context.Context, f privacy.Filter) error {
			v := viewer.FromContext(ctx)

			ownerFilter, ok := f.(OwnerFilter)
			if !ok {
				return privacy.Deny
			}

			viewerID, exists := v.GetID()
			if !exists {
				return privacy.Skip
			}

			ownerFilter.WhereHasOwnerWith(user.ID(viewerID))
			return privacy.Allowf("applied owner filter")
		},
	)
}
