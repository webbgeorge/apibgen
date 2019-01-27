package apib

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

var testResponseBody = `{"foo": "bar"}`

func TestNewResponse_ResponseBodyIsCorrectlyCopied(t *testing.T) {
	body := bytes.NewBuffer([]byte(testRequestBody))
	req, err := http.NewRequest("POST", "http://httpbin.org/post", body)
	assert.Nil(t, err)

	w := httptest.NewRecorder()
	testResponseHandler(w, req)

	apiResp := NewResponse(w)

	assert.Equal(t, testResponseBody, string(apiResp.Body.Content))
}

func TestNewResponse_OriginalResponseBodyDoesNotChange(t *testing.T) {
	body := bytes.NewBuffer([]byte(testRequestBody))
	req, err := http.NewRequest("POST", "http://httpbin.org/post", body)
	assert.Nil(t, err)

	w := httptest.NewRecorder()
	testResponseHandler(w, req)

	_ = NewResponse(w)

	assert.Equal(t, testRequestBody, w.Body.String())
}

func testResponseHandler(w http.ResponseWriter, _ *http.Request) {
	_, _ = io.WriteString(w, testResponseBody)
}
