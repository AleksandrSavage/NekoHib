package main

import (
    "os"

	"NekoSleep/internal/ui"
	"NekoSleep/internal/monitor"
	"NekoSleep/internal/config"
	"NekoSleep/internal/startup"
)


func main() {
    App := ui.NewApp(
        resourceNunitoRegularTtf,
        resourceNunitoBoldTtf,
        resourceKitteniconIco,
        resourceKittengreetPng,
    )
    
    monitor.Init(resourceKittenasleepPng)

    isSilentRun := false
    if len(os.Args) > 1 && os.Args[1] == "-silent" {
        isSilentRun = true
    }

    _, err := config.Load()
    if err == nil {
        monitor.Start()
        
        if !isSilentRun {
            App.Show()
        }
    } else {
        startup.Disable()

        App.Show()
    }
    
    App.Run()
}