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
	defaultForbidden     = "forbiden to access this resource"
	defaultUnauthorised  = "unauthorised to access this resource"
	defaultNotFound      = "resource not found"
	defaultConflict      = "there is a conflict with the resource"
	defaultBadRequest    = "bad request"
	defaultGenericServer = "generic server error"
)

// RespondOK - respond ok
func RespondOK(message *string, headers *map[string]string) events.APIGatewayProxyResponse {
	msg := defaultOk
	if message != nil {
		msg = *message
	}
	return respondJSONSimple(http.StatusOK, msg, headers)
}

// RespondCreated - respond created
func RespondCreated(message *string, headers *map[string]string) events.APIGatewayProxyResponse {
	msg := defaultOk
	if message != nil {
		msg = *message
	}
	return respondJSONSimple(http.StatusCreated, msg, headers)
}

// RespondBadRequest - respond bad request
func RespondBadRequest(message *string, headers *map[string]string) events.APIGatewayProxyResponse {
	msg := defaultBadRequest
	if message != nil {
		msg = *message
	}
	return respondError(
		http.StatusBadRequest,
		EnvelopeError{
			Title:  http.StatusText(http.StatusBadRequest),
			Detail: msg,
		},
		headers,
	)
}

// RespondGenericServer - respond generic error
func RespondGenericServer(message *string, headers *map[string]string) events.APIGatewayProxyResponse {
	msg := defaultGenericServer
	if message != nil {
		msg = *message
	}
	return respondError(
		http.StatusInternalServerError,
		EnvelopeError{
			Title:  http.StatusText(http.StatusInternalServerError),
			Detail: msg,
		},
		headers,
	)
}

// RespondConflict - respond conflict
func RespondConflict(message *string, headers *map[string]string) events.APIGatewayProxyResponse {
	msg := defaultConflict
	if message != nil {
		msg = *message
	}
	return respondError(
		http.StatusConflict,
		EnvelopeError{
			Title:  http.StatusText(http.StatusConflict),
			Detail: msg,
		},
		headers,
	)
}

// RespondNotFound - respond not found
func RespondNotFound(message *string, headers *map[string]string) events.APIGatewayProxyResponse {
	msg := defaultNotFound
	if message != nil {
		msg = *message
	}
	return respondError(
		http.StatusNotFound,
		EnvelopeError{
			Title:  http.StatusText(http.StatusNotFound),
			Detail: msg,
		},
		headers)
}

func RespondUnAuthorised(message *string, headers *map[string]string) events.APIGatewayProxyResponse {
	msg := defaultUnauthorised
	if message != nil {
		msg = *message
	}
	return respondError(
		http.StatusUnauthorized,
		EnvelopeError{
			Title:  http.StatusText(http.StatusUnauthorized),
			Detail: msg,
		},
		headers,
	)
}

func RespondForbidden(message *string, headers *map[string]string) events.APIGatewayProxyResponse {
	msg := defaultForbidden
	if message != nil {
		msg = *message
	}
	return respondError(
		http.StatusForbidden,
		EnvelopeError{
			Title:  http.StatusText(http.StatusForbidden),
			Detail: msg,
		},
		headers,
	)
}

// RespondJSONWith - api gateway proxy response
func RespondJSONWith(status int, payload interface{}, headers *map[string]string) (events.APIGatewayProxyResponse, error) {
	b, err := json.Marshal(&payload)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}
	return events.APIGatewayProxyResponse{
		StatusCode:      status,
		Headers:         *headers,
		Body:            string(b),
		IsBase64Encoded: false,
	}, nil
}

// respondJSONSimple - respond with message and status
func respondJSONSimple(status int, message string, headers *map[string]string) events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{
		StatusCode:      status,
		Headers:         *headers,
		Body:            fmt.Sprintf("{\"message\": \"%s\"}", message),
		IsBase64Encoded: false,
	}
}

func respondError(status int, envelope EnvelopeError, headers *map[string]string) events.APIGatewayProxyResponse {
	d, err := json.Marshal(envelope)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode:      status,
			Headers:         *headers,
			Body:            fmt.Sprintf("{ \"title\": \"%s\" \"detail\": \"%s\"}", http.StatusText(status), envelope.Detail),
			IsBase64Encoded: false,
		}
	}
	return events.APIGatewayProxyResponse{
		StatusCode:      status,
		Headers:         *headers,
		Body:            string(d),
		IsBase64Encoded: false,
	}
}
