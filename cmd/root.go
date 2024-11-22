package cmd

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "app",
	Short: "app is a CLI application",
}

func init() {
	RootCmd.AddCommand(restfulCmd)
}
