package server

import (
	"log"
	"net/http"

	"github.com/datumforge/datum/internal/httpserve/config"
	"github.com/datumforge/echox"
)

type Server struct {
	config config.Server
}

func serve() {
	e := echox.New()
	// add middleware and routes
	// ...
	s := http.Server{
		Addr:    ":8080",
		Handler: e,
		//ReadTimeout: 30 * time.Second, // customize http.Server timeouts
	}
	if err := s.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
