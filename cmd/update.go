/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"thebrag/helpers"
	"thebrag/responses"

	"github.com/mitchellh/mapstructure"
	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update brag data",
	Long:  `This command will let you update content of a brag`,
	Run: func(cmd *cobra.Command, args []string) {
		if os.Getenv("USER_ID") == "" {
			fmt.Println("Please login to update brags")
			return
		}
		bragId, _ := cmd.Flags().GetInt("id")
		bragTitle, _ := cmd.Flags().GetString("title")
		bragDetails, _ := cmd.Flags().GetString("details")
		categoryName, _ := cmd.Flags().GetString("category")
		var categoryId int
		if bragId <= 0 {
			fmt.Println("bragId not given")
			return
		} else {
			existingBragData, statusCode := helpers.GetABrag(bragId)
			if statusCode != 200 {
				fmt.Println(existingBragData.Data)
				return
			}

			var existingBrag responses.Brag
			mapstructure.Decode(existingBragData.Data, &existingBrag)
			if bragTitle == "" {
				bragTitle = existingBrag.Title
			}
			if categoryName == "" {
				categoryName = existingBrag.CategoryName
				categoryId = existingBrag.CategoryID
			}
			if bragDetails == "" {
				bragDetails = existingBrag.Details
			}
		}

		responseBody, statusCode := helpers.UpdateABrag(bragId, bragTitle, bragDetails, categoryName, categoryId)
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
	updateCmd.Flags().StringP("category", "c", "", "updated category of the brag")
}
