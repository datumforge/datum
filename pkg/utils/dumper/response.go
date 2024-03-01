package dumper

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"net"
	"net/http"

	echo "github.com/datumforge/echox"
)

// Dumper is a response writer that captures the response body
type Dumper struct {
	// ResponseWriter is the original response writer
	http.ResponseWriter
	// mw is the multi writer that writes to the original response writer and the buffer
	mw io.Writer
	// buf is the buffer that captures the response body
	buf *bytes.Buffer
}

// NewDumper returns a new Dumper
func NewDumper(resp *echo.Response) *Dumper {
	buf := new(bytes.Buffer)

	return &Dumper{
		ResponseWriter: resp.Writer,
		// multi
		mw:  io.MultiWriter(resp.Writer, buf),
		buf: buf,
	}
}

// Write writes the response body
func (d *Dumper) Write(b []byte) (int, error) {
	// Write to the original response writer and the buffer
	nBytes, err := d.mw.Write(b)
	if err != nil {
		err = fmt.Errorf("error writing response: %w", err)
	}

	return nBytes, err
}

// GetResponse returns the response body out of the buffer
func (d *Dumper) GetResponse() string {
	// Return the response body out of the buffer
	return d.buf.String()
}

// Flush flushes the response writer if it implements http.Flusher interface and is not nil
func (d *Dumper) Flush() {
	if flusher, ok := d.ResponseWriter.(http.Flusher); ok {
		// Flush the response writer
		flusher.Flush()
	}
}

// Hijack hijacks the response writer and returns the connection and read writer
func (d *Dumper) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	// Hijack the response writer and return the connection and read writer; does not work with HTTP/2 so needs to be checked
	if hijacker, ok := d.ResponseWriter.(http.Hijacker); ok {
		conn, rw, err := hijacker.Hijack()

		if err != nil {
			err = fmt.Errorf("error hijacking response: %w", err)
		}
		// close the connection
		defer conn.Close()

		return conn, rw, err
	}

	return nil, nil, nil
}
