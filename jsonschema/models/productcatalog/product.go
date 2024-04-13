package productcatalog

type Config struct {
	Product            Product            `json:"product" yaml:"product"`
	BehaviorConfig     BehaviorConfig     `json:"behavior_config" yaml:"behavior_config"`
	Price              Price              `json:"price" yaml:"price"`
	Tier               Tier               `json:"tier" yaml:"tier"`
	Sku                Sku                `json:"sku" yaml:"sku"`
	AggregationInfo    AggregationInfo    `json:"aggregation_info" yaml:"aggregation_info"`
	Category           Category           `json:"category" yaml:"category"`
	GeoTaxonomy        GeoTaxonomy        `json:"geo_taxonomy" yaml:"geo_taxonomy"`
	Money              Money              `json:"money" yaml:"money"`
	PricingExpression  PricingExpression  `json:"pricing_expression" yaml:"pricing_expression"`
	PricingInfo        PricingInfo        `json:"pricing_info" yaml:"pricing_info"`
	ProjectBillingInfo ProjectBillingInfo `json:"project_billing_info" yaml:"project_billing_info"`
}

type Behavior string

const (
	BasicBehavior   Behavior = "basic"
	CreditBehavior  Behavior = "credits"
	PerSeatBehavior Behavior = "per_seat"
)

type BehaviorConfig struct {
	// CreditAmount is amount of credits that are awarded/consumed when buying/using this feature
	CreditAmount int64 `json:"credit_amount" yaml:"credit_amount"`

	// SeatLimit is the maximum number of seats that can be added to the subscription
	SeatLimit int64 `json:"seat_limit" yaml:"seat_limit"`
}

// Product is an item being sold by the platform and has a corresponding reference
// in the billing engine
type Product struct {
	ID         string   `json:"id" yaml:"id"`
	ProviderID string   `json:"provider_id" yaml:"provider_id"` // in case of stripe, provider id and id are same
	PlanIDs    []string // plans this feature belongs to, this is optional and can be empty

	Name        string `json:"name" yaml:"name"`   // a machine friendly name for the feature
	Title       string `json:"title" yaml:"title"` // a human friendly title for the feature
	Description string `json:"description" yaml:"description"`

	// Type is the type of the feature
	// known types are "credits" and "per_seat". Default is "basic"
	Behavior Behavior `json:"behavior" yaml:"behavior" default:"basic"`

	// Config is the configuration for the behavior
	Config BehaviorConfig `json:"config" yaml:"config"`

	// Prices for the product, return only, shouldn't be set while updating a product
	Prices Price `json:"prices" yaml:"prices"`
	// Features for the product, return only, shouldn't be set while updating a product
	Features Feature `json:"features" yaml:"features"`

	State    string `json:"state" yaml:"state"`
	Metadata string `json:"metadata" yaml:"metadata"`
}

type PriceUsageType string

const (
	PriceUsageTypeLicensed PriceUsageType = "licensed"
	PriceUsageTypeMetered  PriceUsageType = "metered"
)

type BillingScheme string

const (
	BillingSchemeFlat   BillingScheme = "flat"
	BillingSchemeTiered BillingScheme = "tiered"
)

type PriceTierMode string

const (
	PriceTierModeGraduated PriceTierMode = "graduated"
	PriceTierModeVolume    PriceTierMode = "volume"
)

// Price is a product price and has a corresponding price in the billing engine
// when creating a price, the feature must already exist
// when subscribing to a plan, the price must already exist
type Price struct {
	ID         string `json:"id" yaml:"id"`
	ProductID  string `json:"feature_id" yaml:"feature_id"`
	ProviderID string `json:"provider_id" yaml:"provider_id"`

	Name string `json:"name" yaml:"name"` // a machine friendly name for the price

	// BillingScheme specifies the billing scheme for the price
	// known schemes are "tiered" and "flat". Default is "flat"
	BillingScheme BillingScheme `json:"billing_scheme" yaml:"billing_scheme" default:"flat"`

	// Currency Three-letter ISO 4217 currency code in lower case
	// like "usd", "eur", "gbp"
	// https://www.six-group.com/en/products-services/financial-information/data-standards.html
	Currency string `json:"currency" yaml:"currency" default:"usd"`

	// Amount price in the minor currency unit
	// Minor unit is the smallest unit of a currency, e.g. 1 dollar equals 100 cents (with 2 decimals).
	Amount int64 `json:"amount" yaml:"amount"`

	// UsageType specifies the usage type for the price
	// known types are "licensed" and "metered". Default is "licensed"
	UsageType PriceUsageType `json:"usage_type" yaml:"usage_type" default:"licensed"`

	// MeteredAggregate specifies the aggregation method for the price
	// known aggregations are "sum", "last_during_period", "last_ever" and "max". Default is "sum"
	MeteredAggregate string `json:"metered_aggregate" yaml:"metered_aggregate" default:"sum"`

	// Interval is the interval at which the plan is billed
	// e.g. day, week, month, year
	Interval string `json:"interval" yaml:"interval"`

	Metadata string `json:"metadata" yaml:"metadata"`

	// TierMode specifies the tier mode for the price
	// known modes are "graduated" and "volume". Default is "graduated"
	// In volume-based, the maximum quantity within a period determines the per-unit price
	// In graduated, pricing changes as the quantity increases to specific thresholds
	TierMode string `json:"tier_mode" yaml:"tier_mode" default:"graduated"`

	// Tiers specifies the optional tiers for the price
	// only applicable when BillingScheme is "tiered"
	Tiers Tier `json:"tiers" yaml:"tiers"`

	State string `json:"state" yaml:"state"`
}

type Tier struct {
	FlatAmount int64 `json:"flat_amount" yaml:"flat_amount"`
	UpTo       int64 `json:"up_to" yaml:"up_to"`
}

// Feature are part of a product which allows for a more granular control on
// what is packed with the product. It is a platform specific concept and
// doesn't have a corresponding billing engine entity
type Feature struct {
	ID    string `json:"id" yaml:"id"`
	Name  string `json:"name" yaml:"name"`   // a machine friendly name for the feature
	Title string `json:"title" yaml:"title"` // a human friendly title for the feature

	// products this feature belongs to, this is optional and can be empty
	// a product will have at least one feature with the same name as the product
	ProductIDs []string `json:"product_ids" yaml:"product_ids"`

	Metadata string `json:"metadata" yaml:"metadata"`
}

// Sku: Encapsulates a single SKU
type Sku struct {
	// Category: The category hierarchy of this SKU, purely for
	// organizational purpose.
	Category Category `json:"category"`

	// Description: A human readable description of the SKU, has a maximum
	// length of 256 characters.
	Description string `json:"description"`

	// GeoTaxonomy: The geographic taxonomy for this sku.
	GeoTaxonomy GeoTaxonomy `json:"geoTaxonomy"`

	// Name: The resource name for the SKU. Example:
	// "services/DA34-426B-A397/skus/AA95-CD31-42FE"
	Name string `json:"name"`

	// PricingInfo: A timeline of pricing info for this SKU in chronological
	// order
	PricingInfo PricingInfo `json:"pricingInfo"`

	// ServiceProviderName: Identifies the service provider
	ServiceProviderName string `json:"serviceProviderName"`

	// ServiceRegions: List of service regions this SKU is offered at.
	// Example: "asia-east1"
	ServiceRegions []string `json:"serviceRegions"`

	// SkuId: The identifier for the SKU. Example: "AA95-CD31-42FE"
	SkuId string `json:"skuId"`
}

// AggregationInfo: Represents the aggregation level and interval for
// pricing of a single SKU
type AggregationInfo struct {
	// AggregationCount: The number of intervals to aggregate over. Example:
	// If aggregation_level is "DAILY" and aggregation_count is 14,
	// aggregation will be over 14 days.
	AggregationCount int64 `json:"aggregationCount"`

	// Possible values:
	//   "AGGREGATION_INTERVAL_UNSPECIFIED"
	//   "DAILY"
	//   "MONTHLY"
	AggregationInterval string `json:"aggregationInterval"`

	// Possible values:
	//   "AGGREGATION_LEVEL_UNSPECIFIED"
	//   "ACCOUNT"
	//   "PROJECT"
	AggregationLevel string `json:"aggregationLevel"`
}

// Category: Represents the category hierarchy of a SKU
type Category struct {
	ResourceFamily string `json:"resourceFamily"`

	ResourceGroup string `json:"resourceGroup"`

	// ServiceDisplayName: The display name of the service this SKU belongs
	// to.
	ServiceDisplayName string `json:"serviceDisplayName"`

	// UsageType: Represents how the SKU is consumed. Example: "OnDemand",
	// "Preemptible", "Commit1Mo", "Commit1Yr" etc.
	UsageType string `json:"usageType"`
}

// GeoTaxonomy: Encapsulates the geographic taxonomy data for a sku
type GeoTaxonomy struct {
	// Regions: The list of regions associated with a sku. Empty for Global
	// skus, which are associated with all regions
	Regions []string `json:"regions"`

	// Type: The type of Geo Taxonomy: GLOBAL, REGIONAL, or MULTI_REGIONAL.
	//
	// Possible values:
	//   "TYPE_UNSPECIFIED" - The type is not specified.
	//   "GLOBAL" - The sku is global in nature, e.g. a license sku. Global
	// skus are available in all regions, and so have an empty region list.
	//   "REGIONAL" - The sku is available in a specific region, e.g.
	// "us-west2".
	//   "MULTI_REGIONAL" - The sku is associated with multiple regions,
	// e.g. "us-west2" and "us-east1".
	Type string `json:"type"`
}

// Money: Represents an amount of money with its currency type.
type Money struct {
	// CurrencyCode: The three-letter currency code defined in ISO 4217.
	CurrencyCode string `json:"currencyCode"`

	// Nanos: Number of nano (10^-9) units of the amount. The value must be
	// between -999,999,999 and +999,999,999 inclusive. If `units` is
	// positive, `nanos` must be positive or zero. If `units` is zero,
	// `nanos` can be positive, zero, or negative. If `units` is negative,
	// `nanos` must be negative or zero. For example $-1.75 is represented
	// as `units`=-1 and `nanos`=-750,000,000.
	Nanos int64 `json:"nanos"`

	// Units: The whole units of the amount. For example if `currencyCode`
	// is "USD", then 1 unit is one US dollar.
	Units int64 `json:"units,string"`
}

type PricingExpression struct {
	// BaseUnit: The base unit for the SKU which is the unit used in usage
	// exports. Example: "By"
	BaseUnit string `json:"baseUnit"`

	// BaseUnitConversionFactor: Conversion factor for converting from price
	// per usage_unit to price per base_unit, and start_usage_amount to
	// start_usage_amount in base_unit. unit_price /
	// base_unit_conversion_factor = price per base_unit. start_usage_amount
	//  base_unit_conversion_factor = start_usage_amount in base_unit.
	BaseUnitConversionFactor float64 `json:"baseUnitConversionFactor"`

	// BaseUnitDescription: The base unit in human readable form. Example:
	// "byte".
	BaseUnitDescription string `json:"baseUnitDescription"`

	DisplayQuantity float64 `json:"displayQuantity"`

	// UsageUnit: The short hand for unit of usage this pricing is specified
	// in. Example: usage_unit of "GiBy" means that usage is specified in
	// "Gibi Byte".
	UsageUnit string `json:"usageUnit"`

	// UsageUnitDescription: The unit of usage in human readable form.
	// Example: "gibi byte".
	UsageUnitDescription string `json:"usageUnitDescription"`
}

// PricingInfo: Represents the pricing information for a SKU at a single
// point of time.
type PricingInfo struct {
	// AggregationInfo: Aggregation Info. This can be left unspecified if
	// the pricing expression doesn't require aggregation.
	AggregationInfo AggregationInfo `json:"aggregationInfo"`

	CurrencyConversionRate float64 `json:"currencyConversionRate"`

	EffectiveTime string `json:"effectiveTime"`

	// PricingExpression: Expresses the pricing formula. See
	// `PricingExpression` for an example.
	PricingExpression PricingExpression `json:"pricingExpression"`

	// Summary: An optional human readable summary of the pricing
	// information, has a maximum length of 256 characters.
	Summary string `json:"summary"`
}

// ProjectBillingInfo: Encapsulation of billing information
type ProjectBillingInfo struct {
	BillingAccountName string `json:"billingAccountName"`

	BillingEnabled bool `json:"billingEnabled"`

	Name string `json:"name"`

	ProjectId string `json:"projectId"`
}
