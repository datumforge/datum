package server

import (
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/datumforge/echox"

	"github.com/datumforge/datum/internal/httpserve/config"
)

type Server struct {
	config config.Server //nolint: unused
}

func serve() { //nolint: unused
	e := echox.New()
	// add middleware and routes
	// ...
	s := http.Server{
		Addr:    ":8080",
		Handler: e,
		// TODO: customize http.Server timeouts
		ReadTimeout: 30 * time.Second, //nolint:gomnd
	}

	if err := s.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		log.Fatal(err)
	}
}
