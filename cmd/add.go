package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"thebrag/requests"
	"thebrag/responses"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add a new brag",
	Long:  `this command will add a new brag to your work log.`,
	Run: func(cmd *cobra.Command, args []string) {
		bragTitle, _ := cmd.Flags().GetString("title")
		bragDetails, _ := cmd.Flags().GetString("details")
		brag := requests.AddBragRequest{
			Title:   bragTitle,
			Details: bragDetails,
		}
		json_data, err := json.Marshal(brag)
		if err != nil {
			log.Fatal(err)
		}
		response, err := http.Post("http://localhost:8080/brag", "application/json", bytes.NewBuffer(json_data))
		if err != nil {
			fmt.Println(err)
		}

		defer response.Body.Close()
		responseBody, err := io.ReadAll(response.Body)
		if err != nil {
			fmt.Println(err)
		}

		var res responses.PostBragResponse
		json.Unmarshal(responseBody, &res)
		if response.StatusCode == 201 {
			fmt.Println(res.Message)
		} else {
			fmt.Println(res.Data)
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// local flags
	addCmd.Flags().StringP("title", "t", "", "specify title of your brag")
	addCmd.Flags().StringP("details", "d", "", "specify details of your brag")
}
