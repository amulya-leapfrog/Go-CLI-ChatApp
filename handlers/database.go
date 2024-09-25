package handlers

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

var userCollection *mongo.Collection
var chatCollection *mongo.Collection

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	hostName := os.Getenv("DB_HOST")
	userName := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbPort := os.Getenv("DB_PORT")

	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s/", userName, password, hostName, dbPort)

	clientOptions := options.Client().ApplyURI(uri)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("========== Connected to MongoDB! ==========")

	userCollection = client.Database("chat_app").Collection("users")
	chatCollection = client.Database("chat_app").Collection("chats")
}

func RegisterUser(user User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
	}

	user.Password = string(hashedPassword)

	_, err = userCollection.InsertOne(context.TODO(), user)
	return err
}

func LoginUser(username, password string) (*User, error) {
	var user User

	err := userCollection.FindOne(context.TODO(), bson.M{"username": username}).Decode(&user)
	if err != nil {
		return nil, errors.New("invalid username or password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		fmt.Println("Password does not match")
		return nil, errors.New("invalid username or password")
	}

	return &user, nil
}

func SaveChatMessage(message string, sender string) error {
	chat := ChatMessage{
		Message:   message,
		Sender:    sender,
		Timestamp: time.Now(),
	}

	_, err := chatCollection.InsertOne(context.TODO(), chat)
	return err
}

func FetchChatHistory() ([]ChatHistoryOutput, error) {
	pipeline := mongo.Pipeline{
		{
			{Key: "$addFields", Value: bson.D{
				{Key: "senderObjectID", Value: bson.D{
					{Key: "$toObjectId", Value: "$sender"},
				}},
			}},
		},
		{
			{Key: "$lookup", Value: bson.D{
				{Key: "from", Value: "users"},
				{Key: "localField", Value: "senderObjectID"},
				{Key: "foreignField", Value: "_id"},
				{Key: "as", Value: "senderInfo"},
			}},
		},
		{
			{Key: "$unwind", Value: bson.D{
				{Key: "path", Value: "$senderInfo"},
				{Key: "preserveNullAndEmptyArrays", Value: false},
			}},
		},
		{
			{Key: "$project", Value: bson.D{
				{Key: "_id", Value: 1},
				{Key: "message", Value: 1},
				{Key: "timestamp", Value: 1},
				{Key: "senderInfo.username", Value: 1},
				{Key: "senderInfo.email", Value: 1},
			}},
		},
	}

	history, err := chatCollection.Aggregate(context.TODO(), pipeline)
	if err != nil {
		return nil, err
	}

	var chatData []ChatHistoryOutput
	for history.Next(context.TODO()) {
		var chat ChatHistoryOutput
		err := history.Decode(&chat)
		if err != nil {
			return nil, err
		}
		chatData = append(chatData, chat)
	}

	return chatData, nil
}
