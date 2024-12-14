package file

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type Files struct {
	ExeList        []string
	SourceDir      string
	DestinationDir string
}

var (
	userHomeDir, _ = os.UserHomeDir()
	downloadsDir   = userHomeDir + "\\Downloads\\"
	buildDir       = userHomeDir + "\\Documents\\BuildFiles\\"
	docDir         = userHomeDir + "\\Documents\\DocFiles\\"
	themeDir       = userHomeDir + "\\Documents\\ThemeFiles\\"
	picDir         = userHomeDir + "\\Documents\\PicFiles\\"

	BuildFiles = Files{
		[]string{".exe", ".zip", ".msi", ".iso"},
		downloadsDir,
		buildDir,
	}

	DocFiles = Files{
		[]string{".pdf"},
		downloadsDir,
		docDir,
	}
	ThemeFiles = Files{
		[]string{".vsix"},
		downloadsDir,
		themeDir,
	}
	PicFiles = Files{
		[]string{".jpeg", ".jpg", ".gif", ".png"},
		downloadsDir,
		picDir,
	}
)

func (e Files) getFiles() ([]os.DirEntry, error) {
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

func (e Files) moveFiles(files []os.DirEntry) error {
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

func ManageFiles(files []Files) {
	for _, archivedFile := range files {
		movedFiles, err := archivedFile.getFiles()
		if err != nil {
			log.Fatal(err)
		}

		if len(movedFiles) == 0 {
			continue
		}

		err = archivedFile.moveFiles(movedFiles)
		if err != nil {
			log.Fatal(err)
		}
	}
}
