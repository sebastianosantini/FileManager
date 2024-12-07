package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type Files struct {
	ExeList        []string
	SourceDir      string
	DestinationDir string
}

var (
	userHomeDir, err = os.UserHomeDir()
	downloadsDir     = userHomeDir + "\\Downloads\\"
	buildDir         = userHomeDir + "\\Documents\\BuildFiles\\"
	docDir           = userHomeDir + "\\Documents\\DocFiles\\"
	themeDir         = userHomeDir + "\\Documents\\ThemeFiles\\"
	picDir           = userHomeDir + "\\Documents\\PicFiles\\"

	buildFiles = Files{
		[]string{".exe", ".zip", ".msi", ".iso"},
		downloadsDir,
		buildDir,
	}

	docFiles = Files{
		[]string{".pdf"},
		downloadsDir,
		docDir,
	}
	themeFiles = Files{
		[]string{".vsix"},
		downloadsDir,
		themeDir,
	}
	picFiles = Files{
		[]string{".jpeg", ".jpg", ".gif", ".png"},
		downloadsDir,
		picDir,
	}
)

func (e Files) GetFiles() ([]os.DirEntry, error) {
	movedFiles := []os.DirEntry{}
	entries, err := os.ReadDir(e.SourceDir)
	if err != nil {
		fmt.Println("There was an error while reading the sourceDir:", err)
		return nil, err
	}

	for _, entry := range entries {
		for _, exe := range e.ExeList {
			if strings.Contains(entry.Name(), exe) {
				movedFiles = append(movedFiles, entry)
			}
		}
	}

	return movedFiles, nil
}

func (e Files) MoveFiles(files []os.DirEntry) error {
	for _, file := range files {
		filePath := e.SourceDir + file.Name()

		inputFile, err := os.Open(filePath)
		if err != nil {
			return err
		}
		defer inputFile.Close()

		_, err = os.Stat(e.DestinationDir)
		if os.IsNotExist(err) {
			if err := os.Mkdir(e.DestinationDir, os.ModePerm); err != nil {
				return err
			}
		}

		outputFile, err := os.Create(e.DestinationDir + file.Name())
		if err != nil {
			return err
		}
		defer outputFile.Close()

		_, err = io.Copy(outputFile, inputFile)
		if err != nil {
			return err
		}

		inputFile.Close()

		err = os.Remove(filePath)
		if err != nil {
			return err
		}
	}

	return nil
}