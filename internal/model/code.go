package model

// FunctionModel represents a standalone function or a struct method.
type FunctionModel struct {
	Name            string
	ReceiverType    string // Empty ("") for regular functions; populated for methods
	Package         string
	FieldAccesses   []string // Struct fields accessed
	Calls           []string // Other functions/methods invoked
	ComplexityScore int      // Cyclomatic complexity (ifs, loops, switches)
}

// StructModel represents a Go struct found during static analysis.
type StructModel struct {
	Name    string
	Package string
	Fields  []string
	Methods []*FunctionModel
}

// PackageModel represents a Go package containing structs and package-level functions.
type PackageModel struct {
	Name      string
	Path      string
	Structs   map[string]*StructModel
	Functions []*FunctionModel // Funções livres que não pertencem a nenhuma struct
}
