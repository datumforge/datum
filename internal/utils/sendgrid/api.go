package sendgrid

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sendgrid/rest"
	"github.com/sendgrid/sendgrid-go"

	"github.com/datumforge/datum/internal/httpserve/middleware/auth"
)

// constants for working with sendgrid; didn't make them configurable as these are dictated by the vendor and this is a vendor specific package
const (
	Host       = "https://api.sendgrid.com"
	ContactsEP = "/v3/marketing/contacts"
	ListsEP    = "/v3/marketing/lists"
	FieldsEP   = "/v3/marketing/field_definitions"
)

// AddContactData is a struct holding the list(s) to add to as well as a pointer to Contact info to populate
type AddContactData struct {
	ListIDs  []string   `json:"list_ids"`
	Contacts []*Contact `json:"contacts"`
}

// AddContacts to SendGrid marketing list
func AddContacts(apiKey string, data *AddContactData) error {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(data); err != nil {
		return fmt.Errorf("could not encode json sendgrid contact data: %w", err)
	}

	req := sendgrid.GetRequest(apiKey, ContactsEP, Host)
	req.Method = http.MethodPut
	req.Body = buf.Bytes()

	if _, err := doRequest(req); err != nil {
		return err
	}

	return nil
}

// MarketingLists fetches lists of contacts from SendGrid
func MarketingLists(apiKey, pageToken string) (string, error) {
	params := map[string]string{
		"page_size": "100",
	}

	if pageToken != "" {
		params["page_token"] = pageToken
	}

	req := sendgrid.GetRequest(apiKey, ListsEP, Host)
	req.Method = http.MethodGet
	req.QueryParams = params

	return doRequest(req)
}

// FieldDefinitions gets defs from SendGrid - get back all the definitions to re-use them in requests
func FieldDefinitions(apiKey string) (string, error) {
	req := sendgrid.GetRequest(apiKey, FieldsEP, Host)
	req.Method = http.MethodGet

	return doRequest(req)
}

// doRequest is a helper to perform a SendGrid request, handling errors and returning the response body
func doRequest(req rest.Request) (_ string, err error) {
	rep, err := sendgrid.MakeRequest(req)
	if err != nil {
		return "", auth.ErrorResponse(err)
	}

	if rep.StatusCode < 200 || rep.StatusCode >= 300 {
		return "", auth.ErrorResponse(rep.StatusCode)
	}

	return rep.Body, nil
}
