package main

import (
	"context"
	"fmt"
	"log"

	middleware "game/src/middeware"

	"github.com/gdamore/tcell/v3"
	"github.com/tait/go-game/src/middeware/services"
)

func main() {
	ctx := context.Background()
	ctx = middleware.AddServicesToContext(ctx)

	uiService := services.GetUIService(ctx)

	defStyle := tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)
	boxStyle := tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorPurple)

	// Initialize screen
	s, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("%+v", err)
	}
	if err := s.Init(); err != nil {
		log.Fatalf("%+v", err)
	}
	s.SetStyle(defStyle)
	s.EnableMouse()
	s.EnablePaste()
	s.Clear()

	// Draw initial boxes
	uiService.DrawBox(s, 1, 1, 42, 7, boxStyle, "Click and drag to draw a box")
	uiService.DrawBox(s, 5, 9, 32, 14, boxStyle, "Press C to reset")

	quit := func() {
		// You have to catch panics in a defer, clean up, and
		// re-raise them - otherwise your application can
		// die without leaving any diagnostic trace.
		maybePanic := recover()
		s.Fini()
		if maybePanic != nil {
			panic(maybePanic)
		}
	}
	defer quit()

	// Here's how to get the screen size when you need it.
	// xmax, ymax := s.Size()

	// Here's an example of how to inject a keystroke where it will
	// be picked up by a future read of the event queue.  Note that
	// care should be used to avoid blocking writes to the queue if
	// this is done from the same thread that is responsible for reading
	// the queue, or else a single-party deadlock might occur.
	// s.EventQ() <- tcell.NewEventKey(tcell.KeyRune, rune('a'), 0)

	// Event loop
	ox, oy := -1, -1
	for {
		// Update screen
		s.Show()

		// Poll event (this can be in a select statement as well)
		ev := <-s.EventQ()

		// Process event
		switch ev := ev.(type) {
		case *tcell.EventResize:
			s.Sync()
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyEscape || ev.Key() == tcell.KeyCtrlC {
				return
			} else if ev.Key() == tcell.KeyCtrlL {
				s.Sync()
			} else if ev.Str() == "C" || ev.Str() == "c" {
				s.Clear()
			}
		case *tcell.EventMouse:
			x, y := ev.Position()

			switch ev.Buttons() {
			case tcell.Button1, tcell.Button2:
				if ox < 0 {
					ox, oy = x, y // record location when click started
				}

			case tcell.ButtonNone:
				if ox >= 0 {
					label := fmt.Sprintf("%d,%d to %d,%d", ox, oy, x, y)
					uiService.DrawBox(s, ox, oy, x, y, boxStyle, label)
					ox, oy = -1, -1
				}
			}
		}
	}
}
