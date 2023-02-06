package filewalker

import (
	"os"
	"path/filepath"
)

func DirectoryWalk(dir string) []string {
	var files []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() && info.Name() != "DELETIONS.txt" {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		panic(err)
	} else {
		return files
	}
}
