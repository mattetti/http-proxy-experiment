package main

import (
  "io"
  "os"
)

// There is probably already something doing that in a std package
// but oh well, i'm learning :) 

//
type ChunkedWriter struct {
	Wire io.Writer
}

// implement the Writer interface
func (cw *ChunkedWriter) Write(data []byte) (n int, err os.Error) {
	// Don't send 0-length data. It looks like EOF for chunked encoding.
	if len(data) == 0 {
		return 0, nil
	}

	if n, err = cw.Wire.Write(data); err != nil {
		return
	}
	if n != len(data) {
		err = io.ErrShortWrite
		return
	}

	return
}

// implement the Writer interface
func (cw *ChunkedWriter) Close() os.Error {
	_, err := io.WriteString(cw.Wire, "0\r\n")
	return err
}
