package cmd

import (
	"fmt"
	"os"

	"thebrag/helpers"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add a new brag",
	Long:  `this command will add a new brag to your work log.`,
	Run: func(cmd *cobra.Command, args []string) {
		if os.Getenv("USER_ID") == "" {
			fmt.Println("Please login to add brags")
			return
		}
		bragTitle, _ := cmd.Flags().GetString("title")
		bragDetails, _ := cmd.Flags().GetString("details")
		categoryName, _ := cmd.Flags().GetString("category")

		categoryId := helpers.GetCategoryId(categoryName)
		if categoryId == 0 {
			fmt.Println("Invalid category, please check if this category exists or it is correctly spelled")
			return
		}
		response, statusCode := helpers.AddABrag(bragTitle, bragDetails, categoryId)
		if statusCode == 201 {
			fmt.Println(response.Message)
		} else {
			fmt.Println(response.Data)
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// local flags
	addCmd.Flags().StringP("title", "t", "", "specify title of your brag")
	addCmd.Flags().StringP("details", "d", "", "specify details of your brag")
	addCmd.Flags().StringP("category", "c", "", "specify the category of your brag")
}
