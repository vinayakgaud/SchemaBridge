package commands

import (
	"github.com/spf13/cobra"
	"github.com/vinayakgaud/schemabridge/core/filehandling"
)

var analyzeCmd = &cobra.Command{
	Use:   "analyze <input>",
	Short: "Analyze an API contract or sample",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		file := args[0]
		data, err := filehandling.ReadFile(file)

		if err != nil {
			return err
		}

		cmd.Printf("Input: %s\n", file)
		cmd.Printf("Bytes: %d\n", len(data))

		return nil
	},
}

func init() {
	RootCmd.AddCommand(analyzeCmd)
}
