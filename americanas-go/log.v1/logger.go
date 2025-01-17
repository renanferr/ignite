package log

import (
	"github.com/americanas-go/ignite/go.uber.org/zap.v1"
	"github.com/americanas-go/ignite/rs/zerolog.v1"
	"github.com/americanas-go/ignite/sirupsen/logrus.v1"
)

func New() {
	switch Impl() {
	case "ZEROLOG":
		zerolog.NewLogger()
	case "ZAP":
		zap.NewLogger()
	default:
		logrus.NewLogger()
	}
}
