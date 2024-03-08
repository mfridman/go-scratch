package foo

import (
	"context"

	"github.com/bufbuild/protoplugin"
)

// Handle implements the (protoplugin.Handler)[https://pkg.go.dev/github.com/bufbuild/protoplugin#Handler] interface and
// is the entry point for the plugin.
func Handle(
	ctx context.Context,
	responseWriter *protoplugin.ResponseWriter,
	request *protoplugin.Request,
) error {
	return nil
}
