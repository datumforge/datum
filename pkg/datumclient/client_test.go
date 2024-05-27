package datumclient_test

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/require"

	api "github.com/datumforge/datum/pkg/datumclient"
	"github.com/datumforge/datum/pkg/models"
	"github.com/datumforge/datum/pkg/rout"
	"github.com/datumforge/datum/pkg/utils/ulids"
)

func TestClient(t *testing.T) {
	// Create a Test Server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			require.Equal(t, int64(0), r.ContentLength)
			w.Header().Add("Content-Type", "application/json; charset=utf-8")
			w.WriteHeader(http.StatusOK)
			fmt.Fprintln(w, "{\"hello\":\"world\"}")

			return
		}

		require.Equal(t, int64(18), r.ContentLength)
		w.Header().Add("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "{\"error\":\"bad request\"}")
	}))

	defer ts.Close()

	// Create a Client that makes requests to the test server
	client, err := api.New(ts.URL)
	require.NoError(t, err)

	// Ensure that the latest version of the client is returned
	apiv1, ok := client.(*api.APIv1)
	require.True(t, ok)

	// Create a new GET request to a basic path
	req, err := apiv1.NewRequest(context.TODO(), http.MethodGet, "/foo", nil, nil)
	require.NoError(t, err)

	require.Equal(t, "/foo", req.URL.Path)
	require.Equal(t, "", req.URL.RawQuery)
	require.Equal(t, http.MethodGet, req.Method)
	require.Equal(t, "Datum API Client/v1", req.Header.Get("User-Agent"))
	require.Equal(t, "application/json", req.Header.Get("Accept"))
	require.Equal(t, "application/json; charset=utf-8", req.Header.Get("Content-Type"))

	// Create a new GET request with query params
	params := url.Values{}
	params.Add("q", "searching")
	params.Add("key", "open says me")
	req, err = apiv1.NewRequest(context.TODO(), http.MethodGet, "/foo", nil, &params)
	require.NoError(t, err)
	require.Equal(t, "key=open+says+me&q=searching", req.URL.RawQuery)

	data := make(map[string]string)
	rep, err := apiv1.Do(req, &data, true)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, rep.StatusCode)
	require.Contains(t, data, "hello")
	require.Equal(t, "world", data["hello"])

	// Create a new POST request and check error handling
	req, err = apiv1.NewRequest(context.TODO(), http.MethodPost, "/bar", data, nil)
	require.NoError(t, err)
	rep, err = apiv1.Do(req, nil, false)
	require.NoError(t, err)
	require.Equal(t, http.StatusBadRequest, rep.StatusCode)

	req, err = apiv1.NewRequest(context.TODO(), http.MethodPost, "/bar", data, nil)
	require.NoError(t, err)
	_, err = apiv1.Do(req, nil, true)
	require.EqualError(t, err, "[400] bad request")

	// Test supplying an authentication override token in the request context
	ctx := api.ContextWithToken(context.Background(), "newtoken")
	req, err = apiv1.NewRequest(ctx, http.MethodPost, "/bar", data, nil)
	require.NoError(t, err, "could not create request")
	require.Equal(t, "Bearer newtoken", req.Header.Get("Authorization"), "expected the authorization header to be set")

	// Test that default credentials are used if no credentials are supplied in the request context
	defaultCreds := api.Token("default")
	client, err = api.New(ts.URL, api.WithCredentials(defaultCreds))
	require.NoError(t, err, "could not create client")
	apiv1, ok = client.(*api.APIv1)
	require.True(t, ok, "could not cast client to APIv1")
	req, err = apiv1.NewRequest(context.Background(), http.MethodPost, "/bar", data, nil)
	require.NoError(t, err, "could not create request")
	require.Equal(t, "Bearer default", req.Header.Get("Authorization"), "expected the authorization header to be set to default")

	// Test that request credentials override default credentials
	ctx = api.ContextWithToken(context.Background(), "newtoken")
	req, err = apiv1.NewRequest(ctx, http.MethodPost, "/bar", data, nil)
	require.NoError(t, err, "could not create request")
	require.Equal(t, "Bearer newtoken", req.Header.Get("Authorization"), "expected the authorization header to be set to newtoken")
}

func TestRegister(t *testing.T) {
	// Setup the response fixture
	ulid := ulids.New()
	fixture := &models.RegisterReply{
		ID:      ulid.String(),
		Email:   "jb@example.com",
		Message: "Thank you for registering for datum!",
	}

	// Create a test server
	ts := httptest.NewServer(testhandler(fixture, http.MethodPost, "/v1/register"))
	defer ts.Close()

	// Create a client and execute endpoint request
	client, err := api.New(ts.URL)
	require.NoError(t, err, "could not create api client")

	req := &models.RegisterRequest{
		Email:    "jb@example.com",
		Password: "supers3cr4etsquir!!",
	}

	rep, err := client.Register(context.TODO(), req)
	require.NoError(t, err, "could not execute api request")
	require.Equal(t, fixture, rep, "unexpected response returned")
}

func TestLogin(t *testing.T) {
	// Setup the response fixture
	fixture := &models.LoginReply{}

	// Create a test server
	ts := httptest.NewServer(testhandler(fixture, http.MethodPost, "/v1/login"))
	defer ts.Close()

	// Create a client and execute endpoint request
	client, err := api.New(ts.URL)
	require.NoError(t, err, "could not create api client")

	req := &models.LoginRequest{}

	rep, err := client.Login(context.TODO(), req)
	require.NoError(t, err, "could not execute api request")
	require.Equal(t, fixture, rep, "unexpected response returned")
}

func TestRefresh(t *testing.T) {
	// Setup the response fixture
	fixture := &models.RefreshReply{
		Message: "refreshed",
	}

	// Create a test server
	ts := httptest.NewServer(testhandler(fixture, http.MethodPost, "/v1/refresh"))
	defer ts.Close()

	// Create a client and execute endpoint request
	client, err := api.New(ts.URL)
	require.NoError(t, err, "could not create api client")

	req := &models.RefreshRequest{}
	rep, err := client.Refresh(context.TODO(), req)
	require.NoError(t, err, "could not execute api request")
	require.Equal(t, fixture, rep, "unexpected response returned")
}

func TestSwitch(t *testing.T) {
	// Setup the response fixture
	fixture := &models.SwitchOrganizationReply{}

	// Create a test server
	ts := httptest.NewServer(testhandler(fixture, http.MethodPost, "/v1/switch"))
	defer ts.Close()

	// Create a client and execute endpoint request
	client, err := api.New(ts.URL)
	require.NoError(t, err, "could not create api client")

	req := &models.SwitchOrganizationRequest{}
	rep, err := client.Switch(context.TODO(), req)
	require.NoError(t, err, "could not execute api request")
	require.Equal(t, fixture, rep, "unexpected response returned")
}

func TestVerifyEmail(t *testing.T) {
	// Create a test server with a simple response fixture.
	fixture := &models.VerifyReply{
		AccessToken:  "access",
		RefreshToken: "refresh",
	}
	ts := httptest.NewServer(testhandler(fixture, http.MethodPost, "/v1/verify"))
	defer ts.Close()

	// Create a client and execute endpoint request
	client, err := api.New(ts.URL)
	require.NoError(t, err, "could not create api client")

	req := &models.VerifyRequest{Token: "1234567890"}
	rep, err := client.VerifyEmail(context.TODO(), req)
	require.NoError(t, err, "could not execute api request")
	require.Equal(t, fixture, rep, "unexpected response returned")
}

func TestResendEmail(t *testing.T) {
	// Create a test server with a simple response fixture.
	fixture := &models.ResendReply{
		Message: "email sent",
	}
	ts := httptest.NewServer(testhandler(fixture, http.MethodPost, "/v1/resend"))
	defer ts.Close()

	// Create a client and execute endpoint request
	client, err := api.New(ts.URL)
	require.NoError(t, err, "could not create api client")

	req := &models.ResendRequest{Email: "frank@example.com"}
	rep, err := client.ResendEmail(context.Background(), req)
	require.NoError(t, err, "could not execute api request")
	require.Equal(t, fixture, rep, "unexpected response returned")
}

func TestForgotPassword(t *testing.T) {
	// Create a test server with a simple response fixture.
	fixture := &models.ForgotPasswordReply{
		Message: "password reset email sent",
	}

	ts := httptest.NewServer(testhandler(fixture, http.MethodPost, "/v1/forgot-password"))
	defer ts.Close()

	// Create a client and execute endpoint request
	client, err := api.New(ts.URL)
	require.NoError(t, err, "could not create api client")
	req := &models.ForgotPasswordRequest{Email: "leopold.wentzel@gmail.com"}
	rep, err := client.ForgotPassword(context.Background(), req)
	require.NoError(t, err, "could not execute api request")
	require.Equal(t, fixture, rep, "unexpected response returned")
}

func TestResetPassword(t *testing.T) {
	fixture := &models.ResetPasswordReply{
		Message: "password reset",
	}
	ts := httptest.NewServer(testhandler(fixture, http.MethodPost, "/v1/password-reset"))

	defer ts.Close()

	// Create a client and execute endpoint request
	client, err := api.New(ts.URL)
	require.NoError(t, err, "could not create api client")
	req := &models.ResetPasswordRequest{
		Token:    "token",
		Password: "password",
	}
	rep, err := client.ResetPassword(context.Background(), req)
	require.NoError(t, err, "could not execute api request")
	require.Equal(t, fixture, rep, "unexpected response returned")
}

func TestInviteCreate(t *testing.T) {
	// Setup the response fixture
	fixture := &models.InviteReply{
		ID:    ulids.New().String(),
		Email: "leopold.wentzel@gmail.com",
		Role:  "admin",
	}

	// Create a test server
	ts := httptest.NewServer(testhandler(fixture, http.MethodPost, "/v1/invite"))
	defer ts.Close()

	// Create a client and execute endpoint request
	client, err := api.New(ts.URL)
	require.NoError(t, err, "could not create api client")

	req := &models.InviteRequest{
		Token: "foo",
	}
	reply, err := client.Invite(context.TODO(), req)
	require.NoError(t, err, "could not execute api request")
	require.Equal(t, fixture, reply, "unexpected response returned")
}

func TestInviteAccept(t *testing.T) {
	// Setup the response fixture
	fixture := &models.InviteReply{
		ID:          ulids.New().String(),
		Email:       "mitb@datum.net",
		JoinedOrgID: ulids.New().String(),
		Role:        "admin",
		Message:     "Welcome to your new organization!",
	}

	// Create a test server
	ts := httptest.NewServer(testhandler(fixture, http.MethodPost, "/v1/invite"))
	defer ts.Close()

	// Create a client and execute endpoint request
	client, err := api.New(ts.URL)
	require.NoError(t, err)

	req := &models.InviteRequest{
		Token: "foo",
	}
	reply, err := client.Invite(context.Background(), req)
	require.NoError(t, err, "could not execute api request")
	require.Equal(t, fixture, reply, "unexpected response returned")
}

func testhandler(fixture interface{}, expectedMethod, expectedPath string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json; charset=utf-8")

		if r.Method != expectedMethod {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(rout.ErrorResponse("unexpected http method"))
			return
		}

		if r.URL.Path != expectedPath {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(rout.ErrorResponse("unexpected endpoint path"))
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(fixture)
	})
}
