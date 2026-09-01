package main

import (
	"fmt"
	"os"

	"github.com/JoaoVitor615/archscan/internal/ast"
	"github.com/JoaoVitor615/archscan/internal/ui"
	"github.com/spf13/cobra"
)

var showStructure bool

var scanCmd = &cobra.Command{
	Use:   "scan [directory]",
	Short: "Scans a target repository and presents architectural metrics",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		targetPath := "."
		if len(args) > 0 {
			targetPath = args[0]
		}

		parser := ast.NewProjectParser()
		packages, err := parser.ParseProject(targetPath)
		if err != nil {
			fmt.Printf("Error scanning project: %v\n", err)
			os.Exit(1)
		}

		if len(packages) == 0 {
			fmt.Println("No Go packages found.")
			return
		}

		// Calculate overview statistics from parsed AST
		totalStructs := 0
		totalFuncs := 0
		for _, pkg := range packages {
			totalStructs += len(pkg.Structs)
			totalFuncs += len(pkg.Functions)
			for _, s := range pkg.Structs {
				totalFuncs += len(s.Methods)
			}
		}

		// Always render the main architecture metrics dashboard
		ui.RenderDashboard(targetPath, len(packages), totalStructs, totalFuncs)

		// If --show-structure / -s flag is provided, print detailed AST breakdown
		if showStructure {
			fmt.Printf("📦 DETAILED PACKAGE & AST STRUCTURE (%s):\n\n", targetPath)
			for dir, pkg := range packages {
				fmt.Printf("📦 Package: %s (%s)\n", pkg.Name, dir)

				if len(pkg.Structs) > 0 {
					fmt.Println("  Structs:")
					for _, s := range pkg.Structs {
						fmt.Printf("    • %s (Fields: %d, Methods: %d)\n", s.Name, len(s.Fields), len(s.Methods))
						for _, m := range s.Methods {
							fmt.Printf("        -> method: %s()\n", m.Name)
						}
					}
				}

				if len(pkg.Functions) > 0 {
					fmt.Println("  Functions:")
					for _, fn := range pkg.Functions {
						fmt.Printf("    • %s()\n", fn.Name)
					}
				}

				fmt.Println()
			}
		}
	},
}

func init() {
	scanCmd.Flags().BoolVarP(&showStructure, "show-structure", "s", false, "Display detailed package, struct, and function structure")
	rootCmd.AddCommand(scanCmd)
}

