package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"transacai-wms/core/api"
	wms_v1 "transacai-wms/gen/wms/v1"
	"transacai-wms/gen/wms/v1/wms_v1connect"
	"transacai-wms/utils"

	"connectrpc.com/connect"
	"github.com/joho/godotenv"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

type WMSServer struct {}

func (s *WMSServer) GenerateInsights(
	ctx context.Context,
	req *connect.Request[wms_v1.GenerateInsightsRequest],
) (*connect.Response[wms_v1.GenerateInsightsResponse], error) {
	log.Printf("Received GenerateInsights request")
	// Validate authorization
	err := utils.ValidateAuthorization(req.Header())
	if err != nil {
		return nil, err
	}
	// Initiate insights generation request
	generateInsightsResponse, err := api.GenerateInsightsHandler(req.Msg)
	if err != nil {
		errMsg := fmt.Sprintf("Error generating insights: %v", err)
		log.Printf("Error generating insights: %v", errMsg)
		return nil, connect.NewError(connect.CodeInternal, errors.New(errMsg))
	}
	// If insights generation request successfully initiated
	// return `request_id` in response that can be used to track the request
	res := connect.NewResponse(generateInsightsResponse)
	res.Header().Set("WMS-Version", "v1")
	return res, nil
}

func confirmNecessaryEnvVars() error {
	// Check for necessary environment variables
	if os.Getenv("RSS_URL") == "" {
		return errors.New("RSS_URL not found")
	}
	if os.Getenv("IGS_URL") == "" {
		return errors.New("IGS_URL not found")
	}
	if os.Getenv("TRANSAC_AI_RSS_API_KEY") == "" {
		return errors.New("TRANSAC_AI_RSS_API_KEY not found")
	}
	if os.Getenv("TRANSAC_AI_WMS_API_KEY") == "" {
		return errors.New("TRANSAC_AI_WMS_API_KEY not found")
	}
	if os.Getenv("TRANSAC_AI_IGS_API_KEY") == "" {
		return errors.New("TRANSAC_AI_IGS_API_KEY not found")
	}
	return nil
}

func main() {
	// get GO_ENV from environment variable
	goEnv := os.Getenv("GO_ENV")
	// if not production, load environment variables from .env file
	if goEnv != "production" {
			// resolve path to .env file
		envPath, err := filepath.Abs(".env")
		if err != nil {
      	log.Printf("Warning: error resolving absolute path: %v", err)
				return
  	}

		// Load environment variables
		err = godotenv.Load(envPath)
		if err != nil {
			log.Printf("Error loading .env file: %v", err)
			return
		}
	}

	// confirm necessary environment variables are set
	err := confirmNecessaryEnvVars()
	if err != nil {
		log.Printf("Error confirming necessary environment variables: %v", err)
		return
	}

	wmsServer := &WMSServer{}
	mux := http.NewServeMux()
	path, handler := wms_v1connect.NewWMSServiceHandler(wmsServer)
	mux.Handle(path, handler)
	wmsServerAddr := "0.0.0.0:8080"
	log.Printf("WMS server listening on %s", wmsServerAddr)
	http.ListenAndServe(
		wmsServerAddr,
		// Use h2c so we can serve HTTP/2 without TLS.
		h2c.NewHandler(mux, &http2.Server{}),
	)
}

