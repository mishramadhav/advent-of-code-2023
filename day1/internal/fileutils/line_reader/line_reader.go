package linereader

import (
	"io"
)

// FileLineReader reads a file line by line.
type FileLineReader struct {
	// reader is the reader/stream to read from.
	reader io.ReaderAt
	// offset is the current offset in the file.
	offset int64
	// readError is the error that occurred during the last read.
	readError error
}

// NewFileReaderByLine creates a new FileReaderByLine
// for the given reader.
func NewFileReaderByLine(reader io.ReaderAt) *FileLineReader {
	return &FileLineReader{reader: reader}
}

func (f *FileLineReader) ReadLine() (string, error) {
	if f.readError != nil {
		return "", f.readError
	}

	buffer := make([]byte, 1)
	var line []byte

	for {
		// Read one byte at a time.
		_, err := f.reader.ReadAt(buffer, f.offset)
		if err != nil {
			f.readError = err
			return "", err
		}

		// Increment the offset.
		f.offset++

		// If we read a newline, return the line.
		if buffer[0] == '\n' {
			return string(line), nil
		}

		// Otherwise, append the byte to the line.
		line = append(line, buffer[0])
	}
}
