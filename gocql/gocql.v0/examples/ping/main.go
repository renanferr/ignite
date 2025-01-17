package main

import (
	"context"

	"github.com/americanas-go/config"
	"github.com/americanas-go/ignite/gocql/gocql.v0"
	"github.com/americanas-go/ignite/sirupsen/logrus.v1"
)

func main() {

	config.Load()

	logrus.NewLogger()

	session, err := gocql.NewSession(context.Background())
	if err != nil {
		panic(err)
	}

	defer session.Close()

	err = session.Query("void").Exec()
	if err != nil {
		panic(err)
	}

}
