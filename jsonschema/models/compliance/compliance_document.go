package compliance

import "time"

type Config struct {
	ComplianceDocument        ComplianceDocument        `json:"complianceDocument"`
	ComplianceDocumentSummary ComplianceDocumentSummary `json:"complianceDocumentSummary"`
}

// ComplianceDocument A compliance document that exists in the tenancy.
type ComplianceDocument struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compliance document, which is assigned
	// when you create the document as an Oracle Cloud Infrastructure resource and is immutable.
	Id string `json:"id"`

	// A friendly name or title for the compliance document. You cannot update this value later.
	// Avoid entering confidential information.
	Name string `json:"name"`

	// The date and time the compliance document was created, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339) timestamp format.
	TimeCreated time.Time `json:"timeCreated"`

	// The current lifecycle state of the compliance document.
	LifecycleState ComplianceDocumentLifecycleStateEnum `json:"lifecycleState"`

	// The file name of the compliance document.
	DocumentFileName string `json:"documentFileName"`

	// The version number of the compliance document.
	Version int `json:"version"`

	// The type of compliance document. For definitions of supported types of compliance documents, see Types of Compliance Documents
	Type ComplianceDocumentTypeEnum `json:"type"`

	// The information technology infrastructure platform, or set of services, to which the compliance document belongs
	Platform ComplianceDocumentPlatformEnum `json:"platform"`

	// The date and time the compliance document was last updated, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339) timestamp format.
	TimeUpdated time.Time `json:"timeUpdated"`
}

// ComplianceDocumentLifecycleStateEnum Enum with underlying type: string
type ComplianceDocumentLifecycleStateEnum string

// Set of constants representing the allowable values for ComplianceDocumentLifecycleStateEnum
const (
	ComplianceDocumentLifecycleStateActive   ComplianceDocumentLifecycleStateEnum = "ACTIVE"
	ComplianceDocumentLifecycleStateInactive ComplianceDocumentLifecycleStateEnum = "INACTIVE"
)

var mappingComplianceDocumentLifecycleState = map[string]ComplianceDocumentLifecycleStateEnum{
	"ACTIVE":   ComplianceDocumentLifecycleStateActive,
	"INACTIVE": ComplianceDocumentLifecycleStateInactive,
}

// ComplianceDocumentTypeEnum Enum with underlying type: string
type ComplianceDocumentTypeEnum string

// Set of constants representing the allowable values for ComplianceDocumentTypeEnum
const (
	ComplianceDocumentTypeSod          ComplianceDocumentTypeEnum = "SOD"
	ComplianceDocumentTypeAttestation  ComplianceDocumentTypeEnum = "ATTESTATION"
	ComplianceDocumentTypeBridgeletter ComplianceDocumentTypeEnum = "BRIDGELETTER"
	ComplianceDocumentTypePentest      ComplianceDocumentTypeEnum = "PENTEST"
	ComplianceDocumentTypeAudit        ComplianceDocumentTypeEnum = "AUDIT"
	ComplianceDocumentTypeCertificate  ComplianceDocumentTypeEnum = "CERTIFICATE"
	ComplianceDocumentTypeSoc3         ComplianceDocumentTypeEnum = "SOC3"
	ComplianceDocumentTypeOther        ComplianceDocumentTypeEnum = "OTHER"
)

var mappingComplianceDocumentType = map[string]ComplianceDocumentTypeEnum{
	"SOD":          ComplianceDocumentTypeSod,
	"ATTESTATION":  ComplianceDocumentTypeAttestation,
	"BRIDGELETTER": ComplianceDocumentTypeBridgeletter,
	"PENTEST":      ComplianceDocumentTypePentest,
	"AUDIT":        ComplianceDocumentTypeAudit,
	"CERTIFICATE":  ComplianceDocumentTypeCertificate,
	"SOC3":         ComplianceDocumentTypeSoc3,
	"OTHER":        ComplianceDocumentTypeOther,
}

// ComplianceDocumentPlatformEnum Enum with underlying type: string
type ComplianceDocumentPlatformEnum string

// Set of constants representing the allowable values for ComplianceDocumentPlatformEnum
const (
	ComplianceDocumentPlatformOciedgeservices ComplianceDocumentPlatformEnum = "OCIEDGESERVICES"
	ComplianceDocumentPlatformOci             ComplianceDocumentPlatformEnum = "OCI"
	ComplianceDocumentPlatformPaas            ComplianceDocumentPlatformEnum = "PAAS"
	ComplianceDocumentPlatformCloudconsole    ComplianceDocumentPlatformEnum = "CLOUDCONSOLE"
	ComplianceDocumentPlatformOmcs            ComplianceDocumentPlatformEnum = "OMCS"
	ComplianceDocumentPlatformOciCIaas        ComplianceDocumentPlatformEnum = "OCI_C_IAAS"
	ComplianceDocumentPlatformOther           ComplianceDocumentPlatformEnum = "OTHER"
)

var mappingComplianceDocumentPlatform = map[string]ComplianceDocumentPlatformEnum{
	"OCIEDGESERVICES": ComplianceDocumentPlatformOciedgeservices,
	"OCI":             ComplianceDocumentPlatformOci,
	"PAAS":            ComplianceDocumentPlatformPaas,
	"CLOUDCONSOLE":    ComplianceDocumentPlatformCloudconsole,
	"OMCS":            ComplianceDocumentPlatformOmcs,
	"OCI_C_IAAS":      ComplianceDocumentPlatformOciCIaas,
	"OTHER":           ComplianceDocumentPlatformOther,
}

// ComplianceDocumentSummary A summary representation of the compliance document.
type ComplianceDocumentSummary struct {

	// A unique identifier for the document that is assigned when you create
	// the document as an Oracle Cloud Infrastructure resource and is immutable.
	Id string `json:"id"`

	// A friendly name or title for the compliance document. You cannot update this value later.
	// Avoid entering confidential information.
	Name string `json:"name"`

	// The current lifecycle state of the compliance document.
	LifecycleState ComplianceDocumentLifecycleStateEnum `json:"lifecycleState"`

	// The file name of the compliance document.
	DocumentFileName string `json:"documentFileName"`

	// The version number of the compliance document.
	Version int `json:"version"`

	// The type of compliance document.
	Type ComplianceDocumentTypeEnum `json:"type"`

	// The environment, also known as platform or business pillar, to which the compliance document belongs.
	Platform ComplianceDocumentPlatformEnum `json:"platform"`

	// The date and time the compliance document was last updated, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339) timestamp format.
	TimeUpdated time.Time `json:"timeUpdated"`

	// The date and time the compliance document was created, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339) timestamp format.
	TimeCreated time.Time `json:"timeCreated"`
}
