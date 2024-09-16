package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
)

func prepareLaunchPath(basePath, version string) string {
	defaultPath := filepath.Join(basePath, version)

	if !isFolderInUse(defaultPath) {
		return defaultPath
	}

	for i := 1; ; i++ {
		tempFolderName := fmt.Sprintf("%s_%d", version, i)
		tempPath := filepath.Join(basePath, tempFolderName)

		if _, err := os.Stat(tempPath); err == nil {
			if !isFolderInUse(tempPath) {
				return tempPath
			}
		} else {
			if err := os.MkdirAll(tempPath, os.ModePerm); err != nil {
				fmt.Println("Error creating temporary folder:", err)
				return defaultPath
			}
			if err := copyFolder(defaultPath, tempPath); err != nil {
				fmt.Println("Error copying folder for new instance:", err)
				return defaultPath
			}

			return tempPath
		}
	}
}

func createLockFile(path string, pid int) {
	file, err := os.Create(path)
	if err != nil {
		fmt.Println("Error creating lock file:", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(strconv.Itoa(pid))
	if err != nil {
		fmt.Println("Error writing to lock file:", err)
	}
}

func isFolderInUse(folderPath string) bool {
	dcrFiles := []string{"habbo.dcr", "habbo-xl.dcr"}

	for _, dcrFile := range dcrFiles {
		habboFilePath := filepath.Join(folderPath, dcrFile)
		file, err := os.OpenFile(habboFilePath, os.O_RDWR, 0666)
		if err != nil {
			if os.IsNotExist(err) {
				continue
			}
			return true
		}
		file.Close()
	}

	return false
}

func launchApplication(path, habboExe string) {
	exePath := filepath.Join(path, habboExe)
	cmd := exec.Command(exePath)

	err := cmd.Start()
	if err != nil {
		fmt.Println("Error launching application:", err)
		return
	}

	lockFilePath := filepath.Join(path, "instance.lock")
	createLockFile(lockFilePath, cmd.Process.Pid)
}

func copyFolder(src, dest string) error {
	return filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		relPath, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}

		destPath := filepath.Join(dest, relPath)

		if info.IsDir() {
			return os.MkdirAll(destPath, info.Mode())
		}

		inputFile, err := os.Open(path)
		if err != nil {
			return err
		}
		defer inputFile.Close()

		outputFile, err := os.Create(destPath)
		if err != nil {
			return err
		}
		defer outputFile.Close()

		_, err = io.Copy(outputFile, inputFile)
		return err
	})
}
