package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "chat-app",
	Short: "CLI Chat Application",
	Long:  `A CLI chat application with MongoDB and Socket.IO support.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}

func init() {
	rootCmd.AddCommand(registerCmd)
	rootCmd.AddCommand(loginCmd)
	rootCmd.AddCommand(serverCmd)
	rootCmd.AddCommand(historyCmd)
}
