package main

import (
	"fmt"
	"os"
)

func main() {
	config := loadConfig()

	habboExe := getHabboExe(config)

	pathToUse := getLauncherPath(config.Path)

	if _, err := os.Stat(pathToUse); os.IsNotExist(err) {
		err := os.MkdirAll(pathToUse, os.ModePerm)
		if err != nil {
			fmt.Println("Error creating directory:", err)
			return
		}
	}

	highestVersion, err := getHighestFolderNumber(pathToUse)
	if err != nil {
		fmt.Println("Error finding highest version folder:", err)
		return
	}

	clientData, err := fetchClientData()
	if err != nil {
		fmt.Println("Error fetching client data:", err)
		return
	}

	if clientData.ShockwaveWindowsVersion != highestVersion {
		fmt.Println("New version found, downloading update...")
		err = downloadUpdate(clientData.ShockwaveWindows, clientData.ShockwaveWindowsVersion, pathToUse)
		if err != nil {
			fmt.Println("Error downloading update:", err)
			return
		}

		deleteOldFolders(pathToUse, highestVersion, config.DeleteOldVersion)

		highestVersion = clientData.ShockwaveWindowsVersion
	} else {
		fmt.Println("No new updates found.")
	}

	launchPath := prepareLaunchPath(pathToUse, highestVersion)

	launchApplication(launchPath, habboExe)
	fmt.Println("launching: " + launchPath + "\\" + habboExe)
}
