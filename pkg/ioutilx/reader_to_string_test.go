package ioutilx_test

import (
	"strings"
	"testing"

	"github.com/pasdam/go-io-utilx/pkg/ioutilx"
	"github.com/stretchr/testify/assert"
)

func Test_ReaderToString(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "some-value"},
		{name: "some-other-value"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reader := strings.NewReader(tt.name)
			assert.Equal(t, tt.name, ioutilx.ReaderToString(reader))
		})
	}
}
