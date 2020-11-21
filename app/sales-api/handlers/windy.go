package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/leogoesger/goservices/business/data/windy"
	"github.com/leogoesger/goservices/foundation/web"
	"github.com/pkg/errors"
)

// GetWindy ...
func GetWindy(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	v, ok := ctx.Value(web.KeyValues).(*web.Values)
	if !ok {
		return web.NewShutdownError("web value missing from context")
	}

	var nu windy.UserInput
	if err := web.Decode(r, &nu); err != nil {
		return errors.Wrapf(err, "unable to decode payload")
	}
	fmt.Println(nu)
	windy.GetReport(ctx, v.TraceID)
	return nil
}
