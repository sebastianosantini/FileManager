package main

import (
	"filemanager/file"
)

func main() {
	archivedFiles := []file.Files{file.BuildFiles, file.DocFiles, file.ThemeFiles, file.PicFiles}
	file.ManageFiles(archivedFiles)
}
