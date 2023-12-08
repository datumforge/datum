package authtest

import (
	"crypto/rand"
	"crypto/rsa"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/datumforge/datum/internal/tokens"
)

const (
	Audience = "http://127.0.0.1"
	Issuer   = "http://127.0.0.1"
)

// Server implements an endpoint to host JWKS public keys and also provides simple
// functionality to create access and refresh tokens that would be authenticated.
type Server struct {
	srv    *httptest.Server
	mux    *http.ServeMux
	tokens *tokens.TokenManager
	URL    *url.URL
}

// NewServer starts and returns a new authtest server. The caller should call Close
// when finished, to shut it down.
func NewServer() (s *Server, err error) {
	// Setup routes for the mux
	s = &Server{}
	s.mux = http.NewServeMux()
	s.mux.HandleFunc("/.well-known/jwks.json", s.JWKS)

	// Setup httptest Server
	s.srv = httptest.NewServer(s.mux)
	s.URL, _ = url.Parse(s.srv.URL)

	// Create token manager
	var key *rsa.PrivateKey

	if key, err = rsa.GenerateKey(rand.Reader, 2048); err != nil { // nolint: gomnd
		return nil, err
	}

	conf := tokens.TokenConfig{
		Audience:        Audience,
		Issuer:          Issuer,
		AccessDuration:  1 * time.Hour,
		RefreshDuration: 2 * time.Hour, // nolint: gomnd
		RefreshOverlap:  -15 * time.Minute,
	}

	if s.tokens, err = tokens.NewWithKey(key, conf); err != nil {
		return nil, err
	}

	return s, nil
}

func (s *Server) Close() {
	s.srv.Close()
}

func (s *Server) JWKS(w http.ResponseWriter, r *http.Request) {
	keys, err := s.tokens.Keys()
	if err != nil {
		w.Header().Add("Content-Type", "text/plain")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error())) // nolint: errcheck

		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(keys) // nolint: errcheck
}

func (s *Server) KeysURL() string {
	return s.URL.ResolveReference(&url.URL{Path: "/.well-known/jwks.json"}).String()
}

// CreateToken creates a token without overwriting the claims, which is useful for
// creating tokens with specific not before and expiration times for testing.
func (s *Server) CreateToken(claims *tokens.Claims) (tks string, err error) {
	return s.tokens.Sign(s.tokens.CreateToken(claims))
}

func (s *Server) CreateAccessToken(claims *tokens.Claims) (tks string, err error) {
	var token *jwt.Token

	if token, err = s.tokens.CreateAccessToken(claims); err != nil {
		return "", err
	}

	return s.tokens.Sign(token)
}

func (s *Server) CreateTokenPair(claims *tokens.Claims) (accessToken, refreshToken string, err error) {
	return s.tokens.CreateTokenPair(claims)
}
