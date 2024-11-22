package cmd

import (
	"edmk/interface/restful"

	"github.com/spf13/cobra"
)

var restfulCmd = &cobra.Command{
	Use:   "restful",
	Short: "Start the RESTful API server",
	Run: func(cmd *cobra.Command, args []string) {
		restful.NewRestfulServer()
	},
}
