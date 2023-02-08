/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"thebrag/helpers"

	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update brag data",
	Long:  `This command will let you update content of a brag`,
	Run: func(cmd *cobra.Command, args []string) {
		bragId, _ := cmd.Flags().GetInt("id")
		bragTitle, _ := cmd.Flags().GetString("title")
		bragDetails, _ := cmd.Flags().GetString("details")

		if bragId <= 0 {
			fmt.Println("bragId not given")
			return
		} else if bragTitle == "" {
			fmt.Println("brag title not given")
			return
		}

		responseBody, statusCode := helpers.UpdateABrag(bragId, bragTitle, bragDetails)
		if statusCode == 201 {
			fmt.Println(responseBody.Message)
		} else {
			fmt.Println(responseBody.Data)
		}
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)

	// local flags
	updateCmd.Flags().IntP("id", "i", 0, "ID of the brag you want to update")
	updateCmd.Flags().StringP("title", "t", "", "updated title of the brag")
	updateCmd.Flags().StringP("details", "d", "", "updated details of the brag")
}
