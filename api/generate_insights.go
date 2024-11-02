package api

import (
	"log"
	wms_v1 "transacai-wms/gen/wms/v1"
)

func GenerateInsightsHandler(
	params *wms_v1.GenerateInsightsRequest,
) (*wms_v1.GenerateInsightsResponse, error) {
	log.Printf("Received GenerateInsights request: %v", params)
	res := &wms_v1.GenerateInsightsResponse{
		RequestId: "test_request_id",
	}
	return res, nil
}