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

type Dumper struct {
	http.ResponseWriter

	mw  io.Writer
	buf *bytes.Buffer
}

func NewDumper(resp *echo.Response) *Dumper {
	buf := new(bytes.Buffer)

	return &Dumper{
		ResponseWriter: resp.Writer,

		mw:  io.MultiWriter(resp.Writer, buf),
		buf: buf,
	}
}

func (d *Dumper) Write(b []byte) (int, error) {
	nBytes, err := d.mw.Write(b)
	if err != nil {
		err = fmt.Errorf("error writing response: %w", err)
	}

	return nBytes, err
}

func (d *Dumper) GetResponse() string {
	return d.buf.String()
}

func (d *Dumper) Flush() {
	if flusher, ok := d.ResponseWriter.(http.Flusher); ok {
		flusher.Flush()
	}
}

func (d *Dumper) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	if hijacker, ok := d.ResponseWriter.(http.Hijacker); ok {
		conn, rw, err := hijacker.Hijack()
		if err != nil {
			err = fmt.Errorf("error hijacking response: %w", err)
		}

		return conn, rw, err
	}

	return nil, nil, nil
}
