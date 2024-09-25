package cmd

import (
	"log"

	"github.com/amulya-leapfrog/go-chat/handlers"
	"github.com/spf13/cobra"
)

var historyCmd = &cobra.Command{
	Use:     "history",
	Aliases: []string{"h"},
	Short:   "Show chat history",
	Run: func(cmd *cobra.Command, args []string) {
		history, err := handlers.FetchChatHistory()
		if err != nil {
			log.Fatalf("Error Fetching Chat History: %v", err)
			return
		}

		handlers.ShowHistory(history)
	},
}
