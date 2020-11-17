package handlers

import (
	"context"
	"errors"
	"math/rand"
	"net/http"

	"github.com/leogoesger/goservices/foundation/web"
)

func readiness(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	if n := rand.Intn(100); n%2 == 0 {
		err := errors.New("bill an error")
		// return err
		return web.NewRequestError(err, http.StatusBadRequest)
	}

	status := struct {
		Status string
	}{
		Status: "OK",
	}
	return web.Respond(ctx, w, status, http.StatusOK)
}