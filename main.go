package main

import (
	"fmt"
	"os"
)

func main() {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	downloadsDir := userHomeDir + "\\Downloads\\"
	buildDir := userHomeDir + "\\Documents\\BuildFiles\\"
	docDir := userHomeDir + "\\Documents\\DocFiles\\"
	themeDir := userHomeDir + "\\Documents\\ThemeFiles\\"
	picDir := userHomeDir + "\\Documents\\PicFiles\\"

	buildExe := Exestension{
		[]string{".exe", ".zip", ".msi", ".iso"},
		downloadsDir,
		buildDir,
	}
	docExe := Exestension{
		[]string{".pdf"},
		downloadsDir,
		docDir,
	}
	themeExe := Exestension{
		[]string{".vsix"},
		downloadsDir,
		themeDir,
	}
	picExe := Exestension{
		[]string{".jpeg", ".jpg", ".gif", ".png"},
		downloadsDir,
		picDir,
	}

	archivedFiles := []Exestension{buildExe, docExe, themeExe, picExe}
	for _, archivedFile := range archivedFiles {
		movedFiles, err := archivedFile.GetFiles()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		if len(movedFiles) == 0 {
			continue
		}

		err = archivedFile.MoveFiles(movedFiles)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}
