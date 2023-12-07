package entdb

import (
	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	"github.com/datumforge/datum/internal/utils/viperconfig"
)

const (
	defaultDBPrimaryURI   = "datum.db?mode=memory&_fk=1"
	defaultDBSecondaryURI = "backup.db?mode=memory&_fk=1"
)

// RegisterDatabaseFlags registers the flags for the database configuration
func RegisterDatabaseFlags(v *viper.Viper, flags *pflag.FlagSet) error {
	err := viperconfig.BindConfigFlag(v, flags, "db.multi-write", "db-multi-write", false, "write to a primary and secondary database", flags.Bool)
	if err != nil {
		return err
	}

	err = viperconfig.BindConfigFlag(v, flags, "db.primary", "db-primary", defaultDBPrimaryURI, "db primary uri", flags.String)
	if err != nil {
		return err
	}

	err = viperconfig.BindConfigFlag(v, flags, "db.secondary", "db-secondary", defaultDBSecondaryURI, "db secondary uri", flags.String)
	if err != nil {
		return err
	}

	return nil
}
