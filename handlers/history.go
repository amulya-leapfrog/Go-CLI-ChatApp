package handlers

import "fmt"

func ShowHistory(chatHistory []ChatHistoryOutput) {
	for _, chat := range chatHistory {
		username := chat.SenderInfo.Username
		message := chat.Message

		fmt.Printf("%s: %s\n", username, message)
	}
}
