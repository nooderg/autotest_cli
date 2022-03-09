/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/spf13/cobra"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command.`,
	Run: func(cmd *cobra.Command, args []string) {

		if generateFile() != nil {
			fmt.Println("File search error")
		} else {
			fmt.Println("Success")

		}

	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
}

func generateFile() error {

	_, err := http.Get("http://172.19.0.3:8000/users/generate?name=test")

	if err != nil {
		log.Printf("Could not request login. %v", err)
		return err
	}

	var isDocument bool = false
	var minuteCurrent int = time.Now().Minute()

	for minuteCurrent >= minuteCurrent+1 || !isDocument {
		isExistdocument := searchFile()

		if isExistdocument == nil {
			isDocument = true
			return isExistdocument
		}
	}

	return nil
}

func searchFile() error {

	var client = &http.Client{
		Timeout: time.Second * 10,
	}
	_, err := client.Get("http://172.19.0.3:8000/users/search")

	if err != nil {
		log.Printf("Could not request login. %v", err)
		return err
	}
	return nil
}
