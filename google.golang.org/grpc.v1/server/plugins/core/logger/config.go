package logger

import (
	"github.com/americanas-go/config"
	"github.com/americanas-go/ignite/google.golang.org/grpc.v1/server"
)

const (
	root    = server.PluginsRoot + ".logger"
	enabled = root + ".enabled"
	level   = root + ".level"
)

func init() {
	config.Add(enabled, true, "enable/disable logger")
	config.Add(level, "INFO", "sets log level INFO/DEBUG/TRACE")
}

func IsEnabled() bool {
	return config.Bool(enabled)
}

func Level() string {
	return config.String(level)
}
