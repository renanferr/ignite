package main

import (
	"context"

	"github.com/americanas-go/config"
	"github.com/americanas-go/ignite/go.mongodb.org/mongo-driver.v1"
	newrelic "github.com/americanas-go/ignite/go.mongodb.org/mongo-driver.v1/plugins/contrib/newrelic/go-agent.v3"
	"github.com/americanas-go/ignite/sirupsen/logrus.v1"
)

func main() {

	config.Load()
	logrus.NewLogger()

	mongo.NewConn(context.Background(), newrelic.Register)
}
