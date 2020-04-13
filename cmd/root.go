package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// Version of ipfs-cloud
var Version = "development"

var rootCmd = &cobra.Command{
	Use:     "ipfs-cloud",
	Short:   "Private file storage on IPFS",
	Long:    `Store anything you want on IPFS`,
	Version: Version,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
