package tokens

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"
	"io"
	"net/url"
	"os"
	"strings"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwk"
	"github.com/oklog/ulid/v2"

	"github.com/datumforge/datum/config"
)

const DefaultRefreshAudience = "https://auth.datum.net/v1/refresh"

// the signing method should match the value returned by the JWKS
var (
	signingMethod = jwt.SigningMethodRS256
	nilID         = ulid.ULID{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
)

// TokenManager handles the creation and verification of RSA signed JWT tokens. To
// facilitate signing key rollover, TokenManager can accept multiple keys identified by
// a ulid. JWT tokens generated by token managers include a kid ("Key ID") in the header that
// allows the token manager to verify the key with the specified signature. To sign keys
// the token manager will always use the latest private key by ulid.
//
// When the TokenManager creates tokens it will use JWT standard claims as well as
// extended claims based on Datum usage. The standard claims included are exp, nbf
// aud, and sub. On token verification, the exp, nbf, iss and aud claims are validated.

type TokenManager struct {
	validator
	refreshAudience string
	conf            config.TokenConfig
	currentKeyID    ulid.ULID
	currentKey      *rsa.PrivateKey
	keys            map[ulid.ULID]*rsa.PublicKey
	kidEntropy      io.Reader
}

var TimeFunc = time.Now

// Keyfunc will be used by the Parse methods as a callback function to supply the key for verification
type Keyfunc func(*Token) (interface{}, error)

// Token represents a JWT Token.  Different fields will be used depending on whether you're
// creating or parsing/verifying a token
type Token struct {
	// Raw is the raw token; populated when you parse a token
	Raw string
	// Method is the signing metehod of the token
	Method SigningMethod
	// Header is the first segment of the token
	Header map[string]interface{}
	// Claims is the second segment of the token
	Claims           Claims
	ClaimBytes       []byte
	ToBeSignedString string
	// Signature is the third segment of the token; populated when you parse a token
	Signature string
	// Valid is a bool determining if the token is valid; populated when you parse or verify a toekn
	Valid bool
}

// New creates a TokenManager with the specified keys which should be a mapping of ULID
// strings to paths to files that contain PEM encoded RSA private keys. This input is
// specifically designed for the config environment variable so that keys can be loaded
// from k8s or vault secrets that are mounted as files on disk
func New(conf config.TokenConfig) (tm *TokenManager, err error) {
	tm = &TokenManager{
		validator: validator{
			audience: conf.Audience,
			issuer:   conf.Issuer,
		},
		conf: conf,
		keys: make(map[ulid.ULID]*rsa.PublicKey),
		kidEntropy: &ulid.LockedMonotonicReader{
			MonotonicReader: ulid.Monotonic(rand.Reader, 0),
		},
	}
	tm.validator.keyFunc = tm.keyFunc

	for kid, path := range conf.Keys {
		var keyID ulid.ULID

		if keyID, err = ulid.Parse(kid); err != nil {
			return nil, newParseError("path", kid, err)
		}

		// Load the keys from disk
		var data []byte

		if data, err = os.ReadFile(path); err != nil {
			return nil, newParseError("path", kid, err)
		}

		var key *rsa.PrivateKey

		if key, err = jwt.ParseRSAPrivateKeyFromPEM(data); err != nil {
			return nil, newParseError("path", kid, err)
		}

		tm.keys[keyID] = &key.PublicKey

		// Set the current key if it is the latest key
		if tm.currentKey == nil || keyID.Time() > tm.currentKeyID.Time() {
			tm.currentKey = key
			tm.currentKeyID = keyID
		}
	}

	return tm, nil
}

// NewWithKey is a constructor function that creates a new instance of the TokenManager struct
// with a specified RSA private key. It takes in the private key as a parameter and initializes the
// TokenManager with the provided key, along with other configuration settings from the config.TokenConfig
// struct. It returns the created TokenManager instance or an error if there was a problem
// initializing the TokenManager.
func NewWithKey(key *rsa.PrivateKey, conf config.TokenConfig) (tm *TokenManager, err error) {
	tm = &TokenManager{
		validator: validator{
			audience: conf.Audience,
			issuer:   conf.Issuer,
		},
		conf: conf,
		keys: make(map[ulid.ULID]*rsa.PublicKey),
		kidEntropy: &ulid.LockedMonotonicReader{
			MonotonicReader: ulid.Monotonic(rand.Reader, 0),
		},
	}
	tm.validator.keyFunc = tm.keyFunc

	var kid ulid.ULID

	if kid, err = tm.genKeyID(); err != nil {
		return nil, err
	}

	tm.keys[kid] = &key.PublicKey
	tm.currentKey = key
	tm.currentKeyID = kid

	return tm, nil
}

// Sign an access or refresh token and return the token
func (tm *TokenManager) Sign(token *jwt.Token) (string, error) {
	if tm.currentKey == nil || tm.currentKeyID.Compare(nilID) == 0 {
		return "", ErrTokenManagerFailedInit
	}

	// Add the kid to the header
	token.Header["kid"] = tm.currentKeyID.String()

	// Return the signed string
	return token.SignedString(tm.currentKey)
}

// CreateTokenPair returns signed access and refresh tokens for the specified claims in one step since usually you want both access and refresh tokens at the same time
func (tm *TokenManager) CreateTokenPair(claims *Claims) (accessToken, refreshToken string, err error) {
	var atk, rtk *jwt.Token

	if atk, err = tm.CreateAccessToken(claims); err != nil {
		return "", "", fmt.Errorf("could not create access token: %w", err)
	}

	if rtk, err = tm.CreateRefreshToken(atk); err != nil {
		return "", "", fmt.Errorf("could not create refresh token: %w", err)
	}

	if accessToken, err = tm.Sign(atk); err != nil {
		return "", "", fmt.Errorf("could not sign access token: %w", err)
	}

	if refreshToken, err = tm.Sign(rtk); err != nil {
		return "", "", fmt.Errorf("could not sign refresh token: %w", err)
	}

	return
}

// CreateToken from the claims payload without modifying the claims unless the claims
// are missing required fields that need to be updated
func (tm *TokenManager) CreateToken(claims *Claims) *jwt.Token {
	if len(claims.Audience) == 0 {
		claims.Audience = jwt.ClaimStrings{tm.audience}
	}

	if claims.Issuer == "" {
		claims.Issuer = tm.issuer
	}

	return jwt.NewWithClaims(signingMethod, claims)
}

// CreateAccessToken from the credential payload or from an previous token if the access token is being reauthorized from previous credentials or an already issued access token
func (tm *TokenManager) CreateAccessToken(claims *Claims) (_ *jwt.Token, err error) {
	// Create the claims for the access token, using access token defaults
	now := time.Now()
	sub := claims.RegisteredClaims.Subject

	var kid ulid.ULID

	if kid, err = tm.genKeyID(); err != nil {
		return nil, err
	}

	issueTime := jwt.NewNumericDate(now)
	claims.RegisteredClaims = jwt.RegisteredClaims{
		ID:        strings.ToLower(kid.String()), // ID is randomly generated and shared between access and refresh
		Subject:   sub,
		Audience:  jwt.ClaimStrings{tm.audience},
		Issuer:    tm.issuer,
		IssuedAt:  issueTime,
		NotBefore: issueTime,
		ExpiresAt: jwt.NewNumericDate(now.Add(tm.conf.AccessDuration)),
	}

	return tm.CreateToken(claims), nil
}

// CreateRefreshToken from the Access token claims with predefined expiration
func (tm *TokenManager) CreateRefreshToken(accessToken *jwt.Token) (refreshToken *jwt.Token, err error) {
	accessClaims, ok := accessToken.Claims.(*Claims)
	if !ok {
		return nil, ErrFailedRetrieveClaimsFromToken
	}

	audience := accessClaims.Audience

	// Append the refresh token audience to the audience claims
	audience = append(audience, tm.RefreshAudience())

	// Create claims for the refresh token from the access token defaults
	claims := &Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        accessClaims.ID, // ID is randomly generated and shared between access and refresh tokens
			Audience:  audience,
			Issuer:    accessClaims.Issuer,
			Subject:   accessClaims.Subject,
			IssuedAt:  accessClaims.IssuedAt,
			NotBefore: jwt.NewNumericDate(accessClaims.ExpiresAt.Add(tm.conf.RefreshOverlap)),
			ExpiresAt: jwt.NewNumericDate(accessClaims.IssuedAt.Add(tm.conf.RefreshDuration)),
		},
		OrgID:       accessClaims.OrgID,
		ParentOrgID: accessClaims.ParentOrgID,
		Tier:        accessClaims.Tier,
	}

	return tm.CreateToken(claims), nil
}

// Keys returns the JWKS with public keys for use externally
func (tm *TokenManager) Keys() (keys jwk.Set, err error) {
	keys = jwk.NewSet()
	for kid, pubkey := range tm.keys {
		var key jwk.Key

		if key, err = jwk.FromRaw(pubkey); err != nil {
			return nil, err
		}

		if err = key.Set(jwk.KeyIDKey, kid.String()); err != nil {
			return nil, err
		}

		if err = key.Set(jwk.KeyUsageKey, jwk.ForSignature); err != nil {
			return nil, err
		}

		// NOTE: the algorithm should match the signing method of this package
		if err = key.Set(jwk.AlgorithmKey, jwa.RS256); err != nil {
			return nil, err
		}

		if err = keys.AddKey(key); err != nil {
			return nil, err
		}
	}

	return keys, nil
}

// RefreshAudience returns the refresh audience for the token manager; The refresh audience in plain-human-speak is the URL where the refresh token should be sent for validation (which is our datum endpoint)
func (tm *TokenManager) RefreshAudience() string {
	if tm.refreshAudience == "" {
		if tm.conf.RefreshAudience != "" {
			tm.refreshAudience = tm.conf.RefreshAudience
		}

		if aud, err := url.Parse(tm.issuer); err == nil {
			tm.refreshAudience = aud.ResolveReference(&url.URL{Path: "/v1/refresh"}).String()
		} else {
			tm.refreshAudience = DefaultRefreshAudience
		}
	}

	return tm.refreshAudience
}

// CurrentKey returns the ulid of the current key being used to sign tokens - this is just the identifier of the key, not the key itself
func (tm *TokenManager) CurrentKey() ulid.ULID {
	return tm.currentKeyID
}

// keyFunc selects the RSA public key from the list of tokenmanager internal keys based on the kid in the token header - if the kid does not exist an error is returned the token is not validated
func (tm *TokenManager) keyFunc(token *jwt.Token) (key interface{}, err error) {
	// Per JWT security notice: do not forget to validate alg is expected, else haxorz!~
	if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
		return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"]) //nolint:goerr113
	}

	// Fetch that kid
	kid, ok := token.Header["kid"]
	if !ok {
		return nil, ErrTokenMissingKid
	}

	// Parse that kid
	var keyID ulid.ULID

	if keyID, err = ulid.Parse(kid.(string)); err != nil {
		return nil, ErrFailedParsingKid
	}

	// Fetch the key from the list of managed keys
	if key, ok = tm.keys[keyID]; !ok {
		return nil, ErrUnknownSigningKey
	}

	return key, nil
}

// genKeyID generates a ulid for a key (the identifier of the key)
func (tm *TokenManager) genKeyID() (uid ulid.ULID, err error) {
	ms := ulid.Timestamp(time.Now())
	if uid, err = ulid.New(ms, tm.kidEntropy); err != nil {
		return uid, fmt.Errorf("could not generate key id: %w", err)
	}

	return uid, nil
}

// ParseUnverified parses a string of tokens and returns the claims and any error encountered
func ParseUnverified(tks string) (claims *jwt.RegisteredClaims, err error) {
	claims = &jwt.RegisteredClaims{}
	parser := jwt.NewParser(jwt.WithoutClaimsValidation())

	if _, _, err = parser.ParseUnverified(tks, claims); err != nil {
		return nil, err
	}

	return claims, nil
}

// ParseUnverifiedTokenClaims parses token claims from an access token
func ParseUnverifiedTokenClaims(tks string) (claims *Claims, err error) {
	claims = &Claims{}
	parser := jwt.NewParser(jwt.WithoutClaimsValidation())

	if _, _, err = parser.ParseUnverified(tks, claims); err != nil {
		return nil, err
	}

	return claims, nil
}

// ExpiresAt parses a JWT token and returns the expiration time if it exists
func ExpiresAt(tks string) (_ time.Time, err error) {
	var claims *jwt.RegisteredClaims

	if claims, err = ParseUnverified(tks); err != nil {
		return time.Time{}, err
	}

	return claims.ExpiresAt.Time, nil
}

// NotBefore parses a JWT token and returns the "NotBefore" time claim if it exists
func NotBefore(tks string) (_ time.Time, err error) {
	var claims *jwt.RegisteredClaims

	if claims, err = ParseUnverified(tks); err != nil {
		return time.Time{}, err
	}

	return claims.NotBefore.Time, nil
}
