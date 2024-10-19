package main

import (
	"io"
	"log"
)

// LineReader is an interface for reading a file line by line.
type LineReader interface {
	ReadLine() (string, error)
}

func processFile(r LineReader) int {
	var ans int
	for {
		line, err := r.ReadLine()
		if err != nil {
			switch err {
			case io.EOF:
			default:
				log.Fatalf("error while reading line - %v", err)
			}

			break
		}

		firstDigit := getFirstDigitOfLine(line)
		lastDigit := getLastDigitOfLine(line)
		number := firstDigit*10 + lastDigit

		ans += number
	}

	return ans
}

func getFirstDigitOfLine(line string) int {
	var ans int

	for _, c := range line {
		if c >= '0' && c <= '9' {
			ans = int(c - '0')
			break
		}
	}

	return ans
}

func getLastDigitOfLine(line string) int {
	for i := len(line) - 1; i >= 0; i-- {
		if line[i] >= '0' && line[i] <= '9' {
			return int(line[i] - '0')
		}
	}
	return 0
}
