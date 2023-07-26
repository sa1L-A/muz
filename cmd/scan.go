package cmd

import (
	"github.com/sa1L/muz/pkg/logger"
	"github.com/spf13/cobra"
)

var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "hello",
	Long:  "hello there",
	Run: func(cmd *cobra.Command, args []string) {
		logger.Info("handle scan here")
	},
}

func init() {
	rootCmd.AddCommand(scanCmd)
}
