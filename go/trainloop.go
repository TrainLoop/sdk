package trainloop

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Client represents the configuration needed to authenticate and send data
type Client struct {
	APIKey     string
	HTTPClient *http.Client
}

// NewClient initializes a new TrainLoop client with an API key
func NewClient(apiKey string) *Client {
	transport := &http.Transport{
		MaxIdleConns:        100,
		MaxIdleConnsPerHost: 100,
		MaxConnsPerHost:     100,
		IdleConnTimeout:     90 * time.Second,
	}

	return &Client{
		APIKey: apiKey,
		HTTPClient: &http.Client{
			Timeout:   10 * time.Second,
			Transport: transport,
		},
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
// Returns nil on success, or an error if something went wrong.
func (t *Client) SendData(messages []Message, sampleFeedback SampleFeedbackType, datasetID string) error {
	// Construct the payload
	payload := requestPayload{
		Messages:       messages,
		SampleFeedback: sampleFeedback,
		DatasetID:      datasetID,
	}

	// Convert payload to JSON
	body, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	// Create the request
	req, err := http.NewRequest("POST", "https://app.trainloop.ai/api/datasets/collect", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	// Set appropriate headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+t.APIKey)

	// Send the request using the stored HTTPClient
	resp, err := t.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check for successful status code
	if resp.StatusCode != http.StatusOK {
		// Read the response body for more detailed error message
		bodyBytes, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("request returned non-200 status code: %s, body: %s", resp.Status, string(bodyBytes))
	}

	return nil
}
