package main

import (
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "archscan",
	Short: "Static architecture analyzer and design metrics tool",
	Long: `Archscan scans Go repositories to calculate advanced 
architectural metrics such as LCOM, cohesion, and connascence.`,
}
