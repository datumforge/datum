package cloud

type Config struct {
	CloudProduct              CloudProduct              `json:"cloud_product"`
	CloudCustomer             CloudCustomer             `json:"cloud_customer"`
	UserFacingProduct         UserFacingProduct         `json:"user_facing_product"`
	AddOn                     AddOn                     `json:"add_on"`
	StripeSetupIntent         StripeSetupIntent         `json:"stripe_setup_intent"`
	CloudSubscription         CloudSubscription         `json:"cloud_subscription"`
	SubscriptionHistory       SubscriptionHistory       `json:"subscription_history"`
	SubscriptionHistoryChange SubscriptionHistoryChange `json:"subscription_history_change"`
	Invoice                   Invoice                   `json:"invoice"`
	InvoiceLineItem           InvoiceLineItem           `json:"invoice_line_item"`
	PaymentMethod             PaymentMethod             `json:"payment_method"`
	Address                   Address                   `json:"address"`
}

const (
	EventTypeFailedPayment                = "failed-payment"
	EventTypeFailedPaymentNoCard          = "failed-payment-no-card"
	EventTypeSendAdminWelcomeEmail        = "send-admin-welcome-email"
	EventTypeSendUpgradeConfirmationEmail = "send-upgrade-confirmation-email"
	EventTypeSubscriptionChanged          = "subscription-changed"
	EventTypeTriggerDelinquencyEmail      = "trigger-delinquency-email"
)

const UpcomingInvoice = "upcoming"

type CloudBillingScheme string

const (
	BillingSchemePerSeat    = CloudBillingScheme("per_seat")
	BillingSchemeFlatFee    = CloudBillingScheme("flat_fee")
	BillingSchemeSalesServe = CloudBillingScheme("sales_serve")
)

type BillingType string

const (
	BillingTypeLicensed = BillingType("licensed")
	BillingTypeInternal = BillingType("internal")
)

type RecurringInterval string

const (
	RecurringIntervalYearly  = RecurringInterval("year")
	RecurringIntervalMonthly = RecurringInterval("month")
)

type SubscriptionFamily string

const (
	SubscriptionFamilyCloud  = SubscriptionFamily("cloud")
	SubscriptionFamilyOnPrem = SubscriptionFamily("on-prem")
)

type ProductSku string

const (
	SkuStarterGov        = ProductSku("starter-gov")
	SkuProfessionalGov   = ProductSku("professional-gov")
	SkuEnterpriseGov     = ProductSku("enterprise-gov")
	SkuStarter           = ProductSku("starter")
	SkuProfessional      = ProductSku("professional")
	SkuEnterprise        = ProductSku("enterprise")
	SkuCloudStarter      = ProductSku("cloud-starter")
	SkuCloudProfessional = ProductSku("cloud-professional")
	SkuCloudEnterprise   = ProductSku("cloud-enterprise")
)

// CloudProduct model represents a product on the cloud system.
type CloudProduct struct {
	ID                 string             `json:"id"`
	Name               string             `json:"name"`
	Description        string             `json:"description"`
	PricePerSeat       float64            `json:"price_per_seat"`
	AddOns             AddOn              `json:"add_ons"`
	SKU                string             `json:"sku"`
	PriceID            string             `json:"price_id"`
	Family             SubscriptionFamily `json:"product_family"`
	RecurringInterval  RecurringInterval  `json:"recurring_interval"`
	CloudBillingScheme CloudBillingScheme `json:"billing_scheme"`
	CrossSellsTo       string             `json:"cross_sells_to"`
}

type UserFacingProduct struct {
	ID                string            `json:"id"`
	Name              string            `json:"name"`
	SKU               string            `json:"sku"`
	PricePerSeat      float64           `json:"price_per_seat"`
	RecurringInterval RecurringInterval `json:"recurring_interval"`
	CrossSellsTo      string            `json:"cross_sells_to"`
}

// AddOn represents an addon to a product.
type AddOn struct {
	ID           string  `json:"id"`
	Name         string  `json:"name"`
	DisplayName  string  `json:"display_name"`
	PricePerSeat float64 `json:"price_per_seat"`
}

// StripeSetupIntent represents the SetupIntent model from Stripe for updating payment methods.
type StripeSetupIntent struct {
	ID           string `json:"id"`
	ClientSecret string `json:"client_secret"`
}

// ConfirmPaymentMethodRequest contains the fields for the customer payment update API.
type ConfirmPaymentMethodRequest struct {
	StripeSetupIntentID string `json:"stripe_setup_intent_id"`
	SubscriptionID      string `json:"subscription_id"`
}

// Customer model represents a customer on the system.
type CloudCustomer struct {
	CloudCustomerInfo
	ID             string        `json:"id"`
	CreatorID      string        `json:"creator_id"`
	CreateAt       int64         `json:"create_at"`
	BillingAddress Address       `json:"billing_address"`
	CompanyAddress Address       `json:"company_address"`
	PaymentMethod  PaymentMethod `json:"payment_method"`
}

type StartCloudTrialRequest struct {
	Email          string `json:"email"`
	SubscriptionID string `json:"subscription_id"`
}

type ValidateBusinessEmailRequest struct {
	Email string `json:"email"`
}

type ValidateBusinessEmailResponse struct {
	IsValid bool `json:"is_valid"`
}

type SubscriptionLicenseSelfServeStatusResponse struct {
	IsExpandable bool `json:"is_expandable"`
	IsRenewable  bool `json:"is_renewable"`
}

// CloudCustomerInfo represents editable info of a customer.
type CloudCustomerInfo struct {
	Name                  string `json:"name"`
	Email                 string `json:"email,omitempty"`
	ContactFirstName      string `json:"contact_first_name,omitempty"`
	ContactLastName       string `json:"contact_last_name,omitempty"`
	NumEmployees          int    `json:"num_employees"`
	CloudAltPaymentMethod string `json:"monthly_subscription_alt_payment_method"`
}

// Address model represents a customer's address.
type Address struct {
	City       string `json:"city"`
	Country    string `json:"country"`
	Line1      string `json:"line1"`
	Line2      string `json:"line2"`
	PostalCode string `json:"postal_code"`
	State      string `json:"state"`
}

// PaymentMethod represents methods of payment for a customer.
type PaymentMethod struct {
	Type      string `json:"type"`
	LastFour  string `json:"last_four"`
	ExpMonth  int    `json:"exp_month"`
	ExpYear   int    `json:"exp_year"`
	CardBrand string `json:"card_brand"`
	Name      string `json:"name"`
}

// CloudSubscription model represents a subscription on the system.
type CloudSubscription struct {
	ID                      string   `json:"id"`
	CustomerID              string   `json:"customer_id"`
	ProductID               string   `json:"product_id"`
	AddOns                  []string `json:"add_ons"`
	StartAt                 int64    `json:"start_at"`
	EndAt                   int64    `json:"end_at"`
	CreateAt                int64    `json:"create_at"`
	Seats                   int      `json:"seats"`
	Status                  string   `json:"status"`
	DNS                     string   `json:"dns"`
	LastInvoice             Invoice  `json:"last_invoice"`
	UpcomingInvoice         Invoice  `json:"upcoming_invoice"`
	IsFreeTrial             string   `json:"is_free_trial"`
	TrialEndAt              int64    `json:"trial_end_at"`
	DelinquentSince         int64    `json:"delinquent_since"`
	OriginallyLicensedSeats int      `json:"originally_licensed_seats"`
	ComplianceBlocked       string   `json:"compliance_blocked"`
	BillingType             string   `json:"billing_type"`
	CancelAt                int64    `json:"cancel_at"`
	WillRenew               string   `json:"will_renew"`
	SimulatedCurrentTimeMs  int64    `json:"simulated_current_time_ms"`
}

// Subscription History model represents true up event in a yearly subscription
type SubscriptionHistory struct {
	ID             string `json:"id"`
	SubscriptionID string `json:"subscription_id"`
	Seats          int    `json:"seats"`
	CreateAt       int64  `json:"create_at"`
}

type SubscriptionHistoryChange struct {
	SubscriptionID string `json:"subscription_id"`
	Seats          int    `json:"seats"`
	CreateAt       int64  `json:"create_at"`
}

// Invoice model represents a cloud invoice
type Invoice struct {
	ID                 string          `json:"id"`
	Number             string          `json:"number"`
	CreateAt           int64           `json:"create_at"`
	Total              int64           `json:"total"`
	Tax                int64           `json:"tax"`
	Status             string          `json:"status"`
	Description        string          `json:"description"`
	PeriodStart        int64           `json:"period_start"`
	PeriodEnd          int64           `json:"period_end"`
	SubscriptionID     string          `json:"subscription_id"`
	Items              InvoiceLineItem `json:"line_items"`
	CurrentProductName string          `json:"current_product_name"`
}

// InvoiceLineItem model represents a cloud invoice lineitem tied to an invoice.
type InvoiceLineItem struct {
	PriceID      string         `json:"price_id"`
	Total        int64          `json:"total"`
	Quantity     float64        `json:"quantity"`
	PricePerUnit int64          `json:"price_per_unit"`
	Description  string         `json:"description"`
	Type         string         `json:"type"`
	Metadata     map[string]any `json:"metadata"`
	PeriodStart  int64          `json:"period_start"`
	PeriodEnd    int64          `json:"period_end"`
}
