package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var scanCmd = &cobra.Command{
	Use:   "scan [directory]",
	Short: "Scans a target repository",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		targetPath := args[0]

		fmt.Printf("Starting architectural analysis on directory: %s\n", targetPath)

		//future logic...
	},
}

func init() {
	rootCmd.AddCommand(scanCmd)

}
