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
	Long:  `This command will let you login to your account or create a new account if it doesn't exists.`,
	Run: func(cmd *cobra.Command, args []string) {
		// check if .data file exists
		// if yes read the file and find user_id
		// if found no need to login
		// if not found login user
		// if file doesn't exist
		// login user
		if _, err := os.Stat(".data"); err == nil {
			configs.LoadData()
			if os.Getenv("USER_ID") == "" {
				helpers.StartLogin()
			} else {
				fmt.Println("You are already logged in!")
			}
		} else {
			helpers.StartLogin()
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
