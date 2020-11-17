// Package handlers contains the full set of handler functions and routes
// supported by the web api.
package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/leogoesger/goservices/business/mid"
	"github.com/leogoesger/goservices/foundation/web"
)

// API constructs an http.Handler with all application routes defined.
func API(build string, shutdown chan os.Signal, log *log.Logger) *web.App {
	app := web.NewApp(shutdown, mid.Logger(log), mid.Errors(log))

	app.Handle(http.MethodGet, "/readiness", readiness)

	return app
}