package ast

import (
	"go/parser"
	"go/token"
	"path/filepath"

	"github.com/JoaoVitor615/archscan/internal/model"
)

// ProjectParser orchestrates file reading, parsing, and AST inspection.
type ProjectParser struct {
	finder    *FileFinder
	inspector *FileInspector
	fset      *token.FileSet
}

// NewProjectParser creates a new ProjectParser instance.
func NewProjectParser() *ProjectParser {
	return &ProjectParser{
		finder:    NewFileFinder(),
		inspector: NewFileInspector(),
		fset:      token.NewFileSet(),
	}
}

// ParseProject scans the root directory and returns a map of package paths to PackageModels.
func (p *ProjectParser) ParseProject(root string) (map[string]*model.PackageModel, error) {
	files, err := p.finder.FindGoFiles(root)
	if err != nil {
		return nil, err
	}

	packages := make(map[string]*model.PackageModel)

	for _, file := range files {
		dir := filepath.Dir(file)

		node, err := parser.ParseFile(p.fset, file, nil, parser.ParseComments)
		if err != nil {
			return nil, err
		}

		pkgName := node.Name.Name

		if _, exists := packages[dir]; !exists {
			packages[dir] = &model.PackageModel{
				Name:      pkgName,
				Path:      dir,
				Structs:   make(map[string]*model.StructModel),
				Functions: make([]*model.FunctionModel, 0),
			}
		}

		p.inspector.InspectFile(node, packages[dir])
	}

	return packages, nil
}
