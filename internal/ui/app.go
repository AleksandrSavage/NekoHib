package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)


type App struct {
	fyneApp fyne.App
	window  fyne.Window
}

func NewApp(font fyne.Resource, fontBold fyne.Resource, icon fyne.Resource, kitten_greet fyne.Resource) *App {
	a := app.New()

	a.Settings().SetTheme(NewCustomTheme(font, fontBold, icon))

	w := a.NewWindow("NekoSleep")
	w.Resize(fyne.NewSize(640, 400))
	w.SetFixedSize(true)

	content := buildMainLayout(kitten_greet, w)
	w.SetContent(content)

	return &App{
		fyneApp: a,
		window:  w,
	}
}

func (a *App) Run() {
	a.window.ShowAndRun()
}