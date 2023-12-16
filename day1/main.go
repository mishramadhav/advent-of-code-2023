package main

import (
	"io"
	"log"
	"os"

	linereader "github.com/mishramadhav/advent-of-code-2023/day1/internal/fileutils/line_reader"
)

func main() {
	fileName := "input.txt"
	log.Printf("Reading file %s\n", fileName)
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("error while reading file - %v", err)
	}

	defer file.Close()

	var ans int
	lineReader := linereader.NewFileReaderByLine(file)
	for {
		line, err := lineReader.ReadLine()
		if err != nil {
			switch err {
			case io.EOF:
				log.Println("Reached end of file")
			default:
				log.Fatalf("error while reading line - %v", err)
			}

			break
		}

		log.Printf("Read line - %s\n", line)

		firstDigit := getFirstDigitOfLine(line)
		lastDigit := getLastDigitOfLine(line)
		number := firstDigit*10 + lastDigit

		ans += number
	}

	log.Printf("Answer is %d\n", ans)
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
	var ans int

	for _, c := range line {
		if c >= '0' && c <= '9' {
			ans = int(c - '0')
		}
	}

	return ans
}
