package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main_() {
	a := app.New()
	w := a.NewWindow("Mi ventana")

	//Resize la ventana
	w.Resize(fyne.NewSize(400, 400))
	var valor int = 0
	btn := widget.NewButton("Mi boton", func() {
		valor = valor + 1
		fmt.Println("Me apretaron", valor, "veces")
	})

	w.SetContent((btn))

	//w.SetContent(widget.NewLabel("Hello Sofi, esta rico el helado?"))

	w.ShowAndRun()

}
