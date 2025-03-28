package main

import (
	"fmt"
	"os"

	trainloop "github.com/TrainLoop/sdk/go"
)

func main() {
	// Get the API key from the environment variable
	apiKey := os.Getenv("TRAINLOOP_API_KEY")

	if apiKey == "" {
		fmt.Println("TRAINLOOP_API_KEY environment variable is not set")
		return
	}

	// Initialize the client with the API key
	client := trainloop.NewClient(apiKey)

	// Define the messages to send (via inference or other means)
	messages := []trainloop.Message{
		{Role: "user", Content: "Hello, from the user"},
		{Role: "assistant", Content: "Hello, from the assistant"},
	}

	// Send the data to TrainLoop
	// The dataset ID is the ID of the dataset you want to send the data to
	// The sample feedback is the feedback you want to give to the sample
	// The feedback can be either good or bad.
	err := client.SendData(messages, trainloop.GoodFeedback, "example-dataset-id")

	// If the data errored, the above function will return an error
	if err == nil {
		fmt.Println("Data sent successfully!")
	} else {
		fmt.Println("Data was not sent.")
	}
}
