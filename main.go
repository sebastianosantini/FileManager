package main

import (
	"log"
)

func main() {
	archivedFiles := []Exestension{buildExe, docExe, themeExe, picExe}
	for _, archivedFile := range archivedFiles {
		movedFiles, err := archivedFile.GetFiles()
		if err != nil {
			log.Fatal(err)
		}

		if len(movedFiles) == 0 {
			continue
		}

		err = archivedFile.MoveFiles(movedFiles)
		if err != nil {
			log.Fatal(err)
		}
	}
}
