package datadog

import (
	"github.com/americanas-go/config"
	"github.com/americanas-go/ignite/google.golang.org/grpc.v1/server"
)

const (
	root    = server.PluginsRoot + ".datadog"
	enabled = root + ".enabled"
)

func init() {
	config.Add(enabled, true, "enable/disable datadog")
}

func IsEnabled() bool {
	return config.Bool(enabled)
}
