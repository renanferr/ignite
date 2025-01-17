package newrelic

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	newrelic "github.com/americanas-go/ignite/newrelic/go-agent.v3"
	"github.com/americanas-go/log"
	"github.com/gofiber/fiber/v2"
	nr "github.com/newrelic/go-agent/v3/newrelic"
)

func Register(ctx context.Context, instance *fiber.App) error {

	if !IsEnabled() || !newrelic.IsEnabled() {
		return nil
	}

	logger := log.FromContext(ctx)
	logger.Trace("enabling newrelic middleware in fiber")

	instance.Use(middleware(newrelic.Application()))

	logger.Debug("newrelic middleware successfully enabled in fiber")

	return nil
}

func middleware(app *nr.Application) fiber.Handler {

	return func(c *fiber.Ctx) error {

		transactionPattern := fmt.Sprintf("%s - %s ", c.Method(), string(c.Request().URI().Path()))
		txn := app.StartTransaction(transactionPattern)
		defer txn.End()

		// TODO criar whitelist de headers
		c.Request().Header.VisitAll(func(key, value []byte) {
			txn.AddAttribute(strings.ToLower(string(key)), string(value))
		})

		wr := setNewRelicWebRequest(c)
		txn.SetWebRequest(wr)

		ctx := c.Context()
		ctx.SetUserValue(newrelic.NewRelicTransaction, txn)

		return c.Next()
	}
}

func setNewRelicWebRequest(c *fiber.Ctx) nr.WebRequest {
	header := http.Header{}

	c.Request().Header.VisitAll(func(key, value []byte) {
		header.Add(string(key), string(value))
	})

	URL := fmt.Sprintf("%s%s", c.BaseURL(), c.Path())
	parsedURL, _ := url.Parse(URL)

	wr := nr.WebRequest{
		Header:    header,
		URL:       parsedURL,
		Method:    c.Method(),
		Transport: nr.TransportType(c.Protocol()),
		Host:      string(c.Request().Host()),
	}

	return wr
}
