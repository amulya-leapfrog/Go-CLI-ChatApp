package handlers

import "fmt"

func ShowHistory(chatHistory []ChatHistoryOutput) {
	if len(chatHistory) == 0 {
		fmt.Println("No chat history found.")
		return
	}

	for _, chat := range chatHistory {
		username := chat.SenderInfo.Username
		message := chat.Message

		fmt.Printf("%s: %s\n", username, message)
	}
}
