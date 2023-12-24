package cmd

import (
	"github.com/spf13/cobra"
)
 
var rootCmd = &cobra.Command{
	Use:   "cobra",
	Short: "Iptracker created with Ujjwal Mahar ⚓",
	Long: `This is a CLI Iptracker application.`,
}

func Execute() error{
	return rootCmd.Execute()
}