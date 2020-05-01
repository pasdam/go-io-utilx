package ioutilx

import (
	"bytes"
	"io"
)

// ReaderToString reads from the specified Reader and returns the content as a
// string
func ReaderToString(reader io.Reader) string {
	var buffer bytes.Buffer
	b := make([]byte, 8)
	for {
		n, err := reader.Read(b)
		buffer.Write(b[:n])
		if err == io.EOF {
			break
		}
	}
	return buffer.String()
}
