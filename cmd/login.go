package cmd

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"syscall"

	"github.com/spf13/cobra"
	"golang.org/x/term"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		var email string
		reader := bufio.NewReader(os.Stdin)
		email, _ = cmd.Flags().GetString("email")

		if email == "" {
			fmt.Println("Enter username: ")
			email, _ = reader.ReadString('\n')
		}

		fmt.Println("Enter password: ")
		bytePassword, _ := term.ReadPassword(int(syscall.Stdin))
		password := string(bytePassword)

		getUserToken(email, password)
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
	loginCmd.PersistentFlags().String("email", "", "A help for email")
}

func getUserToken(email string, password string) {
	jsonData, _ := json.Marshal(map[string]string{
		"email":    email,
		"password": password,
	})
	body := bytes.NewBuffer(jsonData)

	responseBytes := getLoginData(body)
	fmt.Println("token: " + string(responseBytes))
}

func getLoginData(body *bytes.Buffer) []byte {
	response, err := http.Post(
		"http://172.19.0.3:8000/users/login",
		"application/json",
		body,
	)

	if err != nil {
		log.Printf("Could not request login. %v", err)
	}

	responseBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Printf("Could not read response body. %v", err)
	}

	return responseBytes
}
