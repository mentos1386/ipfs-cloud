package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var decryptCmd = &cobra.Command{
	Use:   "decrypt",
	Short: "decrypt STDIN",
	Long:  ``,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("hello world: decrypt")
	},
}

func init() {
	rootCmd.AddCommand(decryptCmd)
}
