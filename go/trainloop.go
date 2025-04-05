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
	BaseURL    string
	HTTPClient *http.Client
}

// NewClient initializes a new TrainLoop client with an API key
func NewClient(apiKey string, baseURL ...string) *Client {
	url := "https://app.trainloop.ai"
	if len(baseURL) > 0 {
		url = baseURL[0]
	}

	transport := &http.Transport{
		MaxIdleConns:        100,
		MaxIdleConnsPerHost: 100,
		MaxConnsPerHost:     100,
		IdleConnTimeout:     90 * time.Second,
	}

	return &Client{
		APIKey:  apiKey,
		BaseURL: url,
		HTTPClient: &http.Client{
			Timeout:   10 * time.Second,
			Transport: transport,
		},
	}
}

// Message represents a message with string keys and string values.
type Message map[string]string

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
	req, err := http.NewRequest("POST", t.BaseURL+"/api/datasets/collect", bytes.NewBuffer(body))
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
