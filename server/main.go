package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	api "transacai-wms/api"
	wms_v1 "transacai-wms/gen/wms/v1"
	"transacai-wms/gen/wms/v1/wms_v1connect"
	"transacai-wms/utils"

	"connectrpc.com/connect"
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

func main() {
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

