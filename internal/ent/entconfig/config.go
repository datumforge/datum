package entconfig

// Config holds the configuration for the ent server
type Config struct {
	// Flags contains the flags for the server to allow use to test different code paths
	Flags Flags `json:"flags" koanf:"flags" jsonschema:"description=flags for the server"`
}

// Flags contains the flags for the server to allow use to test different code paths
type Flags struct {
	// UseListUserService is a flag to use the list services endpoint for object access, if false, the db is used directly instead
	UseListUserService bool `json:"useListUserService" koanf:"useListUserService" jsonschema:"description=use list services endpoint for object access" default:"true"`
	// UserListObjectService is a flag to use the list object services endpoint for object access, if false, the db is used directly instead
	UseListObjectService bool `json:"useListObjectServices" koanf:"useListObjectServices" jsonschema:"description=use list object services endpoint for object access" default:"false"`
}
