package config

import (
	"encoding/json"
	"os"
	"path/filepath"
	"time"
	"strconv"
)


type SleepData struct {
	Hour         string `json:"hour"`
	Minute       string `json:"minute"`
	Cycles       string `json:"cycles"`
	MaxCycles    string `json:"maxcycles"`      
    LastReset    string `json:"last_reset"` 
}


func getConfigPath() (string, error) {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return  "", err
	}

	appDir := filepath.Join(configDir, "NekoSleep")
	
	os.MkdirAll(appDir, os.ModePerm)

	return filepath.Join(appDir, "config.json"), nil
}


func Save(data *SleepData) error {
	path, err := getConfigPath()
	if err != nil {
		return err
	}

	fileData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(path, fileData, 0644)
}


func Load() (*SleepData, error) {
	path, err := getConfigPath()
	if err != nil {
		return nil, err
	}

	fileData, err := os.ReadFile(path)
	if err != nil {
		return nil, err 
	}

	var data SleepData

	err = json.Unmarshal(fileData, &data)
	if err != nil {
		return nil, err
	}
	
	now := time.Now()
    year, week := now.ISOWeek()
    
    currentWeekStr := strconv.Itoa(year) + "-W" + strconv.Itoa(week)

    if data.LastReset != currentWeekStr {
        data.Cycles = data.MaxCycles
        data.LastReset = currentWeekStr
        Save(&data)
    }

	return &data, nil
}


func Delete() error {
	path, err := getConfigPath()
	if err != nil {
		return err
	}

	return os.Remove(path)
}