package rambda

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

type LambdaEventFn = func(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error)

var defaultHeaders = map[string]string{
	"Content-Type":                "application/json",
	"Access-Control-Allow-Origin": "*",
}

var defaultForbidden = "forbiden to access this resource"
var defaultUnauthorised = "unauthorised to access this resource"
var defaultNotFound = "resource not found"
var defaultConflict = "there is a conflict with the resource"
var defaultBadRequest = "bad request"
var defaultGenericServer = "generic server error"

// RespondOK - respond ok
func RespondOK(message string) events.APIGatewayProxyResponse {
	return respondJSONSimple(http.StatusOK, message)
}

// RespondCreated - respond created
func RespondCreated(message string) events.APIGatewayProxyResponse {
	return respondJSONSimple(http.StatusCreated, message)
}

// RespondBadRequest - respond bad request
func RespondBadRequest(message *string) events.APIGatewayProxyResponse {
	msg := defaultBadRequest
	if message != nil {
		msg = *message
	}
	return respondError(http.StatusBadRequest, EnvelopeError{
		Title:  http.StatusText(http.StatusBadRequest),
		Detail: msg,
	})
}

// RespondGenericServer - respond generic error
func RespondGenericServer(message *string) events.APIGatewayProxyResponse {
	msg := defaultGenericServer
	if message != nil {
		msg = *message
	}
	return respondError(http.StatusInternalServerError, EnvelopeError{
		Title:  http.StatusText(http.StatusInternalServerError),
		Detail: msg,
	})
}

// RespondConflict - respond conflict
func RespondConflict(message *string) events.APIGatewayProxyResponse {
	msg := defaultConflict
	if message != nil {
		msg = *message
	}
	return respondError(http.StatusConflict, EnvelopeError{
		Title:  http.StatusText(http.StatusConflict),
		Detail: msg,
	})
}

// RespondNotFound - respond not found
func RespondNotFound(message *string) events.APIGatewayProxyResponse {
	msg := defaultNotFound
	if message != nil {
		msg = *message
	}
	return respondError(http.StatusNotFound, EnvelopeError{
		Title:  http.StatusText(http.StatusNotFound),
		Detail: msg,
	})
}

func RespondUnAuthorised(message *string) events.APIGatewayProxyResponse {
	msg := defaultUnauthorised
	if message != nil {
		msg = *message
	}
	return respondError(http.StatusUnauthorized, EnvelopeError{
		Title:  http.StatusText(http.StatusUnauthorized),
		Detail: msg,
	})
}

func RespondForbidden(message *string) events.APIGatewayProxyResponse {
	msg := defaultForbidden
	if message != nil {
		msg = *message
	}
	return respondError(http.StatusForbidden, EnvelopeError{
		Title:  http.StatusText(http.StatusForbidden),
		Detail: msg,
	})
}

// RespondJSONWith - api gateway proxy response
func RespondJSONWith(status int, payload interface{}) (events.APIGatewayProxyResponse, error) {
	b, err := json.Marshal(&payload)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}
	return events.APIGatewayProxyResponse{
		StatusCode:      status,
		Headers:         defaultHeaders,
		Body:            string(b),
		IsBase64Encoded: false,
	}, nil
}

// respondJSONSimple - respond with message and status
func respondJSONSimple(status int, message string) events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{
		StatusCode:      status,
		Headers:         defaultHeaders,
		Body:            fmt.Sprintf("{\"message\": \"%s\"}", message),
		IsBase64Encoded: false,
	}
}

func respondError(status int, envelope EnvelopeError) events.APIGatewayProxyResponse {
	d, err := json.Marshal(envelope)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode:      status,
			Headers:         defaultHeaders,
			Body:            fmt.Sprintf("{ \"title\": \"%s\" \"detail\": \"%s\"}", http.StatusText(status), envelope.Detail),
			IsBase64Encoded: false,
		}
	}
	return events.APIGatewayProxyResponse{
		StatusCode:      status,
		Headers:         defaultHeaders,
		Body:            string(d),
		IsBase64Encoded: false,
	}
}
