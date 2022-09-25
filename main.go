package main

import (

	//"github.com/tarm/serial"

	"strings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"go.bug.st/serial"
)

var descr *widget.Label

func catchErr() {
	str := "ERROR [Connection is not established]\n\nInput any key to exit"
	descr.SetText(str)

}

func workCOM1(port serial.Port, mode *serial.Mode, str string) {
	_, err := port.Write([]byte(str))
	if err != nil {
		catchErr()
	}
}

func workCOM2(port serial.Port, mode *serial.Mode) string {

	buff := make([]byte, 100)
	n, err := port.Read(buff)
	if err != nil {
		catchErr()
	}
	return string(buff[:n])
}

func main() {
	//Open the first serial port detected at 9600bps N81
	a := app.New()
	w := a.NewWindow("Serial-port")
	w.Resize(fyne.NewSize(600, 600))
	entry := widget.NewMultiLineEntry()
	entry.Wrapping = fyne.TextWrapBreak
	answer := widget.NewLabel("")
	answer.Wrapping = fyne.TextWrapBreak
	descr = widget.NewLabel("The connection is established")
	contain := container.NewGridWithColumns(3, entry, answer, descr)
	w.SetContent(contain)

	mode := &serial.Mode{
		BaudRate: 9600,
		Parity:   serial.NoParity,
		DataBits: 8,
		StopBits: serial.OneStopBit,
	}
	port1, err := serial.Open("COM1", mode)
	if err != nil {
		catchErr()
	}
	port2, err := serial.Open("COM2", mode)
	if err != nil {
		catchErr()
	}

	go func() {
		var strMain string
		for range time.Tick(time.Second / 8) {
			str := entry.Text
			if len(str) == len(strMain) {
				continue
			}
			if !strings.HasPrefix(str, strMain) {
				entry.SetText(strMain)
			} else {
				_, chars, _ := strings.Cut(str, strMain)
				workCOM1(port1, mode, chars)
				newChars := workCOM2(port2, mode)
				//fmt.Println(newChars, "!!!")
				strMain += newChars
				answer.SetText(strMain)
			}
		}
	}()
	w.ShowAndRun()
}
