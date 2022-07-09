package rambda

import (
	"github.com/aws/jsii-runtime-go"
)

const (
	defaultCorsAllowedOrigins string = "*"
	defaultCorsAllowedMethods string = "OPTIONS,GET,PUT,POST,PATCH,DELETE"
	defaultCorsAllowedHeaders string = "Content-Type,X-Amz-Date,Authorization,X-Api-Key,X-Amz-Security-Token,X-Amz-User-Agent"
)

// DefaultHeaders - provide sane defaults for rest api headers
func DefaultHeaders() *string {
	return jsii.String(defaultCorsAllowedHeaders)
}

// DefaultMethods - provide sane defaults for rest api methods
func DefaultMethods() *string {
	return jsii.String(defaultCorsAllowedMethods)
}
