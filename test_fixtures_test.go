package rambda

import (
	"fmt"
	"net/http"
)

func expectedDefaultResponse(message string) string {
	return fmt.Sprintf("{\"message\": \"%s\"}", message)
}
func expectedDefaultErrorResponse(status int, message string) string {
	return fmt.Sprintf("{\"detail\":\"%s\",\"title\":\"%s\"}", message, http.StatusText(status))
}
