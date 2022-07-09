package rambda

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetDefaultHeaders_should_set_headers(t *testing.T) {

	expected := map[string]string{
		"header": "brand-new",
	}

	SetDefaultHeaders(expected)
	assert.Equal(t, expected, defaultResponseHeaders)
}
