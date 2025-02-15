# TrainLoop SDK

This is the SDK for TrainLoop. It allows you to send data to TrainLoop to be used for finetuning.

## Installation

### Python

```bash
pip install trainloop-sdk
```

### Go

```bash
go get -u github.com/TrainLoop/sdk@v0.0.4
```

## Getting an API key

You can get an API key from the [TrainLoop dashboard](https://app.trainloop.ai/settings).

## Usage

### Python

```python
import os
from trainloop import Trainloop, SampleFeedbackType

def main():
    # Initialize the TrainLoop client with your api key
    client = Trainloop(api_key=os.getenv("TRAINLOOP_API_KEY"))

    # Some example message thread - this should be generated by your system
    messages = [
        {"role": "system", "content": "System message here"},
        {"role": "user", "content": "Hello from the user!"},
    ]

    # Send data to trainloop and let us know if this was a success case or a failure case
    success = client.send_data(
        messages, sample_feedback=SampleFeedbackType.GOOD, dataset_id="test-dataset"
    )
    if success:
        print("Data sent successfully!")
    else:
        print("Data was not sent. Check logs for details.")


if __name__ == "__main__":
    main()
```

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
client := trainloop.NewClient(apiKey)

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
