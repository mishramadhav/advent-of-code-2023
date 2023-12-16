package main

// LineReader is an interface for reading a file line by line.
type LineReader interface {
	ReadLine() (string, error)
}
