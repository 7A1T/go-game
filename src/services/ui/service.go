package ui

import "github.com/gdamore/tcell/v3"

type UIService struct{}

func NewUIService() *UIService {
	return &UIService{}
}

func (u *UIService) DrawBox(s tcell.Screen, x1, y1, x2, y2 int, style tcell.Style, text string) {
	drawBox(s, x1, y1, x2, y2, style, text)
}
