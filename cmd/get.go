package cmd

import (
	"fmt"
	"os"

	"thebrag/helpers"
	"thebrag/responses"

	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "get all your brags",
	Long:  `the get command gets all your saved brags`,
	Run: func(cmd *cobra.Command, args []string) {
		if os.Getenv("USER_ID") == "" {
			fmt.Println("Please login to fetch your brags")
			return
		}
		skip, _ := cmd.Flags().GetInt("skip")
		limit, _ := cmd.Flags().GetInt("limit")
		id, _ := cmd.Flags().GetInt("id")
		if id <= 0 {
			//get all brags
			bragResponse, statusCode := helpers.GetAllBrags(skip, limit)
			if statusCode == 200 {
				brags := bragResponse.Data
				for i := 0; i < len(brags); i++ {
					fmt.Printf("[%d]", brags[i].ID)
					fmt.Println(brags[i].Title)
					fmt.Println(brags[i].Details)
					fmt.Println()
				}
			} else {
				fmt.Println(bragResponse.Data)
			}
		} else {
			//get brag of given id
			bragResponse, statusCode := helpers.GetABrag(id)
			if statusCode == 200 {
				var brag responses.Brag
				// TODO: the below line is not working
				brag, ok := bragResponse.Data.(responses.Brag)
				if ok {
					fmt.Printf("[%d] ", brag.ID)
					fmt.Println(brag.Title)
					fmt.Println(brag.Details)
				} else {
					fmt.Println("error parsing response")
				}
			} else {
				fmt.Println(bragResponse.Data)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(getCmd)

	// local flags
	getCmd.Flags().IntP("limit", "n", 10, "get max no. of brags")
	getCmd.Flags().IntP("skip", "s", 0, "skip no. of brags from start")
	getCmd.Flags().IntP("id", "i", 0, "get brag having ID")
}
