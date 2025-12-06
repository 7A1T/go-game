package services

import (
	"context"

	"github.com/gdamore/tcell/v3"
)

type UIService interface {
	DrawBox(s tcell.Screen, x1, y1, x2, y2 int, style tcell.Style, text string)
}

type uiServiceKey struct{}

func UIServiceInContext(ctx context.Context, svc UIService) context.Context {
	return context.WithValue(ctx, uiServiceKey{}, svc)
}

func GetUIService(ctx context.Context) UIService {
	return ctx.Value(uiServiceKey{}).(UIService)
}
