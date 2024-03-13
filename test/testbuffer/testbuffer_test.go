package testbuffer

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestBuffer is to help with testing and not really needs to be tested itself. However
func TestRead(t *testing.T) {
	tests := []struct {
		name        string
		writeString []byte
		expectedLen int
	}{
		{
			name:        "testOne",
			writeString: []byte("This is the test string"),
			expectedLen: 23,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Filename = fmt.Sprintf("/tmp/testbuffer-%d", rand.Int())
			buf := TestBuffer{}
			length, err := buf.Write(tt.writeString)
			assert.NoError(t, err, "error writing to buffer")
			assert.Equal(t, tt.expectedLen, length)

			val, err := buf.Read()
			assert.NoError(t, err, "error reading from buffer")
			assert.Equal(t, tt.writeString, val)
		})
	}
}
