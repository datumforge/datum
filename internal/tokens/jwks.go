package tokens

import (
	"context"
	"fmt"

	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/lestrrat-go/jwx/v2/jwk"
)

// JWKSValidator provides public verification that JWT tokens have been issued by the
// Datum authentication service by checking that the tokens have been signed using
// public keys from a JSON Web Key Set (JWKS). The validator then returns Datum
// specific claims if the token is in fact valid.
type JWKSValidator struct {
	validator
	keys jwk.Set
}

func NewJWKSValidator(keys jwk.Set, audience, issuer string) *JWKSValidator {
	validator := &JWKSValidator{
		validator: validator{
			audience: audience,
			issuer:   issuer,
		},
		keys: keys,
	}
	validator.validator.keyFunc = validator.keyFunc

	return validator
}

// keyFunc is an jwt.KeyFunc that selects the RSA public key from the list of managed
// internal keys based on the kid in the token header. If the kid does not exist an
// error is returned and the token will not be able to be verified.
func (v *JWKSValidator) keyFunc(token *jwt.Token) (publicKey interface{}, err error) {
	// Fetch the kid from the header
	kid, ok := token.Header["kid"]
	if !ok {
		return nil, ErrTokenMissingKid
	}

	key, found := v.keys.LookupKeyID(kid.(string))
	if !found {
		return nil, ErrUnknownSigningKey
	}

	// Per JWT security notice: do not forget to validate alg is expected
	if token.Method.Alg() != key.Algorithm().String() {
		return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"]) //nolint:goerr113
	}

	// Extract the raw public key from the key material and return it.
	if err = key.Raw(&publicKey); err != nil {
		return nil, fmt.Errorf("could not extract raw key: %w", err)
	}

	return publicKey, nil
}

type CachedJWKSValidator struct {
	JWKSValidator
	cache    *jwk.Cache
	endpoint string
}

func NewCachedJWKSValidator(ctx context.Context, cache *jwk.Cache, endpoint, audience, issuer string) (validator *CachedJWKSValidator, err error) {
	validator = &CachedJWKSValidator{
		cache:    cache,
		endpoint: endpoint,
	}

	var keys jwk.Set

	if keys, err = cache.Refresh(ctx, endpoint); err != nil {
		return nil, err
	}

	validator.JWKSValidator = *NewJWKSValidator(keys, audience, issuer)
	validator.validator.keyFunc = validator.keyFunc

	return validator, nil
}

func (v *CachedJWKSValidator) Refresh(ctx context.Context) (err error) {
	if v.JWKSValidator.keys, err = v.cache.Refresh(ctx, v.endpoint); err != nil {
		return fmt.Errorf("could not refresh cache from %s: %w", v.endpoint, err)
	}

	return nil
}

func (v *CachedJWKSValidator) keyFunc(token *jwt.Token) (publicKey interface{}, err error) {
	if v.JWKSValidator.keys, err = v.cache.Get(context.Background(), v.endpoint); err != nil {
		return nil, fmt.Errorf("could not retrieve JWKS from cache: %w", err)
	}

	return v.JWKSValidator.keyFunc(token)
}
