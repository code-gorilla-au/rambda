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
	}

	resp := respondError(http.StatusOK, payload, nil)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, "{\"detail\":\"bar\",\"title\":\"foo\"}", resp.Body)
}

func Test_respondError_should_return_flash_and_no_content_type_headers(t *testing.T) {
	payload := EnvelopeError{
		Title:  "foo",
		Detail: "bar",
	}

	headers := map[string]string{
		"Content-Type": "application/json",
	}

	resp := respondError(http.StatusOK, payload, headers)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, "{\"detail\":\"bar\",\"title\":\"foo\"}", resp.Body)
	assert.Equal(t, headers, resp.Headers)
}

func Test_respondError_marshal_error_should_respond_generic_template(t *testing.T) {
	payload := EnvelopeError{
		Title:  "foo",
		Detail: "bar",
	}

	headers := map[string]string{
		"Content-Type": "application/json",
	}

	resp := respondError(http.StatusOK, payload, headers)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, "{\"detail\":\"bar\",\"title\":\"foo\"}", resp.Body)
	assert.Equal(t, headers, resp.Headers)
}

func TestRespondOK(t *testing.T) {
	resp := RespondOK("", nil)
	assert.Equal(t, expectedDefaultResponse(defaultOk), resp.Body)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestRespondOK_custom_message(t *testing.T) {
	resp := RespondOK("custom message", nil)
	assert.Equal(t, expectedDefaultResponse("custom message"), resp.Body)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestRespondOK_headers(t *testing.T) {
	headers := map[string]string{
		"Content-Type": "application/json",
	}
	resp := RespondOK("", headers)
	assert.Equal(t, expectedDefaultResponse(defaultOk), resp.Body)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestRespondCreated(t *testing.T) {
	resp := RespondCreated("", nil)
	assert.Equal(t, expectedDefaultResponse(defaultCreated), resp.Body)
	assert.Equal(t, http.StatusCreated, resp.StatusCode)
}

func TestRespondCreated_custom_message(t *testing.T) {
	resp := RespondCreated("custom message", nil)
	assert.Equal(t, "{\"message\": \"custom message\"}", resp.Body)
	assert.Equal(t, http.StatusCreated, resp.StatusCode)
}

func TestRespondCreated_headers(t *testing.T) {
	headers := map[string]string{
		"Content-Type": "application/json",
	}
	resp := RespondCreated("", headers)
	assert.Equal(t, expectedDefaultResponse(defaultCreated), resp.Body)
	assert.Equal(t, http.StatusCreated, resp.StatusCode)
}

func TestRespondBadRequest(t *testing.T) {
	resp := RespondBadRequest("", nil)
	assert.Equal(t, expectedDefaultErrorResponse(http.StatusBadRequest, defaultBadRequest), resp.Body)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

func TestRespondBadRequest_custom_message(t *testing.T) {
	resp := RespondBadRequest("custom message", nil)
	assert.Equal(t, expectedDefaultErrorResponse(http.StatusBadRequest, "custom message"), resp.Body)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

func TestRespondBadRequest_headers(t *testing.T) {
	headers := map[string]string{
		"Content-Type": "application/json",
	}
	resp := RespondBadRequest("", headers)
	assert.Equal(t, expectedDefaultErrorResponse(http.StatusBadRequest, defaultBadRequest), resp.Body)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

func TestRespondGenericServer(t *testing.T) {
	resp := RespondGenericServer("", nil)
	assert.Equal(t, expectedDefaultErrorResponse(http.StatusInternalServerError, defaultGenericServer), resp.Body)
	assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
}

func TestRespondGenericServer_custom_message(t *testing.T) {
	resp := RespondGenericServer("custom message", nil)
	assert.Equal(t, expectedDefaultErrorResponse(http.StatusInternalServerError, "custom message"), resp.Body)
	assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
}

func TestRespondGenericServer_headers(t *testing.T) {
	headers := map[string]string{
		"Content-Type": "application/json",
	}
	resp := RespondGenericServer("", headers)
	assert.Equal(t, expectedDefaultErrorResponse(http.StatusInternalServerError, defaultGenericServer), resp.Body)
	assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
}

func TestRespondConflict(t *testing.T) {
	resp := RespondConflict("", nil)
	assert.Equal(t, expectedDefaultErrorResponse(http.StatusConflict, defaultConflict), resp.Body)
	assert.Equal(t, http.StatusConflict, resp.StatusCode)
}

func TestRespondConflict_custom_message(t *testing.T) {
	resp := RespondConflict("custom message", nil)
	assert.Equal(t, expectedDefaultErrorResponse(http.StatusConflict, "custom message"), resp.Body)
	assert.Equal(t, http.StatusConflict, resp.StatusCode)
}

func TestRespondConflict_headers(t *testing.T) {
	headers := map[string]string{
		"Content-Type": "application/json",
	}
	resp := RespondConflict("", headers)
	assert.Equal(t, expectedDefaultErrorResponse(http.StatusConflict, defaultConflict), resp.Body)
	assert.Equal(t, http.StatusConflict, resp.StatusCode)
}
