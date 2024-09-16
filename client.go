package main

import (
	"archive/zip"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

const (
	clientDataURL  = "https://origins.habbo.com/gamedata/clienturls"
	updateFileName = "temp_update.zip"
)

type ClientData struct {
	ShockwaveWindowsVersion string `json:"shockwave-windows-version"`
	ShockwaveWindows        string `json:"shockwave-windows"`
}

func fetchClientData() (*ClientData, error) {
	resp, err := http.Get(clientDataURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data ClientData
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func downloadUpdate(url, version, path string) error {
	tempUpdatePath := filepath.Join(path, updateFileName)

	out, err := os.Create(tempUpdatePath)
	if err != nil {
		return err
	}

	resp, err := http.Get(url)
	if err != nil {
		out.Close()
		return err
	}
	defer resp.Body.Close()

	_, err = io.Copy(out, resp.Body)
	out.Close()
	if err != nil {
		return err
	}

	err = extractZip(tempUpdatePath, filepath.Join(path, version))
	if err != nil {
		return err
	}

	err = os.Remove(tempUpdatePath)
	if err != nil {
		fmt.Println("Error deleting temp update file:", err)
		return err
	}

	return nil
}

func extractZip(src, dest string) error {
	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer r.Close()

	for _, f := range r.File {
		fpath := filepath.Join(dest, f.Name)
		if f.FileInfo().IsDir() {
			os.MkdirAll(fpath, os.ModePerm)
			continue
		}

		if err := os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
			return err
		}

		outFile, err := os.Create(fpath)
		if err != nil {
			return err
		}
		defer outFile.Close()

		rc, err := f.Open()
		if err != nil {
			return err
		}

		_, err = io.Copy(outFile, rc)
		rc.Close()
		if err != nil {
			return err
		}
	}

	return nil
}

func deleteOldFolders(basePath, previousVersion string, deleteMain bool) {
	previousVer, err := strconv.Atoi(previousVersion)
	if err != nil {
		fmt.Println("Error converting previous version to integer:", err)
		return
	}

	highestInstalledVersion := 0
	folders, err := os.ReadDir(basePath)
	if err != nil {
		fmt.Println("Error reading base path:", err)
		return
	}

	for _, folder := range folders {
		if folder.IsDir() && isNumeric(folder.Name()) {
			folderVersion, _ := strconv.Atoi(folder.Name())
			if folderVersion > highestInstalledVersion && folderVersion <= previousVer {
				highestInstalledVersion = folderVersion
			}
		}
	}

	for _, folder := range folders {
		folderName := folder.Name()
		folderPath := filepath.Join(basePath, folderName)

		if len(folderName) > 2 && folderName[len(folderName)-2:] == "_1" || len(folderName) > 2 && folderName[len(folderName)-2:] == "_2" {
			fmt.Printf("Deleting temporary folder: %s\n", folderPath)
			err := os.RemoveAll(folderPath)
			if err != nil {
				fmt.Printf("Error deleting folder %s: %v\n", folderPath, err)
			}
		} else if deleteMain && isNumeric(folderName) {
			folderVersion, _ := strconv.Atoi(folderName)
			if folderVersion == highestInstalledVersion {
				fmt.Printf("Deleting main version folder: %s\n", folderPath)
				err := os.RemoveAll(folderPath)
				if err != nil {
					fmt.Printf("Error deleting folder %s: %v\n", folderPath, err)
				}
			}
		}
	}
}

// Helper function to check if a string is numeric
func isNumeric(str string) bool {
	_, err := strconv.Atoi(str)
	return err == nil
}
