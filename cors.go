package rambda

import "github.com/aws/jsii-runtime-go"

var defaultCors = &map[string]*string{
	"method.response.header.Access-Control-Allow-Headers": jsii.String("'Content-Type,X-Amz-Date,Authorization,X-Api-Key,X-Amz-Security-Token,X-Amz-User-Agent'"),
	"method.response.header.Access-Control-Allow-Origin":  jsii.String("'*'"),
	"method.response.header.Access-Control-Allow-Methods": jsii.String("'OPTIONS,GET,PUT,POST,DELETE'"),
}
