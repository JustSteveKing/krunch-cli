package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "krunch",
	Short: "Krunch a link",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Printf("Failed to run command: %v\n\n", err)
		os.Exit(1)
	}
}
