package ast

import (
	"os"
	"path/filepath"
	"testing"
)

func TestFileFinderAndParser(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "archscan_test")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	pkgDir := filepath.Join(tmpDir, "sample")
	if err := os.MkdirAll(pkgDir, 0755); err != nil {
		t.Fatalf("failed to create pkg dir: %v", err)
	}

	code := `package sample

type Service struct {
	ID   string
	Name string
}

func (s *Service) Execute() string {
	return s.Name
}

func HelperFunc() int {
	return 42
}
`
	filePath := filepath.Join(pkgDir, "service.go")
	if err := os.WriteFile(filePath, []byte(code), 0644); err != nil {
		t.Fatalf("failed to write test file: %v", err)
	}

	parser := NewProjectParser()
	packages, err := parser.ParseProject(tmpDir)
	if err != nil {
		t.Fatalf("ParseProject failed: %v", err)
	}

	if len(packages) != 1 {
		t.Fatalf("expected 1 package, got %d", len(packages))
	}

	pkg, ok := packages[pkgDir]
	if !ok {
		t.Fatalf("package not found for dir %s", pkgDir)
	}

	if pkg.Name != "sample" {
		t.Errorf("expected package name 'sample', got '%s'", pkg.Name)
	}

	if len(pkg.Structs) != 1 {
		t.Fatalf("expected 1 struct, got %d", len(pkg.Structs))
	}

	svc, ok := pkg.Structs["Service"]
	if !ok {
		t.Fatalf("struct 'Service' not found")
	}

	if len(svc.Fields) != 2 {
		t.Errorf("expected 2 fields, got %d", len(svc.Fields))
	}

	if len(svc.Methods) != 1 || svc.Methods[0].Name != "Execute" {
		t.Errorf("expected method 'Execute', got %+v", svc.Methods)
	}

	if len(pkg.Functions) != 1 || pkg.Functions[0].Name != "HelperFunc" {
		t.Errorf("expected free function 'HelperFunc', got %+v", pkg.Functions)
	}
}
