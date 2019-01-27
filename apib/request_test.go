package apib

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

var testRequestBody = `{"foo": "bar"}`

func TestNewRequest_RequestBodyIsCorrectlyCopied(t *testing.T) {
	body := bytes.NewBuffer([]byte(testRequestBody))
	req, err := http.NewRequest("POST", "http://httpbin.org/post", body)
	assert.Nil(t, err)

	apiReq, err := NewRequest(req)
	assert.Nil(t, err)

	assert.Equal(t, testRequestBody, string(apiReq.Body.Content))
}

func TestNewRequest_OriginalRequestBodyDoesNotChange(t *testing.T) {
	body := bytes.NewBuffer([]byte(testRequestBody))
	req, err := http.NewRequest("POST", "http://httpbin.org/post", body)
	assert.Nil(t, err)

	_, err = NewRequest(req)
	assert.Nil(t, err)

	httpReqBody, err := ioutil.ReadAll(req.Body)
	assert.Nil(t, err)
	assert.Equal(t, testRequestBody, string(httpReqBody))
}
