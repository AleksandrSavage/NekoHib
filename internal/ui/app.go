package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/driver/desktop"

	"NekoSleep/internal/monitor"
)


type App struct {
	fyneApp fyne.App
	window  fyne.Window
}

func NewApp(font fyne.Resource, fontBold fyne.Resource, icon fyne.Resource, kitten_greet fyne.Resource) *App {
	a := app.NewWithID("dev.as.nekosleep")

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

func (a *App) Show() {
    a.window.Show()
}

func (a *App) Run() {
    // Настройка трея (остается как была)
    if desk, ok := a.fyneApp.(desktop.App); ok {
        m := fyne.NewMenu("NekoSleep",
            fyne.NewMenuItem("Show NekoSleep", func() {
                a.window.Show() 
            }),
            fyne.NewMenuItem("Quit", func() {
                monitor.Stop()  
                a.fyneApp.Quit() 
            }),
        )
        desk.SetSystemTrayMenu(m)
    }

    // Перехват закрытия (скрываем вместо выхода)
    a.window.SetCloseIntercept(func() {
        a.window.Hide() 
    })

    // ВАЖНО: Вместо ShowAndRun пишем просто Run
    // Это запустит приложение, но НЕ покажет окно само по себе.
    a.fyneApp.Run()
}