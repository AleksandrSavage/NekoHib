package main

import (
	// Замени NekoSleep на имя твоего модуля из go.mod
	"NekoSleep/internal/ui"
)

func main() {
	// Создаем экземпляр нашего приложения, передавая ресурсы из bundled.go
	// Названия переменных проверь в bundled.go!
	App := ui.NewApp(resourceNunitoRegularTtf, resourceKitteniconIco)

	// Запускаем
	App.Run()
}