# TrainLoop SDK

This is the SDK for TrainLoop. It allows you to send data to TrainLoop to be used for finetuning.

## Installation

### Python

```bash
pip install trainloop-sdk
```

### Go

```bash
go get github.com/TrainLoop/sdk@v0.0.3
```

## Usage

### Go

```go
package main

// 1. Import the SDK
import (
	trainloop "github.com/TrainLoop/sdk/go"
)

// 2. Get the api key from the environment
apiKey := os.Getenv("TRAINLOOP_API_KEY")

if apiKey == "" {
    log.Fatal("TRAINLOOP_API_KEY environment variable is not set")
}

// 3. Initialize the SDK
client := trainloop.NewTrainloop(apiKey)

// 4. Do some inference and generate some message thread
messages := []trainloop.Message{
    {Role: "user", Content: "Hello, from the user"},
    {Role: "assistant", Content: "Hello, from the assistant"},
}

// 3. Send data to TrainLoop (good or bad)
success, err := client.SendData(messages, trainloop.GoodFeedback, "example-dataset-id")
if err != nil {
    log.Printf("Failed to send data to TrainLoop: %v", err)
    return
}

if success {
    log.Printf("Data sent to TrainLoop successfully")
}
```
