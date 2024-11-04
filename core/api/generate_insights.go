package api

import (
	"log"
	igs_v1 "transacai-wms/gen/igs/v1"
	wms_v1 "transacai-wms/gen/wms/v1"
	igs "transacai-wms/services/igs"
	rss "transacai-wms/services/rss"
)

func GenerateInsightsHandler(
	params *wms_v1.GenerateInsightsRequest,
) (*wms_v1.GenerateInsightsResponse, error) {
	log.Printf("Received GenerateInsights request: %v", params)
	// submit a new request to the RSS service and retrieve 
	// the request ID to be returned to the client
	requestId, err := rss.CreateRequest(params)
	if err != nil {
		log.Printf("Failed to submit new request: %v", err)
		return nil, err
	}
	// prepare response with the request ID
	res := &wms_v1.GenerateInsightsResponse{
		RequestId: requestId,
	}
	log.Printf("Successfully submitted new request: %v", requestId)
	// dispatch a goroutine to submit the insights generation request to IGS service
	go func() {
		// prepare parameters for the insights generation request
		igsParams := &igs_v1.GenerateInsightsRequest{
			ReqId: requestId,
			ClientId: params.ClientId,
			PromptId: params.PromptId,
			RecordsSourceId: params.RecordsSourceId,
			PromptTemplatesSourceId: params.PromptTemplatesSourceId,
			FromTime: params.FromTime,
			ToTime: params.ToTime,
		}
		// submit a new request to the IGS service
		success, err := igs.SubmitGenerateInsightsRequest(igsParams)
		if err != nil {
			log.Printf("Failed to submit insights generation request: %v", err)
			return
		}
		if success {
			log.Printf("Successfully submitted insights generation request")
			// update the request status in the RSS service
			err = rss.UpdateRequestStatus(requestId, params.ClientId, "PROCESSING")
		}

	}()
	
	return res, nil
}