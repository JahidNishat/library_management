package cmd

import (
	"github.com/library_management/api"
	"github.com/spf13/cobra"
)

var serve = &cobra.Command{
	Use:   "serve",
	Short: "serve command",
	Long:  "This is serve command",
	Run:   Start,
}

func Start(cmd *cobra.Command, args []string) {
	api.StartServer()
}
