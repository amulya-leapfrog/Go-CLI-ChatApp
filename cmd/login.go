package cmd

import (
	"fmt"
	"log"

	"github.com/amulya-leapfrog/go-chat/handlers"
	"github.com/spf13/cobra"
)

var loginCmd = &cobra.Command{
	Use:     "login",
	Aliases: []string{"l"},
	Short:   "Login a user",
	Run: func(cmd *cobra.Command, args []string) {
		var username, password string
		fmt.Print("Enter username: ")
		fmt.Scan(&username)
		fmt.Print("Enter password: ")
		fmt.Scan(&password)

		user, err := handlers.LoginUser(username, password)
		if err != nil {
			log.Fatalf("Error logging in: %v", err)
			return
		}

		fmt.Println("Logged in successfully!")

		handlers.StartClient(user)
	},
}
