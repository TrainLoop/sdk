package main

import (
	"fmt"
	"log"

	trainloop "github.com/TrainLoop/sdk/go"
)

func main() {
	client := trainloop.NewTrainloop("tl-1234567890")

	messages := []trainloop.Message{
		{Role: "user", Content: "Hello, from the user"},
		{Role: "assistant", Content: "Hello, from the assistant"},
	}

	success, err := client.SendData(messages, trainloop.GoodFeedback, "example-dataset-id")
	if err != nil {
		log.Printf("Failed to send data to TrainLoop: %v", err)
		return
	}

	if success {
		fmt.Println("Data sent successfully!")
	} else {
		fmt.Println("Data was not sent.")
	}
}
