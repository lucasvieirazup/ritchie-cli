package cmd

import "github.com/spf13/cobra"

// NewDeleteCmd create a new delete instance
func NewDeleteCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "delete SUBCOMMAND",
		Short: "Delete objects",
		Long:  `Delete objects like users, etc.`,
	}
}
