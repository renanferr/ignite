package status

import (
	"github.com/americanas-go/config"
	"github.com/americanas-go/ignite/go-chi/chi.v5"
)

const (
	root    = chi.PluginsRoot + ".status"
	enabled = root + ".enabled"
	route   = root + ".route"
)

func init() {
	config.Add(enabled, true, "enable/disable status route")
	config.Add(route, "/resource-status", "define status url")
}

func IsEnabled() bool {
	return config.Bool(enabled)
}

func getRoute() string {
	return config.String(route)
}
