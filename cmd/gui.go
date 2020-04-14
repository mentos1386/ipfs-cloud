package cmd

import (
	"github.com/spf13/cobra"

	"github.com/mentos1386/ipfs-cloud/internal/application"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "gui",
	Short: "Start GUI",
	Long:  ``,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		application.Create(cmd.Version, "space.tjo.ipfs-cloud", args)
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
}