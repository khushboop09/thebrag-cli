/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"thebrag/cmd"
	"thebrag/configs"
)

func main() {
	configs.LoadEnv()
	configs.LoadData()
	cmd.Execute()
}
