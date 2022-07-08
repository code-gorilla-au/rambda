package rambda

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

const (
	defaultOk            = "ok"
	defaultCreated       = "resource created"
	defaultForbidden     = "forbidden to access this resource"
	defaultUnauthorized  = "unauthorized to access this resource"
	defaultNotFound      = "resource not found"
	defaultConflict      = "there is a conflict with the resource"
	defaultBadRequest    = "bad request"
	defaultGenericServer = "generic server error occurred for this resource"
)

// RespondOK - respond ok
func RespondOK(message string, headers map[string]string) events.APIGatewayProxyResponse {
	msg := defaultOk
	if message != "" {
		msg = message
	}
	return respondJSONSimple(http.StatusOK, msg, headers)
}

// RespondCreated - respond created
func RespondCreated(message string, headers map[string]string) events.APIGatewayProxyResponse {
	msg := defaultCreated
	if message != "" {
		msg = message
	}
	merged := mergeHeaders(defaultHeaders, headers)
	return respondJSONSimple(http.StatusCreated, msg, merged)
}

// RespondBadRequest - respond bad request
func RespondBadRequest(message string, headers map[string]string) events.APIGatewayProxyResponse {
	msg := defaultBadRequest
	if message != "" {
		msg = message
	}
	merged := mergeHeaders(defaultHeaders, headers)
	return respondError(
		http.StatusBadRequest,
		EnvelopeError{
			Title:  http.StatusText(http.StatusBadRequest),
			Detail: msg,
		},
		merged,
	)
}

// RespondGenericServer - respond generic error
func RespondGenericServer(message string, headers map[string]string) events.APIGatewayProxyResponse {
	msg := defaultGenericServer
	if message != "" {
		msg = message
	}
	merged := mergeHeaders(defaultHeaders, headers)
	return respondError(
		http.StatusInternalServerError,
		EnvelopeError{
			Title:  http.StatusText(http.StatusInternalServerError),
			Detail: msg,
		},
		merged,
	)
}

// RespondConflict - respond conflict
func RespondConflict(message string, headers map[string]string) events.APIGatewayProxyResponse {
	msg := defaultConflict
	if message != "" {
		msg = message
	}
	merged := mergeHeaders(defaultHeaders, headers)
	return respondError(
		http.StatusConflict,
		EnvelopeError{
			Title:  http.StatusText(http.StatusConflict),
			Detail: msg,
		},
		merged,
	)
}

// RespondNotFound - respond not found
func RespondNotFound(message string, headers map[string]string) events.APIGatewayProxyResponse {
	msg := defaultNotFound
	if message != "" {
		msg = message
	}
	merged := mergeHeaders(defaultHeaders, headers)
	return respondError(
		http.StatusNotFound,
		EnvelopeError{
			Title:  http.StatusText(http.StatusNotFound),
			Detail: msg,
		},
		merged,
	)

}

func RespondUnAuthorised(message string, headers map[string]string) events.APIGatewayProxyResponse {
	msg := defaultUnauthorized
	if message != "" {
		msg = message
	}
	merged := mergeHeaders(defaultHeaders, headers)
	return respondError(
		http.StatusUnauthorized,
		EnvelopeError{
			Title:  http.StatusText(http.StatusUnauthorized),
			Detail: msg,
		},
		merged,
	)
}

func RespondForbidden(message string, headers map[string]string) events.APIGatewayProxyResponse {
	msg := defaultForbidden
	if message != "" {
		msg = message
	}
	merged := mergeHeaders(defaultHeaders, headers)
	return respondError(
		http.StatusForbidden,
		EnvelopeError{
			Title:  http.StatusText(http.StatusForbidden),
			Detail: msg,
		},
		merged,
	)
}

// RespondJSONWith - api gateway proxy response
func RespondJSONWith(status int, payload interface{}, headers map[string]string) (events.APIGatewayProxyResponse, error) {
	b, err := json.Marshal(&payload)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}
	merged := mergeHeaders(defaultHeaders, headers)
	return events.APIGatewayProxyResponse{
		StatusCode:      status,
		Headers:         merged,
		Body:            string(b),
		IsBase64Encoded: false,
	}, nil
}

// respondJSONSimple - respond with message and status
func respondJSONSimple(status int, message string, headers map[string]string) events.APIGatewayProxyResponse {
	merged := mergeHeaders(defaultHeaders, headers)
	return events.APIGatewayProxyResponse{
		StatusCode:      status,
		Headers:         merged,
		Body:            fmt.Sprintf("{\"message\": \"%s\"}", message),
		IsBase64Encoded: false,
	}
}

// respondError - response standardised error response message
func respondError(status int, envelope EnvelopeError, headers map[string]string) events.APIGatewayProxyResponse {
	merged := mergeHeaders(defaultHeaders, headers)
	d, err := json.Marshal(envelope)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode:      status,
			Headers:         merged,
			Body:            fmt.Sprintf("{ \"title\": \"%s\" \"detail\": \"%s\"}", http.StatusText(status), envelope.Detail),
			IsBase64Encoded: false,
		}
	}

	return events.APIGatewayProxyResponse{
		StatusCode:      status,
		Headers:         merged,
		Body:            string(d),
		IsBase64Encoded: false,
	}
}
