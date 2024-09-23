package cmd

import (
	"fmt"
	"log"

	"github.com/amulya-leapfrog/go-chat/handlers"
	"github.com/spf13/cobra"
)

var registerCmd = &cobra.Command{
	Use:     "register",
	Aliases: []string{"r"},
	Short:   "Register a new user",
	Run: func(cmd *cobra.Command, args []string) {
		var email, username, password string
		fmt.Print("Enter email: ")
		fmt.Scan(&email)
		fmt.Print("Enter username: ")
		fmt.Scan(&username)
		fmt.Print("Enter password: ")
		fmt.Scan(&password)

		user := handlers.User{Email: email, Username: username, Password: password}
		err := handlers.RegisterUser(user)
		if err != nil {
			log.Fatalf("Error registering user: %v", err)
		}
		fmt.Println("User registered successfully!")
	},
}
