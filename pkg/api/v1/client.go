package api

import (
	"github.com/datumforge/datum/pkg/datumclient"
)

type DatumClient interface {
	datumclient.DatumClient
	datumclient.DatumGraphClient
}
