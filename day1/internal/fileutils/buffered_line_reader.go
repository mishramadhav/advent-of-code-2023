package fileutils

import (
	"bufio"
	"io"
)

type BufferedFileReader struct {
	r *bufio.Scanner
}

func NewBufferedFileReader(r io.Reader) *BufferedFileReader {
	return &BufferedFileReader{
		r: bufio.NewScanner(r),
	}
}

func (f *BufferedFileReader) ReadLine() (string, error) {
	if !f.r.Scan() {
		return "", io.EOF
	}

	return f.r.Text(), f.r.Err()
}
