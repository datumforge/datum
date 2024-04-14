package invoice

import "time"

type Config struct {
	InvoicingUser            InvoicingUser            `json:"invoicingUser"`
	InvoicingProduct         InvoicingProduct         `json:"invoicingProduct"`
	InvoicingPaymentTerm     InvoicingPaymentTerm     `json:"invoicingPaymentTerm"`
	InvoicingOrganization    InvoicingOrganization    `json:"invoicingOrganization"`
	InvoicingLocation        InvoicingLocation        `json:"invoicingLocation"`
	InvoicingAddress         InvoicingAddress         `json:"invoicingAddress"`
	InvoicingBusinessPartner InvoicingBusinessPartner `json:"invoicingBusinessPartner"`
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
