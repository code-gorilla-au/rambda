package rambda

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRespondJSONWith_should_return_hello_world_and_no_headers(t *testing.T) {
	payload := map[string]string{
		"hello": "world",
	}

	resp, err := RespondJSONWith(http.StatusOK, payload, nil)
	assert.NoError(t, err)
	assert.Equal(t, resp.Body, "{\"hello\":\"world\"}")
}

func TestRespondJSONWith_should_return_hello_world_and_no_content_type_headers(t *testing.T) {
	payload := map[string]string{
		"hello": "world",
	}

	headers := map[string]string{
		"Content-Type": "application/json",
	}

	resp, err := RespondJSONWith(http.StatusOK, payload, headers)
	assert.NoError(t, err)
	assert.Equal(t, resp.Body, "{\"hello\":\"world\"}")
	assert.Equal(t, headers, resp.Headers)
}

func Test_respondJSONSimple(t *testing.T) {
	payload := "flash"

	resp := respondJSONSimple(http.StatusOK, payload, nil)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, "{\"message\": \"flash\"}", resp.Body)
}

func Test_respondJSONSimple_should_return_flash_and_no_content_type_headers(t *testing.T) {
	payload := "flash"

	headers := map[string]string{
		"Content-Type": "application/json",
	}

	resp := respondJSONSimple(http.StatusOK, payload, headers)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, "{\"message\": \"flash\"}", resp.Body)
	assert.Equal(t, headers, resp.Headers)
}

func Test_respondError(t *testing.T) {
	payload := EnvelopeError{
		Title:  "foo",
		Detail: "bar",
		Type:   "baz",
	}

	resp := respondError(http.StatusOK, payload, nil)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, "{\"detail\":\"bar\",\"title\":\"foo\",\"type\":\"baz\"}", resp.Body)
}

func Test_respondError_should_return_flash_and_no_content_type_headers(t *testing.T) {
	payload := EnvelopeError{
		Title:  "foo",
		Detail: "bar",
		Type:   "baz",
	}

	headers := map[string]string{
		"Content-Type": "application/json",
	}

	resp := respondError(http.StatusOK, payload, headers)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, "{\"detail\":\"bar\",\"title\":\"foo\",\"type\":\"baz\"}", resp.Body)
	assert.Equal(t, headers, resp.Headers)
}

func Test_respondError_marshal_error_should_respond_generic_template(t *testing.T) {
	payload := EnvelopeError{
		Title:  "foo",
		Detail: "bar",
		Type:   "baz",
	}

	headers := map[string]string{
		"Content-Type": "application/json",
	}

	resp := respondError(http.StatusOK, payload, headers)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, "{\"detail\":\"bar\",\"title\":\"foo\",\"type\":\"baz\"}", resp.Body)
	assert.Equal(t, headers, resp.Headers)
}

func TestRespondOK(t *testing.T) {
	resp := RespondOK("", nil)
	assert.Equal(t, "{\"message\": \"ok\"}", resp.Body)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestRespondOK_custom_message(t *testing.T) {
	resp := RespondOK("custom message", nil)
	assert.Equal(t, "{\"message\": \"custom message\"}", resp.Body)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestRespondOK_headers(t *testing.T) {
	headers := map[string]string{
		"Content-Type": "application/json",
	}
	resp := RespondOK("", headers)
	assert.Equal(t, "{\"message\": \"ok\"}", resp.Body)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}
