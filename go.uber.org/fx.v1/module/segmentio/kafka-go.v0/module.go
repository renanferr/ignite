package nats

import (
	"sync"

	contextfx "github.com/americanas-go/ignite/go.uber.org/fx.v1/module/context"
	"github.com/americanas-go/ignite/segmentio/kafka-go.v0"
	"go.uber.org/fx"
)

var leaderOnce sync.Once

func LeaderModule() fx.Option {
	options := fx.Options()

	leaderOnce.Do(func() {
		options = fx.Options(
			contextfx.Module(),
			fx.Provide(
				kafka.NewConn,
			),
		)
	})

	return options
}
