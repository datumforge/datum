package rule

import (
	"context"
	"reflect"

	"github.com/datumforge/datum/internal/ent/generated/privacy"
	"github.com/datumforge/datum/internal/ent/privacy/token"
)

// AllowIfContextHasPrivacyTokenOfType allows a mutation
// to proceed if a privacy token of a specific type is found in the
// context. It checks if the actual type of the token in the context
// matches the expected type, and if so, it returns `privacy.Allow`.
// If the types do not match, it returns `privacy.Skipf` with a message
// indicating that no token was found in the context with the expected type
func AllowIfContextHasPrivacyTokenOfType(emptyToken token.PrivacyToken) privacy.QueryMutationRule {
	return privacy.ContextQueryMutationRule(func(ctx context.Context) error {
		actualTokenType := reflect.TypeOf(ctx.Value(emptyToken.GetContextKey()))
		expectedTokenType := reflect.TypeOf(emptyToken)
		if actualTokenType == expectedTokenType {
			return privacy.Allow
		}

		return privacy.Skipf("no token found from context with type %T", emptyToken)
	})
}

// AllowAfterApplyingPrivacyTokenFilter allows the mutation to proceed
// if a privacy token of a specific type is found in the context. It
// also applies a privacy filter to the token before allowing the
// mutation to proceed
func AllowAfterApplyingPrivacyTokenFilter(
	emptyToken token.PrivacyToken,
	applyFilter func(t token.PrivacyToken, filter privacy.Filter),
) privacy.QueryMutationRule {
	return privacy.FilterFunc(
		func(ctx context.Context, filter privacy.Filter) error {
			actualToken := ctx.Value(emptyToken.GetContextKey())
			actualTokenType := reflect.TypeOf(actualToken)
			expectedTokenType := reflect.TypeOf(emptyToken)
			if actualTokenType == expectedTokenType {
				applyFilter(actualToken.(token.PrivacyToken), filter)
				return privacy.Allowf("applied privacy token filter")
			}

			return privacy.Skipf("no token found from context with type %T", emptyToken)
		})
}
