package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var encryptCmd = &cobra.Command{
	Use:   "encrypt",
	Short: "encrypt STDIN",
	Long:  ``,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("hello world: encrypt")
	},
}

func init() {
	rootCmd.AddCommand(encryptCmd)
}
