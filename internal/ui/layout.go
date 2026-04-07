package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// buildMainLayout собирает весь пользовательский интерфейс.
// Функция с маленькой буквы, так как используется только внутри пакета ui (в app.go).
func buildMainLayout() fyne.CanvasObject {
	
	// Создаем виджеты
	title := widget.NewLabel("Добро пожаловать в NekoSleep!")
	
	// Используем хелпер theme.HomeIcon() для вызова иконки, которую мы подменили в theme.go
	btn := widget.NewButtonWithIcon("Кнопка с котенком", theme.HomeIcon(), func() {
		// Здесь будет логика нажатия на кнопку
	})

	// Собираем их в вертикальный контейнер
	content := container.NewVBox(
		title,
		btn,
	)

	return content
}