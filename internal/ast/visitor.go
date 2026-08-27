package ast

import (
	"go/ast"

	"github.com/JoaoVitor615/archscan/internal/model"
)

// FileInspector extracts models from a parsed Go AST file.
type FileInspector struct{}

// NewFileInspector creates a new FileInspector instance.
func NewFileInspector() *FileInspector {
	return &FileInspector{}
}

func (fi *FileInspector) InspectFile(node *ast.File, pkg *model.PackageModel) {
	ast.Inspect(node, func(n ast.Node) bool {
		switch decl := n.(type) {
		case *ast.GenDecl:
			fi.extractStructs(decl, pkg)
		case *ast.FuncDecl:
			fi.extractFunction(decl, pkg)
		}
		return true
	})
}

func (i *FileInspector) extractStructs(decl *ast.GenDecl, pkg *model.PackageModel) {
	for _, spec := range decl.Specs {
		typeSpec, ok := spec.(*ast.TypeSpec)
		if !ok {
			continue
		}

		structType, ok := typeSpec.Type.(*ast.StructType)
		if !ok {
			continue
		}

		structName := typeSpec.Name.Name
		var fieldNames []string

		if structType.Fields != nil {
			for _, field := range structType.Fields.List {
				for _, name := range field.Names {
					fieldNames = append(fieldNames, name.Name)
				}
			}
		}

		pkg.Structs[structName] = &model.StructModel{
			Name:    structName,
			Package: pkg.Name,
			Fields:  fieldNames,
			Methods: make([]*model.FunctionModel, 0),
		}
	}
}

func (i *FileInspector) extractFunction(decl *ast.FuncDecl, pkg *model.PackageModel) {
	funcName := decl.Name.Name
	receiverName := i.extractReceiverType(decl.Recv)

	fn := &model.FunctionModel{
		Name:         funcName,
		ReceiverType: receiverName,
		Package:      pkg.Name,
	}

	if receiverName != "" {
		if targetStruct, exists := pkg.Structs[receiverName]; exists {
			targetStruct.Methods = append(targetStruct.Methods, fn)
		}
	} else {
		pkg.Functions = append(pkg.Functions, fn)
	}
}

func (i *FileInspector) extractReceiverType(recv *ast.FieldList) string {
	if recv == nil || len(recv.List) == 0 {
		return ""
	}

	switch expr := recv.List[0].Type.(type) {
	case *ast.Ident:
		return expr.Name
	case *ast.StarExpr:
		if ident, ok := expr.X.(*ast.Ident); ok {
			return ident.Name
		}
	}

	return ""
}
