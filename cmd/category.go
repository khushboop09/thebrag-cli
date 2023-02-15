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

// categoryCmd represents the category command
var categoryCmd = &cobra.Command{
	Use:   "category",
	Short: "Get or Add Categories",
	Long: `This command gets your existing categories or let's you add new categories. For example:

"thebrag category" command will get all your existing categories.
"thebrag category -c 'Professional Projects'" command will add a new category.`,
	Run: func(cmd *cobra.Command, args []string) {
		if os.Getenv("USER_ID") == "" {
			fmt.Println("Please login to export your brags")
			return
		}
		catName, _ := cmd.Flags().GetString("add")
		if catName == "" {
			// get cats
			categoriesResponse, statusCode := helpers.GetAllCategories()
			if statusCode != 200 {
				fmt.Println("Categories not found, add a new one!")
				return
			}
			categories := categoriesResponse.Data
			if categories == nil {
				fmt.Println("Categories not found, add a new one!")
			}
			for i := range categories {
				fmt.Printf("[%d] ", categories[i].ID)
				fmt.Println(categories[i].Name)
				fmt.Println()
			}
		} else {
			// add new cat
			resp, statusCode := helpers.AddACategory(catName)
			if statusCode == 201 {
				fmt.Println("category created!")
			} else {
				fmt.Println(resp.Data)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(categoryCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// categoryCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	categoryCmd.Flags().StringP("add", "c", "", "Add category name")
}
