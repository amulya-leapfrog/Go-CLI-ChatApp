package cmd

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/amulya-leapfrog/go-chat/handlers"
	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:     "server",
	Aliases: []string{"s"},
	Short:   "Start the server",
	Run: func(cmd *cobra.Command, args []string) {
		serverPort := os.Getenv("PORT")

		go func() {
			handlers.StartServer()
		}()

		log.Printf("Starting server on port %v...", serverPort)

		port := fmt.Sprintf(":%v", serverPort)

		if err := http.ListenAndServe(port, nil); err != nil {
			log.Fatalf("Error starting server: %v", err)
		}
	},
}
