//go:build windows

package startup

import (
	"os"
	"path/filepath"

	"golang.org/x/sys/windows/registry"
)

// Enable прописывает NekoSleep в реестр автозапуска Windows
func Enable() error {
	exePath, err := os.Executable()
	if err != nil {
		return err
	}

	// Оборачиваем путь в кавычки и добавляем флаг невидимого запуска
	cmd := `"` + filepath.Clean(exePath) + `" -silent`

	key, err := registry.OpenKey(registry.CURRENT_USER, `Software\Microsoft\Windows\CurrentVersion\Run`, registry.SET_VALUE)
	if err != nil {
		return err
	}
	defer key.Close()

	// Сохраняем строку с флагом
	return key.SetStringValue("NekoSleep", cmd)
}

func Disable() error {
	key, err := registry.OpenKey(registry.CURRENT_USER, `Software\Microsoft\Windows\CurrentVersion\Run`, registry.SET_VALUE)
	if err != nil {
		return err // Если ключа нет или нет прав, просто возвращаем ошибку
	}
	defer key.Close()

	// Удаляем запись
	return key.DeleteValue("NekoSleep")
}