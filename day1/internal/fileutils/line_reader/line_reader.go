package linereader

import "os"

// FileLineReader reads a file line by line.
type FileLineReader struct {
	// file is the file to read.
	file *os.File
	// offset is the current offset in the file.
	offset int64
	// readError is the error that occurred during the last read.
	readError error
}

// NewFileReaderByLine creates a new FileReaderByLine
// for the given file.
//
// The file is not closed by this function.
func NewFileReaderByLine(file *os.File) *FileLineReader {
	return &FileLineReader{file: file}
}

func (f *FileLineReader) ReadLine() (string, error) {
	if f.readError != nil {
		return "", f.readError
	}

	buffer := make([]byte, 1)
	var line []byte

	for {
		// Read one byte at a time.
		_, err := f.file.ReadAt(buffer, f.offset)
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
