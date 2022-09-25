package main

import (

	//"github.com/tarm/serial"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Serial-port")
	w.Resize(fyne.NewSize(400, 320))
	entry := widget.NewEntry()
	answer := widget.NewLabel("A:")
	descr := widget.NewLabel("B:")
	contain := container.NewHSplit(answer, descr)
	w.SetContent(container.NewHSplit(
		entry,
		contain,
	))
	w.ShowAndRun()
}
