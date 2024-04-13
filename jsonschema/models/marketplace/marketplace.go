package marketplace

import (
	"time"
)

type Config struct {
	Publisher                Publisher                `json:"publisher"`
	PublisherSummary         PublisherSummary         `json:"publisherSummary"`
	Publication              Publication              `json:"publication"`
	Market                   Market                   `json:"market"`
	Listing                  Listing                  `json:"listing"`
	ListingRevision          ListingRevision          `json:"listingRevision"`
	ListingSummary           ListingSummary           `json:"listingSummary"`
	Commitment               Commitment               `json:"commitment"`
	ComputedUsage            ComputedUsage            `json:"computedUsage"`
	InvoicingUser            InvoicingUser            `json:"invoicingUser"`
	InvoicingProduct         InvoicingProduct         `json:"invoicingProduct"`
	InvoicingPaymentTerm     InvoicingPaymentTerm     `json:"invoicingPaymentTerm"`
	InvoicingOrganization    InvoicingOrganization    `json:"invoicingOrganization"`
	InvoicingLocation        InvoicingLocation        `json:"invoicingLocation"`
	InvoicingAddress         InvoicingAddress         `json:"invoicingAddress"`
	InvoicingBusinessPartner InvoicingBusinessPartner `json:"invoicingBusinessPartner"`
	ComputedUsageProduct     ComputedUsageProduct     `json:"computedUsageProduct"`
}

// Publisher The model for a publisher details.
type Publisher struct {

	// Unique OCID identifier for the publisher.
	Id string `json:"id"`

	// The root compartment of the Publisher.
	CompartmentId string `json:"compartmentId"`

	// The namespace for the publisher registry to persist artifacts.
	RegistryNamespace string `json:"registryNamespace"`

	// The name of the publisher.
	DisplayName string `json:"displayName"`

	// The public email address of the publisher for customers.
	ContactEmail string `json:"contactEmail"`

	// The phone number of the publisher in E.164 format.
	ContactPhone string `json:"contactPhone"`

	// publisher type.
	PublisherType string `json:"publisherType"`

	// The time the publisher was created. An RFC3339 formatted datetime string.
	TimeCreated time.Time `json:"timeCreated"`

	// The time the publisher was updated. An RFC3339 formatted datetime string.
	TimeUpdated time.Time `json:"timeUpdated"`

	// publisher status.
	PublisherStatus string `json:"publisherStatus"`

	// Unique legacy service identifier for the publisher.
	LegacyId string `json:"legacyId"`

	// A description of the publisher.
	Description string `json:"description"`

	// The year the publisher's company or organization was founded.
	YearFounded int64 `json:"yearFounded"`

	// The publisher's website.
	WebsiteUrl string `json:"websiteUrl"`

	// The address of the publisher's headquarters.
	HqAddress string `json:"hqAddress"`

	Logo string `json:"logo"`

	// Publisher's Facebook URL
	FacebookUrl string `json:"facebookUrl"`

	// Publisher's Twitter URL
	TwitterUrl string `json:"twitterUrl"`

	// Publisher's LinkedIn URL
	LinkedinUrl string `json:"linkedinUrl"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `json:"definedTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `json:"systemTags"`

	// The private email address of the publisher product team.
	NotificationEmail string `json:"notificationEmail"`
}

// PublisherSummary Summary details about the publisher of the listing.
type PublisherSummary struct {

	// The unique identifier for the publisher.
	Id string `json:"id"`

	// The name of the publisher.
	Name string `json:"name"`

	// A description of the publisher.
	Description string `json:"description"`
}

// Publication The model for an Oracle Cloud Infrastructure Marketplace publication.
type Publication struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment where the publication exists.
	CompartmentId string `json:"compartmentId"`

	// The unique identifier for the publication in Marketplace.
	Id string `json:"id"`

	// The name of the publication, which is also used in the listing.
	Name string `json:"name"`

	// The publisher category to which the publication belongs. The publisher category informs where the listing appears for use.
	ListingType ListingTypeEnum `json:"listingType"`

	// The lifecycle state of the publication.
	LifecycleState PublicationLifecycleStateEnum `json:"lifecycleState,omitempty"`

	// A short description of the publication to use in the listing.
	ShortDescription string `json:"shortDescription"`

	// A long description of the publication to use in the listing.
	LongDescription string `json:"longDescription"`

	// Contact information for getting support from the publisher for the listing.
	SupportContacts SupportContact `json:"supportContacts"`

	Icon string `json:"icon"`

	// The listing's package type.

	// The list of operating systems supported by the listing.
	SupportedOperatingSystems []string `json:"supportedOperatingSystems"`

	// The date and time the publication was created, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339)
	// timestamp format.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated time.Time `json:"timeCreated"`

	// The defined tags associated with this resource, if any. Each key is predefined and scoped to namespaces.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `json:"definedTags"`

	// The freeform tags associated with this resource, if any. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `json:"freeformTags"`

	// The system tags associated with this resource, if any. The system tags are set by Oracle Cloud Infrastructure services. Each key is predefined and scoped to namespaces.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{orcl-cloud: {free-tier-retain: true}}`
	SystemTags map[string]map[string]interface{} `json:"systemTags"`
}

// PublicationLifecycleStateEnum Enum with underlying type: string
type PublicationLifecycleStateEnum string

// Set of constants representing the allowable values for PublicationLifecycleStateEnum
const (
	PublicationLifecycleStateCreating PublicationLifecycleStateEnum = "CREATING"
	PublicationLifecycleStateActive   PublicationLifecycleStateEnum = "ACTIVE"
	PublicationLifecycleStateDeleting PublicationLifecycleStateEnum = "DELETING"
	PublicationLifecycleStateDeleted  PublicationLifecycleStateEnum = "DELETED"
	PublicationLifecycleStateFailed   PublicationLifecycleStateEnum = "FAILED"
)

var mappingPublicationLifecycleStateEnum = map[string]PublicationLifecycleStateEnum{
	"CREATING": PublicationLifecycleStateCreating,
	"ACTIVE":   PublicationLifecycleStateActive,
	"DELETING": PublicationLifecycleStateDeleting,
	"DELETED":  PublicationLifecycleStateDeleted,
	"FAILED":   PublicationLifecycleStateFailed,
}

var mappingPublicationLifecycleStateEnumLowerCase = map[string]PublicationLifecycleStateEnum{
	"creating": PublicationLifecycleStateCreating,
	"active":   PublicationLifecycleStateActive,
	"deleting": PublicationLifecycleStateDeleting,
	"deleted":  PublicationLifecycleStateDeleted,
	"failed":   PublicationLifecycleStateFailed,
}

// ListingTypeEnum Enum with underlying type: string
type ListingTypeEnum string

// Set of constants representing the allowable values for ListingTypeEnum
const (
	ListingTypeCommunity ListingTypeEnum = "COMMUNITY"
	ListingTypePartner   ListingTypeEnum = "PARTNER"
	ListingTypePrivate   ListingTypeEnum = "PRIVATE"
)

var mappingListingTypeEnum = map[string]ListingTypeEnum{
	"COMMUNITY": ListingTypeCommunity,
	"PARTNER":   ListingTypePartner,
	"PRIVATE":   ListingTypePrivate,
}

var mappingListingTypeEnumLowerCase = map[string]ListingTypeEnum{
	"community": ListingTypeCommunity,
	"partner":   ListingTypePartner,
	"private":   ListingTypePrivate,
}

// Market The model for the market details.
type Market struct {

	// The name of the market.
	Name string `json:"name"`

	// The code of the market.
	Code string `json:"code"`

	// The category code of the market.
	CategoryCode string `json:"categoryCode"`

	// bill to countries for the market.
	BillToCountries []string `json:"billToCountries"`

	// The current state for the market.
	LifecycleState string `json:"lifecycleState"`

	// The date and time the market was created, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2022-09-15T21:10:29.600Z`
	TimeCreated time.Time `json:"timeCreated"`

	// The date and time the market was updated, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2022-09-15T21:10:29.600Z`
	TimeUpdated time.Time `json:"timeUpdated"`

	// The realm code of the market.
	RealmCode string `json:"realmCode"`
}

// Listing The model for an Oracle Cloud Infrastructure Marketplace listing.
type Listing struct {

	// The unique identifier for the listing in Marketplace.
	Id string `json:"id"`

	// The name of the listing.
	Name string `json:"name"`

	// The version of the listing.
	Version string `json:"version"`

	// The tagline of the listing.
	Tagline string `json:"tagline"`

	// Keywords associated with the listing.
	Keywords string `json:"keywords"`

	// A short description of the listing.
	ShortDescription string `json:"shortDescription"`

	// Usage information for the listing.
	UsageInformation string `json:"usageInformation"`

	// A long description of the listing.
	LongDescription string `json:"longDescription"`

	// A description of the publisher's licensing model for the listing.
	LicenseModelDescription string `json:"licenseModelDescription"`

	// System requirements for the listing.
	SystemRequirements string `json:"systemRequirements"`

	// The release date of the listing.
	TimeReleased time.Time `json:"timeReleased"`

	// Release notes for the listing.
	ReleaseNotes string `json:"releaseNotes"`

	// Categories that the listing belongs to.
	Categories []string `json:"categories"`

	Publisher Publisher `json:"publisher"`

	// Languages supported by the listing.
	Languages []string `json:"languages"`

	// Screenshots of the listing.
	Screenshots []string `json:"screenshots"`

	// Videos of the listing.
	Videos []string `json:"videos"`

	// Contact information to use to get support from the publisher for the listing.
	SupportContacts SupportContact `json:"supportContacts"`

	// Links to support resources for the listing.
	SupportLinks []string `json:"supportLinks"`

	// Links to additional documentation provided by the publisher specifically for the listing.
	DocumentationLinks []string `json:"documentationLinks"`

	Icon string `json:"icon"`

	Banner string `json:"banner"`

	// The list of compatible architectures supported by the listing
	CompatibleArchitectures []string `json:"compatibleArchitectures,omitempty"`

	// The regions where you can deploy the listing. (Some listings have restrictions that limit their deployment to United States regions only.)
	Regions []string `json:"regions"`

	// The default package version.
	DefaultPackageVersion string `json:"defaultPackageVersion"`

	// Indicates whether the listing is included in Featured Listings.
	IsFeatured bool `json:"isFeatured"`

	// The publisher category to which the listing belongs. The publisher category informs where the listing appears for use.
	ListingType ListingTypeEnum `json:"listingType,omitempty"`

	// List of operating systems supported by the listing.
	SupportedOperatingSystems []string `json:"supportedOperatingSystems"`
}

// ListingRevision The model for an Oracle Cloud Infrastructure Marketplace Publisher listing revision.
type ListingRevision struct {

	// Unique OCID identifier for the listing revision in Marketplace Publisher.
	Id string `json:"id"`

	// The unique identifier for the listing this revision belongs to.
	ListingId string `json:"listingId"`

	// The name for the listing revision.
	DisplayName string `json:"displayName"`

	// Single line introduction for the listing revision.
	Headline string `json:"headline"`

	// The time the listing revision was created. An RFC3339 formatted datetime string.
	TimeCreated time.Time `json:"timeCreated"`

	// The time the listing revision was updated. An RFC3339 formatted datetime string.
	TimeUpdated time.Time `json:"timeUpdated"`

	// The categories for the listing revision.
	Categories []string `json:"categories"`

	// The current status for the Listing revision.
	Status string `json:"status"`

	// The current state of the listing revision.
	LifecycleState string `json:"lifecycleState"`

	// The listing's package type. Populated from the listing.
	PackageType string `json:"packageType"`

	// The pricing model for the listing revision.
	PricingType string `json:"pricingType"`

	// The unique identifier for the compartment.
	CompartmentId string `json:"compartmentId"`

	// The revision number for the listing revision. This is an internal attribute
	RevisionNumber string `json:"revisionNumber"`

	VersionDetails string `json:"versionDetails"`

	// The tagline of the listing revision.
	Tagline string `json:"tagline"`

	// Keywords associated with the listing revision.
	Keywords string `json:"keywords"`

	// A short description for the listing revision.
	ShortDescription string `json:"shortDescription"`

	// Usage information for the listing revision.
	UsageInformation string `json:"usageInformation"`

	// A long description for the listing revision.
	LongDescription string `json:"longDescription"`

	// System requirements for the listing revision.
	SystemRequirements string `json:"systemRequirements"`

	// The markets supported by the listing revision.
	Markets []string `json:"markets"`

	ContentLanguage string `json:"contentLanguage"`

	// Languages supported by the publisher for the listing revision.
	Supportedlanguages []string `json:"supportedlanguages"`

	// Contact information to use to get support from the publisher for the listing revision.
	SupportContacts SupportContact `json:"supportContacts"`

	// Links to support resources for the listing revision.
	SupportLinks []string `json:"supportLinks"`

	Icon string `json:"icon"`

	// Status notes for the listing revision.
	StatusNotes string `json:"statusNotes"`

	// Allowed tenancies provided when a listing revision is published as private.
	AllowedTenancies []string `json:"allowedTenancies"`

	// Identifies whether publisher allows internal tenancy launches for the listing revision.
	AreInternalTenancyLaunchAllowed bool `json:"areInternalTenancyLaunchAllowed"`

	// Additional metadata key/value pairs for the listing revision summary.
	ExtendedMetadata map[string]string `json:"extendedMetadata"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `json:"definedTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `json:"systemTags"`
}

// ListingSummary The model for a summary of the publisher listing.
type ListingSummary struct {

	// The unique OCID of the listing.
	Id string `json:"id"`

	// The unique identifier of the compartment.
	CompartmentId string `json:"compartmentId"`

	// The listing type of the Listing.
	ListingType ListingTypeEnum `json:"listingType"`

	// The name of the listing.
	Name string `json:"name"`

	// The current state for the Listing.
	LifecycleState string `json:"lifecycleState"`

	// The package type for the listing.
	PackageType string `json:"packageType"`

	// The date and time the listing was created, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2023-03-27T21:10:29.600Z`
	TimeCreated time.Time `json:"timeCreated"`

	// The date and time the listing was updated, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2023-03-27T21:10:29.600Z`
	TimeUpdated time.Time `json:"timeUpdated"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `json:"definedTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `json:"systemTags"`
}

// SupportContact Contact information to use to get support.
type SupportContact struct {

	// The name of the contact.
	Name string `json:"name"`

	// The phone number of the contact.
	Phone string `json:"phone"`

	// The email of the contact.
	Email string `json:"email"`

	// The email subject line to use when contacting support.
	Subject string `json:"subject"`
}

type Commitment struct {

	// SPM internal Commitment ID
	Id string `json:"id"`

	// SPM internal Subscribed Service ID
	SubscribedServiceId string `json:"subscribedServiceId"`

	// Commitment start date
	TimeStart time.Time `json:"timeStart"`

	// Commitment end date
	TimeEnd time.Time `json:"timeEnd"`

	// Commitment quantity
	Quantity string `json:"quantity"`

	// Commitment used amount
	UsedAmount string `json:"usedAmount"`

	// Commitment available amount
	AvailableAmount string `json:"availableAmount"`

	// Funded Allocation line value
	// example: 12000.00
	FundedAllocationValue string `json:"fundedAllocationValue"`
}

// ComputedUsage Computed Usage Summary object
type ComputedUsage struct {

	// SPM Internal computed usage Id , 32 character string
	Id string `json:"id"`

	// Computed Usage created time, expressed in RFC 3339 timestamp format.
	TimeCreated time.Time `json:"timeCreated"`

	// Computed Usage updated time, expressed in RFC 3339 timestamp format.
	TimeUpdated time.Time `json:"timeUpdated"`

	// Subscribed service line parent id
	ParentSubscribedServiceId string `json:"parentSubscribedServiceId"`

	ParentProduct ComputedUsageProduct `json:"parentProduct"`

	// Subscription plan number
	PlanNumber string `json:"planNumber"`

	// Currency code
	CurrencyCode string `json:"currencyCode"`

	// References the tier in the ratecard for that usage (OCI will be using the same reference to cross-reference for correctness on the usage csv report), comes from Entity OBSCNTR_IPT_PRODUCTTIER.
	RateCardTierdId string `json:"rateCardTierdId"`

	// Ratecard Id at subscribed service level
	RateCardId string `json:"rateCardId"`

	// SPM Internal compute records source .
	ComputeSource string `json:"computeSource"`

	// Data Center Attribute as sent by MQS to SPM.
	DataCenter string `json:"dataCenter"`

	// MQS Identifier send to SPM , SPM does not transform this attribute and is received as is.
	MqsMessageId string `json:"mqsMessageId"`

	// Total Quantity that was used for computation
	Quantity string `json:"quantity"`

	// SPM Internal usage Line number identifier in SPM coming from Metered Services entity.
	UsageNumber string `json:"usageNumber"`

	// SPM Internal Original usage Line number identifier in SPM coming from Metered Services entity.
	OriginalUsageNumber string `json:"originalUsageNumber"`

	// Subscribed service commitmentId.
	CommitmentServiceId string `json:"commitmentServiceId"`

	// Invoicing status for the aggregated compute usage
	IsInvoiced bool `json:"isInvoiced"`

	// Usage compute type in SPM.
	Type string `json:"type,omitempty"`

	// Usae computation date, expressed in RFC 3339 timestamp format.
	TimeOfArrival time.Time `json:"timeOfArrival"`

	// Metered Service date, expressed in RFC 3339 timestamp format.
	TimeMeteredOn time.Time `json:"timeMeteredOn"`

	// Net Unit Price for the product in consideration, price actual.
	NetUnitPrice string `json:"netUnitPrice"`

	// Computed Line Amount rounded.
	CostRounded string `json:"costRounded"`

	// Computed Line Amount not rounded
	Cost string `json:"cost"`

	Product ComputedUsageProduct `json:"product"`

	// Unit of Measure
	UnitOfMeasure string `json:"unitOfMeasure"`
}

// InvoicingUser User.
type InvoicingUser struct {

	// Name.
	Name string `json:"name"`

	// userName.
	UserName string `json:"userName"`

	// First name.
	FirstName string `json:"firstName"`

	// Last name.
	LastName string `json:"lastName"`

	// Email.
	Email string `json:"email"`

	// TCA contact ID.
	TcaContactId int64 `json:"tcaContactId"`

	// TCA customer account site ID.
	TcaCustAccntSiteId int64 `json:"tcaCustAccntSiteId"`

	// TCA party ID.
	TcaPartyId int64 `json:"tcaPartyId"`
}

// InvoicingProduct Product description
type InvoicingProduct struct {

	// Product part number
	PartNumber string `json:"partNumber"`

	// Product name
	Name string `json:"name"`

	// Unit of Measure
	UnitOfMeasure string `json:"unitOfMeasure"`

	// Rate card part type of Product
	UcmRateCardPartType string `json:"ucmRateCardPartType"`

	// Metered service billing category
	BillingCategory string `json:"billingCategory"`

	// Product category
	ProductCategory string `json:"productCategory"`
}

// InvoicingPaymentTerm Payment Term details
type InvoicingPaymentTerm struct {

	// Payment Term name
	Name string `json:"name"`

	// Payment Term value
	Value string `json:"value"`

	// Payment term Description
	Description string `json:"description"`

	// Payment term active flag
	IsActive bool `json:"isActive"`

	// Payment term last update date
	TimeCreated time.Time `json:"timeCreated"`

	// User that created the Payment term
	CreatedBy string `json:"createdBy"`

	// Payment term last update date
	TimeUpdated time.Time `json:"timeUpdated"`

	// User that updated the Payment term
	UpdatedBy string `json:"updatedBy"`
}

// InvoicingOrganization Organization details
type InvoicingOrganization struct {

	// Organization name
	Name string `json:"name"`

	// Organization ID
	Number float64 `json:"number"`
}

// InvoicingLocation Address location.
type InvoicingLocation struct {

	// Address first line.
	Address1 string `json:"address1"`

	// Address second line.
	Address2 string `json:"address2"`

	// Postal code.
	PostalCode string `json:"postalCode"`

	// City.
	City string `json:"city"`

	// Country.
	Country string `json:"country"`

	// Region.
	Region string `json:"region"`

	// TCA Location identifier.
	TcaLocationId int64 `json:"tcaLocationId"`
}

// InvoicingCurrency Currency details
type InvoicingCurrency struct {

	// Currency Code
	IsoCode string `json:"isoCode"`

	// Currency name
	Name string `json:"name"`

	// Standard Precision of the Currency
	StdPrecision int64 `json:"stdPrecision"`
}

// InvoicingBusinessPartner Business partner.
type InvoicingBusinessPartner struct {

	// Commercial name also called customer name.
	Name string `json:"name"`

	// Phonetic name.
	NamePhonetic string `json:"namePhonetic"`

	// TCA customer account number.
	TcaCustomerAccountNumber string `json:"tcaCustomerAccountNumber"`

	// The business partner is part of the public sector or not.
	IsPublicSector bool `json:"isPublicSector"`

	// The business partner is chain customer or not.
	IsChainCustomer bool `json:"isChainCustomer"`

	// Customer chain type.
	CustomerChainType string `json:"customerChainType"`

	// TCA party number.
	TcaPartyNumber string `json:"tcaPartyNumber"`

	// TCA party ID.
	TcaPartyId int64 `json:"tcaPartyId"`

	// TCA customer account ID.
	TcaCustomerAccountId int64 `json:"tcaCustomerAccountId"`
}

// InvoicingAddress Address.
type InvoicingAddress struct {
	Location InvoicingLocation `json:"location"`

	// Address name identifier.
	Name string `json:"name"`

	// Phone.
	Phone string `json:"phone"`

	// Identify as the customer's billing address.
	IsBillTo bool `json:"isBillTo"`

	// Identify as the customer's shipping address.
	IsShipTo bool `json:"isShipTo"`

	// Bill to site use Id.
	BillSiteUseId int64 `json:"billSiteUseId"`

	// Service to site use Id.
	Service2SiteUseId int64 `json:"service2SiteUseId"`

	// TCA customer account site Id.
	TcaCustAcctSiteId int64 `json:"tcaCustAcctSiteId"`

	// Party site number.
	TcaPartySiteNumber string `json:"tcaPartySiteNumber"`
}

// ComputedUsageProduct Product description
type ComputedUsageProduct struct {

	// Product part number
	PartNumber string `json:"partNumber"`

	// Product name
	Name string `json:"name"`

	// Unit of Measure
	UnitOfMeasure string `json:"unitOfMeasure"`

	// Product provisioning group
	ProvisioningGroup string `json:"provisioningGroup"`

	// Metered service billing category
	BillingCategory string `json:"billingCategory"`

	// Product category
	ProductCategory string `json:"productCategory"`

	// Rate card part type of Product
	UcmRateCardPartType string `json:"ucmRateCardPartType"`
}
