/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/ylanzinhoy/docker_compose_gen/internal/generator"
)

// dockercomposegenCmd represents the dockercomposegen command
var gen = &cobra.Command{
	Use:   "gen",
	Short: "comando para gerar um docker-compose.yml com seu banco de dados",
	Run: func(cmd *cobra.Command, args []string) {
		generator.GenerateDockerFile()
	},
}

func init() {
	rootCmd.AddCommand(gen)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// dockercomposegenCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// dockercomposegenCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
