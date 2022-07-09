package rambda

var defaultResponseHeaders map[string]string = map[string]string{}

// SetDefaultHeaders - set default headers that are applied to all lambda responses
func SetDefaultHeaders(headers map[string]string) {
	defaultResponseHeaders = headers
}

// mergeHeaders - merge headers, this action will overwrite a previously define header.
func mergeHeaders(headersList ...map[string]string) map[string]string {
	mergedHeaders := map[string]string{}
	for _, headers := range headersList {
		for key, value := range headers {
			mergedHeaders[key] = value
		}
	}
	return mergedHeaders
}
