// cmd/root.go
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var insecureSkipVerify bool

var rootCmd = &cobra.Command{
	Use:   "prometheus-to-csv",
	Short: "A tool for exporting Prometheus data to CSV",
	Long:  `Prometheus-to-CSV is a CLI tool for querying Prometheus and exporting the data to CSV format.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&insecureSkipVerify, "insecure", "i", false, "Skip TLS certificate verification")
}
