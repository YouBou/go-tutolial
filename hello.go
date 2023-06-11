package main

import (
	"sync"

	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

// SrData is structure
type SrData struct {
	msg string
	mux sync.Mutex
}

func main() {
	a := app.New()

	w := a.NewWindow("Hello")
	w.SetContent(
		widget.NewLabel("Hello Fyne!"),
	)

	w.ShowAndRun()
}
