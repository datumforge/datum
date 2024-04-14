package infra

type Config struct {
	BGPRoute               BGPRoute               `json:"bgp_route"`
	Server                 Server                 `json:"server"`
	ServerBGP              ServerBGP              `json:"server_bgp"`
	Project                Project                `json:"project"`
	Region                 Region                 `json:"region"`
	Plan                   Plan                   `json:"plan"`
	Plans                  Plans                  `json:"plans"`
	BlockStorage           BlockStorage           `json:"block_storage"`
	AssignedTo             AssignedTo             `json:"assigned_to"`
	RoutedTo               RoutedTo               `json:"routed_to"`
	IPAddresses            IPAddresses            `json:"ip_addresses"`
	SSHKeys                SSHKeys                `json:"ssh_keys"`
	Cpus                   Cpus                   `json:"cpus"`
	Memory                 Memory                 `json:"memory"`
	Nics                   Nics                   `json:"nics"`
	Raid                   Raid                   `json:"raid"`
	Storage                Storage                `json:"storage"`
	Bandwidth              Bandwidth              `json:"bandwidth"`
	IPAddressCreateRequest IPAddressCreateRequest `json:"ip_address_create_request"`
	CreateServer           CreateServer           `json:"create_server"`
}

// BGPRoute single server BGP route
type BGPRoute struct {
	Subnet  string `json:"subnet"`
	Active  bool   `json:"active"`
	Router  string `json:"router"`
	Age     string `json:"age"`
	Updated string `json:"updated"`
}

// ServerBGP status of BGP on a server
type ServerBGP struct {
	Enabled   bool       `json:"enabled"`
	Available bool       `json:"available"`
	Status    string     `json:"status"`
	Routers   int        `json:"routers"`
	Connected int        `json:"connected"`
	Limit     int        `json:"limit"`
	Active    int        `json:"active"`
	Routes    []BGPRoute `json:"routes"`
	Updated   string     `json:"updated"`
}

// Project a CherryServers project
type Project struct {
	ID   int        `json:"id"`
	Name string     `json:"name"`
	Bgp  ProjectBGP `json:"bgp"`
	Href string     `json:"href"`
}

// Region a CherryServers region
type Region struct {
	ID         int       `json:"id"`
	Slug       string    `json:"slug"`
	Name       string    `json:"name"`
	RegionIso2 string    `json:"region_iso_2"`
	BGP        RegionBGP `json:"bgp"`
	Href       string    `json:"href"`
}

// RegionBGP information about BGP in a region
type RegionBGP struct {
	Hosts []string `json:"hosts"`
	Asn   int      `json:"asn"`
}

// ProjectBGP information about BGP on an individual project
type ProjectBGP struct {
	Enabled  bool `json:"enabled"`
	LocalASN int  `json:"local_asn"`
}

// Plan a server plan
type Plan struct {
	ID               int                `json:"id"`
	Slug             string             `json:"slug"`
	Name             string             `json:"name"`
	Custom           bool               `json:"custom"`
	Specs            Specs              `json:"specs"`
	Pricing          []Pricing          `json:"pricing"`
	AvailableRegions []AvailableRegions `json:"available_regions"`
}

// Plans represents a list of Cherry Servers plans
type Plans []Plan

// Pricing price for a specific plan
type Pricing struct {
	Price    float32 `json:"price"`
	Taxed    bool    `json:"taxed"`
	Currency string  `json:"currency"`
	Unit     string  `json:"unit"`
}

// AvailableRegions regions that are available to the user
type AvailableRegions struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	RegionIso2 string `json:"region_iso_2"`
	StockQty   int    `json:"stock_qty"`
}

// AttachedTo what a resource is attached to
type AttachedTo struct {
	Href string `json:"href"`
}

// BlockStorage cloud block storage
type BlockStorage struct {
	ID            int        `json:"id"`
	Name          string     `json:"name"`
	Href          string     `json:"href"`
	Size          int        `json:"size"`
	AllowEditSize bool       `json:"allow_edit_size"`
	Unit          string     `json:"unit"`
	Description   string     `json:"description"`
	AttachedTo    AttachedTo `json:"attached_to"`
	VlanID        string     `json:"vlan_id"`
	VlanIP        string     `json:"vlan_ip"`
	Initiator     string     `json:"initiator"`
	DiscoveryIP   string     `json:"discovery_ip"`
}

// AssignedTo assignment of a network floating IP to a server
type AssignedTo struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Href     string  `json:"href"`
	Hostname string  `json:"hostname"`
	Image    string  `json:"image"`
	Region   Region  `json:"region"`
	State    string  `json:"state"`
	Pricing  Pricing `json:"pricing"`
}

// RoutedTo routing of a floating IP to an underlying IP
type RoutedTo struct {
	ID            string `json:"id"`
	Address       string `json:"address"`
	AddressFamily int    `json:"address_family"`
	Cidr          string `json:"cidr"`
	Gateway       string `json:"gateway"`
	Type          string `json:"type"`
	Region        Region `json:"region"`
}

// IPAddresses individual IP address
type IPAddresses struct {
	ID            string            `json:"id"`
	Address       string            `json:"address"`
	AddressFamily int               `json:"address_family"`
	Cidr          string            `json:"cidr"`
	Gateway       string            `json:"gateway"`
	Type          string            `json:"type"`
	Region        Region            `json:"region"`
	RoutedTo      RoutedTo          `json:"routed_to"`
	AssignedTo    AssignedTo        `json:"assigned_to"`
	TargetedTo    AssignedTo        `json:"targeted_to"`
	Project       Project           `json:"project"`
	PtrRecord     string            `json:"ptr_record"`
	ARecord       string            `json:"a_record"`
	Tags          map[string]string `json:"tags"`
	Href          string            `json:"href"`
}

// Server represents a Cherry Servers server
type Server struct {
	ID               int               `json:"id"`
	Name             string            `json:"name"`
	Href             string            `json:"href"`
	Hostname         string            `json:"hostname"`
	Image            string            `json:"image"`
	SpotInstance     bool              `json:"spot_instance"`
	BGP              ServerBGP         `json:"bgp"`
	Project          Project           `json:"project"`
	Region           Region            `json:"region"`
	State            string            `json:"state"`
	Plan             Plan              `json:"plan"`
	AvailableRegions AvailableRegions  `json:"availableregions"`
	Pricing          Pricing           `json:"pricing"`
	IPAddresses      []IPAddresses     `json:"ip_addresses"`
	SSHKeys          []SSHKeys         `json:"ssh_keys"`
	Tags             map[string]string `json:"tags"`
	Storage          BlockStorage      `json:"storage"`
	Created          string            `json:"created_at"`
	TerminationDate  string            `json:"termination_date"`
}

// SSHKeys an ssh key
type SSHKeys struct {
	ID          int    `json:"id"`
	Label       string `json:"label"`
	Key         string `json:"key"`
	Fingerprint string `json:"fingerprint"`
	Updated     string `json:"updated"`
	Created     string `json:"created"`
	Href        string `json:"href"`
}

// Cpus cpu information for a server
type Cpus struct {
	Count     int     `json:"count"`
	Name      string  `json:"name"`
	Cores     int     `json:"cores"`
	Frequency float32 `json:"frequency"`
	Unit      string  `json:"unit"`
}

// Memory cpu information for a server
type Memory struct {
	Count int    `json:"count"`
	Total int    `json:"total"`
	Unit  string `json:"unit"`
	Name  string `json:"name"`
}

// Nics network interface information for a server
type Nics struct {
	Name string `json:"name"`
}

// Raid raid for block storage on a server
type Raid struct {
	Name string `json:"name"`
}

// Storage amount of storage
type Storage struct {
	Count int     `json:"count"`
	Name  string  `json:"name"`
	Size  float32 `json:"size"`
	Unit  string  `json:"unit"`
}

// Bandwidth total bandwidth available
type Bandwidth struct {
	Name string `json:"name"`
}

// Specs aggregated specs for a server
type Specs struct {
	Cpus      Cpus      `json:"cpus"`
	Memory    Memory    `json:"memory"`
	Storage   []Storage `json:"storage"`
	Raid      Raid      `json:"raid"`
	Nics      Nics      `json:"nics"`
	Bandwidth Bandwidth `json:"bandwidth"`
}

// IPAddressCreateRequest represents a request to create a new IP address within a CreateServer request
type IPAddressCreateRequest struct {
	AddressFamily int  `json:"address_family"`
	Public        bool `json:"public"`
}

// CreateServer represents a request to create a new Cherry Servers server. Used by createNodes
type CreateServer struct {
	ProjectID       int                `json:"project_id"`
	Plan            string             `json:"plan"`
	Hostname        string             `json:"hostname"`
	Image           string             `json:"image"`
	Region          string             `json:"region"`
	SSHKeys         []int              `json:"ssh_keys"`
	IPAddresses     []string           `json:"ip_addresses"`
	UserData        string             `json:"user_data"`
	Tags            *map[string]string `json:"tags"`
	SpotInstance    int                `json:"spot_market"`
	OSPartitionSize int                `json:"os_partition_size"`
}
