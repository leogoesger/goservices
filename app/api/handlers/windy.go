package handlers

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"

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

	params := web.Params(r)

	if params["lat-long"] == "" {
		return web.NewRequestError(fmt.Errorf("invalid lat,lng format: %s", params["lat-long"]), http.StatusBadRequest)
	}
	latLng := strings.Split(params["lat-long"], ",")
	if len(latLng) != 2 {
		return web.NewRequestError(fmt.Errorf("invalid lat,lng format: %s", params["lat-long"]), http.StatusBadRequest)
	}

	hours, err := strconv.Atoi(params["hours"])
	if err != nil {
		return web.NewRequestError(fmt.Errorf("invalid hours format: %s", params["hours"]), http.StatusBadRequest)
	}

	reports, err := windy.GetReport(ctx, v.TraceID, latLng[0], latLng[1], hours)
	if err != nil {
		return errors.Wrap(err, "unable to query for users")
	}
	return web.Respond(ctx, w, reports, http.StatusOK)

}
