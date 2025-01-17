package requestid

import (
	"context"

	"github.com/americanas-go/log"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

func Register(ctx context.Context, client *resty.Client) error {

	if !IsEnabled() {
		return nil
	}

	logger := log.FromContext(ctx)
	logger.Trace("enabling requestID middleware in resty")

	client.OnBeforeRequest(requestId)

	logger.Debug("requestID middleware successfully enabled in resty")

	return nil
}

func requestId(client *resty.Client, request *resty.Request) error {

	ctx := request.Context()

	idValue, ok := ctx.Value("requestId").(string)
	if !ok {
		id, err := uuid.NewUUID()
		if err != nil {
			return err
		}
		idValue = id.String()
	}

	request.SetHeader("X-Request-ID", idValue)

	return nil
}
