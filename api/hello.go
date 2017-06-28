package api

import (
	"github.com/get-ion/ion/context"
)

// Hello response structure
type Hello struct {
	Message string
}

// HelloHandler to be used as Handler for Hello API
func HelloHandler(ctx context.Context) {
	m := Hello{"Welcome to Cloud Native Go."}
	ctx.JSON(m)
}
