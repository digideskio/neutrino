package main

import (
	"github.com/go-neutrino/neutrino-core/log"
	"github.com/go-neutrino/neutrino-core/realtime-service/server"
	"net/http"
	"github.com/go-neutrino/neutrino-core/config"
)

func main() {
	server.Initialize()

	port := config.Get(config.KEY_REALTIME_PORT)
	log.Info("Listening on port: " + port)
	log.Info(http.ListenAndServe(port, nil))
}
