package rule

import (
	"context"

	"github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/ent/generated/privacy"
	"github.com/datumforge/datum/internal/ent/privacy/token"
)

type MutationEmailGetter func(generated.Mutation) (string, error)

// AllowMutationIfContextHasValidEmailSignUpToken is used to determine whether a
// mutation should be allowed or skipped based on the presence and validity of an
// email signup token in the context
func AllowMutationIfContextHasValidEmailSignUpToken(getEmail MutationEmailGetter) privacy.MutationRule {
	return privacy.MutationRuleFunc(
		func(ctx context.Context, mutation generated.Mutation) error {
			emailSignupToken := token.EmailSignUpTokenFromContext(ctx)
			if emailSignupToken == nil {
				return privacy.Skipf("email signup token not found in context")
			}

			email, err := getEmail(mutation)
			if err != nil {
				return privacy.Skipf("unable to obtain email from mutation")
			}

			if email != emailSignupToken.GetEmail() {
				return privacy.Skipf("email sign up token does not match mutation result")
			}

			return privacy.Allow
		},
	)
}
