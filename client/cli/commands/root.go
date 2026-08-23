package commands

import "github.com/spf13/cobra"

var RootCmd = &cobra.Command{
	Use:   "sb",
	Short: "API contract intelligence and migration tool",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}
