package ast

import (
	"io/fs"
	"path/filepath"
	"strings"
)

// FileFinder handles discovering eligible Go files within a target path.
type FileFinder struct{}

func NewFileFinder() *FileFinder {
	return &FileFinder{}
}

func (f *FileFinder) FindGoFiles(root string) ([]string, error) {
	var files []string

	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// Skip hidden directories and vendor folders
		if d.IsDir() {
			name := d.Name()
			if strings.HasPrefix(name, ".") || name == "vendor" {
				return filepath.SkipDir
			}
			return nil
		}

		// Collect only standard Go source files (ignoring tests)
		if strings.HasSuffix(d.Name(), ".go") && !strings.HasSuffix(d.Name(), "_test.go") {
			files = append(files, path)
		}

		return nil
	})

	return files, err
}
