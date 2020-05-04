package cmd

import (
	"context"
	"log"

	"github.com/mentos1386/ipfs-cloud/pkg/ipfs"
	"github.com/spf13/cobra"
)

var nodeCmd = &cobra.Command{
	Use:   "node",
	Short: "Start IPFS Node",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		_, err := ipfs.StartNode(ctx)
		if err != nil {
			log.Panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(nodeCmd)
}
