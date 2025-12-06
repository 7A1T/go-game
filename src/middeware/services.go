package middleware

import "context"

func AddServicesToContext(ctx context.Context) context.Context {
	ctx = services.UIServiceInContext(ctx, ui.NewUIService())

	return ctx
}
