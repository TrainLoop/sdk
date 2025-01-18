package trainloop

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

// Client represents the configuration needed to authenticate and send data
type Client struct {
	APIKey string
}

// NewClient initializes a new Trainloop client with an API key
func NewClient(apiKey string) *Client {
	return &Client{
		APIKey: apiKey,
	}
}

// Message represents the messages you’re sending to TrainLoop.
// Adjust fields as required by your system.
type Message struct {
	Role    string      `json:"role"`
	Content interface{} `json:"content"`
}

type SampleFeedbackType string

const (
	GoodFeedback SampleFeedbackType = "good"
	BadFeedback  SampleFeedbackType = "bad"
)

// requestPayload models the JSON body of the POST request
type requestPayload struct {
	Messages       []Message          `json:"messages"`
	SampleFeedback SampleFeedbackType `json:"sample_feedback"`
	DatasetID      string             `json:"dataset_id"`
}

// SendData sends your data to the TrainLoop API.
// Returns true on success, or false plus an error if something went wrong.
func (t *Client) SendData(messages []Message, sampleFeedback SampleFeedbackType, datasetID string) (bool, error) {
	// Construct the payload
	payload := requestPayload{
		Messages:       messages,
		SampleFeedback: sampleFeedback,
		DatasetID:      datasetID,
	}

	// Convert payload to JSON
	body, err := json.Marshal(payload)
	if err != nil {
		return false, fmt.Errorf("could not marshal payload: %w", err)
	}

	// Create the request
	req, err := http.NewRequest("POST", "https://app.trainloop.ai/api/datasets/collect", bytes.NewBuffer(body))
	if err != nil {
		return false, fmt.Errorf("could not create request: %w", err)
	}

	// Set appropriate headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+t.APIKey)

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		return false, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	// Check for successful status code
	if resp.StatusCode != http.StatusOK {
		return false, errors.New("request returned non-200 status code: " + resp.Status)
	}

	return true, nil
}
