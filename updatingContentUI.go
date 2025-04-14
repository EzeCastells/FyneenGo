package main

import (
	"time"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func updateTime(r *widget.Label) {
	formatted := time.Now().Format("Hora 03:04:05")
	r.SetText(formatted)
}

func main() {
	a := app.New()
	w := a.NewWindow("Reloj")

	reloj := widget.NewLabel("")

	updateTime(reloj)

	w.SetContent(reloj)

	go func() {
		for range time.Tick(time.Second) {
			updateTime(reloj)
			print("Actualizo")
		}
	}()

	w.ShowAndRun()

}
