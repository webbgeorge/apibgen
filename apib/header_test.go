package apib

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContentType_OneContentType(t *testing.T) {
	ct := "text/plain"

	h := http.Header{}
	h.Add("Content-Type", ct)
	assert.Equal(t, ct, NewHeader(h).ContentType)
}

func TestContentType_MultipleContentTypes_Mistakenly(t *testing.T) {
	ct := "text/plain"

	h := http.Header{}
	h.Add("Content-Type", ct)
	h.Add("Content-Type", "application/json")
	assert.Equal(t, ct, NewHeader(h).ContentType)
}

func TestNewHeader_EmptyHeader(t *testing.T) {
	h := http.Header{}
	assert.Nil(t, NewHeader(h))
}
