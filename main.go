package main

import (
	"log"

	"github.com/tarm/serial"
	// "fyne.io/fyne/v2"
	// "fyne.io/fyne/v2/app"
	// "fyne.io/fyne/v2/container"
	// "fyne.io/fyne/v2/widget"
)

func main() {
	c := &serial.Config{Name: "COM1", Baud: 115200}
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}

	n, err := s.Write([]byte("test"))
	if err != nil {
		log.Fatal(err)
	}

	buf := make([]byte, 128)
	n, err = s.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	log.Print("%q", buf[:n])
	// a := app.New()
	// w := a.NewWindow("Serial-port")
	// w.Resize(fyne.NewSize(400, 320))
	// entry := widget.NewEntry()
	// answer := widget.NewLabel("A:")
	// descr := widget.NewLabel("B:")
	// contain := container.NewHSplit(entry, answer)
	// w.SetContent(container.NewHSplit(
	// 	contain,
	// 	descr,
	// ))
	// w.ShowAndRun()
}
