package main

import (
	"fmt"
	"os"

	"github.com/JoaoVitor615/archscan/internal/ast"
	"github.com/spf13/cobra"
)

var showStructure bool

var scanCmd = &cobra.Command{
	Use:   "scan [directory]",
	Short: "Scans a target repository",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		targetPath := args[0]

		fmt.Printf("Analyzing directory: %s\n\n", targetPath)

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

		if showStructure {
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
		} else {
			fmt.Printf("✓ Successfully scanned %d package(s).\n", len(packages))
		}
	},
}

func init() {
	scanCmd.Flags().BoolVarP(&showStructure, "show-structure", "s", false, "Display detailed package, struct, and function structure")
	rootCmd.AddCommand(scanCmd)
}

