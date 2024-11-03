package lib

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"

	wms_v1 "transacai-wms/gen/wms/v1"
)

// Parameters for the GraphQL request to create a new request
type CreateRequestGraphQLRequest struct {
    Query     string                 `json:"query"`
    Variables map[string]interface{} `json:"variables"`
}

// Response from the GraphQL request to create a new request
type CreateRequestGraphQLResponse struct {
    Data struct {
        CreateRequest struct {
            ID string `json:"id"`
        } `json:"createRequest"`
    } `json:"data"`
}

// CreateRequest submits a new request in the RSS GraphQL API
// Returns the request ID if successful, or an error if the request fails
func CreateRequest(params *wms_v1.GenerateInsightsRequest) (string, error) {
		// get RSS URL from environment variable
    url := os.Getenv("RSS_URL")
		if url == "" {
			log.Printf("RSS URL not found")
			return "", errors.New("RSS URL not found")
		}

		// GraphQL query to create a new request and get back the request ID
    query := `
        mutation createNewRequest($data: RequestCreateInput!) {
            createRequest(data: $data) {
                id
            }
        }
    `

		// Prepare request body variables
    variables := map[string]interface{}{
        "data": map[string]interface{}{
            "clientId":                params.ClientId,
            "promptId":                params.PromptId,
            "recordsSourceId":         params.RecordsSourceId,
            "promptTemplatesSourceId": params.PromptTemplatesSourceId,
            "fromTime":                params.FromTime,
            "toTime":                  params.ToTime,
        },
    }

		// Prepare the request body with query and variables
    requestBody, err := json.Marshal(CreateRequestGraphQLRequest{
        Query:     query,
        Variables: variables,
    })
    if err != nil {
				log.Printf("Error marshalling request body: %v", err)
        return "", errors.New("failed to create request")
    }

		// Create the HTTP request
    req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
    if err != nil {
				log.Printf("Error creating request: %v", err)
				return "", errors.New("failed to create request")
    }

		// Prepare request headers
    req.Header.Set("Content-Type", "application/json")

		// Get TRANSAC_AI_RSS_API_KEY from environment variable
		apiKey := os.Getenv("TRANSAC_AI_RSS_API_KEY")
		if apiKey == "" {
			log.Printf("RSS API key not found")
			return "", errors.New("RSS API key not found")
		}
		// Set the Authorization header with the API key
		req.Header.Set("Authorization", "Bearer " + apiKey)

		// Send the request
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
				log.Printf("Error sending request: %v", err)
				return "", errors.New("failed to create request")
    }
    defer resp.Body.Close()

		// Check if the request was unsuccessful
    if resp.StatusCode != http.StatusOK {
				log.Printf("Request failed with status: %v", resp.Status)
        return "", errors.New("failed to complete request")
    }

		// Parse the response body
    var response CreateRequestGraphQLResponse
    if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
				log.Printf("Error decoding response: %v", err)
				return "", errors.New("failed to parse response")
    }

		// Return the request ID
    return response.Data.CreateRequest.ID, nil
}