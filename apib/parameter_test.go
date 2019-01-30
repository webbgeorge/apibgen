package apib

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuoteParameterValue(t *testing.T) {
	val := ParameterValue("param-value")

	quotedVal := val.Quote()
	assert.Equal(t, len(quotedVal), len(val)+2)
	assert.True(t, strings.Contains(quotedVal, string(val)))
	assert.Equal(t, quotedVal[len(quotedVal)-1], quotedVal[0])
}

func TestQuoteParameterValue_EmptyValue(t *testing.T) {
	var val ParameterValue

	quotedVal := val.Quote()
	assert.Equal(t, quotedVal, string(val))
}
