package lib

import (
	"context"
	"crypto/tls"
	"errors"
	"log"
	"net"
	"net/http"
	"os"
	"time"
	igs_v1 "transacai-wms/gen/igs/v1"
	"transacai-wms/gen/igs/v1/igs_v1connect"

	"connectrpc.com/connect"
	"golang.org/x/net/http2"
)

func newInsecureClient() *http.Client {
  return &http.Client{
    Transport: &http2.Transport{
      AllowHTTP: true,
			DialTLSContext: func(ctx context.Context, network, addr string, cfg *tls.Config) (net.Conn, error) {
				return net.Dial(network, addr)
			},
      // Timeouts
			IdleConnTimeout: 30 * time.Second,
			PingTimeout: 30 * time.Second,
			ReadIdleTimeout: 30 * time.Second,
			WriteByteTimeout: 30 * time.Second,
    },
		// Disable redirects
		CheckRedirect: func(_ *http.Request, _ []*http.Request) error {
			return http.ErrUseLastResponse
		},
  }
}

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

	log.Printf("Initializing gRPC client to connect to IGS service")
	// Setup gRPC client to connect to IGS service
	client := igs_v1connect.NewIGSServiceClient(
		newInsecureClient(),
		IGS_URL,
		connect.WithGRPC(),
	)

	// Prepare parameters for the insights generation request
	req := connect.NewRequest(generateInsightsParams)
	// get IGS api key
	igsApiKey := os.Getenv("TRANSAC_AI_IGS_API_KEY")
	// validate api key
	if igsApiKey == "" {
		log.Printf("IGS API key not found")
		return false, errors.New("IGS API key not found")
	}
	// Set the authorization header
	req.Header().Set("Authorization", "Bearer " + igsApiKey)

	log.Printf("Submitting insights generation request to IGS")
	// Create a new request to generate insights
	res, err := client.GenerateInsights(
		context.Background(),
		req,
	)
	if err != nil {
		log.Printf("Error generating insights: %v", err)
		return false, err
	}

	// Return the request ID
	return res.Msg.Received, nil
}