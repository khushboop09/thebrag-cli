/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"thebrag/helpers"

	"github.com/spf13/cobra"
)

// exportCmd represents the export command
var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "Export your brags",
	Long:  `Exports your brags to a CSV file and send it to the registered email id.`,
	Run: func(cmd *cobra.Command, args []string) {
		if os.Getenv("USER_ID") == "" {
			fmt.Println("Please login to export your brags")
			return
		}
		bragsCategory, _ := cmd.Flags().GetString("category")
		bragDates, _ := cmd.Flags().GetStringSlice("date")

		responseBody, _ := helpers.ExportBrags(bragDates[0], bragDates[1], bragsCategory)
		fmt.Println(responseBody.Data)
	},
}

func init() {
	rootCmd.AddCommand(exportCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// exportCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	exportCmd.Flags().StringSliceP("date", "d", nil, "['from', 'to') date range for when the brags were created")
	exportCmd.Flags().StringP("category", "c", "", "category name")
}
