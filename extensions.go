package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type Exestension struct {
	ExeList        []string
	SourceDir      string
	DestinationDir string
}

func (e Exestension) GetFiles() ([]os.DirEntry, error) {
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

func (e Exestension) MoveFiles(files []os.DirEntry) error {
	for _, file := range files {
		filePath := e.SourceDir + file.Name()
		inputFile, err := os.Open(filePath)
		if err != nil {
			fmt.Println("There was an error while opening the file: ", err)
			return err
		}
		defer inputFile.Close()

		outputFile, err := os.Create(e.DestinationDir + file.Name())
		if err != nil {
			fmt.Println("There was an error while creating the output file:", err)
			return err
		}
		defer outputFile.Close()

		_, err = io.Copy(outputFile, inputFile)
		if err != nil {
			fmt.Println("There was an error while copying the input file onto the output file:", err)
			return err
		}

		inputFile.Close()

		err = os.Remove(filePath)
		if err != nil {
			fmt.Println("There was an error while deleting the input file:", err)
			return err
		}
	}

	return nil
}
