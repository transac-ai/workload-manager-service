package lib

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	igs_v1 "transacai-wms/gen/igs/v1"
	"transacai-wms/gen/igs/v1/igs_v1connect"

	"connectrpc.com/connect"
)

/**
 * SubmitGenerateInsightsRequest submits a new insights generation request to the IGS service
 * Returns the true if successful, or an error if the request fails
 */
func SubmitGenerateInsightsRequest(generateInsightsParams *igs_v1.GenerateInsightsRequest) (bool, error) {
	// URL to IGS service (in-cluster internal URL)
	IGS_URL := os.Getenv("IGS_URL")
	if IGS_URL == "" {
		log.Printf("IGS URL not found")
		return false, errors.New("IGS URL not found")
	}

	// Setup gRPC client to connect to IGS service
	client := igs_v1connect.NewInsightsGenerationServiceClient(
		http.DefaultClient,
		IGS_URL,
		connect.WithGRPC(),
	)

	// Create a new request to generate insights
	res, err := client.GenerateInsights(
		context.Background(),
		connect.NewRequest(generateInsightsParams),
	)
	if err != nil {
		log.Printf("Error generating insights: %v", err)
		return false, err
	}

	// Return the request ID
	return res.Msg.Received, nil
}