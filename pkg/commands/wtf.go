package commands

import (
	"fmt"
	"os"

	"github.com/oktopac/wtf/pkg/checker"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "wtf",
	Short: "wtf: find out about files quickly",
	Long:  `A simple command line tool that helps you quickly work out what a file is, and whats in it`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		checker.CheckFile(args[0])
	},
}

// Execute runs the main command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
