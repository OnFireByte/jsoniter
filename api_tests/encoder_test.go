package test

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/onfirebyte/jsoniter"
	"github.com/stretchr/testify/require"
)

// Standard Encoder has trailing newline.
func TestEncoderHasTrailingNewline(t *testing.T) {
	should := require.New(t)
	var buf, stdbuf bytes.Buffer
	enc := jsoniter.ConfigCompatibleWithStandardLibrary.NewEncoder(&buf)
	enc.Encode(1)
	stdenc := json.NewEncoder(&stdbuf)
	stdenc.Encode(1)
	should.Equal(stdbuf.Bytes(), buf.Bytes())
}
