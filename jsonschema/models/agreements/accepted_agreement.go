package agreements

import "time"

type Config struct {
	AcceptedAgreement        AcceptedAgreement        `json:"acceptedAgreement"`
	Agreement                Agreement                `json:"agreement"`
	AgreementSummary         AgreementSummary         `json:"agreementSummary"`
	AcceptedAgreementSummary AcceptedAgreementSummary `json:"acceptedAgreementSummary"`
}

// AcceptedAgreement The model for an accepted terms of use agreement.
type AcceptedAgreement struct {

	// The unique identifier for the acceptance of the agreement within a specific compartment.
	Id string `json:"id"`

	// A display name for the accepted agreement.
	DisplayName string `json:"displayName"`

	// The unique identifier for the compartment where the agreement was accepted.
	CompartmentId string `json:"compartmentId"`

	// The unique identifier for the listing associated with the agreement.
	ListingId string `json:"listingId"`

	// The package version associated with the agreement.
	PackageVersion string `json:"packageVersion"`

	// The unique identifier for the terms of use agreement itself.
	AgreementId string `json:"agreementId"`

	// The time the agreement was accepted.
	TimeAccepted time.Time `json:"timeAccepted"`

	DefinedTags map[string]map[string]interface{} `json:"definedTags"`

	FreeformTags map[string]string `json:"freeformTags"`
}

// AgreementSummary The model for a summary of an end user license agreement.
type AgreementSummary struct {

	// The unique identifier for the agreement.
	Id string `json:"id"`

	// The content URL of the agreement.
	ContentUrl string `json:"contentUrl"`

	// Who authored the agreement.
	Author string `json:"author"`

	// Textual prompt to read and accept the agreement.
	Prompt string `json:"prompt"`
}

// AcceptedAgreementSummary The model for a summary of an accepted agreement.
type AcceptedAgreementSummary struct {

	// The unique identifier for the acceptance of the agreement within a specific compartment.
	Id string `json:"id"`

	// A display name for the accepted agreement.
	DisplayName string `json:"displayName"`

	// The unique identifier for the compartment where the agreement was accepted.
	CompartmentId string `json:"compartmentId"`

	// The unique identifier for the listing associated with the agreement.
	ListingId string `json:"listingId"`

	// The package version associated with the agreement.
	PackageVersion string `json:"packageVersion"`

	// The unique identifier for the terms of use agreement itself.
	AgreementId string `json:"agreementId"`

	// The time the agreement was accepted.
	TimeAccepted time.Time `json:"timeAccepted"`
}

// Agreement The model for an end user license agreement.
type Agreement struct {

	// The unique identifier for the agreement.
	Id string `json:"id"`

	// The content URL of the agreement.
	ContentUrl string `json:"contentUrl"`

	// A time-based signature that can be used to accept an agreement or remove a
	// previously accepted agreement from the list that Marketplace checks before a deployment.
	Signature string `json:"signature"`

	// The unique identifier for the compartment.
	CompartmentId string `json:"compartmentId"`

	// Who authored the agreement.
	Author AgreementAuthorEnum `json:"author"`

	// Textual prompt to read and accept the agreement.
	Prompt string `json:"prompt"`
}

// AgreementAuthorEnum Enum with underlying type: string
type AgreementAuthorEnum string

// Set of constants representing the allowable values for AgreementAuthorEnum
const (
	AgreementAuthorDatum   AgreementAuthorEnum = "DATUM"
	AgreementAuthorPartner AgreementAuthorEnum = "PARTNER"
)

var mappingAgreementAuthorEnum = map[string]AgreementAuthorEnum{
	"DATUM":   AgreementAuthorDatum,
	"PARTNER": AgreementAuthorPartner,
}

var mappingAgreementAuthorEnumLowerCase = map[string]AgreementAuthorEnum{
	"datum":   AgreementAuthorDatum,
	"partner": AgreementAuthorPartner,
}
