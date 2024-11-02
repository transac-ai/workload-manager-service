package utils

import (
	"errors"
	"net/http"
	"os"
	"strings"

	"connectrpc.com/connect"
)

// Method to validate authorization headers (for API key validation)
func ValidateAuthorization(
	reqHeader http.Header,
) error {
	// Confirm environment variable is set
	transacaiWMSAPIKey := os.Getenv("TRANSAC_AI_WMS_API_KEY")
	if transacaiWMSAPIKey == "" {
		return connect.NewError(connect.CodeInternal, errors.New("TRANSAC_AI_WMS_API_KEY environment variable is not set"))
	}
	
	// get auth header
	authHeader := reqHeader.Get("Authorization")
	// Auth header expected format: "Bearer <API key>"
	if authHeader == "" {
		return connect.NewError(connect.CodeInvalidArgument, errors.New("Authorization header is missing"))
	}

	// Split the auth header to get the API key
	authHeaderParts := strings.Split(authHeader, " ")

	// Validate auth header format
	if len(authHeaderParts) != 2 {
		return connect.NewError(connect.CodeInvalidArgument, errors.New("Invalid Authorization header format"))
	}
	if authHeaderParts[0] != "Bearer" {
		return connect.NewError(connect.CodeInvalidArgument, errors.New("Invalid Authorization header format"))
	}

	// Validate API key
	if authHeaderParts[1] != transacaiWMSAPIKey {
		return connect.NewError(connect.CodePermissionDenied, errors.New("Invalid API key"))
	}

	return nil
}