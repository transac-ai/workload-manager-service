package lib

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
)

// Parameters for the GraphQL request to update an existing request status
type UpdateRequestGraphQLRequest struct {
    Query     string                 `json:"query"`
    Variables map[string]interface{} `json:"variables"`
}

// Response from the GraphQL request to update an existing request status
type UpdateRequestGraphQLResponse struct {
    Data struct {
        UpdateRequest struct {
            Status string `json:"status"`
        } `json:"updateRequest"`
    } `json:"data"`
}

// UpdateRequestStatus updates the status of an existing request in the RSS GraphQL API
// Returns an error if the request fails
func UpdateRequestStatus(requestId string, clientId string, status string) error {
    // get RSS URL from environment variable
    url := os.Getenv("RSS_URL")
    if url == "" {
      log.Printf("RSS URL not found")
      return errors.New("RSS URL not found")
    }

    // GraphQL query to update an existing request status
    query := `
        mutation updateExistingRequest($data: RequestUpdateInput!) {
            updateRequest(data: $data) {
                status
            }
        }
    `

    // Prepare request body variables
    variables := map[string]interface{}{
        "data": map[string]interface{}{
            "id":     requestId,
            "clientId": clientId,
            "status": status,
        },
    }

    // Prepare the request body with query and variables
    requestBody, err := json.Marshal(UpdateRequestGraphQLRequest{
        Query:     query,
        Variables: variables,
    })
    if err != nil {
        log.Printf("Error marshalling request body: %v", err)
        return errors.New("failed to update request status")
    }

    // Create the HTTP request
    req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
    if err != nil {
        log.Printf("Error creating request: %v", err)
        return errors.New("failed to update request status")
    }

    // Prepare request headers
    req.Header.Set("Content-Type", "application/json")

    // Get TRANSAC_AI_RSS_API_KEY from environment variable
    apiKey := os.Getenv("TRANSAC_AI_RSS_API_KEY")
    if apiKey == "" {
        log.Printf("RSS API key not found")
        return errors.New("RSS API key not found")
    }

    // Set the Authorization header with the API key
    req.Header.Set("Authorization", "Bearer " + apiKey)

    // print request body
    log.Printf("Request body: %v", bytes.NewBuffer(requestBody))

    // Send the request
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        log.Printf("Error sending request: %v", err)
        return errors.New("failed to update request status")
    }
    defer resp.Body.Close()

    // Check if the request was unsuccessful
    if resp.StatusCode != http.StatusOK {
        log.Printf("Request failed with status: %v", resp.Status)
        return errors.New("failed to update request status")
    }

    // Parse the response body
    var response UpdateRequestGraphQLResponse
    if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
        log.Printf("Error decoding response: %v", err)
        return errors.New("failed to parse response")
    }

    // print response body
    log.Printf("Response body: %v", response)

    // Check if the status was updated successfully
    if response.Data.UpdateRequest.Status != status {
        msg := fmt.Sprintf("Failed to update request status from %s to %s", response.Data.UpdateRequest.Status, status)
        log.Printf(msg)
        return errors.New(msg)
    }

    return nil
}