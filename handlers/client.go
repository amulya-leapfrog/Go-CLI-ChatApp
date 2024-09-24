package handlers

import (
	"bufio"
	"fmt"
	"log"
	"os"

	socketio_client "github.com/zhouhui8915/go-socket.io-client"
)

func StartClient(user *User) {
	serverPort := os.Getenv("PORT")
	serverHost := os.Getenv("DB_HOST")

	opts := &socketio_client.Options{
		Transport: "websocket",
		Query:     make(map[string]string),
	}

	opts.Query["id"] = user.ID
	opts.Query["user"] = user.Username
	opts.Query["pwd"] = user.Password
	opts.Query["room"] = "test"

	uri := fmt.Sprintf("http://%v:%v/socket.io/", serverHost, serverPort)

	client, err := socketio_client.NewClient(uri, opts)
	if err != nil {
		log.Printf("NewClient error:%v\n", err)
		return
	}

	client.On("error", func() {
		log.Printf("Error Occured \n")
	})

	client.On("connection", func() {
		log.Printf("Connected to the server \n")
	})

	client.On("chat_message", func(msg map[string]interface{}) {
		userName := msg["username"].(string)
		content := msg["content"].(string)
		log.Printf("%s: %s\n", userName, content)
		fmt.Print("Send Message: ")
	})

	client.On("disconnection", func() {
		log.Printf("Server Disconnected\n")
	})

	reader := bufio.NewReader(os.Stdin)

	for {
		data, _, _ := reader.ReadLine()
		command := string(data)

		if command != "" {
			client.Emit("chat_message", command)
		}
	}
}
