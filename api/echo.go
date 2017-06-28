package api

import (
	"github.com/get-ion/ion/context"
)

// EchoHandler to be used as Handler for ECHO API
func EchoHandler(ctx context.Context) {
	message := ctx.URLParam("message")
	ctx.Text(message)
}
