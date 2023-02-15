package cmd

import (
	"fmt"
	"os"
	"thebrag/helpers"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a brag",
	Long:  `This command will delete a brag using bragId`,
	Run: func(cmd *cobra.Command, args []string) {
		if os.Getenv("USER_ID") == "" {
			fmt.Println("Please login to delete brags")
			return
		}
		bragId, _ := cmd.Flags().GetInt("id")
		if bragId <= 0 {
			fmt.Println("bragId not given")
			return
		}
		bragResponse := helpers.DeleteABrag(bragId)
		fmt.Println(bragResponse.Data)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	// local flags
	deleteCmd.Flags().IntP("id", "i", 0, "delete brag having ID")
}
