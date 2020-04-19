package cmd

import (
	"os"
	"log"

	"github.com/spf13/cobra"
	"github.com/mentos1386/ipfs-cloud/pkg/app"
)

var guiCmd = &cobra.Command{
	Use:   "gui",
	Short: "Start GUI",
	Long:  ``,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		application, err := app.Create(cmd.Version, "space.tjo.ipfs-cloud")
		if err != nil {
			log.Panicf("Error creating application: %v", err)
		}

		// Launch the application
		os.Exit(application.Run(args))
	},
}

func init() {
	rootCmd.AddCommand(guiCmd)
}
