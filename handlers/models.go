package handlers

import "time"

type User struct {
	ID       string `bson:"_id,omitempty"`
	Email    string `bson:"email"`
	Username string `bson:"username"`
	Password string `bson:"password"`
}

type ChatMessage struct {
	ID        string    `bson:"_id,omitempty"`
	Message   string    `bson:"message"`
	Sender    string    `bson:"sender"`
	Timestamp time.Time `bson:"timestamp"`
}
