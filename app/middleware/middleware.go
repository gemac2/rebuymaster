// middleware package is intended to host the middlewares used
// across the app.
package middleware

import (
	"rebuymaster/app/models"

	csrf "github.com/gobuffalo/mw-csrf"
	paramlogger "github.com/gobuffalo/mw-paramlogger"
	"github.com/wawandco/ox/pkg/buffalotools"
)

var (
	// RequestID pulls the request id from the request and
	// adds it if its not present.
	RequestID = buffalotools.NewRequestIDMiddleware("RequestID")

	// Database middleware adds a `tx` context variable
	// to every request, this tx variates to be a plain connection
	// or a transaction based on the type of request.
	Database = buffalotools.DatabaseMiddleware(models.DB(), nil)

	// ParameterLogger logs out parameters that the app received
	// taking care of sensitive data.
	ParameterLogger = paramlogger.ParameterLogger

	// CSRF middleware protects from CSRF attacks.
	CSRF = csrf.New
)
