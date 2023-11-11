package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	application := app.New()
	mainWindow := application.NewWindow("window on GOLang")
	mainWindow.SetContent(
		widget.NewLabel("asda"),
	)
	mainWindow.ShowAndRun()
}
