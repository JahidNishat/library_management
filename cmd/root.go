package cmd

import (
	"github.com/spf13/cobra"
)

var root = &cobra.Command{
	Use: "root",
	Short: "root command",
	Long: "This is root command",
}

func Execute() error {
	return root.Execute()
}

func init() {
	root.AddCommand(serve)
}