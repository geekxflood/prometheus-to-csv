// cmd/query.go
package cmd

import (
	"context"
	"fmt"
	"time"

	"github.com/geekxflood/prometheus-to-csv/internal/csvutil"
	"github.com/geekxflood/prometheus-to-csv/internal/promclient"
	"github.com/geekxflood/prometheus-to-csv/internal/util"
	"github.com/spf13/cobra"
)

var queryCmd = &cobra.Command{
	Use:   "query",
	Short: "Query Prometheus and export data to CSV",
	Long:  `Query Prometheus metrics and export the results to a CSV file.`,
	Run: func(cmd *cobra.Command, args []string) {
		address, _ := cmd.Flags().GetString("address")
		query, _ := cmd.Flags().GetString("query")
		timeRange, _ := cmd.Flags().GetString("time-range")

		// Validate inputs
		if address == "" || query == "" || timeRange == "" {
			fmt.Println("Address, query, and time range are required")
			return
		}

		// Set the InsecureSkipVerify flag in promclient package
		promclient.InsecureSkipVerify = insecureSkipVerify

		// Create Prometheus client using the promclient package
		v1api, err := promclient.CreatePrometheusClient(address)
		if err != nil {
			fmt.Printf("Error creating Prometheus client: %v\n", err)
			return
		}

		// Parse time range
		start, end, err := util.ParseTimeRange(timeRange)
		if err != nil {
			fmt.Printf("Error parsing time range: %v\n", err)
			return
		}

		// Query Prometheus
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		result, warnings, err := util.QueryPrometheus(ctx, v1api, query, start, end)
		if err != nil {
			fmt.Printf("Error querying Prometheus: %v\n", err)
			return
		}
		if len(warnings) > 0 {
			fmt.Printf("Warnings: %v\n", warnings)
		}

		// Process results to CSV
		csvData, err := csvutil.CreateCSVDataFromResult(result)
		if err != nil {
			fmt.Printf("Error processing results: %v\n", err)
			return
		}

		// Output CSV data
		csvutil.PrintCSVData(csvData)
	},
}

func init() {
	rootCmd.AddCommand(queryCmd)
	queryCmd.Flags().StringP("address", "a", "", "Address of the Prometheus server")
	queryCmd.Flags().StringP("query", "q", "", "PromQL query")
	queryCmd.Flags().StringP("time-range", "t", "", "Time range for the query")
}
