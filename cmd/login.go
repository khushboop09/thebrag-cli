/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"thebrag/configs"
	"thebrag/helpers"

	"github.com/spf13/cobra"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login or create your thebrag account",
	Long: `This command will let you login to your account or create a new account if it doesn't exists.

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// check if .data file exists
		// if yes read the file and find user_id
		// if found no need to login
		// if not found login user
		// if file not found
		// login user
		if _, err := os.Stat(".data"); err == nil {
			configs.LoadData()
			if os.Getenv("USER_ID") == "" {
				fmt.Printf("File does not exist\n")
				name := ""
				var email string
				var password string
				fmt.Println("Name: [Press enter if you don't want to share]")
				fmt.Scanln(&name)
				fmt.Println("Email: [Mandatory]")
				fmt.Scanln(&email)
				fmt.Println("Password: [Mandatory, peeping tom alert!]")
				fmt.Scanln(&password)
				response, statusCode := helpers.LoginUser(name, email, password)
				if statusCode == 201 {
					fmt.Println(response.Message)
					// save to id to .data file
					helpers.SaveUserIdToDataFile(response.Data, "add")
					configs.LoadData()
				} else {
					fmt.Println(response.Data)
				}
			} else {
				fmt.Println("You are already logged in!")
			}
		} else {
			name := ""
			var email string
			var password string
			fmt.Println("Name: [Press enter if you don't want to share]")
			fmt.Scanln(&name)
			fmt.Println("Email: [Mandatory]")
			fmt.Scanln(&email)
			fmt.Println("Password: [Mandatory, peeping tom alert!]")
			fmt.Scanln(&password)
			response, statusCode := helpers.LoginUser(name, email, password)
			if statusCode == 201 {
				fmt.Println(response.Message)
				// save to id to .data file
				helpers.SaveUserIdToDataFile(response.Data, "create")
				configs.LoadData()
			} else {
				fmt.Println(response.Data)
			}
		}

	},
}

func init() {
	rootCmd.AddCommand(loginCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// loginCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// loginCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
