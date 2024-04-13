package workrequests

import (
	"time"
)

type Config struct {
	WorkRequest        WorkRequest        `json:"workRequest"`
	WorkRequestSummary WorkRequestSummary `json:"workRequestSummary"`
}

// WorkRequest A description of workrequest
type WorkRequest struct {

	// Type of the work request
	OperationType string `json:"operationType"`

	// The current status of the work request.
	Status string `json:"status"`

	// The OCID of the work request.
	Id string `json:"id"`

	// The OCID of the compartment that contains the work request.
	CompartmentId string `json:"compartmentId"`

	// How much progress the operation has made.
	PercentComplete float32 `json:"percentComplete"`

	// Date and time the work was accepted, in the format defined by RFC3339. Example: `2016-08-25T21:10:29.600Z`
	TimeAccepted time.Time `json:"timeAccepted"`

	// The resources affected by this work request.
	Resources []string `json:"resources"`

	// Date and time the work started, in the format defined by RFC3339. Example: `2016-08-25T21:10:29.600Z`
	TimeStarted time.Time `json:"timeStarted"`

	// Date and time the work completed, in the format defined by RFC3339. Example: `2016-08-25T21:10:29.600Z`
	TimeFinished time.Time `json:"timeFinished"`

	// The listing id associated with the work request.
	ListingId string `json:"listingId"`

	// The package version associated with the work request.
	PackageVersion string `json:"packageVersion"`
}

// WorkRequestSummary A summary of the status of a work request.
type WorkRequestSummary struct {

	// Type of the work request
	OperationType string `json:"operationType"`

	// Status of current work request.
	Status string `json:"status"`

	// The id of the work request.
	Id string `json:"id"`

	// The ocid of the compartment that contains the work request. Work requests should be scoped to
	// the same compartment as the resource the work request affects.
	CompartmentId string `json:"compartmentId"`

	// Percentage of the request completed.
	PercentComplete float32 `json:"percentComplete"`

	// The date and time the request was created, as described in
	// RFC 3339 (https://tools.ietf.org/rfc/rfc3339), section 14.29.
	TimeAccepted time.Time `json:"timeAccepted"`

	// The resources affected by this work request.
	Resources []string `json:"resources"`

	// The date and time the request was started, as described in RFC 3339 (https://tools.ietf.org/rfc/rfc3339),
	// section 14.29.
	TimeStarted time.Time `json:"timeStarted"`

	// The date and time the object was finished, as described in RFC 3339 (https://tools.ietf.org/rfc/rfc3339).
	TimeFinished time.Time `json:"timeFinished"`

	// The listing id associated with the work request.
	ListingId string `json:"listingId"`

	// The package version associated with the work request.
	PackageVersion string `json:"packageVersion"`
}
