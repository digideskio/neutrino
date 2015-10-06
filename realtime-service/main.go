package main

import (
	"github.com/go-neutrino/neutrino-config"
	"github.com/go-neutrino/neutrino-core/log"
	"github.com/go-neutrino/neutrino-core/realtime-service/server"
	"net/http"
)

func main() {
	c := nconfig.Load()

	server.Initialize(c)

	port := c.GetString(nconfig.KEY_REALTIME_PORT)
	log.Info("Listening on port: " + port)
	log.Info(http.ListenAndServe(port, nil))
}
