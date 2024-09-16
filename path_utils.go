package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

const defaultLauncherSubPath = "Habbo Launcher\\downloads\\shockwave"

func getLauncherPath(customPath string) string {
	if customPath != "" {
		absPath, err := filepath.Abs(customPath)
		if err == nil {
			return absPath
		}
	}

	appDataPath, err := os.UserConfigDir()
	if err != nil {
		fmt.Println("Error getting user config directory, using current directory.")
		return "."
	}

	return filepath.Join(appDataPath, defaultLauncherSubPath)
}

func getHighestFolderNumber(path string) (string, error) {
	dirs, err := os.ReadDir(path)
	if err != nil {
		return "", err
	}

	var highest int
	for _, dir := range dirs {
		if dir.IsDir() {
			num, err := strconv.Atoi(dir.Name())
			if err == nil && num > highest {
				highest = num
			}
		}
	}

	return strconv.Itoa(highest), nil
}
