package subscriptions

import "time"

type Config struct {
	SubscribedService                SubscribedService                `json:"subscribedService"`
	SubscriptionProduct              SubscriptionProduct              `json:"subscriptionProduct"`
	SubscriptionSubscribedService    SubscriptionSubscribedService    `json:"subscriptionSubscribedService"`
	CommitmentService                CommitmentService                `json:"commitmentService"`
	SubscribedServiceAddress         SubscribedServiceAddress         `json:"subscribedServiceAddress"`
	SubscribedServiceBusinessPartner SubscribedServiceBusinessPartner `json:"subscribedServiceBusinessPartner"`
	SubscribedServiceLocation        SubscribedServiceLocation        `json:"subscribedServiceLocation"`
	SubscribedServicePaymentTerm     SubscribedServicePaymentTerm     `json:"subscribedServicePaymentTerm"`
	SubscribedServiceSummary         SubscribedServiceSummary         `json:"subscribedServiceSummary"`
	SubscribedServiceUser            SubscribedServiceUser            `json:"subscribedServiceUser"`
	ComputedUsageProduct             ComputedUsageProduct             `json:"computedUsageProduct"`
	RateCardProduct                  RateCardProduct                  `json:"rateCardProduct"`
	RateCardTier                     RateCardTier                     `json:"rateCardTier"`
}

// SubscribedService Subscribed service contract details
type SubscribedService struct {

	// SPM internal Subscribed Service ID
	Id string `json:"id"`

	// Subscribed Service line type
	Type string `json:"type"`

	// Subscribed service line number
	SerialNumber string `json:"serialNumber"`

	// Subscription ID associated to the subscribed service
	SubscriptionId string `json:"subscriptionId"`

	Product RateCardProduct `json:"product"`

	// Subscribed service start date
	TimeStart time.Time `json:"timeStart"`

	// Subscribed service end date
	TimeEnd time.Time `json:"timeEnd"`

	// Subscribed service quantity
	Quantity string `json:"quantity"`

	// Subscribed service status
	Status string `json:"status"`

	// Subscribed service operation type
	OperationType string `json:"operationType"`

	// Subscribed service net unit price
	NetUnitPrice string `json:"netUnitPrice"`

	// Indicates the period for which the commitment amount can be utilised exceeding which the amount lapses. Also used in calculation of total contract line value
	PricePeriod string `json:"pricePeriod"`

	// Subscribed service line net amount
	LineNetAmount string `json:"lineNetAmount"`

	// Indicates if the commitment lines can have different quantities
	IsVariableCommitment bool `json:"isVariableCommitment"`

	// Indicates if a service can receive usages and consequently have available amounts computed
	IsAllowance bool `json:"isAllowance"`

	// Subscribed service used amount
	UsedAmount string `json:"usedAmount"`

	// Subscribed sercice available or remaining amount
	AvailableAmount string `json:"availableAmount"`

	// Funded Allocation line value
	// example: 12000.00
	FundedAllocationValue string `json:"fundedAllocationValue"`

	// Indicator on whether or not there has been usage for the subscribed service
	IsHavingUsage bool `json:"isHavingUsage"`

	// If true compares rate between ratecard and the active pricelist and minimum rate would be fetched
	IsCapToPriceList bool `json:"isCapToPriceList"`

	// Subscribed service credit percentage
	CreditPercentage string `json:"creditPercentage"`

	// This field contains the name of the partner to which the subscription belongs - depending on which the invoicing may differ
	PartnerTransactionType string `json:"partnerTransactionType"`

	// Used in context of service credit lines
	IsCreditEnabled bool `json:"isCreditEnabled"`

	// Overage Policy of Subscribed Service
	OveragePolicy string `json:"overagePolicy"`

	// Overage Bill To of Subscribed Service
	OverageBillTo string `json:"overageBillTo"`

	// Pay As You Go policy of Subscribed Service (Can be null - indicating no payg policy)
	PaygPolicy string `json:"paygPolicy"`

	// Not null if this service has an associated promotion line in SPM. Contains the line identifier from Order Management of
	// the associated promo line.
	PromoOrderLineId int64 `json:"promoOrderLineId"`

	// Promotion Pricing Type of Subscribed Service (Can be null - indicating no promotion pricing)
	PromotionPricingType string `json:"promotionPricingType"`

	// Subscribed service Rate Card Discount Percentage
	RateCardDiscountPercentage string `json:"rateCardDiscountPercentage"`

	// Subscribed service Overage Discount Percentage
	OverageDiscountPercentage string `json:"overageDiscountPercentage"`

	BillToCustomer SubscribedServiceBusinessPartner `json:"billToCustomer"`

	BillToContact SubscribedServiceUser `json:"billToContact"`

	BillToAddress SubscribedServiceAddress `json:"billToAddress"`

	// Payment Number of Subscribed Service
	PaymentNumber string `json:"paymentNumber"`

	// Subscribed service payment expiry date
	TimePaymentExpiry time.Time `json:"timePaymentExpiry"`

	PaymentTerm SubscribedServicePaymentTerm `json:"paymentTerm"`

	// Payment Method of Subscribed Service
	PaymentMethod string `json:"paymentMethod"`

	// Subscribed service Transaction Extension Id
	TransactionExtensionId int64 `json:"transactionExtensionId"`

	// Sales Channel of Subscribed Service
	SalesChannel string `json:"salesChannel"`

	// Subscribed service eligible to renew field
	EligibleToRenew string `json:"eligibleToRenew"`

	// SPM renewed Subscription ID
	RenewedSubscribedServiceId string `json:"renewedSubscribedServiceId"`

	// Term value in Months
	TermValue int64 `json:"termValue"`

	// Term value UOM
	TermValueUom string `json:"termValueUom"`

	// Subscribed service Opportunity Id
	RenewalOptyId int64 `json:"renewalOptyId"`

	// Renewal Opportunity Number of Subscribed Service
	RenewalOptyNumber string `json:"renewalOptyNumber"`

	// Renewal Opportunity Type of Subscribed Service
	RenewalOptyType string `json:"renewalOptyType"`

	// Booking Opportunity Number of Subscribed Service
	BookingOptyNumber string `json:"bookingOptyNumber"`

	// Subscribed service Revenue Line Id
	RevenueLineId int64 `json:"revenueLineId"`

	// Revenue Line NUmber of Subscribed Service
	RevenueLineNumber string `json:"revenueLineNumber"`

	// Subscribed service Major Set
	MajorSet int64 `json:"majorSet"`

	// Subscribed service Major Set Start date
	TimeMajorsetStart time.Time `json:"timeMajorsetStart"`

	// Subscribed service Major Set End date
	TimeMajorsetEnd time.Time `json:"timeMajorsetEnd"`

	// Subscribed service System ARR
	SystemArrInLc string `json:"systemArrInLc"`

	// Subscribed service System ARR in Standard Currency
	SystemArrInSc string `json:"systemArrInSc"`

	// Subscribed service System ATR-ARR
	SystemAtrArrInLc string `json:"systemAtrArrInLc"`

	// Subscribed service System ATR-ARR in Standard Currency
	SystemAtrArrInSc string `json:"systemAtrArrInSc"`

	// Subscribed service Revised ARR
	RevisedArrInLc string `json:"revisedArrInLc"`

	// Subscribed service Revised ARR in Standard Currency
	RevisedArrInSc string `json:"revisedArrInSc"`

	// Subscribed service total value
	TotalValue string `json:"totalValue"`

	// Subscribed service Promotion Amount
	OriginalPromoAmount string `json:"originalPromoAmount"`

	// Sales Order Header associated to the subscribed service
	OrderHeaderId int64 `json:"orderHeaderId"`

	// Sales Order Number associated to the subscribed service
	OrderNumber int64 `json:"orderNumber"`

	// Order Type of Subscribed Service
	OrderType string `json:"orderType"`

	// Sales Order Line Id associated to the subscribed service
	OrderLineId int64 `json:"orderLineId"`

	// Sales Order Line Number associated to the subscribed service
	OrderLineNumber int `json:"orderLineNumber"`

	// Subscribed service commitment schedule Id
	CommitmentScheduleId string `json:"commitmentScheduleId"`

	// Subscribed service sales account party id
	SalesAccountPartyId int64 `json:"salesAccountPartyId"`

	// Subscribed service data center
	DataCenter string `json:"dataCenter"`

	// Subscribed service data center region
	DataCenterRegion string `json:"dataCenterRegion"`

	// Subscribed service admin email id
	AdminEmail string `json:"adminEmail"`

	// Subscribed service buyer email id
	BuyerEmail string `json:"buyerEmail"`

	// Subscribed service source
	SubscriptionSource string `json:"subscriptionSource"`

	// Subscribed service provisioning source
	ProvisioningSource string `json:"provisioningSource"`

	// Subscribed service fulfillment set
	FulfillmentSet string `json:"fulfillmentSet"`

	// Subscribed service intent to pay flag
	IsIntentToPay bool `json:"isIntentToPay"`

	// Subscribed service payg flag
	IsPayg bool `json:"isPayg"`

	// Subscribed service pricing model
	PricingModel string `json:"pricingModel"`

	// Subscribed service program type
	ProgramType string `json:"programType"`

	// Subscribed service start date type
	StartDateType string `json:"startDateType"`

	// Subscribed service provisioning date
	TimeProvisioned time.Time `json:"timeProvisioned"`

	// Subscribed service promotion type
	PromoType string `json:"promoType"`

	ServiceToCustomer SubscribedServiceBusinessPartner `json:"serviceToCustomer"`

	ServiceToContact SubscribedServiceUser `json:"serviceToContact"`

	ServiceToAddress SubscribedServiceAddress `json:"serviceToAddress"`

	SoldToCustomer SubscribedServiceBusinessPartner `json:"soldToCustomer"`

	SoldToContact SubscribedServiceUser `json:"soldToContact"`

	EndUserCustomer SubscribedServiceBusinessPartner `json:"endUserCustomer"`

	EndUserContact SubscribedServiceUser `json:"endUserContact"`

	EndUserAddress SubscribedServiceAddress `json:"endUserAddress"`

	ResellerCustomer SubscribedServiceBusinessPartner `json:"resellerCustomer"`

	ResellerContact SubscribedServiceUser `json:"resellerContact"`

	ResellerAddress SubscribedServiceAddress `json:"resellerAddress"`

	// Subscribed service CSI number
	Csi int64 `json:"csi"`

	// Identifier for a customer's transactions for purchase of ay oracle services
	CustomerTransactionReference string `json:"customerTransactionReference"`

	// Subscribed service partner credit amount
	PartnerCreditAmount string `json:"partnerCreditAmount"`

	// Indicates if the Subscribed service has a single ratecard
	IsSingleRateCard bool `json:"isSingleRateCard"`

	// Subscribed service agreement ID
	AgreementId int64 `json:"agreementId"`

	// Subscribed service agreement name
	AgreementName string `json:"agreementName"`

	// Subscribed service agreement type
	AgreementType string `json:"agreementType"`

	// Subscribed service invoice frequency
	BillingFrequency string `json:"billingFrequency"`

	// Subscribed service welcome email sent date
	TimeWelcomeEmailSent time.Time `json:"timeWelcomeEmailSent"`

	// Subscribed service service configuration email sent date
	TimeServiceConfigurationEmailSent time.Time `json:"timeServiceConfigurationEmailSent"`

	// Subscribed service customer config date
	TimeCustomerConfig time.Time `json:"timeCustomerConfig"`

	// Subscribed service agreement end date
	TimeAgreementEnd time.Time `json:"timeAgreementEnd"`

	// List of Commitment services of a line
	CommitmentServices CommitmentService `json:"commitmentServices"`

	// List of Rate Cards of a Subscribed Service
	RateCards []string `json:"rateCards"`

	// Subscribed service creation date
	TimeCreated time.Time `json:"timeCreated"`

	// User that created the subscribed service
	CreatedBy string `json:"createdBy"`

	// Subscribed service last update date
	TimeUpdated time.Time `json:"timeUpdated"`

	// User that updated the subscribed service
	UpdatedBy string `json:"updatedBy"`

	// SPM Ratecard Type
	RatecardType string `json:"ratecardType"`
}

// SubscriptionProduct Product description
type SubscriptionProduct struct {

	// Product part number
	PartNumber string `mandatory:"true" json:"partNumber"`

	// Product name
	Name string `mandatory:"true" json:"name"`

	// Unit of measure
	UnitOfMeasure string `mandatory:"true" json:"unitOfMeasure"`

	// Product provisioning group
	ProvisioningGroup string `json:"provisioningGroup"`
}

// SubscriptionSubscribedService Subscribed Service summary
type SubscriptionSubscribedService struct {

	// SPM internal Subscribed Service ID
	Id string `mandatory:"true" json:"id"`

	Product SubscriptionProduct `json:"product"`

	// Subscribed service quantity
	Quantity string `json:"quantity"`

	// Subscribed service status
	Status string `json:"status"`

	// Subscribed service operation type
	OperationType string `json:"operationType"`

	// Subscribed service net unit price
	NetUnitPrice string `json:"netUnitPrice"`

	// Subscribed service used amount
	UsedAmount string `json:"usedAmount"`

	// Subscribed sercice available or remaining amount
	AvailableAmount string `json:"availableAmount"`

	// Funded Allocation line value
	// example: 12000.00
	FundedAllocationValue string `json:"fundedAllocationValue"`

	// This field contains the name of the partner to which the subscription belongs - depending on which the invoicing may differ
	PartnerTransactionType string `json:"partnerTransactionType"`

	// Term value in Months
	TermValue int64 `json:"termValue"`

	// Term value UOM
	TermValueUom string `json:"termValueUom"`

	// Booking Opportunity Number of Subscribed Service
	BookingOptyNumber string `json:"bookingOptyNumber"`

	// Subscribed service total value
	TotalValue string `json:"totalValue"`

	// Subscribed service Promotion Amount
	OriginalPromoAmount string `json:"originalPromoAmount"`

	// Sales Order Number associated to the subscribed service
	OrderNumber int64 `json:"orderNumber"`

	// Subscribed service data center region
	DataCenterRegion string `json:"dataCenterRegion"`

	// Subscribed service pricing model
	PricingModel string `json:"pricingModel"`

	// Subscribed service program type
	ProgramType string `json:"programType"`

	// Subscribed service promotion type
	PromoType string `json:"promoType"`

	// Subscribed service CSI number
	Csi int64 `json:"csi"`

	// Subscribed service intent to pay flag
	IsIntentToPay bool `json:"isIntentToPay"`

	// Subscribed service start date
	TimeStart time.Time `json:"timeStart"`

	// Subscribed service end date
	TimeEnd time.Time `json:"timeEnd"`

	// List of Commitment services of a line
	CommitmentServices CommitmentService `json:"commitmentServices"`
}

// CommitmentService Subscribed service commitment details
type CommitmentService struct {

	// Commitment start date
	TimeStart time.Time `json:"timeStart"`

	// Commitment end date
	TimeEnd time.Time `json:"timeEnd"`

	// Commitment quantity
	Quantity string `json:"quantity"`

	// Commitment available amount
	AvailableAmount string `json:"availableAmount"`

	// Commitment line net amount
	LineNetAmount string `json:"lineNetAmount"`

	// Funded Allocation line value
	FundedAllocationValue string `json:"fundedAllocationValue"`
}

// SubscribedServiceAddress Address.
type SubscribedServiceAddress struct {
	Location SubscribedServiceLocation `json:"location"`

	// Address name identifier.
	Name string `json:"name"`

	// Phone.
	Phone string `json:"phone"`

	// Identify as the customer shipping address.
	IsBillTo bool `json:"isBillTo"`

	// Identify as the customer invoicing address.
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

// SubscribedServiceBusinessPartner Business partner.
type SubscribedServiceBusinessPartner struct {

	// Commercial name also called customer name.
	Name string `json:"name"`

	// Phonetic name.
	NamePhonetic string `json:"namePhonetic"`

	// TCA customer account number.
	TcaCustAccountNumber string `json:"tcaCustAccountNumber"`

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

// SubscribedServiceLocation Address location.
type SubscribedServiceLocation struct {

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

	// Region.
	TcaLocationId int64 `json:"tcaLocationId"`
}

// SubscribedServicePaymentTerm Payment Term details
type SubscribedServicePaymentTerm struct {

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

// SubscribedServiceSummary Subscribed service contract details
type SubscribedServiceSummary struct {

	// SPM internal Subscribed Service ID
	Id string `mandatory:"true" json:"id"`

	// Subscribed Service line type
	Type string `json:"type"`

	// Subscribed service line number
	SerialNumber string `json:"serialNumber"`

	// Subscription ID associated to the subscribed service
	SubscriptionId string `json:"subscriptionId"`

	Product RateCardProduct `json:"product"`

	// Subscribed service start date
	TimeStart time.Time `json:"timeStart"`

	// Subscribed service end date
	TimeEnd time.Time `json:"timeEnd"`

	// Subscribed service quantity
	Quantity string `json:"quantity"`

	// Subscribed service status
	Status string `json:"status"`

	// Subscribed service operation type
	OperationType string `json:"operationType"`

	// Subscribed service net unit price
	NetUnitPrice string `json:"netUnitPrice"`

	// Indicates the period for which the commitment amount can be utilised exceeding which the amount lapses. Also used in calculation of total contract line value
	PricePeriod string `json:"pricePeriod"`

	// Subscribed service line net amount
	LineNetAmount string `json:"lineNetAmount"`

	// Indicates if the commitment lines can have different quantities
	IsVariableCommitment bool `json:"isVariableCommitment"`

	// Indicates if a service can receive usages and consequently have available amounts computed
	IsAllowance bool `json:"isAllowance"`

	// Subscribed service used amount
	UsedAmount string `json:"usedAmount"`

	// Subscribed sercice available or remaining amount
	AvailableAmount string `json:"availableAmount"`

	// Funded Allocation line value
	// example: 12000.00
	FundedAllocationValue string `json:"fundedAllocationValue"`

	// Indicator on whether or not there has been usage for the subscribed service
	IsHavingUsage bool `json:"isHavingUsage"`

	// If true compares rate between ratecard and the active pricelist and minimum rate would be fetched
	IsCapToPriceList bool `json:"isCapToPriceList"`

	// Subscribed service credit percentage
	CreditPercentage string `json:"creditPercentage"`

	// This field contains the name of the partner to which the subscription belongs - depending on which the invoicing may differ
	PartnerTransactionType string `json:"partnerTransactionType"`

	// Used in context of service credit lines
	IsCreditEnabled bool `json:"isCreditEnabled"`

	// Overage Policy of Subscribed Service
	OveragePolicy string `json:"overagePolicy"`

	// Overage Bill To of Subscribed Service
	OverageBillTo string `json:"overageBillTo"`

	// Pay As You Go policy of Subscribed Service (Can be null - indicating no payg policy)
	PaygPolicy string `json:"paygPolicy"`

	// Not null if this service has an associated promotion line in SPM. Contains the line identifier from Order Management of
	// the associated promo line.
	PromoOrderLineId int64 `json:"promoOrderLineId"`

	// Promotion Pricing Type of Subscribed Service (Can be null - indicating no promotion pricing)
	PromotionPricingType string `json:"promotionPricingType"`

	// Subscribed service Rate Card Discount Percentage
	RateCardDiscountPercentage string `json:"rateCardDiscountPercentage"`

	// Subscribed service Overage Discount Percentage
	OverageDiscountPercentage string `json:"overageDiscountPercentage"`

	BillToCustomer SubscribedServiceBusinessPartner `json:"billToCustomer"`

	BillToContact SubscribedServiceUser `json:"billToContact"`

	BillToAddress SubscribedServiceAddress `json:"billToAddress"`

	// Payment Number of Subscribed Service
	PaymentNumber string `json:"paymentNumber"`

	// Subscribed service payment expiry date
	TimePaymentExpiry time.Time `json:"timePaymentExpiry"`

	PaymentTerm SubscribedServicePaymentTerm `json:"paymentTerm"`

	// Payment Method of Subscribed Service
	PaymentMethod string `json:"paymentMethod"`

	// Subscribed service Transaction Extension Id
	TransactionExtensionId int64 `json:"transactionExtensionId"`

	// Sales Channel of Subscribed Service
	SalesChannel string `json:"salesChannel"`

	// Subscribed service eligible to renew field
	EligibleToRenew string `json:"eligibleToRenew"`

	// SPM renewed Subscription ID
	RenewedSubscribedServiceId string `json:"renewedSubscribedServiceId"`

	// Term value in Months
	TermValue int64 `json:"termValue"`

	// Term value UOM
	TermValueUom string `json:"termValueUom"`

	// Subscribed service Opportunity Id
	RenewalOptyId int64 `json:"renewalOptyId"`

	// Renewal Opportunity Number of Subscribed Service
	RenewalOptyNumber string `json:"renewalOptyNumber"`

	// Renewal Opportunity Type of Subscribed Service
	RenewalOptyType string `json:"renewalOptyType"`

	// Booking Opportunity Number of Subscribed Service
	BookingOptyNumber string `json:"bookingOptyNumber"`

	// Subscribed service Revenue Line Id
	RevenueLineId int64 `json:"revenueLineId"`

	// Revenue Line NUmber of Subscribed Service
	RevenueLineNumber string `json:"revenueLineNumber"`

	// Subscribed service Major Set
	MajorSet int64 `json:"majorSet"`

	// Subscribed service Major Set Start date
	TimeMajorsetStart time.Time `json:"timeMajorsetStart"`

	// Subscribed service Major Set End date
	TimeMajorsetEnd time.Time `json:"timeMajorsetEnd"`

	// Subscribed service System ARR
	SystemArrInLc string `json:"systemArrInLc"`

	// Subscribed service System ARR in Standard Currency
	SystemArrInSc string `json:"systemArrInSc"`

	// Subscribed service System ATR-ARR
	SystemAtrArrInLc string `json:"systemAtrArrInLc"`

	// Subscribed service System ATR-ARR in Standard Currency
	SystemAtrArrInSc string `json:"systemAtrArrInSc"`

	// Subscribed service Revised ARR
	RevisedArrInLc string `json:"revisedArrInLc"`

	// Subscribed service Revised ARR in Standard Currency
	RevisedArrInSc string `json:"revisedArrInSc"`

	// Subscribed service total value
	TotalValue string `json:"totalValue"`

	// Subscribed service Promotion Amount
	OriginalPromoAmount string `json:"originalPromoAmount"`

	// Sales Order Header associated to the subscribed service
	OrderHeaderId int64 `json:"orderHeaderId"`

	// Sales Order Number associated to the subscribed service
	OrderNumber int64 `json:"orderNumber"`

	// Order Type of Subscribed Service
	OrderType string `json:"orderType"`

	// Sales Order Line Id associated to the subscribed service
	OrderLineId int64 `json:"orderLineId"`

	// Sales Order Line Number associated to the subscribed service
	OrderLineNumber int `json:"orderLineNumber"`

	// Subscribed service commitment schedule Id
	CommitmentScheduleId string `json:"commitmentScheduleId"`

	// Subscribed service sales account party id
	SalesAccountPartyId int64 `json:"salesAccountPartyId"`

	// Subscribed service data center
	DataCenter string `json:"dataCenter"`

	// Subscribed service data center region
	DataCenterRegion string `json:"dataCenterRegion"`

	// Subscribed service admin email id
	AdminEmail string `json:"adminEmail"`

	// Subscribed service buyer email id
	BuyerEmail string `json:"buyerEmail"`

	// Subscribed service source
	SubscriptionSource string `json:"subscriptionSource"`

	// Subscribed service provisioning source
	ProvisioningSource string `json:"provisioningSource"`

	// Subscribed service fulfillment set
	FulfillmentSet string `json:"fulfillmentSet"`

	// Subscribed service intent to pay flag
	IsIntentToPay bool `json:"isIntentToPay"`

	// Subscribed service payg flag
	IsPayg bool `json:"isPayg"`

	// Subscribed service pricing model
	PricingModel string `json:"pricingModel"`

	// Subscribed service program type
	ProgramType string `json:"programType"`

	// Subscribed service start date type
	StartDateType string `json:"startDateType"`

	// Subscribed service provisioning date
	TimeProvisioned time.Time `json:"timeProvisioned"`

	// Subscribed service promotion type
	PromoType string `json:"promoType"`

	ServiceToCustomer SubscribedServiceBusinessPartner `json:"serviceToCustomer"`

	ServiceToContact SubscribedServiceUser `json:"serviceToContact"`

	ServiceToAddress SubscribedServiceAddress `json:"serviceToAddress"`

	SoldToCustomer SubscribedServiceBusinessPartner `json:"soldToCustomer"`

	SoldToContact SubscribedServiceUser `json:"soldToContact"`

	EndUserCustomer SubscribedServiceBusinessPartner `json:"endUserCustomer"`

	EndUserContact SubscribedServiceUser `json:"endUserContact"`

	EndUserAddress SubscribedServiceAddress `json:"endUserAddress"`

	ResellerCustomer SubscribedServiceBusinessPartner `json:"resellerCustomer"`

	ResellerContact SubscribedServiceUser `json:"resellerContact"`

	ResellerAddress SubscribedServiceAddress `json:"resellerAddress"`

	// Subscribed service CSI number
	Csi int64 `json:"csi"`

	// Identifier for a customer's transactions for purchase of ay oracle services
	CustomerTransactionReference string `json:"customerTransactionReference"`

	// Subscribed service partner credit amount
	PartnerCreditAmount string `json:"partnerCreditAmount"`

	// Indicates if the Subscribed service has a single ratecard
	IsSingleRateCard bool `json:"isSingleRateCard"`

	// Subscribed service agreement ID
	AgreementId int64 `json:"agreementId"`

	// Subscribed service agreement name
	AgreementName string `json:"agreementName"`

	// Subscribed service agreement type
	AgreementType string `json:"agreementType"`

	// Subscribed service invoice frequency
	BillingFrequency string `json:"billingFrequency"`

	// Subscribed service welcome email sent date
	TimeWelcomeEmailSent time.Time `json:"timeWelcomeEmailSent"`

	// Subscribed service service configuration email sent date
	TimeServiceConfigurationEmailSent time.Time `json:"timeServiceConfigurationEmailSent"`

	// Subscribed service customer config date
	TimeCustomerConfig time.Time `json:"timeCustomerConfig"`

	// Subscribed service agreement end date
	TimeAgreementEnd time.Time `json:"timeAgreementEnd"`

	// Subscribed service creation date
	TimeCreated time.Time `json:"timeCreated"`

	// User that created the subscribed service
	CreatedBy string `json:"createdBy"`

	// Subscribed service last update date
	TimeUpdated time.Time `json:"timeUpdated"`

	// User that updated the subscribed service
	UpdatedBy string `json:"updatedBy"`

	// SPM Ratecard Type
	RatecardType string `json:"ratecardType"`
}

// SubscribedServiceUser User.
type SubscribedServiceUser struct {

	// Name.
	Name string `json:"name"`

	// Username.
	Username string `json:"username"`

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

// ComputedUsageProduct Product description
type ComputedUsageProduct struct {

	// Product part number
	PartNumber string `mandatory:"true" json:"partNumber"`

	// Product name
	Name string `mandatory:"true" json:"name"`

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

// RateCardProduct Product description
type RateCardProduct struct {

	// Product part number
	PartNumber string `mandatory:"true" json:"partNumber"`

	// Product name
	Name string `mandatory:"true" json:"name"`

	// Unit of measure
	UnitOfMeasure string `mandatory:"true" json:"unitOfMeasure"`

	// Metered service billing category
	BillingCategory string `json:"billingCategory"`

	// Product category
	ProductCategory string `json:"productCategory"`

	// Rate card part type of Product
	UcmRateCardPartType string `json:"ucmRateCardPartType"`
}

// RateCardTier Rate Card Tier details
type RateCardTier struct {

	// Rate card tier quantity range
	UpToQuantity string `json:"upToQuantity"`

	// Rate card tier net unit price
	NetUnitPrice string `json:"netUnitPrice"`

	// Rate card tier overage price
	OveragePrice string `json:"overagePrice"`
}
