package stripslashes

import (
	"github.com/americanas-go/config"
	"github.com/americanas-go/ignite/go-chi/chi.v5"
)

const (
	root    = chi.PluginsRoot + ".stripslashes"
	enabled = root + ".enabled"
)

func init() {
	config.Add(enabled, true, "enable/disable stripSlashes middleware")
}

func IsEnabled() bool {
	return config.Bool(enabled)
}
