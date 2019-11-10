package main

import (
	"log"

	term "github.com/gdamore/tcell/termbox"
)

func main() {
	if err := term.Init(); err != nil {
		log.Fatalf("term.Init(): %v", err)
	}
	term.SetInputMode(term.InputEsc)

	defer term.Close()

	// Fill the
	termH, termW := term.Size()
	for i := 0; i < termH; i++ {
		for j := 0; j < termW; j++ {
			term.SetCell(i, j, '█', term.ColorBlue, term.AttrBold)
		}
	}

	// After we change colors we need to Flush
	// to actually draw something.
	term.Flush()

	// Cursor position for drawing.
	var h, w int
	for {
		ev := term.PollEvent()
		switch ev.Type {
		case term.EventError:
			log.Fatal(ev.Err)
		case term.EventKey:
		default:
			continue
		}

		switch ev.Key {
		case term.KeyEsc:
			// Exit on espace.
			return
		case term.KeyArrowDown:
			h += 2
		case term.KeyArrowLeft:
			w -= 2
		case term.KeyArrowRight:
			w += 2
		case term.KeyArrowUp:
			h -= 2
		default:
			continue
		}
		term.SetCell(w, h, 'x', term.ColorBlack, 0)
		term.Flush()
	}
}
