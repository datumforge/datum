package tokens

import (
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

type Parser struct {
	// If populated, only these methods will be considered valid.
	//
	// Deprecated: In future releases, this field will not be exported anymore and should be set with an option to NewParser instead.
	ValidMethods []string

	// Use JSON Number format in JSON decoder.
	//
	// Deprecated: In future releases, this field will not be exported anymore and should be set with an option to NewParser instead.
	UseJSONNumber bool

	// Skip claims validation during token parsing.
	//
	// Deprecated: In future releases, this field will not be exported anymore and should be set with an option to NewParser instead.
	SkipClaimsValidation bool
}

// NewParser creates a new Parser with the specified options
func NewParser(options ...ParserOption) *Parser {
	p := &Parser{}

	// loop through our parsing options and apply them
	for _, option := range options {
		option(p)
	}

	return p
}

// ParserOption is used to implement functional-style options that modify the behavior of the parser. To add
// new options, just create a function (ideally beginning with With or Without) that returns an anonymous function that
// takes a *Parser type as input and manipulates its configuration accordingly.
type ParserOption func(*Parser)

// WithValidMethods is an option to supply algorithm methods that the parser will check. Only those methods will be considered valid.
// It is heavily encouraged to use this option in order to prevent attacks such as https://auth0.com/blog/critical-vulnerabilities-in-json-web-token-libraries/.
func WithValidMethods(methods []string) ParserOption {
	return func(p *Parser) {
		p.ValidMethods = methods
	}
}

// WithJSONNumber is an option to configure the underlying JSON parser with UseNumber
func WithJSONNumber() ParserOption {
	return func(p *Parser) {
		p.UseJSONNumber = true
	}
}

// WithoutClaimsValidation is an option to disable claims validation. This option should only be used if you exactly know
// what you are doing.
func WithoutClaimsValidation() ParserOption {
	return func(p *Parser) {
		p.SkipClaimsValidation = true
	}
}

func newParserWithoutClaimsValidation() *jwt.Parser {
	return jwt.NewParser(jwt.WithoutClaimsValidation())
}

func ParseUnverified(tks string) (claims *jwt.RegisteredClaims, err error) {
	claims = &jwt.RegisteredClaims{}
	tsparser := newParserWithoutClaimsValidation()

	if _, _, err = tsparser.ParseUnverified(tks, claims); err != nil {
		return nil, err
	}

	return claims, nil
}

func ExpiresAt(tks string) (_ time.Time, err error) {
	var claims *jwt.RegisteredClaims

	if claims, err = ParseUnverified(tks); err != nil {
		return time.Time{}, err
	}

	return claims.ExpiresAt.Time, nil
}

func NotBefore(tks string) (_ time.Time, err error) {
	var claims *jwt.RegisteredClaims

	if claims, err = ParseUnverified(tks); err != nil {
		return time.Time{}, err
	}

	return claims.NotBefore.Time, nil
}
