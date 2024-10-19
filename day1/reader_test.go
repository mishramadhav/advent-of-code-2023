package main

import (
	"os"
	"testing"

	"github.com/mishramadhav/advent-of-code-2023/day1/internal/fileutils"
)

func BenchmarkProcessFileUnbufferedLineReader(b *testing.B) {
	// Open the input file
	file, err := os.Open("input.txt")
	if err != nil {
		b.Fatalf("Failed to open input file: %v", err)
	}
	defer file.Close()

	// Get file info to determine its size
	fileInfo, err := file.Stat()
	if err != nil {
		b.Fatalf("Failed to get file info: %v", err)
	}

	// Reset the timer to exclude setup time
	b.ResetTimer()

	// Run the benchmark
	for i := 0; i < b.N; i++ {
		// Reset the file offset for each iteration
		file.Seek(0, 0)
		reader := fileutils.NewFileReaderByLine(file)

		// Process the file
		result := processFile(reader)

		// Prevent the compiler from optimizing away the result
		if result == 0 {
			b.Fatalf("Unexpected result: %d", result)
		}
	}

	// Report the number of bytes processed
	b.SetBytes(fileInfo.Size())
}
